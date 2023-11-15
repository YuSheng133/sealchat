package model

type MessageModel struct {
	StringPKBaseModel
	Content   string `json:"content"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id" gorm:"null"`
	MemberID  string `json:"member_id" gorm:"null"`
	UserID    string `json:"user_id" gorm:"null"`

	User   *UserModel   `json:"user"`   // 嵌套 User 结构体
	Member *MemberModel `json:"member"` // 嵌套 Member 结构体
}

func (*MessageModel) TableName() string {
	return "messages"
}