package mcsmapi

// --- Common Types ---

type CompressMode int
type FileType int

const (
	CompressModeZip CompressMode = 1
	CompressModeTar CompressMode = 2
	FileTypeFile    FileType     = 0
	FileTypeFolder  FileType     = 1
)

type TargetedFile struct {
	Target string `json:"target"`
}

type FileTargets struct {
	Targets []string `json:"targets"` // List of file/folder paths
}

type FileSourceTargetPairs struct {
	Targets [][2]string `json:"targets"` // source, target
}

type CommonResponse struct {
	Time int64 `json:"time"`
}

// --- Request & Body Structs ---

type GetFileContentsRequestBody TargetedFile

type FileContents struct {
	Target *BaseRequest
	File   *GetFileContentsRequestBody
}

type UpdateFileRequest BaseRequest

type UpdateFileRequestBody struct {
	Target string `json:"target"`
	Text   string `json:"text"`
}

type UpdateFile struct {
	Target   *UpdateFileRequest
	FileData *UpdateFileRequestBody
}

type CopyFile struct {
	Target   *CopyFileRequest
	FileData *FileSourceTargetPairs
}

type MoveFile struct {
	Target   *MoveFileRequest
	FileData *FileSourceTargetPairs
}

type ZipFileRequestBody struct {
	Type    CompressMode `json:"type"`   // 1 for zip
	Code    string       `json:"code"`   // "utf-8"
	Source  string       `json:"source"` // Zip file path
	Targets []string     `json:"targets"`
}

type CompressFile struct {
	Target   *ZipFileRequest
	FileData *ZipFileRequestBody
}

type DeleteFileRequestBody FileTargets

type DeleteFile struct {
	Target   *DeleteFileRequest
	FileData *DeleteFileRequestBody
}

type TouchFileRequestBody TargetedFile

type TouchFile struct {
	Target   *TouchFileRequest
	FileData *TouchFileRequestBody
}

type CreateFolder struct {
	Target *CreateFolderRequest
	Body   *CreateFolderRequestBody
}

type CreateFolderRequestBody TargetedFile

// --- Response Structs ---

type GetFileListResponse struct {
	Status int `json:"status"`
	Data   struct {
		Items        []*FileItem `json:"items"`
		Page         int         `json:"page"`
		PageSize     int         `json:"page_size"`
		Total        int         `json:"total"`
		AbsolutePath string      `json:"absolute_path"`
	} `json:"data"`
	Time int64 `json:"time"`
}

type FileItem struct {
	Name string   `json:"name"`
	Size int      `json:"size"`
	Time string   `json:"time"`
	Mode int      `json:"mode"`
	Type FileType `json:"type"` // 0 = Folder, 1 = File
}

type GetFileContentsResponse struct {
	Data string `json:"data"`
	CommonResponse
}

type UpdateFileResponse BoolResponse
type CopyFileResponse BoolResponse
type MoveFileResponse BoolResponse
type ZipFileResponse BoolResponse
type UnzipFileResponse BoolResponse
type DeleteFileResponse BoolResponse
type TouchFileResponse BoolResponse
type CreateFolderResponse BoolResponse

type DownloadFileResponse struct {
	Data struct {
		Password string `json:"password"`
		Addr     string `json:"addr"`
	} `json:"data"`
	CommonResponse
}

type UploadFileResponse DownloadFileResponse
