package parquet_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/segmentio/parquet-go"
	"github.com/segmentio/parquet-go/deprecated"
	"github.com/segmentio/parquet-go/encoding/plain"
	"github.com/segmentio/parquet-go/internal/bits"
)

func TestPage(t *testing.T) {
	t.Run("BOOLEAN", testPageBoolean)
	t.Run("INT32", testPageInt32)
	t.Run("INT64", testPageInt64)
	t.Run("INT96", testPageInt96)
	t.Run("FLOAT", testPageFloat)
	t.Run("DOUBLE", testPageDouble)
	t.Run("BYTE_ARRAY", testPageByteArray)
	t.Run("FIXED_LEN_BYTE_ARRAY", testPageFixedLenByteArray)
}

func testPageBoolean(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value bool }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []bool{false, true}
				n, err := w.(io.Writer).Write(bits.BoolToBytes(values))
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]bool, 2)
				n, err := r.(io.Reader).Read(bits.BoolToBytes(values))
				return values[:n], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []bool{false, true}
				n, err := w.(parquet.BooleanWriter).WriteBooleans(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]bool, 2)
				n, err := r.(parquet.BooleanReader).ReadBooleans(values)
				return values[:n], err
			},
		})
	})
}

func testPageInt32(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value int32 }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(io.Writer).Write(bits.Int32ToBytes(values))
				return values[:n/4], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]int32, 10)
				n, err := r.(io.Reader).Read(bits.Int32ToBytes(values))
				return values[:n/4], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(parquet.Int32Writer).WriteInt32s(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]int32, 10)
				n, err := r.(parquet.Int32Reader).ReadInt32s(values)
				return values[:n], err
			},
		})
	})
}

func testPageInt64(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value int64 }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(io.Writer).Write(bits.Int64ToBytes(values))
				return values[:n/8], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]int64, 10)
				n, err := r.(io.Reader).Read(bits.Int64ToBytes(values))
				return values[:n/8], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(parquet.Int64Writer).WriteInt64s(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]int64, 10)
				n, err := r.(parquet.Int64Reader).ReadInt64s(values)
				return values[:n], err
			},
		})
	})
}

func testPageInt96(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value deprecated.Int96 }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []deprecated.Int96{{0: 0}, {0: 1}, {0: 2}}
				n, err := w.(io.Writer).Write(deprecated.Int96ToBytes(values))
				return values[:n/12], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]deprecated.Int96, 3)
				n, err := r.(io.Reader).Read(deprecated.Int96ToBytes(values))
				return values[:n/12], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []deprecated.Int96{{0: 0}, {0: 1}, {0: 2}}
				n, err := w.(parquet.Int96Writer).WriteInt96s(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]deprecated.Int96, 3)
				n, err := r.(parquet.Int96Reader).ReadInt96s(values)
				return values[:n], err
			},
		})
	})
}

func testPageFloat(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value float32 }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(io.Writer).Write(bits.Float32ToBytes(values))
				return values[:n/4], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]float32, 10)
				n, err := r.(io.Reader).Read(bits.Float32ToBytes(values))
				return values[:n/4], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(parquet.FloatWriter).WriteFloats(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]float32, 10)
				n, err := r.(parquet.FloatReader).ReadFloats(values)
				return values[:n], err
			},
		})
	})
}

func testPageDouble(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value float64 }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(io.Writer).Write(bits.Float64ToBytes(values))
				return values[:n/8], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]float64, 10)
				n, err := r.(io.Reader).Read(bits.Float64ToBytes(values))
				return values[:n/8], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				n, err := w.(parquet.DoubleWriter).WriteDoubles(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]float64, 10)
				n, err := r.(parquet.DoubleReader).ReadDoubles(values)
				return values[:n], err
			},
		})
	})
}

func testPageByteArray(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value []byte }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []byte{}
				values = plain.AppendByteArray(values, []byte("A"))
				values = plain.AppendByteArray(values, []byte("B"))
				values = plain.AppendByteArray(values, []byte("C"))
				n, err := w.(io.Writer).Write(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]byte, 3+3*plain.ByteArrayLengthSize)
				n, err := r.(io.Reader).Read(values)
				return values[:n], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []byte{}
				values = plain.AppendByteArray(values, []byte("A"))
				values = plain.AppendByteArray(values, []byte("B"))
				values = plain.AppendByteArray(values, []byte("C"))
				_, err := w.(parquet.ByteArrayWriter).WriteByteArrays(values)
				return values, err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]byte, 3+3*plain.ByteArrayLengthSize)
				n, err := r.(parquet.ByteArrayReader).ReadByteArrays(values)
				return values[:n+n*plain.ByteArrayLengthSize], err
			},
		})
	})
}

func testPageFixedLenByteArray(t *testing.T) {
	schema := parquet.SchemaOf(struct{ Value [3]byte }{})

	t.Run("io", func(t *testing.T) {
		testBufferPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []byte("123456789")
				n, err := w.(io.Writer).Write(values)
				return values[:n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]byte, 3*3)
				n, err := r.(io.Reader).Read(values)
				return values[:n], err
			},
		})
	})

	t.Run("parquet", func(t *testing.T) {
		testPage(t, schema, pageTest{
			write: func(w parquet.ValueWriter) (interface{}, error) {
				values := []byte("123456789")
				n, err := w.(parquet.FixedLenByteArrayWriter).WriteFixedLenByteArrays(values)
				return values[:3*n], err
			},

			read: func(r parquet.ValueReader) (interface{}, error) {
				values := make([]byte, 3*3)
				n, err := r.(parquet.FixedLenByteArrayReader).ReadFixedLenByteArrays(values)
				return values[:3*n], err
			},
		})
	})
}

