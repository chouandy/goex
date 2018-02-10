package jsonex

import "github.com/mailru/easyjson/jlexer"

// UnmarshalArrayJSONBytes unmarshal array JSON
func UnmarshalArrayJSONBytes(data []byte, out *[][]byte) error {
	r := jlexer.Lexer{Data: data}
	unmarshalArrayJSONBytes(&r, out)
	return r.Error()
}

func unmarshalArrayJSONBytes(in *jlexer.Lexer, out *[][]byte) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('[')
	for !in.IsDelim(']') {
		item := in.UnsafeBytes()
		*out = append(*out, item)
		in.WantComma()
	}
	if isTopLevel {
		in.Consumed()
	}
}
