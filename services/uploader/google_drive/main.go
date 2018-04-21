package google_drive

type (
	Config struct {
		ConfPath string
	}

	Uploader struct {
	}
)

func NewUploader(c Config) *Uploader {
	return &Uploader{}
}

func (u *Uploader) Upload(src, dst string) error {
	return nil
}
