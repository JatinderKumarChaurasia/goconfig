package goconfig

import "log"

import (
	_ "github.com/jatinderkumarchaurasia/goconfig/filelooker"
)

func init() {
	println("setting go configurations")
}

func GoConfig() {
	log.Println("hello world")
}
