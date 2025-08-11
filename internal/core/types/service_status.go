// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package types

import (
	"encoding/json"

	v1 "github.com/ra1n6ow/opsx/pkg/api/core/v1"
)

// ServiceStatusWrapper 包装ServiceStatus以提供自定义JSON序列化
type ServiceStatusWrapper struct {
	v1.ServiceStatus
}

// MarshalJSON 自定义JSON序列化，确保即使值为0也会被序列化
func (s ServiceStatusWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(s.ServiceStatus))
}

// UnmarshalJSON 自定义JSON反序列化
func (s *ServiceStatusWrapper) UnmarshalJSON(data []byte) error {
	var val int32
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	s.ServiceStatus = v1.ServiceStatus(val)
	return nil
}

// HealthzResponseWithStatus 使用自定义状态包装器的健康检查响应
type HealthzResponseWithStatus struct {
	Status    ServiceStatusWrapper `json:"status"` // 不使用omitempty
	Timestamp string               `json:"timestamp,omitempty"`
	Message   string               `json:"message,omitempty"`
}

// FromProto 从protobuf消息转换
func (h *HealthzResponseWithStatus) FromProto(pb *v1.HealthzResponse) {
	h.Status = ServiceStatusWrapper{ServiceStatus: pb.GetStatus()}
	h.Timestamp = pb.GetTimestamp()
	h.Message = pb.GetMessage()
}
