package main

import (
	"bufio"
	//	"errors"
	"fmt"
	"io"
	//"math"
	"os"
	//	"path/filepath"
)


func main() {
	has_args := len(os.Args) > 1

	chars := defaultCharacters()
	if has_args {
		setCharsFromArgs(os.Args[1:], chars)
	}

	reader := bufio.NewReader(os.Stdin)
	wrapStream(reader, chars)
}

func defaultCharacters() map[CharacterKey]string {
	chars := map[CharacterKey]string{
		SetSplit: "",  // TODO?
		SetJoin: "\n",  // TODO?
		SetWrapOpen: "(",
		SetWrapClose: ")",
		RecordSplit: "\n",
		RecordJoin: "\n",
		RecordWrapOpen: "(",
		RecordWrapClose: ")",
		FieldSplit: "\t",
		FieldJoin: " ",
		FieldWrapOpen: "\"",
		FieldWrapClose: "\"",
	}
	return chars
}

func setCharsFromArgs(args []string, chars map[CharacterKey]string) {
	for x := 0 ; x < len(args) ; {
		arg := args[x]
		if ((arg == "-ss") && ((x + 1) < len(args))) {
			chars[SetSplit] = args[(x + 1)]
			x++
		} else if ((arg == "-sj") && ((x + 1) < len(args))) {
			chars[SetJoin] = args[(x + 1)]
			x++
		} else if ((arg == "-sw") && ((x + 2) < len(args))) {
			chars[SetWrapOpen] = args[(x + 1)]
			chars[SetWrapClose] = args[(x + 2)]
			x += 2
		} else if ((arg == "-swo") && ((x + 1) < len(args))) {
			chars[SetWrapOpen] = args[(x + 1)]
			x++
		} else if ((arg == "-swc") && ((x + 1) < len(args))) {
			chars[SetWrapClose] = args[(x + 1)]
			x++
		} else if ((arg == "-rs") && ((x + 1) < len(args))) {
			chars[RecordSplit] = args[(x + 1)]
			x++
		} else if ((arg == "-rj") && ((x + 1) < len(args))) {
			chars[RecordJoin] = args[(x + 1)]
			x++
		} else if ((arg == "-rw") && ((x + 2) < len(args))) {
			chars[RecordWrapOpen] = args[(x + 1)]
			chars[RecordWrapClose] = args[(x + 2)]
			x += 2
		} else if ((arg == "-rwo") && ((x + 1) < len(args))) {
			chars[RecordWrapOpen] = args[(x + 1)]
			x++
		} else if ((arg == "-rwc") && ((x + 1) < len(args))) {
			chars[RecordWrapClose] = args[(x + 1)]
			x++
		} else if ((arg == "-fs") && ((x + 1) < len(args))) {
			chars[FieldSplit] = args[(x + 1)]
			x++
		} else if ((arg == "-fj") && ((x + 1) < len(args))) {
			chars[FieldJoin] = args[(x + 1)]
			x++
		} else if ((arg == "-fw") && ((x + 2) < len(args))) {
			chars[FieldWrapOpen] = args[(x + 1)]
			chars[FieldWrapClose] = args[(x + 2)]
			x += 2
		} else if ((arg == "-fwo") && ((x + 1) < len(args))) {
			chars[FieldWrapOpen] = args[(x + 1)]
			x++
		} else if ((arg == "-fwc") && ((x + 1) < len(args))) {
			chars[FieldWrapClose] = args[(x + 1)]
			x++
		} else if (arg == "-sj-") {
			chars[SetJoin] = ""
		} else if (arg == "-sw-") {
			chars[SetWrapOpen] = ""
			chars[SetWrapClose] = ""
		} else if (arg == "-swo-") {
			chars[SetWrapOpen] = ""
		} else if (arg == "-swc-") {
			chars[SetWrapClose] = ""
		} else if (arg == "-rj-") {
			chars[RecordJoin] = ""
		} else if (arg == "-rw-") {
			chars[RecordWrapOpen] = ""
			chars[RecordWrapClose] = ""
		} else if (arg == "-rwo-") {
			chars[RecordWrapOpen] = ""
		} else if (arg == "-rwc-") {
			chars[RecordWrapClose] = ""
		} else if (arg == "-fj-") {
			chars[FieldJoin] = ""
		} else if (arg == "-fw-") {
			chars[FieldWrapOpen] = ""
			chars[FieldWrapClose] = ""
		} else if (arg == "-fwo-") {
			chars[FieldWrapOpen] = ""
		} else if (arg == "-fwc-") {
			chars[FieldWrapClose] = ""
		} else if (arg == "--csv") {
			chars[SetWrapOpen] = ""
			chars[SetWrapClose] = ""
			chars[RecordWrapOpen] = ""
			chars[RecordWrapClose] = ""
			chars[FieldJoin] = ","
		} else if (arg == "--json") {
			chars[SetWrapOpen] = "["
			chars[SetWrapClose] = "]"
			chars[RecordWrapOpen] = "["
			chars[RecordWrapClose] = "]"
			chars[RecordJoin] = ","
			chars[FieldJoin] = ","
		} else {
			fmt.Fprintf(os.Stderr, "wrap: invalid option '%v' or insufficient options to that argument\n", arg)
		}
		x++
	}
}

