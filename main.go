package main

import (
	"fmt"
	// "maps"
	"strings"
)

type ContactCard struct {
	version string
}

type ContactName struct {
	given    string
	middle   string
	phonetic string
	prefix   string
	suffix   string
	sur      string
}

func first(c ContactName) string { return c.given }
func last(c ContactName)  string { return c.sur }

func main() {
	ifile := "in/sample.vcf"  // The input file

	buff := Buffer{ content: make([]byte, 256), cursor: 0, line: make([]byte, 256) }  // The in-memory portion of the file, with a cursor index
	err  := buff.readChunk(ifile)
	check(err)

	confirmVCF(&buff, ifile)  // Exit unless the target file starts 'BEGIN:VCARD'

  	// NOTE: Debug: How many bytes were read and their contents
	fmt.Printf("%d bytes:\n%s\n\nCursor: %d\nLine: %s\n", len(buff.content), string(buff.content), buff.cursor, string(buff.line))
	buff.nextLine()
	parseLine(string(buff.line))
}

func parseLine(line string) map[string]string {
	components := strings.Split(line, ":")

//	var properties = map[string]func(string, string) struct{} {
//		"N": parseName,
//		"FN": "FullName",
//		"X-PHONETIC-LAST-NAME": "PhoneticSurname",
//		"TEL": "PhoneNumber",
//		"VERSION": "vCardVersion",
//	}

	prop := strings.Split(components[0], ";")[0]

	switch prop {
	case "N":
		fmt.Println("Name: ", parseName(components[0], components[1]))
	default:
		fmt.Printf("Property not found: %s", prop)
	}

	for p := range properties {
		if components[0] == p {
			fmt.Printf("%s\n", p)
			properties[p](components[0], components[1])
		}
	}

	return properties
}

func parseName(props string, values string) ContactName {
	components := strings.Split(values, ";")
	modifiers  := strings.Split(props, ";")

	return ContactName{
		sur: components[0],
		given: components[1],
		middle: components[2],
		prefix: components[3],
		suffix: components[4],
	}
}
