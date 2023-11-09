package types

type NET_DVR_TIME_EX struct {
	WYear    WORD
	ByMonth  byte
	ByDay    byte
	ByHour   byte
	ByMinute byte
	BySecond byte
	ByRes    byte
}

type LPNET_DVR_TIME_EX *NET_DVR_TIME_EX
