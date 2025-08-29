package model

import (
	"github.com/google/uuid"

	"server/internal/utils/errors"
	"pkg/datetime"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoCreateTagReq wrapper for proto request
type ProtoCreateTagReq struct {
	*proto.CreateTagRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoCreateTagReq) ConvertToModel() (CreateTagReq, error) {
	var res CreateTagReq

	if p.CreateTagRequest == nil {
		return res, errors.BadRequest.New("CreateTagRequest is required")
	}

	// Parse ID from bytes
	id, err := uuid.FromBytes(p.Id)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Parse AccountGroupID
	accountGroupID, err := uuid.FromBytes(p.AccountGroupID)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Convert datetime
	var datetimeCreate datetime.Time
	if p.DatetimeCreate != nil {
		datetimeCreate = datetime.Time{Time: p.DatetimeCreate.AsTime()}
	}

	return CreateTagReq{
		ID:             id,
		Name:           p.Name,
		AccountGroupID: accountGroupID,
		DatetimeCreate: datetimeCreate,
	}, nil
}
