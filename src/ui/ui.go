package ui

// #cgo LDFLAGS: -L.. -lGoTest
//
// #include <stdlib.h>
// #include "../../GoTest/AppMain.h"
//
// extern void GoTest_ReceiveMessageFromUI(char* message);
//
// static void setReceiver() {
//   GoTest_SetSendMessageToGoFunc(GoTest_ReceiveMessageFromUI);
// }
//
import "C"
import (
	"os"
	"unsafe"
)

var receiver = make(chan string)

//export GoTest_ReceiveMessageFromUI
func GoTest_ReceiveMessageFromUI(message *C.char) {
	receiveMessageFromUI(C.GoString(message))
}

func receiveMessageFromUI(message string) {
	receiver<- message
}

func MessageReceiver() <-chan string {
	return receiver
}

func MainLoop() {
	C.setReceiver()

	cargs := []*C.char{}
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	defer func() {
		for _, carg := range cargs {
			C.free(unsafe.Pointer(carg))
		}
		cargs = cargs[0:0]
	}()

	C.GoTest_AppMain(C.int(len(cargs)), &cargs[0])
}
