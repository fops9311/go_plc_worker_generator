package main

import (
	"encoding/json"
	"log"
	"os"
)

func makeexamplejson() {

	TD := TemplateData{
		TypeName: "test_plc",
		Inputs: []PlcInterface{
			{
				Name:            "In",
				Type:            "cif.CIbool",
				Value:           "false",
				ComDriverLinkId: "metric/plc1/plc/in/pv",
			},
		},
		Outputs: []PlcInterface{
			{
				Name:            "Out",
				Type:            "cif.CIbool",
				Value:           "false",
				ComDriverLinkId: "metric/plc1/plc/out/pv",
			},
			{
				Name:            "State",
				Type:            "cif.CIint",
				Value:           "0",
				ComDriverLinkId: "metric/plc1/plc/state/pv",
			},
		},
		States: []PlcState{
			{
				Id:      0,
				Comment: "Out off, wait for In == true",
				OutputVector: []OutputDescription{
					{
						Name:  "Out",
						Value: "false",
					},
				},
				StateChangeCondition: []StateChange{
					{
						Condition:   "plc.In.Value",
						Destination: 1,
					},
				},
			},
			{
				Id:      1,
				Comment: "Out on, wait for In == false",
				OutputVector: []OutputDescription{
					{
						Name:  "Out",
						Value: "true",
					},
				},
				StateChangeCondition: []StateChange{
					{
						Condition:   "!plc.In.Value",
						Destination: 2,
					},
				},
			},
			{
				Id:      2,
				Comment: "Out on, wait for In == true",
				OutputVector: []OutputDescription{
					{
						Name:  "Out",
						Value: "true",
					},
				},
				StateChangeCondition: []StateChange{
					{
						Condition:   "plc.In.Value",
						Destination: 3,
					},
				},
			},
			{
				Id:      3,
				Comment: "Out off, wait for In == false",
				OutputVector: []OutputDescription{
					{
						Name:  "Out",
						Value: "false",
					},
				},
				StateChangeCondition: []StateChange{
					{
						Condition:   "!plc.In.Value",
						Destination: 0,
					},
				},
			},
		},
	}
	b, err := json.MarshalIndent(TD, "", "	")
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	err = os.WriteFile("./plc.json", b, 0644)
	if err != nil {
		log.Fatalf(" [error] %v\n", err)
	}
	log.Println(string(b))
}
