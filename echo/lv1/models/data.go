package models

type LoginRes struct {
	Token string `json:"token" form:"token"`
}

type LoginReq struct {
	User string `json:"user" form:"user"`
	Pass string `json:"pass" form:"pass"`
}

type X struct {
	Text string `json:"text" form:"text"`
}
