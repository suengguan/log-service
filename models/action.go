package models

type Action struct {
	Id        int64  `json:"id" orm:"column(ID)"`
	Time      int64  `json:"time" orm:"column(TIME)"`
	SessionId string `json:"sessionId" orm:"column(SESSION_ID)"`
	User      *User  `json:"user" orm:"column(USER_ID);rel(fk)"` //设置一对多关系
	DevType   int    `json:"devType" orm:"column(DEV_TYPE)"`
	Type      int    `json:"type" orm:"column(TYPE)"`
	Content   string `json:"content" orm:"column(CONTENT)"` // log内容
}
