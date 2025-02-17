// File:		doubao.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package grpcService

import (
	"context"

	"github.com/yazl-tech/ai-bot/internal/service"
	"github.com/yazl-tech/ai-bot/pkg/proto/doubao"
)

type DoubaoGrpcService struct {
	doubao.UnimplementedDoubaoHandlerServer
}

func NewDoubaoGrpcService(s *service.AiBotService) *DoubaoGrpcService {
	return &DoubaoGrpcService{}
}

func (ds *DoubaoGrpcService) ChatCompletions(ctx context.Context, req *doubao.ChatRequest) (*doubao.ChatResponse, error) {
	return &doubao.ChatResponse{}, nil
}
