// File:		doubao_test.go
// Created by:	Hoven
// Created on:	2025-02-17
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package doubao_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yazl-tech/ai-bot/internal/domain/provider/doubao"

	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"
)

const (
	API_KEY_ENV        = "API_KEY_DOUBAO"
	MODEL_ENDPOINT_ENV = "MODEL_ENDPOINT_DOUBAO"
)

func getTestImageBase64(t *testing.T) string {
	imagePath := fmt.Sprintf("%s/src/github.com/yazl-tech/ai-bot/content/test.webp", os.Getenv("GOPATH"))
	fmt.Println(imagePath)
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		t.Fatalf("Failed to read image file: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(imageData)

	return fmt.Sprintf("data:image/webp;base64,%s", encoded)
}

func TestDoubaoChat(t *testing.T) {
	apiKey := os.Getenv(API_KEY_ENV)
	if apiKey == "" {
		t.Skip("API_KEY_DOUBAO is not set")
	}
	model := os.Getenv(MODEL_ENDPOINT_ENV)
	if model == "" {
		t.Skip("MODEL_ENDPOINT_DOUBAO is not set")
	}

	t.Logf("apiKey: %v, model: %v", apiKey, model)

	conf := &botpb.BotConfig{
		Api:    "https://ark.cn-beijing.volces.com/api/v3",
		ApiKey: apiKey,
	}

	provider := doubao.NewDoubaoProvider(conf)

	testCases := []struct {
		name     string
		messages []*botpb.Message
		options  *botpb.ChatOptions
		wantErr  bool
	}{
		{
			name: "基本对话测试",
			messages: []*botpb.Message{
				{
					Role: botpb.Message_user,
					Content: &botpb.Message_StringContent{
						StringContent: "你好",
					},
				},
			},
			options: &botpb.ChatOptions{
				Model:       model,
				MaxToken:    1000,
				Temperature: 0.7,
			},
			wantErr: false,
		},
		{
			name: "系统消息必须在第一位",
			messages: []*botpb.Message{
				{
					Role: botpb.Message_user,
					Content: &botpb.Message_StringContent{
						StringContent: "你好",
					},
				},
				{
					Role: botpb.Message_system,
					Content: &botpb.Message_StringContent{
						StringContent: "系统消息",
					},
				},
			},
			options: &botpb.ChatOptions{
				Model:       model,
				MaxToken:    1000,
				Temperature: 0.7,
			},
			wantErr: true,
		},
		{
			name: "图片消息测试",
			messages: []*botpb.Message{
				{
					Role: botpb.Message_user,
					Content: &botpb.Message_TypeContent{
						TypeContent: &botpb.TypeMessage{
							Type: botpb.TypeMessage_image,
							ImageUrl: &botpb.TypeMessage_ImageUrl{
								Url: getTestImageBase64(t),
							},
						},
					},
				},
				{
					Role: botpb.Message_user,
					Content: &botpb.Message_TypeContent{
						TypeContent: &botpb.TypeMessage{
							Type: botpb.TypeMessage_string,
							Text: "这是什么logo",
						},
					},
				},
			},
			options: &botpb.ChatOptions{
				Model:       model,
				MaxToken:    1000,
				Temperature: 0.7,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := provider.Chat(context.Background(), tc.messages, tc.options)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, resp)

			t.Logf("chat resp: %v", resp)
			if resp != nil {
				assert.NotEmpty(t, resp.Id)
				assert.NotEmpty(t, resp.Model)
				assert.NotEmpty(t, resp.Choices)
			}
		})
	}
}
