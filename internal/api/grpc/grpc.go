// File:		grpc.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package grpcInterface

import (
	"github.com/yazl-tech/ai-bot/internal/service"
	"google.golang.org/grpc"

	grpcService "github.com/yazl-tech/ai-bot/internal/api/grpc/service"
	doubaopb "github.com/yazl-tech/ai-bot/pkg/proto/doubao"
)

type grpcInitFunc func(srv *grpc.Server)

func setupGrpcServerFn(s *service.AiBotService) grpcInitFunc {
	return func(srv *grpc.Server) {
		doubaoSrv := grpcService.NewDoubaoGrpcService(s)
		doubaopb.RegisterDoubaoHandlerServer(srv, doubaoSrv)
	}
}

func SetupGrpcServer(service *service.AiBotService) grpcInitFunc {
	return setupGrpcServerFn(service)
}
