package uploader

type UploaderConfig struct {
	Name string
}

func NewUploaderConfig(name string) *UploaderConfig {
	return &UploaderConfig{
		Name: name,
	}
}
