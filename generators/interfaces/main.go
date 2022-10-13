package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

var IFTYPE string

func main() {
	if len(os.Args) < 2 {
		log.Fatalf(" [error] %s\n", "Not enough args")
	}
	IFTYPE = os.Args[1]

	b, err := os.ReadFile("./generators/interfaces/default.template")
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	t := template.New("interfaces")
	t, err = t.Parse(string(b))
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	buf := bytes.NewBuffer([]byte{})
	t.Execute(buf, IF{TypeName: IFTYPE})
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	result := buf.String()

	f, err := os.Create("./cachedinterface/CI" + IFTYPE + ".go")
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	defer f.Close()
	_, err = f.WriteString(result)
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	/*
		os.WriteFile("./cachedinterface/CI"+IFTYPE+".go", []byte(result), 0644)
		if err != nil {
			log.Fatalf(" [error] %v\n", err)
		}*/
	log.Printf("generated for %s\n", IFTYPE)
}

type IF struct {
	TypeName string
}
