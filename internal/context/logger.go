package context

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	color "github.com/fatih/color"
)

type Logger struct {
	Logger *log.Logger
}

func LoadLogger() error {

	if err := os.MkdirAll(App.Config.Path.Log(), 0777); err != nil {
		return err
	}

	file := filepath.Join(App.Config.Path.Log(), "log_"+time.Now().Format("2006")+".txt")
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
	case "debug":
		l.Debug(msg...)
	case "warn":
		l.Warn(msg...)
	}
}

func (l *Logger) Info(msg ...interface{}) {
	// 背景色，但是字体颜色就没法修改了
	// c:=color.New(color.FgCyan).Add(color.BgRed)
	l.Logger.SetPrefix(color.BlueString("[Info]"))
	l.Logger.Println(msg...)
	log.SetPrefix(color.BlueString("[Info]"))
	log.Println(msg...)

}

func (l *Logger) Warn(msg ...interface{}) {
	l.Logger.SetPrefix(color.YellowString("[WARN]"))
	l.Logger.Println(msg...)
	log.SetPrefix(color.YellowString("[WARN]"))
	log.Println(msg...)
}

func (l *Logger) Debug(msg ...interface{}) {
	l.Logger.SetPrefix(color.GreenString("[DEBUG]"))
	l.Logger.Println(msg...)
	log.SetPrefix(color.GreenString("[DEBUG]"))
	log.Println(msg...)
}

func (l *Logger) Error(err error, msg ...interface{}) {
	if err != nil {
		l.Logger.SetFlags(log.Llongfile)
		l.Logger.SetPrefix(color.RedString("[ERROR]"))
		l.Logger.Println(msg...)
		log.SetFlags(log.Llongfile)
		log.SetPrefix(color.RedString("[ERROR]"))
		log.Fatalln(msg...)
	}
}
