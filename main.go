package main

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log"
	"plugin"
)

type Hello struct {
}

func (h Hello) Say() {

}

func main() {

}

func kafkaListener() {
	p, err := plugin.Open("/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/go-sonic-kafka-connector.so")

	s, err := p.Lookup("INSTANCE_KAFKA_LISTENER")

	out := s.(core.SonicSource)
	if err != nil {
		log.Fatal(err)
	}
	out.Init()
	log.Println(out)
}
