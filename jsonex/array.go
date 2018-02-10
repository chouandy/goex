package jsonex

import "github.com/mailru/easyjson/jlexer"

// UnmarshalArrayJSON unmarshal array JSON to []string
func UnmarshalArrayJSON(data []byte, out *[]string) error {
	r := jlexer.Lexer{Data: data}
	unmarshalArrayJSON(&r, out)
	return r.Error()
}

func unmarshalArrayJSON(in *jlexer.Lexer, out *[]string) {
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
		*out = append(*out, string(item))
		in.WantComma()
	}
	if isTopLevel {
		in.Consumed()
	}
}

// UnmarshalArrayJSONBytes unmarshal array JSON to [][]byte
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