type pageTest struct {
	write func(parquet.ValueWriter) (interface{}, error)
	read  func(parquet.ValueReader) (interface{}, error)
}

func testPage(t *testing.T, schema *parquet.Schema, test pageTest) {
	t.Run("buffer", func(t *testing.T) { testBufferPage(t, schema, test) })
	t.Run("file", func(t *testing.T) { testFilePage(t, schema, test) })
}

func testBufferPage(t *testing.T, schema *parquet.Schema, test pageTest) {
	buffer := parquet.NewBuffer(schema)
	column := buffer.Column(0).(parquet.ColumnBuffer)

	w, err := test.write(column)
	if err != nil {
		t.Fatal("writing page values:", err)
	}

	r, err := test.read(column.Page().Values())
	if err != io.EOF {
		t.Errorf("expected io.EOF after reading all values but got %v", err)
	}
	if !reflect.DeepEqual(w, r) {
		t.Errorf("wrong values read from the page: got=%+v want=%+v", r, w)
	}
}

func testFilePage(t *testing.T, schema *parquet.Schema, test pageTest) {
	buffer := parquet.NewBuffer(schema)
	column := buffer.Column(0).(parquet.ColumnBuffer)

	w, err := test.write(column)
	if err != nil {
		t.Fatal("writing page values:", err)
	}

	output := new(bytes.Buffer)
	writer := parquet.NewWriter(output)
	n, err := writer.WriteRowGroup(buffer)
	if err != nil {
		t.Fatal("writing parquet file:", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal("writing parquet file:", err)
	}
	if n != buffer.NumRows() {
		t.Fatalf("number of rows written mismatch: got=%d want=%d", n, buffer.NumRows())
	}

	reader := bytes.NewReader(output.Bytes())
	f, err := parquet.OpenFile(reader, reader.Size())
	if err != nil {
		t.Fatal("opening parquet file:", err)
	}

	p, err := f.RowGroup(0).Column(0).Pages().ReadPage()
	if err != nil {
		t.Fatal("reading parquet page:", err)
	}

	values := p.Values()
	r, err := test.read(values)
	if err != io.EOF && err != nil {
		t.Errorf("expected io.EOF after reading all values but got %v", err)
	}
	if !reflect.DeepEqual(w, r) {
		t.Errorf("wrong values read from the page: got=%+v want=%+v", r, w)
	}
	if r, err := test.read(values); reflect.ValueOf(r).Len() != 0 || err != io.EOF {
		t.Errorf("expected no data and io.EOF after reading all values but got %d and %v", r, err)
	}
}

type testStruct struct {
	Value *string
}

func TestOptionalPageTrailingNulls(t *testing.T) {
	schema := parquet.SchemaOf(&testStruct{})
	buffer := parquet.NewBuffer(schema)

	str := "test"
	rows := []testStruct{{
		Value: nil,
	}, {
		Value: &str,
	}, {
		Value: nil,
	}}

	for _, row := range rows {
		if err := buffer.WriteRow(schema.Deconstruct(nil, row)); err != nil {
			t.Fatal("writing row:", err)
		}
	}

	resultRows := []parquet.Row{}
	reader := buffer.Rows()
	for {
		row, err := reader.ReadRow(nil)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal("reading rows:", err)
		}
		resultRows = append(resultRows, row)
	}

	if len(resultRows) != len(rows) {
		t.Errorf("wrong number of rows read: got=%d want=%d", len(resultRows), len(rows))
	}
}

func TestOptionalPagePreserveIndex(t *testing.T) {
	schema := parquet.SchemaOf(&testStruct{})
	buffer := parquet.NewBuffer(schema)

	if err := buffer.WriteRow(schema.Deconstruct(nil, &testStruct{Value: nil})); err != nil {
		t.Fatal("writing row:", err)
	}

	row, err := buffer.Rows().ReadRow(nil)
	if err != nil {
		t.Fatal("reading rows:", err)
	}

	if row[0].Column() != 0 {
		t.Errorf("wrong index: got=%d want=%d", row[0].Column(), 0)
	}
}

func TestRepeatedPageTrailingNulls(t *testing.T) {
	type testStruct struct {
		A []string `parquet:"a"`
	}

	s := parquet.SchemaOf(&testStruct{})

	records := []*testStruct{
		{A: nil},
		{A: []string{"test"}},
		{A: nil},
	}

	buf := parquet.NewBuffer(s)
	for _, rec := range records {
		row := s.Deconstruct(nil, rec)
		err := buf.WriteRow(row)
		if err != nil {
			t.Fatal(err)
		}
	}

	resultRows := []parquet.Row{}
	reader := buf.Rows()
	for {
		row, err := reader.ReadRow(nil)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal("reading rows:", err)
		}
		resultRows = append(resultRows, row)
	}

	if len(resultRows) != len(records) {
		t.Errorf("wrong number of rows read: got=%d want=%d", len(resultRows), len(records))
	}
}
