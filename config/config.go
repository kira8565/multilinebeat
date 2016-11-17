// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period          time.Duration `config:"period"`
	ListenPort      string `config:"listenport"`
	GroupKey        string `config:"groupkey"`
	MultilineRegx   string `config:"multilineregx"`
	MessageFieldKey string `config:"messagefiledkey"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	ListenPort:"9999",
	GroupKey:"",
	MultilineRegx:"\n",
	MessageFieldKey:"message",
}
