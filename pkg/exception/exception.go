// File:		exception.go
// Created by:	Hoven
// Created on:	2025-02-17
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package exception

import "github.com/go-errors/errors"

type Exception struct {
	code  int
	cause error
}

func NewException(code int, cause error) *Exception {
	return &Exception{
		code:  code,
		cause: cause,
	}
}

func NewExceptionWithMessage(code int, message string) *Exception {
	return NewException(code, errors.New(message))
}

func (e *Exception) Error() string {
	return e.cause.Error() + " (code: " + string(e.code) + ")"
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) Cause() error {
	return e.cause
}

func (e *Exception) Message() string {
	return e.Error()
}

var (
	ErrMissingApiKey                = NewExceptionWithMessage(500, "Missing API KEY")
	ErrSystemMessageMustInIndexZero = NewExceptionWithMessage(400, "系统消息必须位于消息列表第一个")
	ErrInvalidMessageType           = NewExceptionWithMessage(400, "不支持的消息类型")
	ErrInvalidContentType           = NewExceptionWithMessage(400, "不支持的信息类型")
	ErrResponseNoChoice             = NewExceptionWithMessage(400, "响应中没有输出")
)
