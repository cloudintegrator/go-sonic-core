package utility

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"log/slog"
	"os"
	"plugin"
	"strings"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

const (
	SO_PATH = "so/"
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

func FindComponentInSO(symbol string, kind string) {
	path := GetSOPath(kind)

	p, err := plugin.Open(path)
	if err != nil {

	}
	s, err := p.Lookup(symbol)
	if err != nil {

	}
	fmt.Println(s)
}

func GetSOPath(kind string) string {
	path := "../../" + SO_PATH
	so := strings.Split(kind, ":")
	if len(so) != 0 {
		path = path + so[0] + so[1] + ".so"
	}
	return path
}
