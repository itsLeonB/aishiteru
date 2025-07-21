package llm_test

import (
	"context"
	"testing"

	"github.com/itsLeonB/aishiteru/internal/appconstant"
	"github.com/itsLeonB/aishiteru/internal/config"
	"github.com/itsLeonB/aishiteru/internal/logging"
	"github.com/itsLeonB/aishiteru/internal/service/llm"
	"github.com/itsLeonB/ezutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockLLMService is a mock implementation of LLMService
type MockLLMService struct {
	mock.Mock
}

func (m *MockLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	args := m.Called(ctx, prompt)
	return args.String(0), args.Error(1)
}

// LLMServiceTestSuite defines the test suite for LLM services
type LLMServiceTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (suite *LLMServiceTestSuite) SetupTest() {
	logging.InitLogger(&ezutil.App{})
	suite.ctx = context.Background()
}

// TestNewLLMService_SingleProvider tests creating a single provider service
func (suite *LLMServiceTestSuite) TestNewLLMService_SingleProvider() {
	configs := &config.LLM{
		Providers:        []string{appconstant.OpenRouter},
		OpenRouterApiKey: "test-openrouter-key",
		OpenRouterModel:  "anthropic/claude-3-sonnet",
	}

	service := llm.NewLLMService(configs)
	assert.NotNil(suite.T(), service)
}

// Run all test suites
func TestLLMServiceTestSuite(t *testing.T) {
	suite.Run(t, new(LLMServiceTestSuite))
}

// Additional unit tests for specific scenarios
func TestLLMService_Integration(t *testing.T) {
	t.Run("should create Open Router service when single provider configured", func(t *testing.T) {
		configs := &config.LLM{
			Providers:        []string{appconstant.OpenRouter},
			OpenRouterApiKey: "test-openrouter-key",
			OpenRouterModel:  "anthropic/claude-3-sonnet",
		}

		// This would need proper mocking to avoid hitting real APIs
		service := llm.NewLLMService(configs)
		assert.NotNil(t, service)
	})
}

// Benchmark tests
func BenchmarkLLMService_GetResponse(b *testing.B) {
	mockService := new(MockLLMService)
	ctx := context.Background()
	prompt := "benchmark test prompt"

	mockService.On("GetResponse", ctx, prompt).Return("benchmark response", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mockService.GetResponse(ctx, prompt)
	}
}
