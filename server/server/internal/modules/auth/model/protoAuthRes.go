package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToSignInProto converts internal auth response to proto SignIn response
func (r AuthRes) ConvertToSignInProto() *proto.SignInResponse {
	response := &proto.SignInResponse{}

	if r.ID != nil {
		idBytes := (*r.ID)[:]
		response.Id = idBytes
	}

	if r.Token != nil {
		response.Token = &proto.Tokens{
			AccessToken:  r.Token.AccessToken,
			RefreshToken: r.Token.RefreshToken,
		}
	}

	return response
}

// ConvertToSignUpProto converts internal auth response to proto SignUp response  
func (r AuthRes) ConvertToSignUpProto() *proto.SignUpResponse {
	response := &proto.SignUpResponse{}

	if r.ID != nil {
		idBytes := (*r.ID)[:]
		response.Id = idBytes
	}

	if r.Token != nil {
		response.Token = &proto.Tokens{
			AccessToken:  r.Token.AccessToken,
			RefreshToken: r.Token.RefreshToken,
		}
	}

	return response
}
