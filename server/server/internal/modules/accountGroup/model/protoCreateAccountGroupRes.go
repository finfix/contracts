package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToProto converts internal response to proto response
func (r CreateAccountGroupRes) ConvertToProto() *proto.CreateAccountGroupResponse {
	response := &proto.CreateAccountGroupResponse{}

	if r.SerialNumber != nil {
		response.SerialNumber = r.SerialNumber
	}

	return response
}
