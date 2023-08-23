package datamodel

type User struct {
	UUID        string `json:"uuid"`
	LineUUID    string `json:"line_uuid"`
	DisplayName string `json:"display_name"`
	AvatarPath  string `json:"avatar_path"`
}
