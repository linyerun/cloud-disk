package dto

type UserFile struct {
	ID       uint   `json:"id"`
	FileId   uint   `json:"file_id,omitempty"`
	Filename string `json:"filename"`
	FileType uint8  `json:"file_type"`
	ParentId uint   `json:"parent_id"`
}
