package types

import "C"

type NET_DVR_NETCFG_V30 struct {
	DwSize               uint32
	StruEtherNet         [MAX_ETHERNET]NET_DVR_ETHERNET_V30
	StruRes1             [2]NET_DVR_IPADDR
	StruAlarmHostIpAddr  NET_DVR_IPADDR //Listening service IP address
	WRes2                [4]BYTE
	WAlarmHostIpPort     WORD //Listening service port No.
	ByUseDhcp            BYTE
	ByIPv6Mode           BYTE
	StruDnsServer1IpAddr NET_DVR_IPADDR
	StruDnsServer2IpAddr NET_DVR_IPADDR
	ByIpResolver         [MAX_DOMAIN_NAME]BYTE
	WIpResolverPort      WORD
	WHttpPortNo          WORD
	StruMulticastIpAddr  NET_DVR_IPADDR
	StruGatewayIpAddr    NET_DVR_IPADDR //Gateway address
	StruPPPoE            NET_DVR_PPPOECFG
	ByRes                [64]BYTE
}

func (c NET_DVR_NETCFG_V30) releaseDvrCfg() {
	//TODO implement me
	panic("implement me")
}

type LPNET_DVR_NETCFG_V30 *NET_DVR_NETCFG_V30
