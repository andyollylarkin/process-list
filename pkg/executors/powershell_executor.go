package executors

type PowershellExecutor struct{}

func NewPowershellExecutor() *PowershellExecutor {
	return &PowershellExecutor{}
}

func (e *PowershellExecutor) Exec() ([]byte, error) {
	panic("not implementer yet")
}
