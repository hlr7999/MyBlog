package app

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("main")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLogger() {
	logBackend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatted := logging.NewBackendFormatter(logBackend, format)
	logging.SetBackend(backendFormatted)
}
