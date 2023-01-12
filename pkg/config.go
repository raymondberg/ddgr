package pkg

import (
	"fmt"
	"github.com/bigkevmcd/go-configparser"
)

type Config struct {
	apikey string
	appkey string
}

func ConfigFromFile(filepath string) (Config, error) {
	p, err := configparser.NewConfigParserFromFile(filepath)

	if err != nil {
		fmt.Println("Unable to load config: ", filepath)
		return Config{}, err
	}

	apikey, err1 := p.Get("Connection", "apikey")
	appkey, err2 := p.Get("Connection", "appkey")

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid config: ", filepath)
		return Config{}, err
	}

	return Config{apikey: apikey, appkey: appkey}, nil
}
