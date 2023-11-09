package types

type NET_DVR_CARD_CFG struct {
	DwSize            DWORD
	DwModifyParamType DWORD
	// 需要修改的卡参数，设置卡参数时有效，按位表示，每位代表一种参数，1为需要修改，0为不修改
	ByCardNo        [ACS_CARD_NO_LEN]byte //卡号
	ByCardValid     byte                  //卡是否有效，0-无效，1-有效（用于删除卡，设置时置为0进行删除，获取时此字段始终为1）
	ByCardType      byte                  //卡类型，1-普通卡，2-残疾人卡，3-黑名单卡，4-巡更卡，5-胁迫卡，6-超级卡，7-来宾卡，8-解除卡，默认普通卡
	ByLeaderCard    byte                  //是否为首卡，1-是，0-否
	ByRes1          byte
	DwDoorRight     DWORD                                       //门权限，按位表示，1为有权限，0为无权限，从低位到高位表示对门1-N是否有权限
	StruValid       NET_DVR_VALID_PERIOD_CFG_BYDOC              //有效期参数
	DwBelongGroup   DWORD                                       //所属群组，按位表示，1-属于，0-不属于，从低位到高位表示是否从属群组1-N
	ByCardPassword  [CARD_PASSWORD_LEN]byte                     //卡密码
	ByCardRightPlan [MAX_DOOR_NUM][MAX_CARD_RIGHT_PLAN_NUM]byte //卡权限计划，取值为计划模板编号，同个门不同计划模板采用权限或的方式处理
	DwMaxSwipeTime  DWORD                                       //最大刷卡次数，0为无次数限制
	DwSwipeTime     DWORD                                       //已刷卡次数
	WRoomNumber     WORD                                        //房间号
	WFloorNumber    WORD                                        //层号
	ByRes2          [20]byte
}

type LPNET_DVR_CARD_CFG *NET_DVR_CARD_CFG

type NET_DVR_VALID_PERIOD_CFG_BYDOC struct {
	ByEnable      byte //使能有效期，0-不使能，1使能
	ByRes1        [2]byte
	ByTimeType    byte
	StruBeginTime NET_DVR_TIME_EX //有效期起始时间
	StruEndTime   NET_DVR_TIME_EX //有效期结束时间
	ByRes2        [32]byte
}

type LPNET_DVR_VALID_PERIOD_CFG_BYDOC *NET_DVR_VALID_PERIOD_CFG_BYDOC
