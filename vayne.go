package Vayne

type TestVayne struct {
	ConfigPath string
	FilePath   string
}

// NewVayne constructor
func NewVayne(configPath, filePath string) *TestVayne {
	return &TestVayne{
		FilePath:   filePath,
		ConfigPath: configPath,
	}
}

func (n *TestVayne) Process() bool {
	return true
}
