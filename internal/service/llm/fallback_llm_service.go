package llm

import (
	"context"

	"github.com/itsLeonB/aishiteru/internal/config"
	"github.com/rotisserie/eris"
)

type fallbackLLMService struct {
	services []LLMService
}

func newFallbackLLMService(configs *config.LLM) LLMService {
	services := make([]LLMService, len(configs.Providers))
	for i, llmProvider := range configs.Providers {
		services[i] = newSingleLLMService(llmProvider, configs)
	}
	return &fallbackLLMService{services: services}
}

func (f *fallbackLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	var lastErr error
	for _, service := range f.services {
		response, err := service.GetResponse(ctx, prompt)
		if err == nil {
			return response, nil
		}
		lastErr = err
	}
	return "", eris.Wrap(lastErr, "all LLM services failed")
}
