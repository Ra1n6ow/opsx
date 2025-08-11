// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package grpc

import (
	corev1 "github.com/ra1n6ow/opsx/pkg/api/core/v1"
)

// Handler 负责处理博客模块的请求.
type Handler struct {
	corev1.UnimplementedCoreServer
}

// NewHandler 创建一个新的 Handler 实例.
func NewHandler() *Handler {
	return &Handler{}
}
