package main

import (
	"flag"
	"fmt"
	"os"

	"super-sayuri.github.com/setu_trans/conf"
	"super-sayuri.github.com/setu_trans/ims/telegram"
)

var (
	confPath string
	keyPath  string
)

func init() {
	flag.StringVar(&confPath, "c", "", "config file path")
	flag.StringVar(&keyPath, "k", "keys.json", "config file path")
}

func main() {
	flag.Parse()
	var err error
	err = conf.Init(confPath, keyPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(200)
	}

	errChan := make(chan struct{})
	chanMap := make(map[string]chan string, 0)
	chanMap["test"] = make(chan string)

	go telegram.StartReceiver(conf.GetConf().Telegram, chanMap, errChan)
	registerSender(chanMap, errChan)
}

func registerSender(msgChan map[string]chan string, errChan chan struct{}) {
	for {
		select {
		case <-errChan:
			return
		case m := <-msgChan["test"]:
			fmt.Println(m)
		}
	}
}
