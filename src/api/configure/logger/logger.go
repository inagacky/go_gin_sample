package logger

import (
	"github.com/inagacky/go_gin_sample/src/api/util"
	l "github.com/sirupsen/logrus"
	"os"
)
const (
	logFileName = "go_sample_output.log"
)

var logger = l.New()

// ロガーの設定
func Init() (*os.File, error) {
	//ログフォーマットの指定
	logger.Formatter = new(l.TextFormatter)

	logger.Level = l.InfoLevel
	logger.ReportCaller = true

	// 環境変数から取得
	logPath := util.Getenv("GO_GIN_SAMPLE_LOG_PATH", "/Users/daisuke/go/src/github.com/inagacky/go_gin_sample/apps/log")

	// ログファイルの作成・追記の準備
	logfile, err := os.OpenFile(logPath + "/" + logFileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		logger.Info(err)
	} else {
		logger.SetOutput(logfile)
	}

	// ログの書き込み先
	logger.Out = logfile

	return logfile, err
}

// Loggerを返却
// TODO: もっと良い書き方を探す
func GetLogger() *l.Logger {
	return logger
}