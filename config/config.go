package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Config is the bot configuration representation, read
// from a configuration file.
type Config struct {
	//Telegram stuff
	TelegramTokenBot string
	ChatID           int64

	//List of subreddits
	Sources []string

	//Video related config
	VideoDownload bool
	DownloadPath  string
}

// ReadConfig loads the values from the config file
func ReadConfig(path string) (Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return Config{}, err
	}

	if conf.TelegramTokenBot == "" {
		return newErr("missing Bot token")
	} else if conf.ChatID == 0 {
		return newErr("missing ID")
	} else if len(conf.Sources) == 0 {
		return newErr("missing reddit sources")
	} else if conf.VideoDownload == true { //download of the videos is an optional feature
		if conf.DownloadPath == "" {
			return newErr("missing download path")
		}
	}

	return conf, nil
}

func newErr(message string) (Config, error) {
	return Config{}, errors.New(message)
}
