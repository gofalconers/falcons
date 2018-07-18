package main

import (
	// "time"
	"os"
	"github.com/mattn/go-redmine"
	"fmt"
	"github.com/sirupsen/logrus"
	// "github.com/zbindenren/logrus_mail"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func fatal(format string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, format, err)
	} else {
		fmt.Fprint(os.Stderr, format)
	}
	os.Exit(1)
}

// func Email() {
// 	logger := logrus.New()
// 	hook, err := logrus_mail.NewMailAuthHook("testapp", "smtp.163.com", 25, "panzhengming91@163.com", "793595755@qq.com", "panzhengming91", "123456@pzm")
// 	if err != nil {
// 		fmt.Println(err)
// 		logger.Hooks.Add(hook)
// 	}
// 	fmt.Println("sendding mail now ...")
// 	var filename = "123.txt"
// 	contextLogger := logger.WithFields(logrus.Fields{
// 		"file": filename,
// 		"content": "Gg",
// 	})
// 	contextLogger.Time = time.Now()
// 	contextLogger.Message = "this a email form hook"
// 	contextLogger.Level = logrus.FatalLevel

// 	hook.Fire(contextLogger)
// 	fmt.Println("sendding mail done .")
// }

func main() {
	client := redmine.NewClient("http://192.168.200.20", "4e60f8ea564f3c0f5cfbc9b68a6632237834a46f")
	issues, err := client.Projects()
	if err != nil {
		fatal("Fail to list projects: %s\n", err)
	}

	for _, i := range issues {
		logrus.Printf("%d: %s\n", i.Id, i.Name)
		logrus.Printf("oops!")
	}
	// logrus.Println("dsfklsdf")
}