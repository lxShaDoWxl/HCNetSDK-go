package types

type NET_DVR_ZONEANDDST struct {
	DwSize         DWORD
	DwZoneIndex    DWORD //
	ByRes1         [12]byte
	DwEnableDST    DWORD
	ByDSTBias      BYTE
	ByRes2         [3]byte
	StruBeginPoint NET_DVR_TIMEPOINT
	StruEndPoint   NET_DVR_TIMEPOINT
}
type LPNET_DVR_ZONEANDDST *NET_DVR_ZONEANDDST

func (n NET_DVR_ZONEANDDST) releaseDvrCfg() {
	//TODO implement me
	panic("implement me")
}
