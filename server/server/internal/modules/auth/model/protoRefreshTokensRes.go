package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToProto converts internal refresh tokens response to proto response
func (r RefreshTokensRes) ConvertToProto() *proto.RefreshTokensResponse {
	response := &proto.RefreshTokensResponse{}

	if r.AccessToken != nil {
		response.AccessToken = r.AccessToken
	}

	if r.RefreshToken != nil {
		response.RefreshToken = r.RefreshToken
	}

	return response
}
