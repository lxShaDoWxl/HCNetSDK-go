package types

type NET_DVR_VALID_PERIOD_CFG struct {
	ByEnable         byte            //使能有效期，0-不使能，1使能
	ByBeginTimeFlag  byte            //是否限制起始时间的标志，0-不限制，1-限制
	ByEnableTimeFlag byte            //是否限制终止时间的标志，0-不限制，1-限制
	ByTimeDurationNo byte            //有效期索引,从0开始（时间段通过SDK设置给锁，后续在制卡时，只需要传递有效期索引即可，以减少数据量）
	StruBeginTime    NET_DVR_TIME_EX //有效期起始时间
	StruEndTime      NET_DVR_TIME_EX //有效期结束时间
	ByTimeType       byte
	ByRes2           [31]byte
}

type LPNET_DVR_VALID_PERIOD_CFG *NET_DVR_VALID_PERIOD_CFG
