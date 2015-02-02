// Package asm provides convenience functions for decoding the RISC dialect
// described in risc/op.
package asm

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/mewmew/playground/archive/cs/risc/op"
)

// Decode decodes and returns the instructions read from r.
func Decode(r io.Reader) (insts []interface{}, err error) {
	var buf uint16
	for {
		err = binary.Read(r, binary.BigEndian, &buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("asm.Decode: read failure; %s", err)
		}
		inst, err := op.Decode(buf)
		if err != nil {
			return nil, err
		}
		insts = append(insts, inst)
	}

	return insts, nil
}

// DecodeSlice decodes and returns the instructions of the byte slice.
func DecodeSlice(p []byte) (insts []interface{}, err error) {
	if len(p)%op.InstSize != 0 {
		return nil, fmt.Errorf("asm.DecodeSlice: p len (%d) not evenly dividable by %d", len(p), op.InstSize)
	}

	insts = make([]interface{}, 0, len(p)/op.InstSize)
	for i := 0; i < len(p); i += op.InstSize {
		buf := binary.BigEndian.Uint16(p[i:])
		inst, err := op.Decode(buf)
		if err != nil {
			return nil, err
		}
		insts = append(insts, inst)
	}

	return insts, nil
}

// EncodeSlice encodes and returns the instructions as a byte slice.
func EncodeSlice(insts []interface{}) (p []byte, err error) {
	p = make([]byte, len(insts)*op.InstSize)
	for _, inst := range insts {
		buf, err := op.Encode(inst)
		if err != nil {
			return nil, err
		}
		binary.BigEndian.PutUint16(p[:op.InstSize], buf)
		p = p[op.InstSize:]
	}
	return p, nil
}
