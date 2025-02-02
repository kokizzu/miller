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
# Streaming processing, and memory usage

## What does streaming mean?

When we say that Miller is **streaming**, we mean that most operations need
only a single record in memory at a time, rather than ingesting all input
before producing any output.

This is contrast to, say, the dataframes approach where you ingest all data,
wait for **end of file**, then start manipulating the data.

Both approaches have their advantages: the dataframes approach requires that
all data fit in **system memory** (which, as hardware gets larger over time, is
less and less of a constraint); the streaming approach requires that you
sometimes need to accumulate results on records (rows) **as they arrive**
rather than looping through them explicitly.

Since Miller takes the streaming approach when possible (see below for
exceptions), you can often operate on files which are larger than your system's
memory . It also means you can do `tail -f some-file | mlr --some-flags` and
Miller will operate on records as they arrive one at a time.  You don't have to
wait for and end-of-file marker (which never arrives with `tail-f`) to start
seeing partial results. This also means if you pipe Miller's output to other
streaming tools (like `cat`, `grep`, `sed`, and so on), they will also output
partial results as data arrives.

The statements in the [Miller programming language](miller-programming-language.md)
(outside of optional `begin`/`end` blocks which execute before and after all
records have been read, respectively) are implicit callbacks which are executed
once per record. For example, using `mlr --csv put '$z = $x + $y' myfile.csv`,
the statement `$z = $x + $y` will be executed 10,000 times if you `myfile.csv`
has 10,000 records.

If you do wish to accumulate all records into memory and loop over them
explicitly, you can do so -- see the page on [operating on all
records](operating-on-all-records.md).

## Streaming and non-streaming verbs

Most verbs, including [`cat`](reference-verbs.md#cat),
[`cut`](reference-verbs.md#cut), etc. operate on each record independently.
They have no state to retain from one record to the next, and are streaming.

For those operations which require deeper retention, Miller retains only as
much data as needed.  For example, the [`sort`](reference-verbs.md#sort) and
[`tac`](reference-verbs.md#tac) (stream-reverse, backward spelling of
[`cat`](reference-verbs.md#cat)) must ingest and retain all records in memory
before emitting any -- the last input record may well end up being the first
one to be emitted.

[`stats1`](reference-verbs.md#stats1) Other verbs, such as
[`tail`](reference-verbs.md#tail) and [`top`](reference-verbs.md#top), need to
retain only a fixed number of records -- 10, perhaps, even if the input data
has a million records.

Yet other verbs, such as [`stats1`](reference-verbs.md#stats1) and
[`stats2`](reference-verbs.md#stats2), retain only summary arithmetic on the
records they visit. These are memory-friendly: memory usage is bounded. However,
they only produce output at the end of the record stream.

## Fully streaming verbs

These don't retain any state from one record to the next.
They are memory-friendly, and they don't wait for end of input to produce their output.

* [altkv](#altkv)
* [bar](#bar) -- if not auto-mode
* [cat](#cat)
* [check](#check)
* [clean-whitespace](#clean-whitespace)
* [cut](#cut)
* [decimate](#decimate)
* [fill-down](#fill-down)
* [fill-empty](#fill-empty)
* [flatten](#flatten)
* [format-values](#format-values)
* [gap](#gap)
* [grep](#grep)
* [having-fields](#having-fields)
* [head](#head)
* [json-parse](#json-parse)
* [json-stringify](#json-stringify)
* [label](#label)
* [merge-fields](#merge-fields)
* [nest](#nest) -- if not `implode-values-across-records`
* [nothing](#nothing)
* [regularize](#regularize)
* [rename](#rename)
* [reorder](#reorder)
* [repeat](#repeat)
* [reshape](#reshape) -- if not long-to-wide
* [sec2gmt](#sec2gmt)
* [sec2gmtdate](#sec2gmtdate)
* [seqgen](#seqgen)
* [skip-trivial-records](#skip-trivial-records)
* [sort-within-records](#sort-within-records)
* [step](#step)
* [tee](#tee)
* [template](#template)
* [unflatten](#unflatten)
* [unsparsify](#unsparsify) if invoked with `-f`

## Non-streaming, retaining all records

These retain all records from one record to the next.
They are memory-unfriendly, and they wait for end of input to produce their output.

* [bar](#bar) -- if auto-mode
* [bootstrap](#bootstrap)
* [count-similar](#count-similar)
* [fraction](#fraction)
* [group-by](#group-by)
* [group-like](#group-like)
* [least-frequent](#least-frequent)
* [most-frequent](#most-frequent)
* [nest](#nest) -- if `implode-values-across-records`
* [remove-empty-columns](#remove-empty-columns)
* [reshape](#reshape) -- if long-to-wide
* [sample](#sample)
* [shuffle](#shuffle)
* [sort](#sort)
* [tac](#tac)
* [uniq](#uniq) -- if `mlr uniq -a -c`
* [unsparsify](#unsparsify) if invoked without `-f`

## Non-streaming, retaining some records

These retain a bounded number of records from one record to the next.
They are memory-friendly, but they wait for end of input to produce their output.

* [tail](#tail)
* [top](#top)

## Non-streaming, retaining some state

These retain an amount of state from one record to the next, but less than if
they were to retain all records in memory.  They are variably memory-friendly
-- depending on how many distinct values for the group-by keys exist in the
input data -- and they wait for end of input to produce their output.

* [count-distinct](#count-distinct)
* [count](#count)
* [histogram](#histogram)
* [stats1](#stats1) -- except `mlr stats1 -s` for incremental stats before end of stream
* [stats2](#stats2)
* [uniq](#uniq) -- if not `mlr uniq -a -c`

## Variable

Any `end` blocks you provide will not be executed until end of stream; otherwise these
don't want for end of stream. Similarly, if you write logic to retain all records
(see also the page on [operating on all records](operating-on-all-records.md.in))
these will be memory-unfriendly; otherwhise they are memory-friendly.

Most simple operations such as `mlr put '$z = $x + $y'` are fully streaming.

* [filter](#filter)
* [put](#put)

## Half-streaming

The main input files are streamed, but the join file (using `-f`) is loaded into memory at the start.
