package main

import (
	"io"
	"os"
)

type Buffer struct {
	cursor  int
	content []byte
	line    []byte
}

func (b Buffer) readChunk(path string) error {
	handle, err := os.Open(path)
	if err != nil { return err }

	defer handle.Close()

	_, err  = handle.Seek(int64(b.cursor), io.SeekStart)
	if err != nil { return err }

	_, err = handle.Read(b.content)
	if err != nil { return err }

	return nil
}

func (b *Buffer) nextLine() {
	var start int
	var end   int

	if b.cursor == len(b.content) { b.cursor = 0 }  // If the cursor reached the end of the buffer start from the begining

	// Save the character index if line-return, else skip
	for cdx, char := range b.content[b.cursor:] {
		if char == '\r' {
			end = b.cursor + cdx
			break
		}
	}

	start    = b.cursor  // Record where we started from
	b.cursor = end + 2   // Move the cursor to the next line

	b.line = b.content[start:end]  // Return just the next line of the buffer
}
