package nested

type (

	// ConfigContextSchema : Models a subset on the ConfigContextSchema model for nested responses.
	ConfigContextSchema struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}

	// Job : Models a subset on the Job model for nested responses.
	Job struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// JobQueue : Models a subset on the JobQueue model for nested responses.
	JobQueue struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// JobResult : Models a subset on the JobResult model for nested responses.
	JobResult struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// SecretsGroup : Models a subset on the SecretsGroup model for nested responses.
	SecretsGroup struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}

	// User : Models a subset on the User model for nested responses.
	User struct {
		ID       string `json:"id"`
		Display  string `json:"display"`
		Username string `json:"username"`
		URL      string `json:"url"`
	}
)
