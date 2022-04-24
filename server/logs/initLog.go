package logs

import (
	"log"
	"os"
)

func InitLog() {
	f, err := os.OpenFile("./logs/myblog.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
}
