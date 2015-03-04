//go:generate usagen usagen

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

var (
	// flagOut specifies the output path.
	flagOut string
	// flagPlain specifies if the output should be plain text or Go source code.
	flagPlain bool
)

const use = `
Usage: usagen [OPTION]... CMD
Generate usage documentation for a given command.

Flags:`

func init() {
	flag.StringVar(&flagOut, "o", "z_usage.go", "Output path.")
	flag.BoolVar(&flagPlain, "plain", false, "Plain text output (default: Go source code)")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, use[1:])
		flag.PrintDefaults()
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
			if len(parts) < 2 {
				continue
			}
			flagSpec := strings.TrimSpace(parts[0]) + ":"
			flagDesc := strings.TrimSpace(strings.Join(parts[1:], ":"))
			flagSpecs = append(flagSpecs, flagSpec)
			flagDescs = append(flagDescs, flagDesc)
		}
	}
	if len(usage) == 0 {
		return fmt.Errorf("unable to locate usage message for %q", cmdName)
	}

	// Print usage message.
	out := new(bytes.Buffer)
	const format = `
Usage:

    %s

Flags:

`
	fmt.Fprintf(out, format[1:], usage)

	// Print command line flags using padding.
	max := 0
	for i := 0; i < len(flagSpecs); i++ {
		if n := len(flagSpecs[i]); n > max {
			max = n
		}
	}
	for i := 0; i < len(flagSpecs); i++ {
		fmt.Fprintf(out, "    %-*s %s\n", max, flagSpecs[i], flagDescs[i])
	}
	return output(out.String())

	return nil
}

// output writes the usage message to the specified output file.
func output(usage string) error {
	f, err := os.Create(flagOut)
	if err != nil {
		return err
	}
	defer f.Close()

	// Early exit for plain text output.
	if flagPlain {
		_, err = fmt.Fprint(f, usage)
		return err
	}

	// Go source code output.
	lines := strings.Split(usage, "\n")
	for i, line := range lines {
		// Skip trailing newline.
		if i == len(lines)-1 && len(line) == 0 {
			break
		}

		pre := "//"
		if len(line) > 0 {
			pre = "// "
		}
		_, err = fmt.Fprintln(f, pre+line)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(f, "package main")
	return err
}
