package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"` //указываем имя тега заранее
	Id      int      `xml:"id,attr"` // id будет атрибутом тега, а не полем
	Name    string   `xml:"name"` // а вот имя будет полем
	Origin  []string `xml:"origin"` //одинаковые теги соберутся в слайс
}

func (p Plant) String() string { //определим метод String для нашей структуры, чтобы ее можно было удобно выводить
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func CreateXml() string {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", "  ") //xml тоже можно форматировать как и JSON для удобства вывода
	return string(out)
}

func DecodeXml(input string) {
	var p Plant
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
		panic(err)
	} //декодирование производится аналогично JSON - структура и анмаршалинг
	fmt.Println(p)
}

func NestedXml() {
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"` // если нам нет необходимости проверять вложенности, мы можем для сокращения пути указать вложенность, не конструируя вложенные обьекты
		//в JSON к сожалению так не прокатит
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ := xml.MarshalIndent(nesting, "", "  ")
	fmt.Println(string(out))

	res := &Nesting{}

	xml.Unmarshal(out, res)

	fmt.Println(res)
}