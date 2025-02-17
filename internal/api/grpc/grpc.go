// File:		grpc.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package grpcInterface

import (
	grpcService "github.com/yazl-tech/ai-bot/internal/api/grpc/service"
	"github.com/yazl-tech/ai-bot/internal/service"
	"github.com/yazl-tech/ai-bot/pkg/proto/doubao"
	"google.golang.org/grpc"
)

type grpcInitFunc func(srv *grpc.Server)

func setupGrpcServerFn(s *service.AiBotService) grpcInitFunc {
	return func(srv *grpc.Server) {
		doubaoSrv := grpcService.NewDoubaoGrpcService(s)
		doubao.RegisterDoubaoHandlerServer(srv, doubaoSrv)
	}
}

func SetupGrpcServer(service *service.AiBotService) grpcInitFunc {
	return setupGrpcServerFn(service)
}
