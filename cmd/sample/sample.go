package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var checkLog = logrus.New()

type ISimpleFormatter struct{}

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

func TS() {
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
	// SampleA()
}

type FaSession struct {
	Value int
}

type FaReq struct {
	FaSession      FaSession
	FaSessionPoint *FaSession
}

func Fa() {
	req := FaReq{}
	fmt.Printf("%v\n", req.FaSession.Value)
	fmt.Printf("%v\n", req.FaSessionPoint == nil)
}

var randGen *rand.Rand

func init() {
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Fb() {
	same := 0
	latest := 0
	for i := 0; i < 1e11; i++ {
		cur := randGen.Intn(1e5)
		if latest == cur {
			same += 1
			if same > 10 {
				fmt.Printf("same value, value=%v, count=%v\n", cur, same)
			}
		} else {
			same = 0
		}
		latest = cur
	}
}

func main() {
	Fb()
}
