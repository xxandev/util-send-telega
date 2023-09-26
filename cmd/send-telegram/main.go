package main

import (
	"flag"
	"io"
	"log"
	"os"
	"util-send-telega/internal/telegram"
	"util-send-telega/internal/utils"
)

var (
	config                  Config
	status, msg, photo, doc string
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[SEND-TELEGRAM] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

	if ok, err := config.Init(); ok && err != nil {
		log.Printf("warning read config file: %v", err)
	}

	info := flag.Bool("info", false, "status info")
	warn := flag.Bool("warn", false, "status warning")
	err := flag.Bool("err", false, "status error")
	crit := flag.Bool("critic", false, "status critical")

	example := flag.Bool("example", false, "create example config (*.yaml)")
	mute := flag.Bool("mute", false, "mute log")
	chatID := flag.Int64("id", 0, "telegram chat id")
	token := flag.String("bot", "", "telegram bot token")
	flag.StringVar(&msg, "m", "", "message/caption text")
	flag.StringVar(&photo, "p", "", "photo/image path")
	flag.StringVar(&doc, "d", "", "document/file path")
	flag.Parse()

	if *example {
		config.Example()
		content, err := config.Marshal(YAML)
		if err != nil {
			log.Fatalf("error marshal example-config: %v\n", err)
		}
		if err := utils.CreateFile("send-telegram.yaml", content); err != nil {
			log.Fatalf("error create example-config: %v\n", err)
		}
		log.Println("example-config created")
		os.Exit(0)
	}

	config.Telegram.Set(*chatID, *token)
	if err := config.Check(); err != nil {
		log.Fatalf("error check config: %v\n", err)
	}

	if *mute {
		log.SetOutput(io.Discard)
	}

	switch {
	case *info:
		status = "✅ "
	case *warn:
		status = "⚠️ "
	case *err:
		status = "⛔️ "
	case *crit:
		status = "🆘 "
	}
}

func main() {
	telega := telegram.New(&config.Telegram)
	if len(photo) > 0 {
		if err := telega.SendPhoto(photo, status+msg); err != nil {
			log.Fatalf("error send photo: %v\n", err)
			return
		}
		log.Println("photo sent ok.")
		return
	}
	if len(doc) > 0 {
		if err := telega.SendDoc(doc, status+msg); err != nil {
			log.Fatalf("error send document: %v\n", err)
			return
		}
		log.Println("document sent ok.")
		return
	}
	if err := telega.SendMsg(status + msg); err != nil {
		log.Fatalf("error send message: %v\n", err)
		return
	}
	log.Println("message sent ok.")
}
