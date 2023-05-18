package system

type VerificationCodeDao struct {
	ImageBase string
	BgBase    string
	X         int
	Y         string
}

type VerificationCodeVo struct {
	ImageBase string `json:"imageBase,omitempty"`
	BgBase    string `json:"bgBase,omitempty"`
	Y         string `json:"y,omitempty"`
}
