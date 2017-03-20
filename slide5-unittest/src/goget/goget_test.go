package goget

import (
	"reflect"
	"testing"
)

func TestGoGet(t *testing.T) {
	mock := HttpMock()
	mock.Body = []byte("<html>stuff</html>")
	httpInterface = mock

	body := GoGet("http://gopher.com")
	if ! reflect.DeepEqual(mock.Body, body) {
		t.Errorf("Did not get expected response %v != %v\n", mock.Body, body)
	}
}
