package models

type FileUpload struct {
	FileID   int64  `json:"file_id"db:"file_id"`
	UserId   int64  `json:"user_id"db:"user_id"`
	FileName string `json:"file_name"db:"filename"`
}
