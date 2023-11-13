package types

type NET_DVR_NETAPPCFG struct {
	DwSize              DWORD
	SDNSIp              [16]byte
	StruNtpClientParam  NET_DVR_NTPPARA
	StruDDNSClientParam NET_DVR_DDNSPARA
	Res                 [464]byte
}
type LPNET_DVR_NETAPPCFG *NET_DVR_NETAPPCFG

func (n NET_DVR_NETAPPCFG) releaseDvrCfg() {
	//TODO implement me
	panic("implement me")
}
