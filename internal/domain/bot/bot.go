package bot

import (
	"context"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type Provider interface {
	Chat(ctx context.Context, messages []*botpb.Message, options *botpb.ChatOptions) (*botpb.ChatResponse, error)
}

type ProviderFactory struct {
	providerMap map[botpb.ProviderType]Provider
}

func NewBotFactory() *ProviderFactory {
	return &ProviderFactory{
		providerMap: make(map[botpb.ProviderType]Provider),
	}
}

func (f *ProviderFactory) RegisterProvider(pt botpb.ProviderType, p Provider) {
	f.providerMap[pt] = p
}

func (f *ProviderFactory) GetProvider(pt botpb.ProviderType) (Provider, bool) {
	p, exists := f.providerMap[pt]
	return p, exists
}
