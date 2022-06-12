package telegram

import (
	"context"
	"encoding/json"

	"super-sayuri.github.com/setu_trans/conf"
	"super-sayuri.github.com/setu_trans/util"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartReceiver(config *conf.TelegramConfig, chanMap map[string]chan string, errChan chan struct{}) {
	ctx := context.WithValue(context.Background(), "job", "tgReceiver")
	l := conf.GetLog(ctx)

	bot, err := tgbot.NewBotAPI(config.BotToken)
	if err != nil {
		l.Errorf("cannot init telegram bot")
		errChan <- struct{}{}
		return
	}

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			sendMessage(ctx, chanMap, update.Message)
		}
	}
}

func sendMessage(ctx context.Context, chanMap map[string]chan string, msg *tgbot.Message) {
	ctx = context.WithValue(ctx, "msgId", msg.MessageID)
	l := conf.GetLog(ctx)
	msgStr, _ := json.Marshal(msg)
	userId := msg.From.ID
	if !util.ContainsIn(conf.GetConf().Telegram.TrustedUsers, userId) {
		l.Warn("untrusted user: ", userId, "msg: ", string(msgStr))
		return
	}
	for _, setuChan := range chanMap {
		setuChan <- string(msgStr)
	}
}
