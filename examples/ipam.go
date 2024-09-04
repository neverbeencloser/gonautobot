package main

import (
	"fmt"
	"net/url"
)

// IPAddressExample : Example usage of the IPAddress Nautobot methods.
func (c *ex) IPAddressExample() {
	r0, err := c.Ipam.IPAddressFilter(&url.Values{"tenant": {"6202c221-2a49-4700-a879-32a6f912428b"}})
	if err != nil {
		panic(err)
	}
	fmt.Println("Filtered ", r0)
}
