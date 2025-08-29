package model

import (
	proto "github.com/finfix/go-server-grpc/proto"
)

// ConvertToProto converts internal response to proto response
func (r CreateAccountRes) ConvertToProto() *proto.CreateAccountResponse {
	response := &proto.CreateAccountResponse{}

	if r.SerialNumber != nil {
		response.SerialNumber = r.SerialNumber
	}

	if r.BalancingAccountID != nil {
		balancingAccountIDBytes := (*r.BalancingAccountID)[:]
		response.BalancingAccountID = balancingAccountIDBytes
	}

	if r.BalancingAccountSerialNumber != nil {
		response.BalancingAccountSerialNumber = r.BalancingAccountSerialNumber
	}

	if r.BalancingTransactionID != nil {
		balancingTransactionIDBytes := (*r.BalancingTransactionID)[:]
		response.BalancingTransactionID = balancingTransactionIDBytes
	}

	return response
}
