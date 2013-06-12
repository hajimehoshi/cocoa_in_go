package main

// #cgo LDFLAGS: -L. -lGoTest
//
// extern int GoTest_AppMain(int argc, const char** argv);
//
import "C"
import "fmt"

func main() {
	go func() {
		fmt.Println("Output from Go!")
		fmt.Println("Can you see this message?")
	}()

	args := []*C.char{C.CString("foo")}
	C.GoTest_AppMain(C.int(len(args)), &args[0])
}
