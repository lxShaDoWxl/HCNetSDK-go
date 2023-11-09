package types

type NET_DVR_DEVICECFG_V40 struct {
	DvrCfg
	ST_dwSize          DWORD
	ST_sDVRName        [NAME_LEN]BYTE //DVR名称
	ST_dwDVRID         DWORD          //DVR ID,用于遥控器 //V1.4(0-99), V1.5(0-255)
	ST_dwRecycleRecord DWORD          //是否循环录像,0:不是; 1:是
	//以下不可更改
	ST_sSerialNumber          [SERIALNO_LEN]BYTE //序列号
	ST_dwSoftwareVersion      DWORD              //软件版本号,高16位是主版本,低16位是次版本
	ST_dwSoftwareBuildDate    DWORD              //软件生成日期,0xYYYYMMDD
	ST_dwDSPSoftwareVersion   DWORD              //DSP软件版本,高16位是主版本,低16位是次版本
	ST_dwDSPSoftwareBuildDate DWORD              // DSP软件生成日期,0xYYYYMMDD
	ST_dwPanelVersion         DWORD              // 前面板版本,高16位是主版本,低16位是次版本
	ST_dwHardwareVersion      DWORD              // 硬件版本,高16位是主版本,低16位是次版本
	ST_byAlarmInPortNum       BYTE               //DVR报警输入个数
	ST_byAlarmOutPortNum      BYTE               //DVR报警输出个数
	ST_byRS232Num             BYTE               //DVR 232串口个数
	ST_byRS485Num             BYTE               //DVR 485串口个数
	ST_byNetworkPortNum       BYTE               //网络口个数
	ST_byDiskCtrlNum          BYTE               //DVR 硬盘控制器个数
	ST_byDiskNum              BYTE               //DVR 硬盘个数
	ST_byDVRType              BYTE               //DVR类型, 1:DVR 2:ATM DVR 3:DVS ......
	ST_byChanNum              BYTE               //DVR 通道个数
	ST_byStartChan            BYTE               //起始通道号,例如DVS-1,DVR - 1
	ST_byDecordChans          BYTE               //DVR 解码路数
	ST_byVGANum               BYTE               //VGA口的个数
	ST_byUSBNum               BYTE               //USB口的个数
	ST_byAuxoutNum            BYTE               //辅口的个数
	ST_byAudioNum             BYTE               //语音口的个数
	ST_byIPChanNum            BYTE               //最大数字通道数 低8位，高8位见byHighIPChanNum
	ST_byZeroChanNum          BYTE               //零通道编码个数
	ST_bySupport              BYTE               //能力，位与结果为0表示不支持，1表示支持，
	//bySupport & 0x1, 表示是否支持智能搜索
	//bySupport & 0x2, 表示是否支持备份
	//bySupport & 0x4, 表示是否支持压缩参数能力获取
	//bySupport & 0x8, 表示是否支持多网卡
	//bySupport & 0x10, 表示支持远程SADP
	//bySupport & 0x20, 表示支持Raid卡功能
	//bySupport & 0x40, 表示支持IPSAN搜索
	//bySupport & 0x80, 表示支持rtp over rtsp
	ST_byEsataUseage BYTE //Esata的默认用途，0-默认备份，1-默认录像
	ST_byIPCPlug     BYTE //0-关闭即插即用，1-打开即插即用
	ST_byStorageMode BYTE //0-盘组模式,1-磁盘配额, 2抽帧模式, 3-自动
	ST_bySupport1    BYTE //能力，位与结果为0表示不支持，1表示支持
	//bySupport1 & 0x1, 表示是否支持snmp v30
	//bySupport1 & 0x2, 支持区分回放和下载
	//bySupport1 & 0x4, 是否支持布防优先级
	//bySupport1 & 0x8, 智能设备是否支持布防时间段扩展
	//bySupport1 & 0x10, 表示是否支持多磁盘数（超过33个）
	//bySupport1 & 0x20, 表示是否支持rtsp over http
	ST_wDevType      WORD                    //设备型号
	ST_byDevTypeName [DEV_TYPE_NAME_LEN]BYTE //设备型号名称
	ST_bySupport2    BYTE                    //能力集扩展，位与结果为0表示不支持，1表示支持
	//bySupport2 & 0x1, 表示是否支持扩展的OSD字符叠加(终端和抓拍机扩展区分)
	ST_byAnalogAlarmInPortNum BYTE    //模拟报警输入个数
	ST_byStartAlarmInNo       BYTE    //模拟报警输入起始号
	ST_byStartAlarmOutNo      BYTE    //模拟报警输出起始号
	ST_byStartIPAlarmInNo     BYTE    //IP报警输入起始号
	ST_byStartIPAlarmOutNo    BYTE    //IP报警输出起始号
	ST_byHighIPChanNum        BYTE    //数字通道个数，高8位
	ST_byEnableRemotePowerOn  BYTE    //是否启用在设备休眠的状态下远程开机功能，0-不启用，1-启用
	ST_wDevClass              WORD    //设备大类备是属于哪个产品线，0 保留，1-50 DVR，51-100 DVS，101-150 NVR，151-200 IPC，65534 其他，具体分类方法见《设备类型对应序列号和类型值.docx》
	ST_byRes2                 [6]BYTE //保留
}
