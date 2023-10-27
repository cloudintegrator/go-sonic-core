package utility

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func YamlToStruct(file string, in interface{}) (err error) {
	buf, e := os.ReadFile(file)
	if e != nil {
		log.Fatal(e)
	}
	e = yaml.Unmarshal(buf, in)
	if e != nil {
		log.Fatal(e)
	}
	return e
}
