package sftp

type ListRequest struct {
	baseRequest
}

type ListResponse struct {
	Files []*FileInfo `json:"files"`
}

type VerifyResponse struct {
	Connected bool   `json:"connected"`
	Message   string `json:"message"`
}

type UploadRequest struct {
	baseRequest
	FileType FileType `json:"fileType"`
	File     *File    `json:"file"`
}

type UploadResponse struct {
	Uploaded bool   `json:"uploaded"`
	FileID   string `json:"fileId"`
}

type DownloadRequest struct {
	baseRequest
	FileType FileType `json:"fileType"`
}

type DownloadResponse struct {
	File   *File  `json:"file"`
	FileID string `json:"fileId"`
}
