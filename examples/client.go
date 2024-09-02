package main

import (
	"fmt"
	nautobot "github.com/josh-silvas/gonautobot"
	"github.com/josh-silvas/gonautobot/circuits"
	"github.com/josh-silvas/gonautobot/core"
	"net/url"
)

type ex struct {
	*nautobot.Client
}

func main() {
	c := ex{
		nautobot.New(
			nautobot.WithToken("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			nautobot.WithEndpoint("https://demo.nautobot.com/"),
		),
	}

	c.circuit()
}

func (c *ex) circuit() {
	r0, err := c.Circuits.CircuitCreate(circuits.CircuitRequest{
		CircuitID:   "abc12345903",
		Status:      "Active",
		Description: "Test Circuit",
		Provider:    "8e384236-9ff5-4f71-8418-1d607177e10e",
		Type:        "8bbf5595-ec96-4fb4-a3f7-87f010031d4e",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created", r0.CircuitID)

	r1, err := c.Circuits.CircuitFilter(&url.Values{"cid": {r0.CircuitID}})
	if err != nil {
		panic(err)
	}
	first, _ := core.First[circuits.Circuit](r1)
	fmt.Println(first.CircuitID)

	r2, err := c.Circuits.CircuitGet(first.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(r2.CircuitID)

	if err := c.Circuits.CircuitDelete(first.ID); err != nil {
		panic(err)
	}
}
