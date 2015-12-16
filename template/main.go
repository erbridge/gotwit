package main

import (
	"os"

	"github.com/erbridge/gotwit"
	"github.com/erbridge/gotwit/twitter"
)

func main() {
	var (
		con twitter.ConsumerConfig
		acc twitter.AccessConfig
	)

	f := "secrets.json"
	if _, err := os.Stat(f); err == nil {
		con, acc, _ = twitter.LoadConfigFile(f)
	} else {
		con, acc, _ = twitter.LoadConfigEnv()
	}

	b := gotwit.NewBot("BotScreenName", con, acc)

	// Register callbacks here.

	if err := b.Start(); err != nil {
		panic(err)
	}

	if err := b.Stop(); err != nil {
		panic(err)
	}
}
