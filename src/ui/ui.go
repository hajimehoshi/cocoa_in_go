package ui

// #cgo LDFLAGS: -L.. -lGoTest
//
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
	"fmt"
	"os"
)

//export GoTest_ReceiveMessageFromUI
func GoTest_ReceiveMessageFromUI(message *C.char) {
	fmt.Printf("Message from UI: %s\n", C.GoString(message))
}

func MainLoop() {
	C.setReceiver()
	cargs := []*C.char{}
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	C.GoTest_AppMain(C.int(len(cargs)), &cargs[0])
}
