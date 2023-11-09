package types

import "C"

// адрес хоста
type NET_DVR_IPADDR struct {
	SIpV4  [16]byte  /* IPv4地址 */
	ByIPv6 [128]BYTE /* 保留 */
}
type LPNET_DVR_IPADDR *NET_DVR_IPADDR
