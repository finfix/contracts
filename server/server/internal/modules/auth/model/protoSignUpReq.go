package model

import (
	"server/internal/utils/errors"
	"server/internal/enum/osType"

	proto "github.com/finfix/go-server-grpc/proto"
)

// ProtoSignUpReq wrapper for proto request
type ProtoSignUpReq struct {
	*proto.SignUpRequest
}

// ConvertToModel converts proto request to internal model
func (p ProtoSignUpReq) ConvertToModel() (SignUpReq, error) {
	var res SignUpReq

	if p.SignUpRequest == nil {
		return res, errors.BadRequest.New("SignUpRequest is required")
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

	return SignUpReq{
		Name:        p.Name,
		Email:       p.Email,
		Password:    p.Password,
		Application: application,
		Device:      device,
	}, nil
}
