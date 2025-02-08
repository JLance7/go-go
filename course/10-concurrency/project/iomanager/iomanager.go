package iomanager

type IOManager interface{
	WriteResult(data any) error
	ReadLines() ([]string, error)
}