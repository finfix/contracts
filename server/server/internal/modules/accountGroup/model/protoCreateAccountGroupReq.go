package model

import (
	"github.com/google/uuid"

	"server/internal/utils/errors"
	"pkg/datetime"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoCreateAccountGroupReq wrapper for proto request
type ProtoCreateAccountGroupReq struct {
	*proto.CreateAccountGroupRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoCreateAccountGroupReq) ConvertToModel() (CreateAccountGroupReq, error) {
	var res CreateAccountGroupReq

	if p.CreateAccountGroupRequest == nil {
		return res, errors.BadRequest.New("CreateAccountGroupRequest is required")
	}

	// Parse ID from bytes
	id, err := uuid.FromBytes(p.Id)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Convert datetime
	var datetimeCreate datetime.Time
	if p.DatetimeCreate != nil {
		datetimeCreate = datetime.Time{Time: p.DatetimeCreate.AsTime()}
	}

	return CreateAccountGroupReq{
		ID:             id,
		Name:           p.Name,
		Currency:       p.Currency,
		DatetimeCreate: datetimeCreate,
	}, nil
}
