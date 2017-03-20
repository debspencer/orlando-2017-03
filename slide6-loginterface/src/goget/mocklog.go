// +build !prod

package goget

type logMock struct {
	Logged bool
}

func LogMock() *logMock {
	return &logMock{}
}

func (m *logMock) Fatal(...interface{}) {
	m.Logged = true
}
