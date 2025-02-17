// File:		config.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package config

import "github.com/yazl-tech/ai-bot/pkg/exception"

type DoubaoConfig struct {
	Api    string
	ApiKey string
}

func (dc *DoubaoConfig) SetDefault() {
	if dc.Api == "" {
		dc.Api = "https://ark.cn-beijing.volces.com/api/v3/chat/completions"
	}
}

func (dc *DoubaoConfig) Validate() error {
	if dc.ApiKey == "" {
		return exception.ErrMissingApiKey
	}
	return nil
}
