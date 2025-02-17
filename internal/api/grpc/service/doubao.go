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
	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
	doubaopb "github.com/yazl-tech/ai-bot/pkg/proto/doubao"
)

type DoubaoGrpcService struct {
	doubaopb.UnimplementedDoubaoHandlerServer
}

func NewDoubaoGrpcService(s *service.AiBotService) *DoubaoGrpcService {
	return &DoubaoGrpcService{}
}

func (ds *DoubaoGrpcService) ChatCompletions(ctx context.Context, req *botpb.ChatRequest) (*botpb.ChatResponse, error) {
	return &botpb.ChatResponse{}, nil
}
