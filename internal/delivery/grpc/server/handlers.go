package server

import (
	"context"

	pb "github.com/itsLeonB/aishiteru/gen/go/llm"
	"github.com/itsLeonB/aishiteru/internal/logging"
)

func (s *Server) Prompt(ctx context.Context, req *pb.PromptRequest) (*pb.PromptResponse, error) {
	response, err := s.llmService.GetResponse(ctx, req.GetPrompt())
	if err != nil {
		logging.Logger.Errorf("error handling request: prompt: %s; err: %v", req.GetPrompt(), err)
		return nil, err
	}

	return &pb.PromptResponse{
		Response: response,
	}, nil
}
