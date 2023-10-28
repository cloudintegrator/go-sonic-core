package main

import (
	"fmt"
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

type IHello interface {
	Say()
}
type Hello struct {
	name string
}

func (h *Hello) Say() {
	fmt.Println(h.name)
}

func main() {
	p, _ := os.Getwd()
	p = filepath.Base(p)
	fmt.Println(p)
}
func InterfaceExample() {
	var x IHello
	x = &Hello{
		name: "anupam",
	}
	fmt.Println(x.(*Hello).name)
}
func kafkaListener() {
	p, err := plugin.Open("/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/kafka-listener.so")

	s, err := p.Lookup("INSTANCE_KAFKA_LISTENER")

	out := s.(core.SonicSource)
	if err != nil {
		log.Fatal(err)
	}
	out.Init()
	log.Println(out)
}
