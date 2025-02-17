// File:		main.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package main

import (
	"github.com/go-puzzles/puzzles/cores"
	"github.com/go-puzzles/puzzles/pflags"
	"github.com/go-puzzles/puzzles/plog"
	"github.com/yazl-tech/ai-bot/internal/domain/bot"
	"github.com/yazl-tech/ai-bot/internal/domain/bot/doubao"
	"github.com/yazl-tech/ai-bot/internal/service"
	botpb "github.com/yazl-tech/ai-bot/pkg/proto/bot"

	consulpuzzle "github.com/go-puzzles/puzzles/cores/puzzles/consul-puzzle"
	grpcpuzzle "github.com/go-puzzles/puzzles/cores/puzzles/grpc-puzzle"
	grpcuipuzzle "github.com/go-puzzles/puzzles/cores/puzzles/grpcui-puzzle"
	grpcInterface "github.com/yazl-tech/ai-bot/internal/api/grpc"
)

var (
	port           = pflags.Int("port", 28089, "ai bot port to listen on")
	doubaoConfFlag = pflags.Struct("doubao", (*doubao.DoubaoConfig)(nil), "doubao config")
)

func main() {
	pflags.Parse()

	doubaoConf := new(doubao.DoubaoConfig)
	plog.PanicError(doubaoConfFlag(doubaoConf))

	botFactory := bot.NewBotFactory()
	doubaoProvider := doubao.NewDoubaoProvider(doubaoConf)
	botFactory.RegisterProvider(botpb.ProviderType_Doubao, doubaoProvider)

	botService := service.NewAiBotService(botFactory)
	grpcSrv := grpcInterface.SetupGrpcServer(botService)

	srv := cores.NewPuzzleCore(
		cores.WithService(pflags.GetServiceName()),
		consulpuzzle.WithConsulRegister(),
		grpcpuzzle.WithCoreGrpcPuzzle(grpcSrv),
		grpcuipuzzle.WithCoreGrpcUI(),
	)

	plog.PanicError(cores.Start(srv, port()))
}
