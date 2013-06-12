package main

// #cgo LDFLAGS: -L. -lGoTest
//
// extern int GoTest_AppMain(int argc, const char** argv);
//
import "C"
import "fmt"
import "os"

func main() {
	go func() {
		fmt.Println("Output from Go!")
		fmt.Println("Can you see this message?")
	}()

	cargs := []*C.char{}
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	C.GoTest_AppMain(C.int(len(cargs)), &cargs[0])
}
