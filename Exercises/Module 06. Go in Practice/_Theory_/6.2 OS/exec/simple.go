package exec

import (
	"log"
	"os/exec"
)

func RunSimpleApp() {

	cmd := exec.Command("firefox") //определяем команду и имя приложения доступное из PATH

	err := cmd.Run() //запускаем и собираем ошибки

	if err != nil {
		log.Fatal(err)
	}
}
