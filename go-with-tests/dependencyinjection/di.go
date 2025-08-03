package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Seb")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
	// writer.Write([]byte("Hello, " + name))
}

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "seb")
// }

// func main() {
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
// }
