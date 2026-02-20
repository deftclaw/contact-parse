package main

import (
	"bytes"
	"fmt"
	"os"
)

func checkEmpty(b Buffer) bool {
	if bytes.Equal(b.line, make([]byte, 256)) { return true }

	return false
}

func checkInit(b Buffer) bool {
	if string(b.line) == "BEGIN:VCARD" { return true }

	return false
}

func confirmVCF(b *Buffer, fname string) {
	for {
		if checkEmpty(*b) { b.nextLine() }
	
		if checkInit(*b) { break } else {
			fmt.Printf("%s is not a valid vcf contact file\n", fname)
			os.Exit(3)
		}
	}
}
