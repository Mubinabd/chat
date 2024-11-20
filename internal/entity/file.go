package entity

type File struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type GetFileByIDReq struct {
	ID string `json:"id"`
}

type FileSave struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
}
type FileResponse struct {
	Message  string `json:"message"`
	FilePath string `json:"file_path"`
}