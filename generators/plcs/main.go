package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/fops9311/go_plc_worker_generator/expression"
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
		//translate expressions
		plc.Inputs, err = TranslateIO(plc.Inputs)
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		plc.Outputs, err = TranslateIO(plc.Outputs)
		if err != nil {
			log.Printf(" [error] %v\n", err)
			return
		}
		plc.States, err = TranslateStates(plc.States)
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
func TranslateIO(s []PlcInterface) (res []PlcInterface, err error) {
	for i := range s {
		s[i].Value, err = expression.Translate(s[i].Expr)
		if err != nil {
			log.Printf(" [error] %s\n", s[i].Expr)
			return s, err
		}
	}
	return s, nil
}
func TranslateStates(s []PlcState) (res []PlcState, err error) {
	for i := range s {
		for j := range s[i].OutputVector {
			s[i].OutputVector[j].Value, err = expression.Translate(s[i].OutputVector[j].Expr)
			if err != nil {
				log.Printf(" [error] %s\n", s[i].OutputVector[j].Expr)
				return s, err
			}
		}
		for k := range s[i].StateChangeCondition {
			s[i].StateChangeCondition[k].Condition, err = expression.Translate(s[i].StateChangeCondition[k].Expr)
			if err != nil {
				log.Printf(" [error] %s\n", s[i].StateChangeCondition[k].Expr)
				return s, err
			}
		}
	}
	return s, err
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
	Expr            string `json:"expr"`
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
	Expr        string `json:"expr"`
}
type OutputDescription struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Expr  string `json:"expr"`
}
