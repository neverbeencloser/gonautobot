package types

type (
	// ConsolePort : Data representation of a Console Port in Nautobot.
	ConsolePort struct {
		ID           string                 `json:"id"`
		URL          string                 `json:"url"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Device       struct {
			ID          string `json:"id"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
		Name                  string            `json:"name"`
		Label                 string            `json:"label"`
		Description           string            `json:"description"`
		ConnectedEndpointType string            `json:"connected_endpoint_type"`
		ConnectedEndpoint     ConnectedEndpoint `json:"connected_endpoint"`
		Cable                 struct {
			ID    string `json:"id"`
			URL   string `json:"url"`
			Label string `json:"label"`
		} `json:"cable"`
		Tags []Tag `json:"tags"`
	}

	// ConnectedEndpoint : Struct representing the far-side of the console connection.
	ConnectedEndpoint struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Device struct {
			ID          string `json:"id"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
		Name  string `json:"name"`
		Cable string `json:"cable"`
		Port  int
	}
)
