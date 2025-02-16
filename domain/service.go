// File:		service.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package doubao

import (
	"context"

	doubao "github.com/yazl-tech/ai-bot/pkg/dto"
)

type Service interface {
	ChatCompletions(ctx context.Context, req *doubao.ChatRequest) (*doubao.ChatResponse, error)
}

type DefaultDoubaoAiService struct{}

func NewDefaultDoubaoAiService() Service {
	return &DefaultDoubaoAiService{}
}

func (s *DefaultDoubaoAiService) ChatCompletions(ctx context.Context, req *doubao.ChatRequest) (*doubao.ChatResponse, error) {
	return nil, nil
}
