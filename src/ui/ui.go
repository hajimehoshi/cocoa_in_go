package ui

// #cgo LDFLAGS: -L. -lGoTest
//
// #include <stdlib.h>
// #include "../../GoTest/AppMain.h"
//
// typedef const char* cstring;
//
// void GoTest_ReceiveMessageFromUI(const char* message);
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
var sender = make(chan string)

//export GoTest_ReceiveMessageFromUI
func GoTest_ReceiveMessageFromUI(message C.cstring) {
	receiver <- C.GoString(message)
}

func MessageReceiver() <-chan string {
	return receiver
}

func MessageSender() chan<- string {
	return sender
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
		cargs = nil
	}()

	go func() {
		for message := range sender {
			cmessage := C.CString(message)
			defer C.free(unsafe.Pointer(cmessage))

			C.GoTest_SendMessageToUI(cmessage)
		}
	}()

	C.GoTest_AppMain(C.int(len(cargs)), &cargs[0])
}
