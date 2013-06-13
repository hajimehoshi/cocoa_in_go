package ui

// #cgo LDFLAGS: -L.. -lGoTest
//
// #include "../../GoTest/AppMain.h"
//
import "C"
import (
	"os"
)

//export GoTest_ReceiveMessageFromUI
func GoTest_ReceiveMessageFromUI(message *C.char) {
}

func MainLoop() {
	cargs := []*C.char{}
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	C.GoTest_AppMain(C.int(len(cargs)), &cargs[0])
}
