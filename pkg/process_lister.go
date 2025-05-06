package pkg

type ProcessLister interface {
	ListProcess() ([]Process, error)
}
