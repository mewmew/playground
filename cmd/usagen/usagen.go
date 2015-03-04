// usagen generates usage documentation for a given command. It does so by
// executing the command with the "--help" flag and parsing the output.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const use = `
Usage: usagen CMD
Generate usage documentation for a given command.
`

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, use[1:])
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := genUsage(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// genUsage generates usage documentation for a given command.
func genUsage(cmdName string) error {
	// Capture output from running:
	//    foo --help
	buf := new(bytes.Buffer)
	cmd := exec.Command(cmdName, "--help")
	cmd.Stderr = buf
	cmd.Run()

	// Parse output to locate the usage message and flag definitions.
	s := buf.String()
	if len(s) == 0 {
		return fmt.Errorf("unable to capture output of %q", cmdName)
	}
	lines := strings.Split(s, "\n")
	inFlags := false
	var usage string
	var flagSpecs []string
	var flagDescs []string
	for _, line := range lines {
		if strings.HasPrefix(line, "Usage: ") {
			usage = line[len("Usage: "):]
			continue
		}
		if strings.HasPrefix(line, "Flags:") {
			inFlags = true
			continue
		}
		if inFlags {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				continue
			}
			flagSpec := strings.TrimSpace(parts[0]) + ":"
			flagDesc := strings.TrimSpace(parts[1])
			flagSpecs = append(flagSpecs, flagSpec)
			flagDescs = append(flagDescs, flagDesc)
		}
	}
	if len(usage) == 0 {
		return fmt.Errorf("unable to locate usage message for %q", cmdName)
	}

	// Print usage message.
	const format = `
// Usage:
//
//     %s
//
// Flags:
//
`
	fmt.Printf(format[1:], usage)

	// Print command line flags using padding.
	max := 0
	for i := 0; i < len(flagSpecs); i++ {
		if n := len(flagSpecs[i]); n > max {
			max = n
		}
	}
	for i := 0; i < len(flagSpecs); i++ {
		fmt.Printf("//    %-*s %s\n", max, flagSpecs[i], flagDescs[i])
	}

	// Print package clause.
	fmt.Println("package main")

	return nil
}
