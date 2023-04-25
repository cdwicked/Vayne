package Vayne

type TestVayne struct {
	FilePath string
}

// NewVayne constructor
func NewVayne(filePath string) *TestVayne {
	return &TestVayne{
		FilePath: filePath,
	}
}

func (n *TestVayne) Process() bool {
	return true
}
