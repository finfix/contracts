package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToProto converts internal response to proto response
func (r CreateTransactionRes) ConvertToProto() *proto.CreateTransactionResponse {
	response := &proto.CreateTransactionResponse{}

	// Только error возвращается, ID больше не нужен
	return response
}
