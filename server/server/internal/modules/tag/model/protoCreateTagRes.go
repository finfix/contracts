package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToProto converts internal response to proto response
func (r CreateTagRes) ConvertToProto() *proto.CreateTagResponse {
	response := &proto.CreateTagResponse{}

	// Только error возвращается, ID больше не нужен
	return response
}
