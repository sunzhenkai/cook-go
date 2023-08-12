package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
)

var checkLog = logrus.New()

type ISimpleFormatter struct {
}

func (f *ISimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}

func InitCheckLog() {
	file, err := os.OpenFile("tmp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	checkLog.SetFormatter(new(ISimpleFormatter))
	if err == nil {
		checkLog.Out = file
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
	InitCheckLog()
	m := make(map[string]interface{})
	m["a"] = 1
	m["b"] = "c"
	bs, _ := json.Marshal(m)
	fmt.Println(string(bs))
	checkLog.Println(string(bs))
	SampleA()
}
