package main

import (
	"fmt"

	"github.com/neverbeencloser/gonautobot/types"
	"github.com/rs/zerolog/log"
)

// RoleExample : Example usage of the Role Nautobot methods.
func (c *ex) RoleExample() {
	// Create a new role
	role, err := c.Extras.RoleCreate(types.NewRole{
		Name:         "example-role",
		ContentTypes: []string{"dcim.device"},
		Color:        "9e9e9e",
		Description:  "Example role for demonstration",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create role")
	}
	fmt.Println("Created role:", role.Name)

	// Get all roles
	roles, err := c.Extras.RoleAll()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get all roles")
	}
	fmt.Println("Total roles:", len(roles))

	// Get a specific role by ID
	retrieved, err := c.Extras.RoleGet(role.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get role")
	}
	fmt.Println("Retrieved role:", retrieved.Name)

	// Update the role
	updated, err := c.Extras.RoleUpdate(role.ID, map[string]any{
		"description": "Updated description",
		"weight":      100,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update role")
	}
	fmt.Println("Updated role description:", updated.Description)

	// Delete the role
	if err := c.Extras.RoleDelete(role.ID); err != nil {
		log.Fatal().Err(err).Msg("Failed to delete role")
	}
	fmt.Println("Deleted role:", role.Name)
}
