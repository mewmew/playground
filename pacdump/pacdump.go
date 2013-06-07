package main

import "flag"
import "fmt"
import "log"
import "os"
import "os/exec"
import "strings"

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: pacdump PKG...")
	fmt.Fprintf(os.Stderr, "Create an archive (%q) containing the files of the specified Arch Linux packages.\n", output)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, `  Create an archive ("boll.tar.gz") containing the files of the mesa package.`)
	fmt.Fprintln(os.Stderr, "    pacdump mesa")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	var paths Paths
	for _, pkg := range flag.Args() {
		err := paths.GetFromPkg(pkg)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err := paths.Dump()
	if err != nil {
		log.Fatalln(err)
	}
}

type Paths []string

func (paths *Paths) GetFromPkg(pkg string) (err error) {
	cmd := exec.Command("pacman", "-Ql", pkg)
	buf, err := cmd.Output()
	if err != nil {
		return err
	}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if len(line) < len(pkg)+1 {
			continue
		}
		// ignore "pkg " prefix.
		line = line[len(pkg)+1:]

		if strings.HasSuffix(line, "/") {
			// skip folders.
			continue
		}
		*paths = append(*paths, line)
	}
	return nil
}

const output = "boll.tar.gz"

func (paths Paths) Dump() (err error) {
	if len(paths) < 1 {
		return nil
	}
	var args []string
	args = append(args, "-czvf", output)
	args = append(args, paths...)
	cmd := exec.Command("tar", args...)
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("created: %s (containing %d files).\n", output, len(paths))
	return nil
}
