package pkg

type ProcessLister interface {
	ListProcess(DoneLookupFunc) ([]Process, error)
}
