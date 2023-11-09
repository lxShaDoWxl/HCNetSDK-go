package types

import "C"

type NET_DVR_PPPOECFG struct {
	DwPPPOE        uint32
	SPPPoEUser     [NAME_LEN]byte
	SPPPoEPassword [PASSWD_LEN]byte
	StruPPPoEIP    NET_DVR_IPADDR
}
type LPNET_DVR_PPPOECFG *NET_DVR_PPPOECFG
