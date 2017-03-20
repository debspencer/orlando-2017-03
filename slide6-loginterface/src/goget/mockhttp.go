// +build !prod

package goget

type Mock struct {
	Error         error
	Body          []byte
}

func HttpMock() *Mock {
	m := &Mock{}
	return m
}

func (m *Mock) Get(url string) ([]byte, error) {
	return m.Body, m.Error
}
