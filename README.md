### Go File Signature Validator

small library to validate Files by reading each magic number from a file

#### Install

```shell
$ go get github.com/telkomdev/go-filesig
```

#### Usage

Simple

```go
package main

import (
	"fmt"
	"github.com/telkomdev/go-filesig"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("required input file")
		os.Exit(1)
	}

	inputFileArg := args[1]
	inFile, err := os.Open(inputFileArg)

	if err != nil {
		fmt.Println("error open input file ", err)
		os.Exit(1)
	}

	defer func() { inFile.Close() }()

	valid := filesig.Is3gp(inFile)
	fmt.Println(valid)
```

One Of
```go
package main

import (
	"fmt"
	"github.com/telkomdev/go-filesig"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("required input file")
		os.Exit(1)
	}

	inputFileArg := args[1]
	inFile, err := os.Open(inputFileArg)

	if err != nil {
		fmt.Println("error open input file ", err)
		os.Exit(1)
	}

	defer func() { inFile.Close() }()

	valid := filesig.IsOneOf(inFile, filesig.Is3gp, filesig.IsPng, filesig.IsJpeg)
	fmt.Println(valid)

}

```

HTTP Form File
```go
package main

import (
	"fmt"
	"github.com/telkomdev/go-filesig"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// validate file
	validPdf := filesig.IsPdf(file)
	if !validPdf {
		fmt.Print("file is not valid PDF file \n")
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}

```