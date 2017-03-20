package goget

var (
	httpInterface = GetHttpInterface()
	logInterface = GetLogInterface()
)

func GoGet(url string) []byte {
	body, err := httpInterface.Get(url)
	if err != nil {
		logInterface.Fatal("Get Failed:", err)
	}
	return body
}
