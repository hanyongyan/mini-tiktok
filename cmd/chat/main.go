package main

import (
	"log"
	chatservice "mini_tiktok/kitex_gen/service/ChatService/messageservice"
)

func main() {
	svr := chatservice.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
