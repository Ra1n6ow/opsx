// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package grpc

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	ucv1 "github.com/ra1n6ow/opsx/pkg/api/usercenter/v1"
)

func (h *Handler) Healthz(ctx context.Context, req *emptypb.Empty) (*ucv1.HealthzResponse, error) {
	return &ucv1.HealthzResponse{
		// 使用 Enum() 方法直接获取常量指针
		Status:    ucv1.ServiceStatus_Healthy.Enum(),
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
