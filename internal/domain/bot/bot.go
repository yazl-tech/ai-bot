// File:		bot.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package bot

import botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"

// Bot 代表着一个机器人实体
// 可以给每个机器人指定参数和Provider
type Bot struct {
	ID       string
	Name     string
	Provider string
	Options  *botpb.ChatOptions
}
