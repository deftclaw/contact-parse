package main

import "fmt"

func main() {
	ifile := "in/sample.vcf"  // The input file

	buff := Buffer{ content: make([]byte, 256), cursor: 0, line: make([]byte, 256) }  // The in-memory portion of the file, with a cursor index
	err  := buff.readChunk(ifile)
	check(err)

	fmt.Printf("%d bytes:\n%s\nCursor: %d\nLine: %s\n\n", len(buff.content), string(buff.content), buff.cursor, string(buff.line))  // NOTE: Debug: How many bytes were read and their contents
	buff.nextLine()
	fmt.Println("GetLine: ", string(buff.line), buff.cursor)
	buff.nextLine()
	fmt.Println("GetNextLine: ", string(buff.line), buff.cursor)
}
