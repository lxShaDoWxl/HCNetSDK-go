package types

import "C"

type NET_DVR_NETCFG_V50 struct {
	DvrCfg
	DwSize                            uint32
	StruEtherNet                      NET_DVR_ETHERNET_V30
	StruRes1                          [2]NET_DVR_IPADDR
	StruAlarmHostIpAddr               NET_DVR_IPADDR
	ByRes2                            [4]BYTE
	WAlarmHostIpPort                  WORD
	ByUseDhcp                         BYTE
	ByIPv6Mode                        BYTE
	StruDnsServer1IpAddr              NET_DVR_IPADDR
	StruDnsServer2IpAddr              NET_DVR_IPADDR
	ByIpResolver                      [MAX_DOMAIN_NAME]BYTE
	WIpResolverPort                   WORD
	WHttpPortNo                       WORD
	StruMulticastIpAddr               NET_DVR_IPADDR
	StruGatewayIpAddr                 NET_DVR_IPADDR
	StruPPPoE                         NET_DVR_PPPOECFG
	ByEnablePrivateMulticastDiscovery BYTE
	byEnableOnvifMulticastDiscovery   BYTE
	StruAlarmHost2IpAddr              NET_DVR_IPADDR
	byEnableDNS                       BYTE
	ByRes                             [599]BYTE
}
type LPNET_DVR_NETCFG_V50 *NET_DVR_NETCFG_V50
