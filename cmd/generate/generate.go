// File:		generate.go
// Created by:	Hoven
// Created on:	2025-02-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package main

//go:generate ./generate_dto_proto.sh dto/doubao.proto

//go:generate ./generate_service_proto.sh doubao.proto doubaopb
func main() {

}
