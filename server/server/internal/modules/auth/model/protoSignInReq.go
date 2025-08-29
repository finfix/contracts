package model

import (
	"server/internal/utils/errors"
	"server/internal/enum/applicationType"
	"server/internal/enum/osType"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoSignInReq wrapper for proto request
type ProtoSignInReq struct {
	*proto.SignInRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoSignInReq) ConvertToModel() (SignInReq, error) {
	var res SignInReq

	if p.SignInRequest == nil {
		return res, errors.BadRequest.New("SignInRequest is required")
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

	return SignInReq{
		Email:       p.Email,
		Password:    p.Password,
		Application: application,
		Device:      device,
	}, nil
}
