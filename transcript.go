package transcript

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

func fromHexChar(c byte) (b byte, skip bool, ok bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', false, true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, false, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, false, true
	case c == ' ' || c == '\t':
		return 0, true, false
	}

	return 0, false, false
}

func Parse(r io.Reader) (sections [][]byte, meta map[interface{}]interface{}, err error) {
	scanner := bufio.NewScanner(r)

	buf := new(bytes.Buffer)

	parts := make([][]byte, 0, 2)

	metadata := make(map[interface{}]interface{})
	var metaBuf bytes.Buffer

	for scanner.Scan() {
		data := scanner.Bytes()

		if i := bytes.IndexByte(data, ';'); i != -1 {
			data = data[:i]
		}

		if len(data) == 0 {
			continue
		}

		if bytes.HasPrefix(data, []byte("# ")) {
			if len(parts) != 0 || buf.Len() != 0 {
				err = errors.New("invalid format: metadata must preceed data")
				return
			}

			metaBuf.Write(data[2:])
			metaBuf.Write([]byte("\n"))
			continue
		}

		if metaBuf.Len() != 0 {
			if err = yaml.Unmarshal(metaBuf.Bytes(), &metadata); err != nil {
				return
			}

			metaBuf.Reset()
		}

		if bytes.Equal(data, []byte{'-', '-', '-'}) {
			parts, buf = append(parts, buf.Bytes()), new(bytes.Buffer)
			continue
		}

		for i := 0; i < len(data); i++ {
			a, skip, ok := fromHexChar(data[i])
			if skip {
				continue
			} else if !ok {
				err = fmt.Errorf("invalid format: expected hex or space, got %c", data[i])
				return
			}

			if i++; i == len(data) {
				err = errors.New("invalid format: expected hex, got EOF")
				return
			}

			b, _, ok := fromHexChar(data[i])
			if !ok {
				err = fmt.Errorf("invalid format: expected hex, got %c", data[i])
				return
			}

			buf.WriteByte((a << 4) | b)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return append(parts, buf.Bytes()), metadata, nil
}
