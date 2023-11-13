package types

type NET_DVR_NTPPARA struct {
	SNTPServer       [64]byte
	WInterval        WORD
	ByEnableNTP      byte
	CTimeDifferenceH int8
	CTimeDifferenceM int8
	Res1             byte
	WNtpPort         WORD
	Res2             [8]byte
}
type LPNET_DVR_NTPPARA *NET_DVR_NTPPARA

func (n NET_DVR_NTPPARA) releaseDvrCfg() {
	//TODO implement me
	panic("implement me")
}
