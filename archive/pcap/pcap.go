// Package pcap provides support for reading pcap files.
package pcap

import (
	"encoding/binary"
	"io"
	"os"
	"time"
)

// A File represents a pcap file.
type File struct {
	hdr FileHeader
	io.ReadCloser
}

// A FileHeader structure is present at the beginning of each pcap file.
type FileHeader struct {
	// Magic number.
	Magic uint32
	// Major version number.
	MajorVer uint16
	// Minor version number.
	MinorVer uint16
	// GMT to local correction.
	ThisZone int32
	// Accuracy of timestamps.
	SigFigs uint32
	// Max length of captured packets.
	SnapLen uint32
	// Data link type.
	Network uint32
}

// Magic specifies if little or big endian encoding has been used.
const (
	MagicLittleEndian = 0xA1B2C3D4
	MagicBigEndian    = 0xD4C3B2A1
)

// Open opens the named pcap file for reading.
func Open(filePath string) (f *File, err error) {
	fr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	f = &File{
		ReadCloser: fr,
	}
	err = binary.Read(f, binary.LittleEndian, &f.hdr)
	if err != nil {
		fr.Close()
		return nil, err
	}
	return f, nil
}

// A Package represents a network packet.
type Package struct {
	Hdr PackageHeader
	Buf []byte
}

// A PackageHeader structure is present at the beginning of each packet stored
// in the pcap file.
type PackageHeader struct {
	// Timestamp seconds.
	Sec uint32
	// Timestamp microseconds.
	Usec uint32
	// Packet length saved in file.
	Len uint32
	// Original packet length.
	OrigLen uint32
}

// ReadPackage reads and returns the next packet in the pcap file.
func (f *File) ReadPackage() (pkg *Package, err error) {
	pkg = new(Package)
	err = binary.Read(f, binary.LittleEndian, &pkg.Hdr)
	if err != nil {
		return nil, err
	}
	pkg.Buf = make([]byte, pkg.Hdr.Len)
	_, err = f.Read(pkg.Buf)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

// Bytes returns the package's content as a byte slice.
func (pkg *Package) Bytes() (buf []byte) {
	return pkg.Buf
}

// Time returns the time when the package was sent.
func (pkg *Package) Time() (t time.Time) {
	return time.Unix(int64(pkg.Hdr.Sec), int64(pkg.Hdr.Usec*1000))
}
