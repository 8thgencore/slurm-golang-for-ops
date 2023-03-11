package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

type SlicesTags []string

func (tags *SlicesTags) UnmarshalYAML(value *yaml.Node) error { //существует возможность добавить кастомные разборы yaml для определенных полей
	if value != nil {
		*tags = strings.Split(value.Value,",") //разделяем строку по запятым, чтобы сформировать слайс тегов
	}
	return nil
}

type Messages struct {
	Tags SlicesTags `yaml:"tags"`
}

type Subs struct {
	Messages Messages `yaml:"messages"`
}

func ParseYamlWithCustomStruct() {
	config := &Subs{}

	yamlFile, err := ioutil.ReadFile("examples/yaml_subs_example.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)

}