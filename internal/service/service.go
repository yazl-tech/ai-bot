// File:		service.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package service

import (
	"context"

	"github.com/go-puzzles/puzzles/plog"
	"github.com/yazl-tech/ai-bot/internal/domain/bot"
	"github.com/yazl-tech/ai-bot/pkg/exception"
	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type AiBotService struct {
	botFactory *bot.ProviderFactory
}

func NewAiBotService(botFactory *bot.ProviderFactory) *AiBotService {
	return &AiBotService{
		botFactory: botFactory,
	}
}

func (s *AiBotService) GetBot(pt botpb.ProviderType) (bot.Provider, error) {
	provider, exists := s.botFactory.GetProvider(pt)
	if !exists {
		return nil, exception.ErrProviderNotFound
	}

	return provider, nil
}

func (s *AiBotService) Chat(ctx context.Context, pt botpb.ProviderType, req *botpb.ChatRequest) (*botpb.ChatResponse, error) {
	bot, err := s.GetBot(pt)
	if err != nil {
		return nil, err
	}

	resp, err := bot.Chat(ctx, req.GetMessages(), req.GetOptions())
	if err != nil {
		plog.Errorc(ctx, "%v chat error: %v", pt.String(), err)
		return nil, err
	}

	usage := resp.GetUsage()
	plog.Infoc(
		ctx, "%v chat usage -> (TotalToken: %d, promptToken: %d, completionTokens: %d)",
		pt.String(), usage.GetTotalTokens(), usage.GetPromptTokens(), usage.GetCompletionTokens(),
	)

	return resp, nil
}
