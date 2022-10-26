package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
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

	err = makeTemplatedFile(templateDir+"plcs.go", plcDir+"plcs.go", data)
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
	Id                string `json:"id"`
	Type              string `json:"type"`
	TickInterval100ms int    `json:"tick_interval_100ms"`
}
type InterfaceConnections map[string][]ComDriverLinkRedirect
type ComDriverLinkRedirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (c InterfaceConnections) UpdateLinks(td *TemplateData) {
	if links, ok := c[td.TypeName]; ok {
		log.Println("redirect record found")
		for i := range links {
			for j := range td.Inputs {
				if td.Inputs[j].ComDriverLinkId == links[i].From {
					td.Inputs[j].ComDriverLinkId = links[i].To
					log.Println("redirect match " + td.Inputs[j].ComDriverLinkId)
				}
			}
			for j := range td.Outputs {
				if td.Outputs[j].ComDriverLinkId == links[i].From {
					log.Println("redirect match")
					td.Outputs[j].ComDriverLinkId = links[i].To
				}
			}
		}
	}
}

func makeTemplatedFile(templateFile string, resultFile string, data interface{}) error {
	b, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	//parsing template
	t := template.New(resultFile).Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	})
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
