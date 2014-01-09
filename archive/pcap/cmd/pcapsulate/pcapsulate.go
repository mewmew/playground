// pcapsulate encapsulates the provided files as packets in a pcap file.
//
//      Usage: pcapsulate [FILE]...
//      Encapsulate the provided files as packets in a pcap file.
//
//        -o="pcapsulate.pcap": Output path.
//
//      With no FILE, or when FILE is -, read standard input.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/mewmew/playground/archive/pcap"
)

// flagOutput corresponds to the output path.
var flagOutput string

func init() {
	flag.StringVar(&flagOutput, "o", "pcapsulate.pcap", "Output path.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: pcapsulate [FILE]...")
	fmt.Fprintln(os.Stderr, "Encapsulate the provided files as packets in a pcap file.")
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "With no FILE, or when FILE is -, read standard input.")
}

// StdinFileName is a reserved file name used for standard input.
const StdinFileName = "-"

func main() {
	flag.Parse()

	var filePaths []string
	if flag.NArg() == 0 {
		// Read from stdin when no FILE has been provided.
		filePaths = []string{StdinFileName}
	} else {
		filePaths = flag.Args()
	}

	err := pcapsulate(filePaths)
	if err != nil {
		log.Fatalln(err)
	}
}

// pcapsulate encapsulates the provided files as packets in a pcap file.
func pcapsulate(filePaths []string) (err error) {
	f, err := os.Create(flagOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write pcap header.
	fh := pcap.FileHeader{
		Magic:    pcap.MagicLittleEndian,
		MajorVer: 2,
		MinorVer: 4,
		ThisZone: 0,
		SigFigs:  0,
		SnapLen:  65535,
		Network:  1,
	}
	err = binary.Write(f, binary.LittleEndian, fh)
	if err != nil {
		return err
	}

	// Encapsulate each file as a packet and write them to the pcap file.
	var r io.Reader
	for _, filePath := range filePaths {
		if filePath == StdinFileName {
			r = os.Stdin
		} else {
			f, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer f.Close()
			r = f
		}

		// Generate and write the packet to the pcap file.
		pkg, err := genPacket(r)
		if err != nil {
			return err
		}
		err = binary.Write(f, binary.LittleEndian, pkg.Hdr)
		if err != nil {
			return err
		}
		_, err = f.Write(pkg.Buf)
		if err != nil {
			return err
		}
	}

	return nil
}

// genPacket generates a pcap packet by encapsulating the content read from r.
func genPacket(r io.Reader) (pkg *pcap.Package, err error) {
	pkg = new(pcap.Package)
	pkg.Buf, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	pkg.Hdr = pcap.PackageHeader{
		Sec:     0,
		Usec:    0,
		Len:     uint32(len(pkg.Buf)),
		OrigLen: uint32(len(pkg.Buf)),
	}
	return pkg, nil
}
