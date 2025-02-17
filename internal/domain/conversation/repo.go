// File:		repo.go
// Created by:	Hoven
// Created on:	2025-02-17
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package conversation

import "context"

type ConversationRepository interface {
	Create(ctx context.Context, conv *Conversation) error
	Get(ctx context.Context, id string) (*Conversation, error)
	Update(ctx context.Context, conv *Conversation) error
	List(ctx context.Context, opts ListOptions) ([]*Conversation, error)
	Delete(ctx context.Context, id string) error
}

type ListOptions struct {
	BotId  int
	Status ConversationStatus
	Limit  int
	Offset int
}
