package types

type NET_DVR_DDNSPARA struct {
	SUserName    [NAME_LEN]byte
	SPassword    [PASSWD_LEN]byte
	SDomainName  [64]byte
	ByEnableDDNS byte
	Res          [15]byte
}
type LPNET_DVR_DDNSPARA *NET_DVR_DDNSPARA
