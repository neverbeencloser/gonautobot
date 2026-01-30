package main

import (
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
	"github.com/rs/zerolog/log"
)

// SecretExample : Example usage of the Secret Nautobot methods.
func (c *ex) SecretExample() {
	// Create a new secret
	secret, err := c.Extras.SecretCreate(types.NewSecret{
		Name:        "example-secret",
		Provider:    "environment-variable",
		Description: "Example secret for demonstration",
		Parameters: map[string]any{
			"variable": "MY_SECRET_VAR",
		},
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create secret")
	}
	log.Info().
		Str("uuid", secret.ID.String()).
		Str("name", secret.Name).
		Str("provider", secret.Provider).
		Str("description", secret.Description).
		Str("url", secret.URL).
		Time("created", secret.Created).
		Msg("Created secret")

	// Filter secrets by provider
	secrets, err := c.Extras.SecretFilter(&url.Values{"provider": {"environment-variable"}})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to filter secrets")
	}
	log.Info().
		Int("count", len(secrets)).
		Str("filter", "provider=environment-variable").
		Msg("Filtered secrets")

	first, _ := core.First[types.Secret](secrets)
	log.Info().
		Str("uuid", first.ID.String()).
		Str("name", first.Name).
		Str("provider", first.Provider).
		Msg("First secret from filter results")

	// Get a specific secret by ID
	retrieved, err := c.Extras.SecretGet(secret.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get secret")
	}
	log.Info().
		Str("uuid", retrieved.ID.String()).
		Str("name", retrieved.Name).
		Str("provider", retrieved.Provider).
		Str("description", retrieved.Description).
		Str("natural_slug", retrieved.NaturalSlug).
		Time("last_updated", retrieved.LastUpdated).
		Msg("Retrieved secret by UUID")

	// Update the secret
	updated, err := c.Extras.SecretUpdate(secret.ID, map[string]any{
		"description": "Updated description",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update secret")
	}
	log.Info().
		Str("uuid", updated.ID.String()).
		Str("name", updated.Name).
		Str("description", updated.Description).
		Time("last_updated", updated.LastUpdated).
		Msg("Updated secret")

	// Delete the secret
	if err := c.Extras.SecretDelete(secret.ID); err != nil {
		log.Fatal().Err(err).Msg("Failed to delete secret")
	}
	log.Info().
		Str("uuid", secret.ID.String()).
		Str("name", secret.Name).
		Msg("Deleted secret")
}
