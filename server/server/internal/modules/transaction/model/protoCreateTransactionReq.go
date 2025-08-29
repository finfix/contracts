package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"server/internal/utils/errors"
	"pkg/datetime"
	"server/internal/enum/transactionType"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoCreateTransactionReq wrapper for proto request
type ProtoCreateTransactionReq struct {
	*proto.CreateTransactionRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoCreateTransactionReq) ConvertToModel() (CreateTransactionReq, error) {
	var res CreateTransactionReq

	if p.CreateTransactionRequest == nil {
		return res, errors.BadRequest.New("CreateTransactionRequest is required")
	}

	// Parse ID from bytes
	id, err := uuid.FromBytes(p.Id)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Parse AccountFromID
	accountFromID, err := uuid.FromBytes(p.AccountFromID)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Parse AccountToID
	accountToID, err := uuid.FromBytes(p.AccountToID)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Parse AccountGroupID
	accountGroupID, err := uuid.FromBytes(p.AccountGroupID)
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	// Convert transaction type
	transactionType, err := transactionType.ProtoTransactionType{TransactionType: p.Type}.ConvertToModel()
	if err != nil {
		return res, err
	}

	// Convert datetime
	var datetimeCreate datetime.Time
	if p.DatetimeCreate != nil {
		datetimeCreate = datetime.Time{Time: p.DatetimeCreate.AsTime()}
	}

	// Convert tag IDs
	var tagIDs []uuid.UUID
	for _, tagIDBytes := range p.TagIDs {
		tagID, err := uuid.FromBytes(tagIDBytes)
		if err != nil {
			return res, errors.BadRequest.Wrap(err)
		}
		tagIDs = append(tagIDs, tagID)
	}

	var note *string
	if p.Note != nil {
		note = p.Note
	}

	return CreateTransactionReq{
		ID:                 id,
		AccountFromID:      accountFromID,
		AccountToID:        accountToID,
		AccountGroupID:     accountGroupID,
		AmountFrom:         decimal.NewFromFloat(p.AmountFrom),
		AmountTo:           decimal.NewFromFloat(p.AmountTo),
		DateTransaction:    p.DateTransaction,
		DatetimeCreate:     datetimeCreate,
		IsExecuted:         p.IsExecuted,
		AccountingInCharts: p.AccountingInCharts,
		Type:               transactionType,
		Note:               note,
		TagIDs:             tagIDs,
	}, nil
}
