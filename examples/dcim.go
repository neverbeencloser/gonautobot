package main

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/dcim"
	"github.com/rs/zerolog/log"
	"net/url"
)

// DeviceExample : Example usage of the Device Nautobot methods.
func (c *ex) DeviceExample() {
	req := dcim.NewDevice{
		Name:       "tst01-device-01",
		Role:       "edge",
		Status:     "Active",
		DeviceType: "e3fc837a-2474-4a94-b099-bd90a2d39468", // DCS-7280CR2-60
		Location:   "9e39051b-e968-4016-b0cf-63a5607375de", // ams01
	}
	res0, err := c.Dcim.DeviceCreate(req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create device")
	}
	fmt.Printf("Created device %s with ID %s\n", res0.Name, res0.ID)

	res1, err := c.Dcim.DeviceUpdate(res0.ID, map[string]any{"name": "tst01-device-99"})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update device")
	}
	fmt.Printf("Updated device %s with ID %s\n", res1.Name, res1.ID)

	r0, err := c.Dcim.DeviceFilter(&url.Values{"name__ic": {res1.Name}})
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Printf("Found a device %s in search results\n", r0[0].Name)

	if err := c.Dcim.DeviceDelete(res1.ID); err != nil {
		log.Fatal().Err(err).Msg("Failed to delete device")
	}
	fmt.Printf("Successfully deleted device %s\n", res1.Name)
}
