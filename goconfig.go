package goconfig

import (
	"github.com/jatinderkumarchaurasia/goconfig/filelooker"
	_ "github.com/jatinderkumarchaurasia/goconfig/filelooker"
)

func init() {
	filelooker.FileLooker()
}
