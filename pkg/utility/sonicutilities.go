package utility

import (
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

func FindComponentInSO(sopath string, symbol string, kind string) interface{} {
	path := GetSOPath(sopath, kind)

	p, err := plugin.Open(path)
	if err != nil {

	}
	s, err := p.Lookup(symbol)
	if err != nil {

	}
	return s
}

func GetSOPath(sopath string, kind string) string {
	so := strings.Split(kind, ":")

	soname := ""
	if len(so) != 0 {
		soname = soname + so[0] + "-" + so[1] + ".so"
	}
	sopath = sopath + soname
	return sopath
}
