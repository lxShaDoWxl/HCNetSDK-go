package types

type NET_DVR_CARD_CFG_V50 struct {
	DwSize            DWORD
	DwModifyParamType DWORD
	// 需要修改的卡参数，设置卡参数时有效，按位表示，每位代表一种参数，1为需要修改，0为不修改
	ByCardNo           [ACS_CARD_NO_LEN]byte                           //卡号
	ByCardValid        byte                                            //Действительна ли карта, 0-недействительна, 1-действительна (используется для удаления карты, при настройке удаления выставляется на 0, при получении это поле всегда равно 1)
	ByCardType         byte                                            //卡类型，1-普通卡，2-残疾人卡，3-黑名单卡，4-巡更卡，5-胁迫卡，6-超级卡，7-来宾卡，8-解除卡，9-员工卡，10-应急卡，11-应急管理卡（用于授权临时卡权限，本身不能开门），默认普通卡
	ByLeaderCard       byte                                            //是否为首卡，1-是，0-否
	ByUserType         byte                                            // 0 – 普通用户1 - 管理员用户;
	ByDoorRight        [MAX_DOOR_NUM_256]byte                          //门权限(楼层权限、锁权限)，按位表示，1为有权限，0为无权限，从低位到高位表示对门（锁）1-N是否有权限
	StruValid          NET_DVR_VALID_PERIOD_CFG                        //有效期参数
	ByBelongGroup      [MAX_GROUP_NUM_128]byte                         //所属群组，按字节表示，1-属于，0-不属于
	ByCardPassword     [CARD_PASSWORD_LEN]byte                         //卡密码
	WCardRightPlan     [MAX_DOOR_NUM_256][MAX_CARD_RIGHT_PLAN_NUM]WORD //卡权限计划，取值为计划模板编号，同个门（锁）不同计划模板采用权限或的方式处理
	DwMaxSwipeTime     DWORD                                           //最大刷卡次数，0为无次数限制（开锁次数）
	DwSwipeTime        DWORD                                           //已刷卡次数
	WRoomNumber        WORD                                            //房间号
	WFloorNumber       SHORT                                           //层号
	DwEmployeeNo       DWORD                                           //工号（用户ID）
	ByName             [NAME_LEN]byte                                  //姓名
	WDepartmentNo      WORD                                            //部门编号
	WSchedulePlanNo    WORD                                            //排班计划编号
	BySchedulePlanType byte                                            //排班计划类型：0-无意义、1-个人、2-部门
	ByRightType        byte                                            //下发权限类型：0-普通发卡权限、1-二维码权限、2-蓝牙权限（可视对讲设备二维码权限配置项：房间号、卡号（虚拟卡号）、最大刷卡次数（开锁次数）、有效期参数；蓝牙权限：卡号（萤石APP账号）、其他参数配置与普通发卡权限一致）
	ByRes2             [2]byte
	DwLockID           DWORD                   //锁ID
	ByLockCode         [MAX_LOCK_CODE_LEN]byte //锁代码
	ByRoomCode         [MAX_DOOR_CODE_LEN]byte //房间代码
	//按位表示，0-无权限，1-有权限
	//第0位表示：弱电报警
	//第1位表示：开门提示音
	//第2位表示：限制客卡
	//第3位表示：通道
	//第4位表示：反锁开门
	//第5位表示：巡更功能
	DwCardRight     DWORD                  //卡权限
	DwPlanTemplate  DWORD                  //计划模板(每天)各时间段是否启用，按位表示，0--不启用，1-启用
	DwCardUserId    DWORD                  //持卡人ID
	ByCardModelType byte                   //0-空，1- MIFARE S50，2- MIFARE S70，3- FM1208 CPU卡，4- FM1216 CPU卡，5-国密CPU卡，6-身份证，7- NFC
	BySIMNum        [NAME_LEN] /*32*/ byte //SIM卡号（手机号）
	ByRes3          [51]byte
}
type LPNET_DVR_CARD_CFG_V50 *NET_DVR_CARD_CFG_V50
