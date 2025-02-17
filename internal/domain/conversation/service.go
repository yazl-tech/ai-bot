package conversation

import (
	"context"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type ConversationService interface {
	// 创建新对话
	Create(ctx context.Context, userID, botID string) (*Conversation, error)
	// 发送消息
	SendMessage(ctx context.Context, convID string, content string) (*botpb.Message, error)
	// 获取对话历史
	GetHistory(ctx context.Context, convID string) (*Conversation, error)
	// 归档对话
	Archive(ctx context.Context, convID string) error
}
