// File:		service.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package doubao

import (
	"context"
	"fmt"
	"time"

	"github.com/go-puzzles/puzzles/plog"
	"github.com/go-puzzles/puzzles/putils"
	"github.com/pkg/errors"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/yazl-tech/ai-bot/pkg/exception"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

type ProviderDoubao struct {
	doubaoClient *arkruntime.Client
	conf         *botpb.BotConfig
}

func NewDoubaoProvider(conf *botpb.BotConfig) *ProviderDoubao {
	dc := arkruntime.NewClientWithApiKey(
		conf.ApiKey,
		arkruntime.WithBaseUrl(conf.GetApi()),
		arkruntime.WithTimeout(2*time.Minute),
		arkruntime.WithRetryTimes(2),
	)
	return &ProviderDoubao{
		doubaoClient: dc,
	}
}

type doubaoRequest struct {
	Model       string          `json:"model"`
	Messages    []doubaoMessage `json:"messages"`
	Stream      bool            `json:"stream"`
	Temperature float32         `json:"temperature"`
	MaxToken    int             `json:"max_token"`
}

type doubaoMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
}

func (d *ProviderDoubao) parseStirngMessage(idx int, msg *botpb.Message, content *botpb.Message_StringContent) (dbMsg *model.ChatCompletionMessage, err error) {
	if msg.GetRole() == botpb.Message_system && idx != 0 {
		return dbMsg, exception.ErrSystemMessageMustInIndexZero
	}

	str := content.StringContent
	dbMsg = &model.ChatCompletionMessage{
		Role: msg.GetRole().String(),
		Content: &model.ChatCompletionMessageContent{
			StringValue: &str,
		},
	}

	return dbMsg, nil
}

func (d *ProviderDoubao) parseTypeMessage(msg *botpb.Message, content *botpb.Message_TypeContent) (dbMsg *model.ChatCompletionMessage, err error) {
	var (
		dbContent  *model.ChatCompletionMessageContentPart
		dbContents []*model.ChatCompletionMessageContentPart
	)

	t := content.TypeContent.GetType()
	switch t {
	case botpb.TypeMessage_string:
		dbContent = &model.ChatCompletionMessageContentPart{
			Type: model.ChatCompletionMessageContentPartTypeText,
			Text: content.TypeContent.GetText(),
		}
	case botpb.TypeMessage_image:
		dbContent = &model.ChatCompletionMessageContentPart{
			Type: model.ChatCompletionMessageContentPartTypeImageURL,
			ImageURL: &model.ChatMessageImageURL{
				URL: content.TypeContent.GetImageUrl().GetUrl(),
			},
		}
	default:
		return dbMsg, exception.ErrInvalidContentType
	}

	dbContents = append(dbContents, dbContent)
	dbMsg = &model.ChatCompletionMessage{
		Role: msg.GetRole().String(),
		Content: &model.ChatCompletionMessageContent{
			ListValue: dbContents,
		},
	}
	return dbMsg, nil
}

func (d *ProviderDoubao) convertMessages(messages []*botpb.Message) (reqMessages []*model.ChatCompletionMessage, err error) {
	reqMessages = make([]*model.ChatCompletionMessage, 0, len(messages))

	for idx, msg := range messages {
		var dbMsg *model.ChatCompletionMessage

		switch c := msg.Content.(type) {
		case *botpb.Message_StringContent:
			dbMsg, err = d.parseStirngMessage(idx, msg, c)
		case *botpb.Message_TypeContent:
			dbMsg, err = d.parseTypeMessage(msg, c)
		default:
			return nil, exception.ErrInvalidMessageType
		}
		if err != nil {
			return nil, errors.Wrap(err, "parseMessage")
		}

		reqMessages = append(reqMessages, dbMsg)
	}
	return reqMessages, nil
}

func (d *ProviderDoubao) convertResponse(resp *model.ChatCompletionResponse) (*botpb.ChatResponse, error) {
	if len(resp.Choices) == 0 {
		return nil, errors.New("response choices is empty")
	}

	botChoices := putils.Convert(resp.Choices, func(choice *model.ChatCompletionChoice) *botpb.Choice {
		return &botpb.Choice{
			Index: int64(choice.Index),
			Message: &botpb.Message{
				Role: botpb.Message_assistant,
				Content: &botpb.Message_StringContent{
					StringContent: *choice.Message.Content.StringValue,
				},
			},
			FinishReason: string(choice.FinishReason),
		}
	})

	botResp := &botpb.ChatResponse{
		Id:      resp.ID,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: botChoices,
	}

	return botResp, nil
}

func (d *ProviderDoubao) Chat(ctx context.Context, messages []*botpb.Message, options *botpb.ChatOptions) (*botpb.ChatResponse, error) {
	msgs, err := d.convertMessages(messages)
	if err != nil {
		return nil, errors.Wrap(err, "convertMessages")
	}

	req := &model.ChatCompletionRequest{
		Model:       options.GetModel(),
		Messages:    msgs,
		Stream:      options.GetStream(),
		MaxTokens:   int(options.GetMaxToken()),
		Temperature: options.GetTemperature(),
	}

	fmt.Println(plog.Jsonify(req))

	resp, err := d.doubaoClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "createChatCompletion")
	}

	return d.convertResponse(&resp)
}
