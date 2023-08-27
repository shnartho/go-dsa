package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Writing to a writer (os.Stdout)
	fmt.Fprintln(os.Stdout, "Hello, io pakcage")

	// Reading from a Reader (strings.Reader)
	reader := strings.NewReader("Simple input for reading.")
	buf := make([]byte, 20)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}

	// Copying data from a reader to a writer:
	src := strings.NewReader("Source data for copying.")
	dest := os.Stdout
	bytesCopied, copyErr := io.Copy(dest, src)
	if copyErr != nil {
		fmt.Println("Error copying:", copyErr)
	} else {
		fmt.Printf("\nCopied %d bytes.\n", bytesCopied)
	}

	//creating a new Reader from a string
	inputString := "Hello from io package!"
	newReader := strings.NewReader(inputString)
	newBuf := make([]byte, len(inputString))
	newReader.Read(newBuf)
	fmt.Println("New Reader:", string(newBuf))

	// writing to a multi-writer (stdout and file)
	file, _ := os.Create("IO.txt")
	defer file.Close()
	multiWriter := io.MultiWriter(os.Stdout, file)
	io.WriteString(multiWriter, "Hello to MultiWriter")

}
