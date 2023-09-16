package structs

import (
	"encoding/json"
	"fmt"
)

type Payload struct {
	Id      int    `json: ID`
	Status  bool   `json: STATUS`
	Created string `json: DATE`
}

func ShowPayload() {
	marshalled, _ := json.MarshalIndent(Payload{1, true, "10.09.23"}, "", " ")

	fmt.Printf("payload is: %s\n", marshalled)
}
