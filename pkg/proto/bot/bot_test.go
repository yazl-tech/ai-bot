// File:		bot_test.go
// Created by:	Hoven
// Created on:	2025-02-17
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package botpb

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBotPb(t *testing.T) {
	typeMsg := &Message{
		Role: Message_user,
		Content: &Message_TypeContent{
			TypeContent: &TypeMessage{
				Type: TypeMessage_image,
				ImageUrl: &TypeMessage_ImageUrl{
					Url: "https://example.com/image.jpg",
				},
			},
		},
	}

	b, err := json.Marshal(typeMsg)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(b))
}
