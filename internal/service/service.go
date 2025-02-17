// File:		service.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package service

import (
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
