package config

type LLM struct {
	Providers        []string `split_words:"true" required:"true"`
	GoogleApiKey     string   `split_words:"true"`
	GoogleModel      string   `split_words:"true"`
	OpenRouterApiKey string   `split_words:"true"`
	OpenRouterModel  string   `split_words:"true"`
}

func (l LLM) Prefix() string {
	return "LLM"
}
