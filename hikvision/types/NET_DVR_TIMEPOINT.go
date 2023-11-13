package types

type NET_DVR_TIMEPOINT struct {
	DwMonth    DWORD // Month . 0-11stands for Jan~Dec
	DwWeekNo   DWORD // Week : 0-1st week; 1-2nd week; 2-3rd week; 3-4th week; 4-last week
	DwWeekDate DWORD // Week number : 0-Sun; 1-Mon; 2-Tue; 3-Wed; 4-Thu; 5-Fri; 6-Sat
	DwHour     DWORD // begin/end hour, begin hour ranges from 0-23, and end hour ranges from 1-23
	DwMin      DWORD // Minute: 0 - 59
}
type LPNET_DVR_TIMEPOINT *NET_DVR_TIMEPOINT
