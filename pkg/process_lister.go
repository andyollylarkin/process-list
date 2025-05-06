package pkg

type ProcessLister interface {
	ListProcess(match func(int, string) bool) ([]Process, error)
}
