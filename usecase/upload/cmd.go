package upload

type Command struct {
	ConfPath,
	Src,
	ZipPath,
	ZipName,
	DropboxToken,
	UseUploader string
}
