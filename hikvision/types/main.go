package types

type (
	LLONG = int64
	LONG  = int32
	DWORD = uint32
	WORD  = uint16
	SHORT = int16
	BYTE  = byte
	char  = byte
	BOOL  = int32
	HWND  = uint32
)

type DvrCfg interface {
	releaseDvrCfg()
}
