// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package grpc

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	corev1 "github.com/ra1n6ow/opsx/pkg/api/core/v1"
)

func (h *Handler) Healthz(ctx context.Context, req *emptypb.Empty) (*corev1.HealthzResponse, error) {
	return &corev1.HealthzResponse{
		// 使用 Enum() 方法直接获取常量指针
		Status:    corev1.ServiceStatus_Healthy.Enum(),
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
