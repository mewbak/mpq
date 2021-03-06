package mpq

import (
	"encoding/binary"
	"io"
)

// HiBlockTable is a strange legacy thing that allows the
// BlockTable to reference offsets greater than 16 bits.
// Although this code is present and it exists I doubt it works
// be careful :D
type HiBlockTable struct {
	Table []uint16
}

func (m *MPQ) readHiBlockTable(r io.Reader) error {
	h := &HiBlockTable{}

	offset := 0
	buffer := make([]byte, m.Header.BlockTableSize*2)
	if _, err := r.Read(buffer); err != nil {
		return err
	}

	h.Table = make([]uint16, m.Header.BlockTableSize)

	for i := 0; i < m.Header.BlockTableSize; i++ {
		h.Table[i] = binary.LittleEndian.Uint16(buffer[offset : offset+2])
		offset += 2
	}

	m.HiBlockTable = h
	return nil
}
