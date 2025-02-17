package provider

import (
	"context"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type Provider interface {
	Chat(ctx context.Context, messages []*botpb.Message, options *botpb.ChatOptions) (*botpb.ChatResponse, error)
}
