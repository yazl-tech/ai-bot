package bot

import (
	"context"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type BotService interface {
	Chat(ctx context.Context, botID string, message *botpb.Message) (*botpb.ChatResponse, error)
	CreateBot(ctx context.Context, bot *Bot) error
	GetBot(ctx context.Context, botID string) (*Bot, error)
}
