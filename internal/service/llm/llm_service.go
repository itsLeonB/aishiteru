package llm

import (
	"context"

	"github.com/itsLeonB/aishiteru/internal/appconstant"
	"github.com/itsLeonB/aishiteru/internal/config"
	"github.com/itsLeonB/aishiteru/internal/logging"
)

type LLMService interface {
	GetResponse(ctx context.Context, prompt string) (string, error)
}

func NewLLMService(configs *config.LLM) LLMService {
	if len(configs.Providers) > 1 {
		return newFallbackLLMService(configs)
	}

	return newSingleLLMService(configs.Providers[0], configs)
}

func newSingleLLMService(provider string, configs *config.LLM) LLMService {
	switch provider {
	case appconstant.GoogleLLM:
		return newGoogleLLMService(configs.GoogleApiKey, configs.GoogleModel)
	case appconstant.OpenRouter:
		return newOpenRouterService(configs.OpenRouterApiKey, configs.OpenRouterModel)
	default:
		logging.Logger.Warn("no LLM provider configured")
		return nil
	}
}
