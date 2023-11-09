package device

import "C"
import (
	"context"
	"encoding/binary"
	"errors"
	hik "github.com/lxShaDoWxl/HCNetSDK-go/hikvision"
	"github.com/lxShaDoWxl/HCNetSDK-go/hikvision/types"
	"unsafe"
)

type Device interface {
	Login() (int, error)
	Logout() error
	GetIP() string
	GetDeviceInfo() hik.NET_DVR_DEVICEINFO_V40
	AlarmArmingMode(ctx context.Context, msgCallback hik.MSGCallBack_V31) error
	AlarmListeningMode(ctx context.Context, msgCallback hik.MSGCallBack, port uint16, ip string) error
	GetCard(ctx context.Context, cardId string) (types.NET_DVR_CARD_CFG_V50, error)
	AddCard(ctx context.Context, cardId string) error
	RemoveCard(ctx context.Context, cardId string) error
	//SetAlarmCallBack() error
	//StartListenAlarmMsg() error
	//StopListenAlarmMsg() error
}

type Info struct {
	IP       string
	Port     int
	UserName string
	Password string
}

type HKDevice struct {
	ip          string
	port        int
	username    string
	password    string
	loginId     int
	alarmHandle int
	deviceInfo  hik.NET_DVR_DEVICEINFO_V40
}

// NewHKDevice new hk-device instance
func NewHKDevice(info Info) Device {
	return &HKDevice{
		ip:       info.IP,
		port:     info.Port,
		username: info.UserName,
		password: info.Password,
	}
}
func (d *HKDevice) GetIP() string {
	return d.ip
}
func (d *HKDevice) GetDeviceInfo() hik.NET_DVR_DEVICEINFO_V40 {
	return d.deviceInfo
}

// Login hk device loin
func (d *HKDevice) Login() (int, error) {

	loginInfo := hik.NET_DVR_USER_LOGIN_INFO{}

	copy(loginInfo.SDeviceAddress[:], d.ip)
	loginInfo.WPort = uint16(d.port)
	loginInfo.ByUseTransport = 0
	loginInfo.BUseAsynLogin = 0
	copy(loginInfo.SUserName[:], d.username)
	copy(loginInfo.SPassword[:], d.password)

	d.loginId = hik.NET_DVR_Login_V40(&loginInfo, &d.deviceInfo)
	if d.loginId < 0 {
		return -1, hik.HKErr("login", d.ip)
	}
	return d.loginId, nil
}

