package iomanager

// I/O 작업을 처리하는 인터페이스 (abstract)
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(interface{}) error
}
