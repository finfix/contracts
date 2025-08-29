package model

import (
	"server/internal/utils/errors"
	"server/internal/enum/osType"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoRefreshTokensReq wrapper for proto request
type ProtoRefreshTokensReq struct {
	*proto.RefreshTokensRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoRefreshTokensReq) ConvertToModel() (RefreshTokensReq, error) {
	var res RefreshTokensReq

	if p.RefreshTokensRequest == nil {
		return res, errors.BadRequest.New("RefreshTokensRequest is required")
	}

	var application *ApplicationInformation
	if p.Application != nil {
		application = &ApplicationInformation{
			Version:  p.Application.Version,
			Build:    p.Application.Build,
			BundleID: p.Application.BundleId,
		}
	}

	var device *DeviceInformation
	if p.Device != nil {
		osType, err := osType.ProtoOSType{OSType: p.Device.NameOS}.ConvertToModel()
		if err != nil {
			return res, err
		}

		device = &DeviceInformation{
			DeviceName: p.Device.DeviceName,
			ModelName:  p.Device.ModelName,
			NameOS:     osType,
			VersionOS:  p.Device.VersionOS,
		}
	}

	return RefreshTokensReq{
		Token:       p.Token,
		Application: application,
		Device:      device,
	}, nil
}
