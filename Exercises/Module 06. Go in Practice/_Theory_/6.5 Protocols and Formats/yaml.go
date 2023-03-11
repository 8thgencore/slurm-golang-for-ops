package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type BuildConf struct {
	Definitions map[string]interface{} `yaml:"definitions"`
	Pipelines   map[string]interface{} `yaml:"pipelines"`
}

type Conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func GetConf() {
	config := &Conf{} // в целом обьект можно и создать заранее, он так же проинициализируется

	yamlFile, err := ioutil.ReadFile("examples/example.yml") // читаем файл
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config) // аналогично JSON и XML, поддерживаются даже анкоры yaml
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)
}

func CreateConf() {
	config := &Conf{
		Hits: 11,
		Time: 1635859067,
	}

	out, _ := yaml.Marshal(config) //все по аналогии, интерфейс единообразный

	fmt.Println(string(out))
}

func ReadBigConf() {
	config := &BuildConf{} //когда мы точно не знаем что у нас будет внутри, мы можем прочитать yaml в map[string]interface{}

	yamlFile, err := ioutil.ReadFile("examples/anchor_example.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)
}
