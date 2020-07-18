package hikvision

/*
// 定义系统:
// 1 = Windows
// 2 = Linux
// 3 = Darwin
#cgo windows CFLAGS: -DCGO_OS=1
#cgo linux 	 CFLAGS: -DCGO_OS=2
#cgo darwin  CFLAGS: -DCGO_OS=3

// 头文件
#cgo CFLAGS: -I../include
#include "HCNetSDK.h"

// 引入动态链接库
#if 1==CGO_OS
    #cgo LDFLAGS: -L${SRCDIR}/../lib/win64 -lHCNetSDK
#elif 2==CGO_OS || 3==CGO_OS
    #cgo LDFLAGS: -L${SRCDIR}/../lib/linux64 -lhcnetsdk
#else
    #error("Unknow OS")
#endif
*/
import "C"
import (
	"unsafe"
)

/************************* SDK 初始化 *************************/

// 设置SDK初始化参数。
// exapmle:
//    p := hikvision.NET_DVR_LOCAL_SDK_PATH{}
//    copy(p.SPath[:], "/usr/lib/hikvision")
//    result := hik.NET_DVR_SetSDKInitCfg(hikvision.NET_SDK_INIT_CFG_SDK_PATH, unsafe.Pointer(&p))
func NET_DVR_SetSDKInitCfg(enumType NET_SDK_INIT_CFG_TYPE, lpInBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_SetSDKInitCfg(
		(C.NET_SDK_INIT_CFG_TYPE)(enumType),
		lpInBuff,
	))
}

// 初始化SDK，调用其他SDK函数的前提。
func NET_DVR_Init() bool {
	return goBOOL(C.NET_DVR_Init())
}

// 释放SDK资源，在程序结束之前调用。
func NET_DVR_Cleanup() bool {
	return goBOOL(C.NET_DVR_Cleanup())
}

/************************* SDK 本地功能 *************************/

//------------------------ 设置/获取SDK本地参数配置 ------------------------

// 设置SDK本地参数配置
// example:
//    p := hik.NET_DVR_LOCAL_TCP_PORT_BIND_CFG{}
//    p.WLocalBindTcpMinPort = 2000
//    p.WLocalBindTcpMaxPort = 3000
//    result := hik.NET_DVR_SetSDKLocalCfg(
//        hik.NET_SDK_LOCAL_CFG_TYPE_TCP_PORT_BIND,
//        unsafe.Pointer(&p),
//    )
func NET_DVR_SetSDKLocalCfg(enumType NET_SDK_LOCAL_CFG_TYPE, lpInBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_SetSDKLocalCfg(
		(C.NET_SDK_LOCAL_CFG_TYPE)(enumType),
		lpInBuff,
	))
}

// 获取SDK本地参数配置
// example:
//    p := hik.NET_DVR_LOCAL_TCP_PORT_BIND_CFG{}
//    result := hik.NET_DVR_GetSDKLocalCfg(
//        hik.NET_SDK_LOCAL_CFG_TYPE_TCP_PORT_BIND,
//        unsafe.Pointer(&p),
//    )
func NET_DVR_GetSDKLocalCfg(enumType NET_SDK_LOCAL_CFG_TYPE, lpOutBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_GetSDKLocalCfg(
		(C.NET_SDK_LOCAL_CFG_TYPE)(enumType),
		lpOutBuff,
	))
}

//------------------------ 连接、接收超时时间、重连设置 ------------------------

// 设置网络连接超时时间和连接尝试次数
func NET_DVR_SetConnectTime(dwWaitTime int32, dwTryTimes int32) bool {
	return goBOOL(C.NET_DVR_SetConnectTime(
		C.uint(dwWaitTime),
		C.uint(dwTryTimes),
	))
}

// 设置接收超时时间
func NET_DVR_SetRecvTimeOut(nRecvTimeOut int32) bool {
	return goBOOL(C.NET_DVR_SetRecvTimeOut(C.uint(nRecvTimeOut)))
}

// 设置重连功能
func NET_DVR_SetReconnect(dwInterval int32, bEnableRecon bool) bool {
	var reconnect C.BOOL = 1
	if !bEnableRecon {
		reconnect = 0
	}
	return goBOOL(C.NET_DVR_SetReconnect(C.uint(dwInterval), reconnect))
}

//------------------------ 多网卡绑定 ------------------------