func (d *HKDevice) GetCard(ctx context.Context, cardId string) (types.NET_DVR_CARD_CFG_V50, error) {
	cardCfg := hik.NET_DVR_CARD_CFG_COND{}
	cardCfg.DwSize = uint32(unsafe.Sizeof(cardCfg))
	// Чтобы получить всю информацию о карте, нет необходимости вызывать интерфейс NET_DVR_StartRemoteConfig для указания условий запроса.
	// cardCfg.DwCardNum = 0xffffffff

	// Необходимо настроить интерфейс NET_DVR_StartRemoteConfig и указать условия запроса.
	cardCfg.DwCardNum = 1
	cardCfg.ByCheckCardNo = 1
	var sendRemoteFinished = make(chan error, 1)
	var result = make(chan unsafe.Pointer, 1)
	defer func() {
		close(sendRemoteFinished)
		close(result)
	}()
	// Перезвонить
	cb := makeRemoteConfigCallback(sendRemoteFinished, result)
	lHandle := hik.NET_DVR_StartRemoteConfig(
		d.loginId,
		types.NET_DVR_GET_CARD_CFG_V50,
		unsafe.Pointer(&cardCfg),
		unsafe.Sizeof(cardCfg),
		cb,
		nil,
	)
	var resultCardInfo types.NET_DVR_CARD_CFG_V50
	if -1 == lHandle {
		return resultCardInfo, hik.HKErr("GetCard.NET_DVR_StartRemoteConfig", d.ip)
	}
	// Отправить условия запроса
	queryBuf := hik.NET_DVR_CARD_CFG_SEND_DATA{}
	structSize := uint32(unsafe.Sizeof(queryBuf))
	queryBuf.DwSize = structSize
	copy(queryBuf.ByCardNo[:], cardId)
	success := hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&queryBuf)), structSize)
	if !success {
		return resultCardInfo, hik.HKErr("GetCard.NET_DVR_SendRemoteConfig", d.ip)

	}
	//Отправка параметров завершена?
	err := <-sendRemoteFinished
	//TODO слушать контекст и изящьно выходить если контекст закрыт
	if err != nil {
		return resultCardInfo, err
	}
	lpBuffer := <-result
	resultCardInfo = *(types.LPNET_DVR_CARD_CFG_V50)(lpBuffer)
	test := []byte("127")
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		return resultCardInfo, hik.HKErr("GetCard.NET_DVR_StopRemoteConfig", d.ip)
	}
	_ = test
	return resultCardInfo, nil
}
func (d *HKDevice) AddCard(ctx context.Context, cardId string) error {
	cardCfg := hik.NET_DVR_CARD_CFG_COND{}
	cardCfg.DwSize = uint32(unsafe.Sizeof(cardCfg))
	cardCfg.DwCardNum = 1
	cardCfg.ByCheckCardNo = 1
	var sendRemoteFinished = make(chan error, 1)
	var result = make(chan unsafe.Pointer, 1)
	defer func() {
		close(sendRemoteFinished)
		close(result)
	}()
	// Перезвонить

	lHandle := hik.NET_DVR_StartRemoteConfig(
		d.loginId,
		types.NET_DVR_SET_CARD_CFG_V50,
		unsafe.Pointer(&cardCfg),
		unsafe.Sizeof(cardCfg),
		makeRemoteConfigCallback(sendRemoteFinished, result),
		nil,
	)
	if -1 == lHandle {
		return hik.HKErr("AddCard.NET_DVR_StartRemoteConfig", d.ip)
	}
	// Отправить условия запроса
	queryBuf := types.NET_DVR_CARD_CFG_V50{}
	copy(queryBuf.ByCardNo[:], cardId)
	queryBuf.ByCardValid = 1
	queryBuf.ByCardType = 1
	queryBuf.ByLeaderCard = 0
	//queryBuf.DwDoorRight = 1
	queryBuf.ByDoorRight[0] = 1
	queryBuf.ByDoorRight[1] = 1
	queryBuf.WCardRightPlan[0][0] = 1
	queryBuf.WCardRightPlan[1][0] = 1
	structSize := uint32(unsafe.Sizeof(queryBuf))
	queryBuf.DwSize = structSize
	// Битовая маска какие параметры изменять: https://open.hikvision.com/hardware/structures/NET_DVR_CARD_CFG_V50.html
	queryBuf.DwModifyParamType = types.CARD_PARAM_REQUIRED_BY_CREATE
	success := hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&queryBuf)), structSize)
	if !success {
		return hik.HKErr("AddCard.NET_DVR_SendRemoteConfig", d.ip)

	}
	//Отправка параметров завершена?
	//TODO слушать контекст и изящьно выходить если контекст закрыт
	err := <-sendRemoteFinished
	if err != nil {
		return err
	}
	return nil
}
func (d *HKDevice) RemoveCard(ctx context.Context, cardId string) error {
	cardCfg := hik.NET_DVR_CARD_CFG_COND{}
	cardCfg.DwSize = uint32(unsafe.Sizeof(cardCfg))
	cardCfg.DwCardNum = 1
	cardCfg.ByCheckCardNo = 1
	var sendRemoteFinished = make(chan error, 1)
	var result = make(chan unsafe.Pointer, 1)
	defer func() {
		close(sendRemoteFinished)
		close(result)
	}()
	lHandle := hik.NET_DVR_StartRemoteConfig(
		d.loginId,
		types.NET_DVR_SET_CARD_CFG_V50,
		unsafe.Pointer(&cardCfg),
		unsafe.Sizeof(cardCfg),
		makeRemoteConfigCallback(sendRemoteFinished, result),
		nil,
	)
	if -1 == lHandle {
		return hik.HKErr("RemoveCard.NET_DVR_StartRemoteConfig", d.ip)
	}
	// Отправить условия запроса
	queryBuf := types.NET_DVR_CARD_CFG_V50{}
	structSize := uint32(unsafe.Sizeof(queryBuf))
	queryBuf.DwSize = structSize
	copy(queryBuf.ByCardNo[:], cardId)
	queryBuf.ByCardValid = 0
	// Битовая маска какие параметры изменять: https://open.hikvision.com/hardware/structures/NET_DVR_CARD_CFG_V50.html
	//queryBuf.DwModifyParamType = 0x00000001 | 0x00000004 | 0x00000008 | 0x00000010 | 0x00000100
	queryBuf.DwModifyParamType = 0x00000001
	success := hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&queryBuf)), structSize)
	if !success {
		return hik.HKErr("RemoveCard.NET_DVR_SendRemoteConfig", d.ip)

	}
	//Отправка параметров завершена?
	//TODO слушать контекст и изящьно выходить если контекст закрыт
	err := <-sendRemoteFinished
	if err != nil {
		return err
	}
	return nil
}
func makeRemoteConfigCallback(sendRemoteFinished chan error, result chan unsafe.Pointer) hik.FRemoteConfigCallback {
	return func(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
		//fmt.Println("dwType: ", dwType)
		if dwType == 0 {
			dwStatusArr := *(*[4]byte)(lpBuffer)
			dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
			//fmt.Println("dwStatus: ", dwStatus)
			//fmt.Println("dwBufLen: ", dwBufLen)

			switch dwStatus {

			case uint32(hik.NET_SDK_CALLBACK_STATUS_PROCESSING):
				//dwStatusArr := *(*[4 + 32]byte)(lpBuffer)
				//fmt.Println("Processing, cardNo: ", string(dwStatusArr[4:]))
				return
			case uint32(hik.NET_SDK_CALLBACK_STATUS_FAILED):

				//fmt.Println("dwErrCode: ", binary.LittleEndian.Uint32(dwErrCode[4:]))
				if binary.LittleEndian.Uint32((*(*[4 + 4]byte)(lpBuffer))[4:]) == types.NET_ERR_CARDNO_NOT_EXIST {
					sendRemoteFinished <- hik.ErrCardNotFound

				}
				sendRemoteFinished <- errors.New("callback failed")
				return
			}

			// 断开长链接
			sendRemoteFinished <- nil
		} else if dwType == 1 {
			//progress := *(*int)(lpBuffer)
			//fmt.Println("progress: ", progress)
		} else if dwType == 2 {
			//cardCfg := *(hik.LPNET_DVR_CARD_CFG_V50)(lpBuffer)
			//fmt.Println("byCardNo: ", string(cardCfg.ByCardNo[:]))
			//fmt.Println("byCardValid: ", cardCfg.ByCardValid)
			//fmt.Println("byCardType: ", cardCfg.ByCardType)
			//fmt.Println("dwCardUserId: ", cardCfg.DwCardUserId)
			result <- lpBuffer
		}
	}
}

