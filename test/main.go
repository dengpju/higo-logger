package main

import (
	"fmt"
	"github.com/dengpju/higo-logger/logger"
	"github.com/dengpju/higo-utils/utils"
)

func main()  {
	defer func(){

		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
			logger.LoggerStack(r, utils.GoroutineID())
		}
	}()
	logger.Logrus.Init()
	logger.Logrus.Info("dddd1111")
	err()
}

func err() {
	panic("panic test")
}
