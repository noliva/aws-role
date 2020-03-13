// Copyright © 2018 Ryan French <ryan.french@youview.com>

package main

import (
	"github.com/ryanfrench/aws-role/cmd"
	"github.com/sirupsen/logrus"
)

func init(){
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	cmd.Execute()
}
