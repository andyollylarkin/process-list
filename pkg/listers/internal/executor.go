package internal

type Executor interface {
	Exec() ([]byte, error)
}
