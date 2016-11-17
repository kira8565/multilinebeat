package beater

import (
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/kira8565/multilinebeat/config"
	"github.com/firstrow/tcp_server"
	"strings"
	"regexp"
	"sync"
	"encoding/json"
	"fmt"
)

type Multilinebeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Multilinebeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Multilinebeat) Run(b *beat.Beat) error {

	bt.client = b.Publisher.Connect()

	listenPort := bt.config.ListenPort
	server := tcp_server.New("127.0.0.1:" + listenPort)
	server.OnNewClient(func(c *tcp_server.Client) {
		println("A New Clinet")
	})

	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		processLine(bt.client, bt.config.GroupKey,
			bt.config.MessageFieldKey,
			bt.config.MultilineRegx,
			message)
	})

	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		println("Client Close")
	})

	go server.Listen()

	sw := sync.WaitGroup{}
	sw.Add(1)

	select {
	case <-bt.done:
		sw.Done()
		return nil
	}

}

var globalQueue map[string]string = make(map[string]string)

func processLine(client publisher.Client, groupKey string, messageKey string, regex string, jsonLog string) {
	println(jsonLog)
	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(jsonLog), &jsonObj)
	log := ""
	if v, ok := jsonObj[messageKey].(string); ok {
		log = v
	} else {
		return
	}

	group := ""
	if v, ok := jsonObj[groupKey]; ok {
		group = v.(string)
	}

	if value, rs := globalQueue[group]; rs {
		combineContent := value + log

		//TODO:这里可以提升下性能的
		firtstLine := strings.Split(combineContent, "\n")[0]

		lastLineMatch, err := regexp.MatchString(regex, log)
		if err != nil {
			panic(err)
		}

		firstLineMatch, err := regexp.MatchString(regex, firtstLine)
		if err != nil {
			panic(err)
		}

		if lastLineMatch && firstLineMatch&& log != firtstLine {
			globalQueue[group] = log

			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       "multilineBeat",
				"message":    value,
			}
			client.PublishEvent(event)
			logp.Info("Event sent")
			println("SendEvent")

		} else {
			globalQueue[group] = combineContent
		}

	} else {
		globalQueue[group] = log
	}
}

func (bt *Multilinebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
