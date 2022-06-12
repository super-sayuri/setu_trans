package conf

import (
	"github.com/BurntSushi/toml"
)

var _conf *Config

type Config struct {
	Telegram *TelegramConfig `toml:"telegram"`
	Log      *LogConfig      `toml:"log"`
}

type TelegramConfig struct {
	BotToken     string  `toml:"bot_token"`
	TrustedUsers []int64 `toml:"trusted_users"`
}

type LogConfig struct {
	Format string `toml:"format"`
	Output string `toml:"output"`
	Level  string `toml:"level"`
	Path   string `toml:"path"`
}

func Init(path, keyPath string) error {

	_conf = &Config{}

	_, err := toml.DecodeFile(path, _conf)
	if err != nil {
		return err
	}
	err = initLog(_conf.Log)
	if err != nil {
		return err
	}

	return nil
}

func GetConf() *Config {
	return _conf
}
