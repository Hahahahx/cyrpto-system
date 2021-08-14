package context

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	color "github.com/fatih/color"
)

type Logger struct {
	Logger *log.Logger
}

func LoadLogger() error {
	file := filepath.Join(App.Config.Path.Log(), "log_"+time.Now().Format("2006")+".txt")
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	App.Logger.Logger = log.New(mw, "", log.LstdFlags|log.LUTC) // 将文件设置为loger作为输出
	return nil
}

func (l *Logger) Log(msg ...interface{}) {

	switch strings.TrimSpace(strings.ToLower(App.Config.Log)) {
	default:
		fallthrough
	case "info":
		l.Info(msg...)
	case "debug":
		l.Debug(msg...)
	case "warn":
		l.Warn(msg...)
	}
}

func backGround(cBg color.Attribute, format string, a ...interface{}) string {
	c := color.New(color.FgCyan).Add(cBg)
	return c.Sprint(color.BlackString(format, a...))
}

func (l *Logger) Info(msg ...interface{}) {

	l.Logger.SetPrefix(backGround(color.BgBlue, " INFO ") + "  ")
	l.Logger.Println(msg...)

}

func (l *Logger) Warn(msg ...interface{}) {
	l.Logger.SetPrefix(backGround(color.BgYellow, " WARN ") + "  ")
	l.Logger.Println(msg...)
}

func (l *Logger) Debug(msg ...interface{}) {
	l.Logger.SetPrefix(backGround(color.BgGreen, " DEBUG ") + " ")
	l.Logger.Println(msg...)
}

func (l *Logger) Error(err error, msg ...interface{}) {
	if err != nil {
		l.Logger.SetPrefix(backGround(color.BgRed, " ERROR ") + " ")
		if App.Config.Log == "debug" {
			_, file, line, _ := runtime.Caller(1)
			l.Logger.Println(file+":"+strconv.Itoa(line), err, fmt.Sprint(msg...))
		} else {
			l.Logger.Println(err, fmt.Sprint(msg...))
		}
		os.Exit(3)
	}
}