func (d *HKDevice) AlarmArmingMode(ctx context.Context, msgCallback hik.MSGCallBack_V31) error {
	success := hik.NET_DVR_SetDVRMessageCallBack_V31(msgCallback, nil)
	if !success {
		return hik.HKErr("AlarmArmingMode.NET_DVR_SetDVRMessageCallBack_V31", d.ip)
	}
	setupParam := hik.NET_DVR_SETUPALARM_PARAM{}
	setupParam.DwSize = uint32(unsafe.Sizeof(setupParam))
	setupParam.ByLevel = 1
	lAlarmHandle := hik.NET_DVR_SetupAlarmChan_V41(d.loginId, &setupParam)
	if lAlarmHandle < 0 {
		return hik.HKErr("AlarmArmingMode.NET_DVR_SetupAlarmChan_V41", d.ip)

	}
	<-ctx.Done()
	success = hik.NET_DVR_CloseAlarmChan_V30(lAlarmHandle)
	if !success {
		return hik.HKErr("AlarmArmingMode.NET_DVR_CloseAlarmChan_V30", d.ip)

	}
	return nil
}
func (d *HKDevice) AlarmListeningMode(ctx context.Context, msgCallback hik.MSGCallBack, port uint16, ip string) error {
	success := hik.NET_DVR_SetDVRMessageCallBack_V30(msgCallback, nil)
	if !success {
		return hik.HKErr("AlarmArmingMode.NET_DVR_SetDVRMessageCallBack_V30", d.ip)
	}
	lAlarmHandle := hik.NET_DVR_StartListen_V30(ip, port, msgCallback, nil)
	if lAlarmHandle < 0 {
		return hik.HKErr("AlarmListeningMode.NET_DVR_StartListen_V30", d.ip)

	}
	<-ctx.Done()
	success = hik.NET_DVR_StopListen_V30(lAlarmHandle)
	if !success {
		return hik.HKErr("AlarmListeningMode.NET_DVR_StopListen_V30", d.ip)

	}
	return nil
}

// Logout hk device logout
func (d *HKDevice) Logout() error {
	hik.NET_DVR_Logout(d.loginId)
	if err := hik.HKErr("NVRLogout", d.ip); err != nil {
		return err
	}
	return nil
}
