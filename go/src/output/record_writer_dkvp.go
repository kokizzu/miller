package output

import (
	"bytes"
	"io"

	"mlr/src/cli"
	"mlr/src/colorizer"
	"mlr/src/types"
)

type RecordWriterDKVP struct {
	writerOptions *cli.TWriterOptions
}

func NewRecordWriterDKVP(writerOptions *cli.TWriterOptions) *RecordWriterDKVP {
	return &RecordWriterDKVP{
		writerOptions: writerOptions,
	}
}

func (writer *RecordWriterDKVP) Write(
	outrec *types.Mlrmap,
	ostream io.WriteCloser,
	outputIsStdout bool,
) {
	// End of record stream: nothing special for this output format
	if outrec == nil {
		return
	}

	if outrec.IsEmpty() {
		ostream.Write([]byte(writer.writerOptions.ORS))
		return
	}

	var buffer bytes.Buffer // 5x faster than fmt.Print() separately
	for pe := outrec.Head; pe != nil; pe = pe.Next {
		buffer.WriteString(colorizer.MaybeColorizeKey(pe.Key, outputIsStdout))
		buffer.WriteString(writer.writerOptions.OPS)
		buffer.WriteString(colorizer.MaybeColorizeValue(pe.Value.String(), outputIsStdout))
		if pe.Next != nil {
			buffer.WriteString(writer.writerOptions.OFS)
		}
	}
	buffer.WriteString(writer.writerOptions.ORS)
	ostream.Write(buffer.Bytes())
}
