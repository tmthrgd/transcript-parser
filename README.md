# transcript-parser

[![GoDoc](https://godoc.org/github.com/tmthrgd/transcript-parser?status.svg)](https://godoc.org/github.com/tmthrgd/transcript-parser)
[![Build Status](https://travis-ci.org/tmthrgd/transcript-parser.svg?branch=master)](https://travis-ci.org/tmthrgd/transcript-parser)

## Format

```
; this is a comment
# meta: data
# format:
#  yaml: true
de ad be ef ; another comment
---
00 01 02 03
04 05 06 07    ; this is also a comment
---
08 09
0a 0b 0c
0d 0e 0f 10  ; this is valid
; as is this
```

## License

Unless otherwise noted, the transcript-parser source files are distributed under the Modified BSD License found in the LICENSE file.
