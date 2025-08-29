package model

import (
	"github.com/google/uuid"

	"server/internal/utils/errors"
	"server/internal/enum/accountType"
	"pkg/datetime"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoGetAccountsReq wrapper for proto request
type ProtoGetAccountsReq struct {
	*proto.GetAccountsRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoGetAccountsReq) ConvertToModel() (GetAccountsReq, error) {
	var res GetAccountsReq

	if p.GetAccountsRequest == nil {
		return res, errors.BadRequest.New("GetAccountsRequest is required")
	}

	// Convert account group IDs
	var accountGroupIDs []uuid.UUID
	for _, idBytes := range p.AccountGroupIDs {
		id, err := uuid.FromBytes(idBytes)
		if err != nil {
			return res, errors.BadRequest.Wrap(err)
		}
		accountGroupIDs = append(accountGroupIDs, id)
	}

	// Convert account type if provided
	var accountType *accountType.Type
	if p.Type != nil {
		t, err := accountType.ProtoAccountType{AccountType: *p.Type}.ConvertToModel()
		if err != nil {
			return res, err
		}
		accountType = &t
	}

	// Convert dates
	var dateFrom *datetime.Time
	if p.DateFrom != nil {
		dateFrom = &datetime.Time{Time: p.DateFrom.AsTime()}
	}

	var dateTo *datetime.Time
	if p.DateTo != nil {
		dateTo = &datetime.Time{Time: p.DateTo.AsTime()}
	}

	return GetAccountsReq{
		AccountGroupIDs:    accountGroupIDs,
		AccountingInCharts: p.AccountingInCharts,
		AccountingInHeader: p.AccountingInHeader,
		DateFrom:           dateFrom,
		DateTo:             dateTo,
		Type:               accountType,
		Visible:            p.Visible,
	}, nil
}
