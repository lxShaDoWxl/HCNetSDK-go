package types

import "C"

type NET_DVR_ETHERNET_V30 struct {
	StruDVRIP        NET_DVR_IPADDR
	StruDVRIPMask    NET_DVR_IPADDR
	DwNetInterface   DWORD
	WDVRPort         WORD
	WMTU             WORD
	ByMACAddr        [MACADDR_LEN]BYTE
	ByEthernetPortNo BYTE
	ByRes            [1]BYTE
}
type LPNET_DVR_ETHERNET_V30 *NET_DVR_ETHERNET_V30
