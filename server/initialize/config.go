package initialize

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"wutool.cn/chat/server/global"
)

func Viper() *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	fmt.Println("config:" + config)
	v := viper.New()

	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
