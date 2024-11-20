package entity

type File struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type GetFileByIDReq struct {
	ID string `json:"id"`
}

type SaveFile struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
}
