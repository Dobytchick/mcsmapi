package mcsmapi

type GetFileListRequest struct {
	BaseRequest
	Target   string `url:"target"`
	Page     int    `url:"page"`
	PageSize int    `url:"page_size"`
	FileName string `url:"file_name"`
}

func (q *GetFileListRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type GetFileContentsRequest struct {
	BaseRequest
}

func (q *GetFileContentsRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type DownloadFileRequest struct {
	BaseRequest
	FileName string `url:"file_name"`
}

func (q *DownloadFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type UploadFileRequest DownloadFileRequest

func (q *UploadFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type CopyFileRequest struct {
	BaseRequest
}

func (q *CopyFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type MoveFileRequest struct {
	BaseRequest
}

func (q *MoveFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type ZipFileRequest struct {
	BaseRequest
}

func (q *ZipFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type DeleteFileRequest struct {
	BaseRequest
}

func (q *DeleteFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type TouchFileRequest struct {
	BaseRequest
}

func (q *TouchFileRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type CreateFolderRequest struct {
	BaseRequest
}

func (q *CreateFolderRequest) BuildQueryString() string {
	return BuildQueryString(q)
}
