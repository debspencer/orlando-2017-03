package goget

var (
	httpInterface = GetHttpInterface()
	ftpInterface  = GetFtpInterface()
	logInterface  = GetLogInterface()
)

func GoGet(url string) []byte {
	var iface DownloadInterface
	if strings.HasPrefix("http") { iface = httpInterface }
	else if strings.HasPrefix("ftp") { iface = ftpInterface }
	else {
		logInterface.Fatal("Unknown scheme", url)
	}

	body, err := iface.Get(url)
	if err != nil {
		logInterface.Fatal("Get Failed:", err)
	}
	return body
}
