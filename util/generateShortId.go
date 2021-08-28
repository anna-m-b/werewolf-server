package util

import (
	"fmt"

	"github.com/teris-io/shortid"
)

func GenerateShortId() string {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		fmt.Printf("there was an error configuring a SID: %v", err)
	}
	id, err := sid.Generate()
	if err != nil {
		fmt.Printf("there was an error generating a SID: %v", err)
	}
	return id
}