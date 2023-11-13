package types

import "time"

type NET_DVR_TIME struct {
	DwYear   int32
	DwMonth  int32
	DwDay    int32
	DwHour   int32
	DwMinute int32
	DwSecond int32
}
type LPNET_DVR_TIME *NET_DVR_TIME

func (n NET_DVR_TIME) releaseDvrCfg() {
	//TODO implement me
	panic("implement me")
}
func (n NET_DVR_TIME) ToTime(tz *time.Location) time.Time {
	if tz == nil {
		tz = time.UTC
	}
	return time.Date(
		int(n.DwYear),
		time.Month(n.DwMonth),
		int(n.DwDay),
		int(n.DwHour),
		int(n.DwMinute),
		int(n.DwSecond),
		0, tz)
}
