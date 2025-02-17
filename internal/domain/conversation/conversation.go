package conversation

import (
	"time"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type Conversation struct {
	ID       string
	BotID    string
	Title    string
	Messages []*botpb.Message
	Status   *ConversationStatus
	Metadata map[string]interface{}
	CreateAt time.Time
	UpdateAt time.Time
}

type ConversationStatus string

const (
	StatusActive   ConversationStatus = "active"
	StatusArchived ConversationStatus = "archived"
)
