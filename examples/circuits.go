package main

import (
	"fmt"
	"github.com/neverbeencloser/gonautobot/circuits"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/rs/zerolog/log"
	"net/url"
)

// CircuitExample : Example usage of the Circuit Nautobot methods.
func (c *ex) CircuitExample() {
	r0, err := c.Circuits.CircuitCreate(circuits.CircuitRequest{
		CircuitID:   "abc12345903",
		Status:      "Active",
		Description: "Test Circuit",
		Provider:    "8e384236-9ff5-4f71-8418-1d607177e10e",
		Type:        "8bbf5595-ec96-4fb4-a3f7-87f010031d4e",
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println("Created", r0.CircuitID)

	r1, err := c.Circuits.CircuitFilter(&url.Values{"cid": {r0.CircuitID}})
	if err != nil {
		log.Fatal().Err(err)
	}
	first, _ := core.First[circuits.Circuit](r1)
	fmt.Println(first.CircuitID)

	r2, err := c.Circuits.CircuitGet(first.ID)
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println(r2.CircuitID)

	if err := c.Circuits.CircuitDelete(first.ID); err != nil {
		log.Fatal().Err(err)
	}
}
