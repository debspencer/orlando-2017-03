package goget
import 	"fmt"
import 	"reflect"
import 	"testing"

func TestGoGet(t *testing.T) {
	mock := HttpMock()
	mock.Body = []byte("<html>stuff</html>")
	httpInterface = mock
	log := LogMock()
	logInterface = log

	body := GoGet("http://gopher.com")
	if ! reflect.DeepEqual(mock.Body, body) {
		t.Errorf("Did not get expected response %v != %v\n", string(mock.Body), string(body))
	}
	mock.Body = []byte{}
	mock.Error = fmt.Errorf("Fail")
	body = GoGet("http://gopher.com")
	if len(body) != 0 {
		t.Errorf("Did not get expected response %v\n", string(body))
	}
	if log.Logged != true { t.Errorf("Expected a log call") }
}
