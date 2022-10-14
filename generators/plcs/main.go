package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"
)

var IFTYPE string
var TEMPLATEDIR string = "./generators/plcs/"
var RESULTDIR string = "./plc/"

func main() {
	func() {
		//read plc json file name
		if len(os.Args) < 2 {
			log.Fatalf(" [error] %s\n", "Not enough args")
		}
		IFTYPE = os.Args[1]
		//reading the file
		j, err := os.ReadFile(IFTYPE + ".json")
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		//unmarshaling json
		plc := &TemplateData{}
		err = json.Unmarshal(j, plc)
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		//reading template file
		b, err := os.ReadFile(TEMPLATEDIR + "default.template")
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		//parsing template
		t := template.New("template").Funcs(template.FuncMap{
			"ToUpper": strings.ToUpper,
		})
		t, err = t.Parse(string(b))
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		//executing template
		buf := bytes.NewBuffer([]byte{})
		t.Execute(buf, plc)
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		result := buf.String()
		//creating result file
		f, err := os.Create(RESULTDIR + "plc_" + IFTYPE + ".go")
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		defer f.Close()
		//writing result to result file
		_, err = f.WriteString(result)
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		//stdio result summary
		log.Printf("generated for %s\n", IFTYPE)
	}()
}

type TemplateData struct {
	TypeName       string         `json:"typename"`
	Inputs         []PlcInterface `json:"inputs"`
	Outputs        []PlcInterface `json:"outputs"`
	States         []PlcState     `json:"states"`
	StartStateId   int            `json:"start_state_id"`
	DefaultStateId int            `json:"default_state_id"`
	ComDriver      string         `json:"com_driver"`
}
type PlcInterface struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	ComDriverLinkId string `json:"com_driver_link_id"`
}
type PlcState struct {
	Id                   int                 `json:"id"`
	Comment              string              `json:"comment"`
	OutputVector         []OutputDescription `json:"output_vector"`
	StateChangeCondition []StateChange       `json:"state_change_condition"`
}
type StateChange struct {
	Destination int    `json:"destination"`
	Condition   string `json:"condition"`
}
type OutputDescription struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
