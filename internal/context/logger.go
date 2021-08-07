package context

import (
	"log"
	"os"
	"strings"
	"time"

	color "github.com/fatih/color"
)

type Logger struct {
	Logger *log.Logger
}

func LoadLogger() error {
	file := App.Config.Path.Config + time.Now().Format("2006") + "_log" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	App.Logger.Logger = log.New(logFile, "", log.LstdFlags|log.LUTC) // 将文件设置为loger作为输出

	return nil
}

func (l *Logger) Log(msg ...interface{}) {

	switch strings.TrimSpace(strings.ToLower(App.Config.Log)) {
	default:
		fallthrough
	case "info":
		l.Info(msg...)
	case "error":
		l.Error(msg...)
	case "debug":
		l.Debug(msg...)
	case "warn":
		l.Warn(msg...)
	}
}

func (l *Logger) Info(msg ...interface{}) {

	l.Logger.SetPrefix(color.BlueString("[Info]"))
	log.Println(msg...)
}

func (l *Logger) Warn(msg ...interface{}) {
	l.Logger.SetPrefix(color.YellowString("[WARN]"))
	log.Println(msg...)
}

func (l *Logger) Debug(msg ...interface{}) {
	l.Logger.SetPrefix(color.GreenString("[DEBUG]"))
	log.Println(msg...)
}

func (l *Logger) Error(msg ...interface{}) {
	log.SetFlags(log.Llongfile)
	l.Logger.SetPrefix(color.RedString("[ERROR]"))
	log.Fatalln(msg...)
}