func isNext(reader *bufio.Reader, check string) bool {
	if (len(check) == 0) {
		return false
	}
	bytes, err := reader.Peek(len(check))
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Error peeking %v bytes: %s\n", len(check), err)
		return false
	}
//fmt.Fprintf(os.Stderr, "CHECKING '%s' vs '%s'\n", string(bytes), check)
	return string(bytes) == check
}

func skipNext(reader *bufio.Reader, amount int) bool {
	n, err := reader.Discard(amount)
	if (err != nil) || (n != amount) {
		fmt.Fprintf(os.Stderr, "Error discarding %v bytes: %s\n", amount, err)
		return false
	}
	return true
}

func wrapStream(reader *bufio.Reader, wraps map[CharacterKey]string) {
	in_set := false
	in_record := false
	in_field := false
	first_set := true
	first_record := true
	first_field := true

	openSet := func () {
		if !first_set {
			fmt.Fprintf(os.Stdout, wraps[SetJoin])
		}
		fmt.Fprintf(os.Stdout, wraps[SetWrapOpen])
		in_set = true
		first_record = true
	}

	openRecord := func () {
		if !in_set {
			openSet()
		}
		if !first_record {
			fmt.Fprintf(os.Stdout, wraps[RecordJoin])
		}
		fmt.Fprintf(os.Stdout, wraps[RecordWrapOpen])
		in_record = true
		first_field = true
	}

	openField := func () {
		if !in_record {
			openRecord()
		}
		if !first_field {
			fmt.Fprintf(os.Stdout, wraps[FieldJoin])
		}
		fmt.Fprintf(os.Stdout, wraps[FieldWrapOpen])
		in_field = true
	}

	closeField := func () {
		fmt.Fprintf(os.Stdout, wraps[FieldWrapClose])
		in_field = false
		first_field = false
	}

	closeRecord := func () {
		fmt.Fprintf(os.Stdout, wraps[RecordWrapClose])
		in_record = false
		first_record = false
	}

	closeSet := func () {
		fmt.Fprintf(os.Stdout, wraps[SetWrapClose])
		in_set = false
		first_set = false
	}

	closeAll := func() {
		if in_field {
			closeField()
		}
		if in_record {
			closeRecord()
		}
		if in_set {
			closeSet()
		}
	}

	for {
		if (isNext(reader, wraps[SetSplit])) {
//fmt.Fprintf(os.Stderr, "Next char is set split\n", wraps[SetSplit])
			skipNext(reader, len(wraps[SetSplit]))
			if in_field {
				closeField()
			}
			if in_record {
				closeRecord()
			}
			if !in_set {
				openSet()
			}
			closeSet()
		} else if (isNext(reader, wraps[RecordSplit])) {
//fmt.Fprintf(os.Stderr, "Next char is record split\n", wraps[RecordSplit])
			skipNext(reader, len(wraps[RecordSplit]))
			if in_field {
				closeField()
			}
			if !in_record {
				openRecord()
			}
			closeRecord()
		} else if (isNext(reader, wraps[FieldSplit])) {
//fmt.Fprintf(os.Stderr, "Next char is field split\n", wraps[FieldSplit])
			skipNext(reader, len(wraps[FieldSplit]))
			if !in_field {
				openField()
			}
			closeField()
		}

		rune, _, err := reader.ReadRune()
		if err == io.EOF {
			closeAll()
			fmt.Fprintf(os.Stdout, "\n")
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "wrap: Error while reading: %s\n", err)
			break
		}
		char := string(rune)
		if !in_set {
			openSet()
		}
		if !in_record {
			openRecord()
		}
		if !in_field {
			openField()
		}
		if (char == wraps[FieldWrapOpen]) || (char == wraps[FieldWrapClose]) {
			fmt.Fprintf(os.Stdout, "\\")
		}
		fmt.Fprintf(os.Stdout, char)
	}
}

func printUsage(program_name string) {
	fmt.Printf("Usage: %s [TODO]\n", program_name)
}
