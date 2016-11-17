package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/kira8565/multilinebeat/beater"
)

func main() {
	err := beat.Run("multilinebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
