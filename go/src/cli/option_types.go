// ================================================================
// Items which might better belong in miller/cli, but which are placed in a
// deeper package to avoid a package-dependency cycle between miller/cli and
// miller/transforming.
// ================================================================

package cli

import (
	"regexp"

	"mlr/src/lib"
)

// ----------------------------------------------------------------
//typedef struct _generator_opts_t {
//	char* field_name;
//	// xxx to do: convert to mv_t
//	long long start;
//	long long stop;
//	long long step;
//} generator_opts_t;

type TCommentHandling int

const (
	CommentsAreData TCommentHandling = iota
	SkipComments
	PassComments
)
const DEFAULT_COMMENT_STRING = "#"

// ----------------------------------------------------------------
type TReaderOptions struct {
	InputFileFormat string
	IFS             string
	IPS             string
	IRS             string
	AllowRepeatIFS  bool
	AllowRepeatIPS  bool
	IFSRegex        *regexp.Regexp
	IPSRegex        *regexp.Regexp

	// If unspecified on the command line, these take input-format-dependent
	// defaults.  E.g. default FS is comma for DKVP but space for NIDX;
	// default AllowRepeatIFS is false for CSV but true for PPRINT.
	IFSWasSpecified            bool
	IPSWasSpecified            bool
	IRSWasSpecified            bool
	AllowRepeatIFSWasSpecified bool
	AllowRepeatIPSWasSpecified bool

	UseImplicitCSVHeader bool
	AllowRaggedCSVInput  bool

	CommentHandling TCommentHandling
	CommentString   string

	//	// Fake internal-data-generator 'reader'
	//	generator_opts_t generator_opts;

	// For out-of-process handling of compressed data, via popen
	Prepipe string
	// For most things like gunzip we do 'gunzip < filename | mlr ...' if
	// filename is present, else 'gunzip | mlr ...' if reading from stdin.
	// However some commands like 'unzip -qc' are weird so this option lets
	// people give the command and we won't insert the '<'.
	PrepipeIsRaw bool
	// For in-process gunzip/bunzip2/zcat (distinct from prepipe)
	FileInputEncoding lib.TFileInputEncoding
}

// ----------------------------------------------------------------
type TWriterOptions struct {
	OutputFileFormat string
	ORS              string
	OFS              string
	OPS              string
	FLATSEP          string

	// If unspecified on the command line, these take input-format-dependent
	// defaults.  E.g. default FS is comma for DKVP but space for NIDX.
	OFSWasSpecified bool
	OPSWasSpecified bool
	ORSWasSpecified bool

	HeaderlessCSVOutput      bool
	BarredPprintOutput       bool
	RightAlignedPprintOutput bool

	//	right_justify_xtab_value bool;
	//	right_align_pprint bool;
	WrapJSONOutputInOuterList bool
	JSONOutputMultiline       bool // Not using miller/types enum to avoid package cycle
	//	json_quote_int_keys bool;
	//	json_quote_non_string_values bool;
	//
	//	quoting_t oquoting;

	// TODO: centralize comments here & mlrcli_parse.go & repl/verbs.go; maybe to a README.md
	// When we read things like
	//
	//   x:a=1,x:b=2
	//
	// which is how we write out nested data structures for non-nested formats
	// (all but JSON), the default behavior is to unflatten them back to
	//
	//   {"x": {"a": 1}, {"b": 2}}
	//
	// unless the user explicitly asks to suppress that.
	AutoUnflatten bool

	// The default behavior is to flatten nested data structures like
	//
	//   {"x": {"a": 1}, {"b": 2}}
	//
	// down to
	//
	//   x:a=1,x:b=2
	//
	// which is how we write out nested data structures for non-nested formats
	// (all but JSON) -- unless the user explicitly asks to suppress that.
	AutoFlatten bool

	// For floating-point numbers: "" means use the Go default.
	FPOFMT string
}

// ----------------------------------------------------------------
type TOptions struct {
	ReaderOptions TReaderOptions
	WriterOptions TWriterOptions

	// Data files to be operated on: e.g. given 'mlr cat foo.dat bar.dat', this
	// is ["foo.dat", "bar.dat"].
	FileNames []string

	// DSL files to be loaded for every put/filter operation -- like 'put -f'
	// or 'filter -f' but specified up front on the command line, suitable for
	// .mlrrc. Use-case is someone has DSL functions they always want to be
	// defined.
	//
	// Risk of CVE if this is in .mlrrc so --load and --mload are explicitly
	// denied in the .mlrrc reader.
	DSLPreloadFileNames []string

	NRProgressMod int
	DoInPlace     bool // mlr -I
	NoInput       bool // mlr -n

	HaveRandSeed bool
	RandSeed     int
}

// ----------------------------------------------------------------
func DefaultOptions() TOptions {
	return TOptions{
		ReaderOptions: DefaultReaderOptions(),
		WriterOptions: DefaultWriterOptions(),

		FileNames:           make([]string, 0),
		DSLPreloadFileNames: make([]string, 0),
		NoInput:             false,
	}
}

func DefaultReaderOptions() TReaderOptions {
	return TReaderOptions{
		InputFileFormat:   "dkvp", // xxx constify at top
		IRS:               "\n",
		IFS:               ",",
		IPS:               "=",
		CommentHandling:   CommentsAreData,
		FileInputEncoding: lib.FileInputEncodingDefault,
	}
}

func DefaultWriterOptions() TWriterOptions {
	return TWriterOptions{
		OutputFileFormat: "dkvp",
		ORS:              "\n",
		OFS:              ",",
		OPS:              "=",
		FLATSEP:          ".",

		HeaderlessCSVOutput:       false,
		WrapJSONOutputInOuterList: false,
		JSONOutputMultiline:       true,
		AutoUnflatten:             true,
		AutoFlatten:               true,
		FPOFMT:                    "",
	}
}
