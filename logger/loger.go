package logger

//import "github.com/fatih/color"

//import "github.com/fatih/color"
import (
	"fmt"
	"github.com/gookit/color"
)

func Info(a ...interface{}) {
	color.White.Println(a...)
}

func Success(a ...interface{}) {
	color.Green.Println(a...)
}

func Warn(a ...interface{}) {
	color.Yellow.Println(a...)
}

func Error(format string, a ...interface{}) {
	color.Red.Printf(format, a...)
	fmt.Println()
}

func Notice(format string, a ...interface{}) {
	d := color.New(color.FgBlue, color.Bold)
	d.Printf(format, a...)
}

//
//
//func Info(format string, a ...interface{}) {
//	color.White(format, a...)
//}
//
//func Success(format string, a ...interface{}) {
//	color.Green.Println(format, a...)
//}
//
//func Warn(format string, a ...interface{}) {
//	color.Yellow(format, a...)
//}
//
//func Error(format string, a ...interface{}) {
//	color.Red(format, a...)
//}
//
//func Notice(format string, a ...interface{}) {
//	d := color.New(color.FgBlue, color.Bold)
//	d.Printf(format, a...)
//}
