package feather

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"encoding/binary"

	"github.com/lppgo/feather/fbs"
	"golang.org/x/exp/mmap"
)

const alignment = 8

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func paddedlength(x int64) int64 {
	return ((x + alignment - 1) / alignment) * alignment
}

// getoutputlength gets the length of the output part of a variable encoded
// size buffer -- depending on version of Feather file. Version >= 2 has 8 byte
// alignment, whereas earlier versions don't force alignment.
func getoutputlength(version int32, x int64) int64 {
	if version < 2 {
		return x
	}
	return paddedlength(x)
}

// DictEncoding is used to encode categorical variables...
type DictEncoding struct {
	refs, pool interface{}
}

// Column holds all information necessary to construct slices of values from Feather formatted columns
type Column struct {
	Name     string
	Values   *fbs.PrimitiveArray
	Meta     FeatherColumn
	UserMeta string
	TypE     int8
}

// Source holds the source for the feather file
type Source struct {
	Data    *mmap.ReaderAt
	Ctable  *fbs.CTable
	Columns []FeatherColumn
}

func (src Source) String() string {
	return fmt.Sprintf("Feather file with %d rows and %d columns", src.NumRows(), src.NumCols())
}

// NumRows returns the number of rows in the feather file
func (src *Source) NumRows() int {
	return int(src.Ctable.NumRows())
}

// NumCols returns the number of columns in the feather file
func (src *Source) NumCols() int {
	return src.Ctable.ColumnsLength()
}

// Read reads a feather file, parses metadata, and returns a Source
func Read(fn string) (*Source, error) {
	src := new(Source)
	file, err := mmap.Open(fn)
	if err != nil {
		return new(Source), err
	}
	src.Data = file
	length := int64(file.Len())

	magic := make([]byte, 4)

	file.ReadAt(magic, 0)
	if string(magic) != "FEA1" {
		return nil, errors.New("didn't find magic bytes in header")
	}

	file.ReadAt(magic, length-4)
	if string(magic) != "FEA1" {
		return nil, errors.New("didn't find magic bytes in footer")
	}

	metadataSizeBuf := make([]byte, 4)
	file.ReadAt(metadataSizeBuf, length-8)
	var metadataSize int32
	binary.Read(bytes.NewBuffer(metadataSizeBuf), binary.LittleEndian, &metadataSize)

	ctableBuf := make([]byte, metadataSize)
	file.ReadAt(ctableBuf, length-8-int64(metadataSize))
	ctable := fbs.GetRootAsCTable(ctableBuf, 0)
	src.Ctable = ctable

	numColumns := ctable.ColumnsLength()
	cols := make([]FeatherColumn, numColumns)

	wg := &sync.WaitGroup{}
	for ix := 0; ix < numColumns; ix++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			col := new(fbs.Column)
			ctable.Columns(col, ii)
			// TODO how does NewFullColumnCol fit in here? Should we check
			// the number of nulls and choose for the user??
			cols[ii] = NewColumnFbsColumn(src, col)
		}(ix)
	}
	wg.Wait()
	src.Columns = cols

	return src, nil
}

func (src *Source) getoutputlength(x int64) int64 {
	return getoutputlength(src.Ctable.Version(), max(x, 1))
}

func (src *Source) getBitmaskLength() int64 {
	nrow := int64(src.NumRows())
	out := nrow / 8
	if nrow%8 != 0 {
		out++
	}
	return src.getoutputlength(out)
}
