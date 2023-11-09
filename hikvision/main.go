package hikvision

import (
	"fmt"
	"github.com/go-errors/errors"
)

// InitHikSDK hk sdk init
func InitHikSDK() error {

	success := NET_DVR_Init()
	success = !success || NET_DVR_SetConnectTime(2000, 5)
	success = !success || NET_DVR_SetReconnect(10000, true)
	if !success {
		return errors.New("Initial failed")
	}
	return nil
}

// HKErr Detect success of operation
func HKErr(operation string, ip string) error {
	errno := NET_DVR_GetLastError()
	if errno > 0 {
		reMsg := fmt.Sprintf("%s:%s device failed failure code number：%d, message:%s", ip, operation, errno, NET_DVR_GetErrorMsg(errno))
		return errors.New(reMsg)
	}
	return nil
}
