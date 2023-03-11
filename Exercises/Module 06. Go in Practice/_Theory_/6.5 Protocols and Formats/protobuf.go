package main

import (
	"6_5/example/gen"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProtobufCase() {

	elliot := gen.Person{
		Name: "Elliot",
		Age: 24,
	}

	data, err := proto.Marshal(&elliot) //по уже сгенерированной структуре, мы маршалим сообщение и получем масив байт
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// печатаем наш обьект
	fmt.Println(data)


	newElliot := &gen.Person{}
	err = proto.Unmarshal(data, newElliot) // обратный ход, анмаршалинг аналогичен остальным протоколам
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	//печатаем поля обьекта
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}