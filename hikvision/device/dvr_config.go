package device

import (
	hik "github.com/lxShaDoWxl/HCNetSDK-go/hikvision"
	"github.com/lxShaDoWxl/HCNetSDK-go/hikvision/types"
	"time"
)

func SetTime(d Device, t time.Time) error {
	cfg := types.NET_DVR_TIME{
		DwYear:   int32(t.Year()),
		DwMonth:  int32(t.Month()),
		DwDay:    int32(t.Day()),
		DwHour:   int32(t.Hour()),
		DwMinute: int32(t.Minute()),
		DwSecond: int32(t.Second()),
	}
	success := hik.NET_DVR_SetDVRConfig(
		d.GetLoginId(),
		types.NET_DVR_SET_TIMECFG,
		0,
		&cfg,
	)
	if !success {
		return hik.HKErr("SetTime", d.GetIP())
	}
	return nil
}

type SetTimeZoneDTO struct {
	Hours     int8
	Minutes   int8
	EnableNTP bool
	IP        string
}

func SetTimeZone(d Device, dto *SetTimeZoneDTO) error {
	cfg := types.NET_DVR_NTPPARA{}
	success := hik.NET_DVR_GetDVRConfig(
		d.GetLoginId(),
		types.NET_DVR_GET_NTPCFG,
		0,
		&cfg,
	)
	if !success {
		return hik.HKErr("SetTimeZone", d.GetIP())
	}
	cfg.CTimeDifferenceH = dto.Hours
	cfg.CTimeDifferenceM = dto.Minutes
	//Включаем NTP и выставляем сервер NTP
	cfg.SNTPServer = [64]byte{}
	cfg.ByEnableNTP = 0
	if dto.EnableNTP {
		cfg.ByEnableNTP = 1
		copy(cfg.SNTPServer[:], dto.IP)
	}

	success = hik.NET_DVR_SetDVRConfig(
		d.GetLoginId(),
		types.NET_DVR_SET_NTPCFG,
		0,
		&cfg,
	)
	if !success {
		return hik.HKErr("SetTimeZone", d.GetIP())
	}
	return nil
}
