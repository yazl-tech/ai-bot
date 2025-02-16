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
	"github.com/yazl-tech/ai-bot/config"
)

var (
	port           = pflags.Int("port", 28089, "ai bot port to listen on")
	doubaoConfFlag = pflags.Struct("doubao", (*config.DoubaoConfig)(nil), "doubao config")
)

func main() {
	pflags.Parse()

	doubaoConf := new(config.DoubaoConfig)
	plog.PanicError(doubaoConfFlag(doubaoConf))

	srv := cores.NewPuzzleCore()

	plog.PanicError(cores.Start(srv, port()))
}
