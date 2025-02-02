<!---  PLEASE DO NOT EDIT DIRECTLY. EDIT THE .md.in FILE PLEASE. --->
<div>
<span class="quicklinks">
Quick links:
&nbsp;
<a class="quicklink" href="../reference-main-flag-list/index.html">Flag list</a>
&nbsp;
<a class="quicklink" href="../reference-verbs/index.html">Verb list</a>
&nbsp;
<a class="quicklink" href="../reference-dsl-builtin-functions/index.html">Function list</a>
&nbsp;
<a class="quicklink" href="../glossary/index.html">Glossary</a>
&nbsp;
<a class="quicklink" href="https://github.com/johnkerl/miller" target="_blank">Repository ↗</a>
</span>
</div>
# In-place mode

Use the `mlr -I` flag to process files in-place. For example, `mlr -I --csv cut -x -f unwanted_column_name mydata/*.csv` will remove `unwanted_column_name` from all your `*.csv` files in your `mydata/` subdirectory.

By default, Miller output goes to the screen (or you can redirect a file using `>` or to another process using `|`). With `-I`, for each file name on the command line, output is written to a temporary file in the same directory. Miller writes its output into that temp file, which is then renamed over the original.  Then, processing continues on the next file. Each file is processed in isolation: if the output format is CSV, CSV headers will be present in each output file; statistics are only over each file's own records; and so on.

Since this replaces your data with modified data, it's often a good idea to back up your original files somewhere
first, to protect against keystroking errors.

Situations in which the input can't be updated in place:

* If the input file is a URL of the form `http://...`, `https://...`, or `file://...`.
* If a [`--prepipe` or `--prepipex` flag](reference-main-compressed-data.md#external-decompressors-on-input) is being used.
* If [in-place compression](reference-main-compressed-data.md) is being used and the format is BZIP2. For technical reasons, this can't be recompressed in place. (GZIP and ZLIB, however, are recompressable in place).

Additional note: `gzip` supports various compression levels, from 1 to 9. If you do `mlr -I ... yourfile.gz` then Miller will produce compressed output using GZIP, but, it makes no attempt to determine, or mimic, the original compression level of the input.

Please see [Choices for printing to files](10min.md#choices-for-printing-to-files) for examples.
