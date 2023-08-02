package config

import (
	"fmt"
	"os"
	"path/filepath"

	"git.huoys.com/vn/go-common/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper             //viper结构
	changeHandlerFunc func() //监听到文件修改处理方法
	fileName          string
}

func NewConfig(fileName string) (*Config, error) {
	conf, err := load(fileName)
	if err != nil {
		return nil, err
	}

	c := &Config{
		Viper:    conf,
		fileName: fileName,
	}

	c.changeMonitor()

	return c, nil
}

func (p *Config) SetChangeHandlerFunc(f func()) {
	p.changeHandlerFunc = f
}

func (p *Config) changeMonitor() {
	//配置文件修改监听
	p.OnConfigChange(func(event fsnotify.Event) {
		if event.Op != fsnotify.Write {
			return
		}
		fmt.Println("enven:", event)
		err1 := p.ReadInConfig()
		if err1 != nil {
			log.Critical("server config was changed,err:", err1)
		}
		log.Info("config change:", event.Name)
		if p.changeHandlerFunc != nil {
			p.changeHandlerFunc()
		}
	})
	p.WatchConfig()
}

func load(fileName string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigFile(fileName)
	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

//获取app进程运行的绝对路径
func GetAppAbsPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("failed get app abs path,err:", err)
		return
	}
	return
}
