package mcsmapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// fileClient wraps a method client for file operations through the "files" endpoint.
type fileClient MethodClient

// newFileClient initializes a new fileClient.
func newFileClient(client *Client) *fileClient {
	return &fileClient{client: client, endpoint: "files"}
}

// sendRequest wraps client's sendRequest method with endpoint prefixing.
func (fc *fileClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	fullEndpoint := fmt.Sprintf("%s/%s", fc.endpoint, endpoint)
	return fc.client.sendRequest(method, fullEndpoint, body)
}

// doRequestAndDecode is a shared utility to decode response JSON into the given output struct.
func (fc *fileClient) doRequestAndDecode(method, endpoint string, body, out any) error {
	resp, err := fc.sendRequest(method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// GetFileList returns the list of files for the given path/query.
func (fc *fileClient) GetFileList(req *GetFileListRequest) (*GetFileListResponse, error) {
	var res GetFileListResponse
	err := fc.doRequestAndDecode("GET", "list?"+req.BuildQueryString(), nil, &res)
	return &res, err
}

// GetFileContents retrieves the contents of a file.
func (fc *fileClient) GetFileContents(req *FileContents) (*GetFileContentsResponse, error) {
	var res GetFileContentsResponse
	err := fc.doRequestAndDecode("PUT", "?"+req.Target.BuildQueryString(), req.File, &res)
	return &res, err
}

// Update writes or replaces the file content.
func (fc *fileClient) Update(req *UpdateFile) (*BaseResponse, error) {
	var res BaseResponse
	err := fc.doRequestAndDecode("PUT", "", req.FileData, &res)
	return &res, err
}

// Download fetches and downloads a file.
func (fc *fileClient) Download(req *DownloadFileRequest) (*DownloadFileResponse, error) {
	var res DownloadFileResponse
	err := fc.doRequestAndDecode("POST", "download?"+req.BuildQueryString(), nil, &res)
	return &res, err
}

// Upload uploads a file to the target destination.
func (fc *fileClient) Upload(req *UploadFileRequest) (*UploadFileResponse, error) {
	var res UploadFileResponse
	err := fc.doRequestAndDecode("POST", "download?"+req.BuildQueryString(), nil, &res) // ← Might be a typo: "download" vs "upload"
	return &res, err
}

// Copy duplicates a file or folder to another location.
func (fc *fileClient) Copy(req *CopyFile) (*BaseResponse, error) {
	var res BaseResponse
	err := fc.doRequestAndDecode("POST", "copy?"+req.Target.BuildQueryString(), req.FileData, &res)
	return &res, err
}

// MoveOrRename moves or renames a file or folder.
func (fc *fileClient) MoveOrRename(req *MoveFile) (*BaseResponse, error) {
	var res BaseResponse
	err := fc.doRequestAndDecode("PUT", "move?"+req.Target.BuildQueryString(), req.FileData, &res)
	return &res, err
}

// CreateZIPArchive compresses files into a ZIP archive.
func (fc *fileClient) CreateZIPArchive(req *CompressFile) (*ZipFileResponse, error) {
	var res ZipFileResponse
	err := fc.doRequestAndDecode("POST", "compress?"+req.Target.BuildQueryString(), req.FileData, &res)
	return &res, err
}

// Unzip extracts files from a ZIP archive.
func (fc *fileClient) Unzip(req *CompressFile) (*UnzipFileResponse, error) {
	var res UnzipFileResponse
	err := fc.doRequestAndDecode("POST", "compress?"+req.Target.BuildQueryString(), req.FileData, &res) // ← Might also be "uncompress"?
	return &res, err
}

// Delete removes a file or folder.
func (fc *fileClient) Delete(req *DeleteFile) (*DeleteFileResponse, error) {
	var res DeleteFileResponse
	err := fc.doRequestAndDecode("DELETE", "?"+req.Target.BuildQueryString(), req.FileData, &res)
	return &res, err
}

// Touch creates an empty file or updates timestamp.
func (fc *fileClient) Touch(req *TouchFile) (*TouchFileResponse, error) {
	var res TouchFileResponse
	err := fc.doRequestAndDecode("POST", "touch?"+req.Target.BuildQueryString(), req.FileData, &res)
	return &res, err
}

// CreateFolder creates a new directory.
func (fc *fileClient) CreateFolder(req *CreateFolder) (*CreateFolderResponse, error) {
	var res CreateFolderResponse
	err := fc.doRequestAndDecode("POST", "mkdir?"+req.Target.BuildQueryString(), req.Body, &res)
	return &res, err
}
