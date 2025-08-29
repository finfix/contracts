package model

import (
	"server/internal/utils/errors"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoSignOutReq wrapper for proto request
type ProtoSignOutReq struct {
	*proto.SignOutRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoSignOutReq) ConvertToModel() (SignOutReq, error) {
	var res SignOutReq

	if p.SignOutRequest == nil {
		return res, errors.BadRequest.New("SignOutRequest is required")
	}

	return SignOutReq{
		AccessToken: p.AccessToken,
	}, nil
}
