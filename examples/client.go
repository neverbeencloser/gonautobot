package main

import (
	nautobot "github.com/neverbeencloser/gonautobot"
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
	c.DeviceExample()
	//c.IPAddressExample()
	//c.CircuitExample()
}
