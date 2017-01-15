package transcript

import (
	"bytes"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	const testData = `
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
`

	sections, meta, err := Parse(strings.NewReader(testData))
	if err != nil {
		t.Fatal(err)
	}

	if len(sections) != 3 {
		t.Fatalf("parsing failed, expected 3 sections, got %d", len(sections))
	}

	if len(meta) != 2 {
		t.Fatalf("parsing failed, expected meta to have length 2, got %d", len(meta))
	}

	if meta["meta"] != "data" {
		t.Fatalf(`parsing failed, expected key meta to have value "data", got %#v`, meta["meta"])
	}

	if val, ok := meta["format"].(map[interface{}]interface{}); ok {
		if len(val) != 1 {
			t.Fatalf("parsing failed, expected meta format key to have length 2, got %d", len(val))
		}

		if val["yaml"] != true {
			t.Fatalf("parsing failed, expected key meta format:yaml to have value true, got %#v", val["yaml"])
		}
	} else {
		t.Fatalf("parsing failed, expected key format to be map[interface{}]interface{}, got %T", meta["format"])
	}

	if !bytes.Equal(sections[0], []byte{0xde, 0xad, 0xbe, 0xef}) {
		t.Fatalf("parsing failed, expected section #0 to be deadbeef, got %x", sections[0])
	}

	if !bytes.Equal(sections[1], []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}) {
		t.Fatalf("parsing failed, expected section #1 to be 0001020304050607, got %x", sections[1])
	}

	if !bytes.Equal(sections[2], []byte{0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}) {
		t.Fatalf("parsing failed, expected section #1 to be 08090a0b0c0d0e0f10, got %x", sections[2])
	}
}
