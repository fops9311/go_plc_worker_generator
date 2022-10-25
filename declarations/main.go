package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func main() {
	var plcsFile = os.Args[1]
	var templateDir = os.Args[2]
	var resultDir = os.Args[3]
	var plcDir = os.Args[5]
	b, err := os.ReadFile(plcsFile)
	if err != nil {
		log.Printf("[error] %v", err)
		return
	}
	data := &PLCJsonData{}
	err = json.Unmarshal(b, data)
	if err != nil {
		log.Printf("[error] %v", err)
		return
	}
	for _, v := range data.Plcs {
		err := makeTemplatedFile(templateDir+v.Type+".json", resultDir+v.Type+"_"+v.Id+".json", v)
		if err != nil {
			log.Printf("[error] %v", err)
			return
		}
		err = DeclarationToCode(resultDir+v.Type+"_"+v.Id+".json", plcDir+v.Type+"_"+v.Id+".go")
		if err != nil {
			log.Printf("[error] %v", err)
			return
		}
	}
}

type PLCJsonData struct {
	Plcs []plcinfo `json:"plcs"`
}
type plcinfo struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func makeTemplatedFile(templateFile string, resultFile string, data interface{}) error {
	b, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	//parsing template
	t := template.New(resultFile)
	t, err = t.Parse(string(b))
	if err != nil {
		return err
	}
	//executing template
	buf := bytes.NewBuffer([]byte{})
	t.Execute(buf, data)
	if err != nil {
		return err
	}
	result := buf.String()
	//creating result file
	f, err := os.Create(resultFile)
	if err != nil {
		return err
	}
	defer f.Close()
	//writing result to result file
	_, err = f.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}
