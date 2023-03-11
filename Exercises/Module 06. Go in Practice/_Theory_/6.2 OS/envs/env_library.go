package envs

//https://github.com/qiangxue/go-env

import (
	"fmt"
	"github.com/qiangxue/go-env"
	"os"
)

type Config struct { //Определяем структуру для выгрузки конфигураций
	Host string
	Port int
}

func Env_library() {
	_ = os.Setenv("APP_HOST", "127.0.0.1") // задаем переменные окружения
	_ = os.Setenv("APP_PORT", "8080")

	var cfg Config
	if err := env.Load(&cfg); err != nil { // сканим в Config
		panic(err)
	}
	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
}