// 获取所有IP，用于支持多网卡接口
// example:
//    ips := [16][16]byte{}
//    var num uint32 = 0
//    bind := false
//    result := hik.NET_DVR_GetLocalIP(&ips, &num, &bind)
func NET_DVR_GetLocalIP(strIP *[16][16]byte, pValidNum *uint32, pEnableBind *bool) bool {
	// GO byte 矩阵转成 C char 矩阵
	buf := [16][16]C.char{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			buf[i][j] = C.char(strIP[i][j])
		}
	}

	// 需要变量才能取地址
	var _pEnableBind C.BOOL = cBOOL(*pEnableBind)

	// 调用C
	result := goBOOL(C.NET_DVR_GetLocalIP(
		&buf[0],
		(*C.DWORD)(pValidNum),
		&_pEnableBind,
	))

	// 结果复制到 GO byte 矩阵中
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			strIP[i][j] = byte(buf[i][j])
		}
	}

	// 结果复制到变量中
	*pEnableBind = goBOOL(_pEnableBind)

	return result
}

// 选择使用哪个IP
func NET_DVR_SetValidIP(dwIPIndex uint32, bEnableBind bool) bool {
	return goBOOL(C.NET_DVR_SetValidIP(C.DWORD(dwIPIndex), cBOOL(bEnableBind)))
}

// 获取所有IP_V6，用于支持多网卡接口
// example:
// 参考 NET_DVR_GetLocalIP
func NET_DVR_GetLocalIPv6(strIP *[16][16]byte, pValidNum *uint32, pEnableBind *bool) bool {
	// GO byte 矩阵转成 C char 矩阵
	buf := [16][16]C.BYTE{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			buf[i][j] = C.BYTE(strIP[i][j])
		}
	}

	// 需要变量才能取地址
	var _pEnableBind C.BOOL = cBOOL(*pEnableBind)

	// 调用C
	result := goBOOL(C.NET_DVR_GetLocalIPv6(
		&buf[0],
		(*C.DWORD)(pValidNum),
		&_pEnableBind,
	))

	// 结果复制到 GO byte 矩阵中
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			strIP[i][j] = byte(buf[i][j])
		}
	}

	// 结果复制到变量中
	*pEnableBind = goBOOL(_pEnableBind)

	return result
}

// 选择使用哪个IP_V6
func NET_DVR_SetValidIPv6(dwIPIndex uint32, bEnableBind bool) bool {
	return goBOOL(C.NET_DVR_SetValidIPv6(C.DWORD(dwIPIndex), cBOOL(bEnableBind)))
}

//------------------------ SDK版本、状态、能力 ------------------------

// 获取SDK的版本信息
func NET_DVR_GetSDKVersion() uint32 {
	return uint32(C.NET_DVR_GetSDKVersion())
}

// 获取SDK的版本号和build信息
func NET_DVR_GetSDKBuildVersion() uint32 {
	return uint32(C.NET_DVR_GetSDKBuildVersion())
}

// 获取当前SDK的状态信息
// example:
//    p := hik.NET_DVR_SDKSTATE{}
//    result := hik.NET_DVR_GetSDKState(&p)
func NET_DVR_GetSDKState(pSDKState LPNET_DVR_SDKSTATE) bool {
	return goBOOL(C.NET_DVR_GetSDKState(
		C.LPNET_DVR_SDKSTATE(unsafe.Pointer(pSDKState)),
	))
}

// 获取当前SDK的功能信息
// example:
//    p := hik.NET_DVR_SDKABL{}
//    result := hik.NET_DVR_GetSDKAbility(&p)
func NET_DVR_GetSDKAbility(pSDKAbl LPNET_DVR_SDKABL) bool {
	return goBOOL(C.NET_DVR_GetSDKAbility(
		C.LPNET_DVR_SDKABL(unsafe.Pointer(pSDKAbl)),
	))
}

//------------------------ SDK启动写日志 ------------------------

// 启用写日志文件
// example:
//    usr, _ := user.Current()
//    result := hik.NET_DVR_SetLogToFile(3, usr.HomeDir + "/hik.log", true)
func NET_DVR_SetLogToFile(nLogLevel uint32, strLogDir string, bAutoDel bool) bool {
	b := []byte(strLogDir)
	return goBOOL(C.NET_DVR_SetLogToFile(
		C.DWORD(nLogLevel),
		(*C.char)(unsafe.Pointer(&b[0])),
		cBOOL(bAutoDel),
	))
}

//------------------------ 获取错误信息 ------------------------

// 返回最后操作的错误码
func NET_DVR_GetLastError() int32 {
	return int32(C.NET_DVR_GetLastError())
}

// 返回最后操作的错误码信息
func NET_DVR_GetErrorMsg(pErrorNo int) string {
	code := C.LONG(pErrorNo)
	return string(C.GoString(
		C.NET_DVR_GetErrorMsg(&code),
	))
}

// TODO: split to another source file, but missing C.BOOL when build.
/******************* golang and cgo type convert each other *******************/
// C.BOOL --> go bool
func goBOOL(flag C.BOOL) bool {
	if flag == 1 {
		return true
	}
	return false
}

// go bool --> C.BOOL
func cBOOL(flag bool) C.BOOL {
	if flag {
		return C.BOOL(1)
	}
	return C.BOOL(0)
}
