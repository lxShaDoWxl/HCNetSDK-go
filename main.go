package main

import (
	hik "HCNetSDK-go/hikvision"
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// 初始化
	success := hik.NET_DVR_Init()
	if !success {
		// 初始化失败
		errCode := hik.NET_DVR_GetLastError()
		fmt.Println("Not success, error code: " + strconv.FormatInt(int64(errCode), 10))
		return
	}

	// 登陆参数
	loginInfo := hik.NET_DVR_USER_LOGIN_INFO{}
	deviceInfo := hik.NET_DVR_DEVICEINFO_V40{}

	copy(loginInfo.SDeviceAddress[:], "192.168.8.110")
	loginInfo.WPort = 8000
	loginInfo.ByUseTransport = 0
	loginInfo.BUseAsynLogin = 1
	copy(loginInfo.SUserName[:], "<your username>")
	copy(loginInfo.SPassword[:], "<your password>")
	// 异步登陆回调函数
	loginInfo.CbLoginResult = func(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
		if 1 != dwResult {
			fmt.Println("异步登陆失败")
			return
		}

		fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
		fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

		// 获取卡
		cardInfo := hik.NET_DVR_CARD_CFG_COND{}
		cardInfo.DwSize = uint32(unsafe.Sizeof(cardInfo))
		cardInfo.DwCardNum = 0xffffffff
		cardInfo.ByCheckCardNo = 1
		cb := func(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
			fmt.Println("dwType: ", dwType)
			if dwType == 0 {
				dwStatusArr := *(*[4]byte)(lpBuffer)
				dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
				fmt.Println("dwStatus: ", dwStatus)
				fmt.Println("dwBufLen: ", dwBufLen)
			} else if dwType == 1 {
				progress := *(*int)(lpBuffer)
				fmt.Println("progress: ", progress)
			} else if dwType == 2 {
				cardCfg := (hik.LPNET_DVR_CARD_CFG_V50)(lpBuffer)
				fmt.Println("byCardNo: ", string(cardCfg.ByCardNo[:]))
				fmt.Println("byCardValid: ", cardCfg.ByCardValid)
				fmt.Println("byCardType: ", cardCfg.ByCardType)
			}
		}

		// 获取卡参数
		hik.NET_DVR_StartRemoteConfig(
			lUserID,
			hik.NET_DVR_GET_CARD_CFG_V50,
			unsafe.Pointer(&cardInfo),
			unsafe.Sizeof(cardInfo),
			cb,
			nil,
		)
	}
	// 登陆
	hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)

	hangOn := make(chan int)
	<-hangOn
}