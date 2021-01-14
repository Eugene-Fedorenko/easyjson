package tests

import (
	"github.com/eugene-fedorenko/easyjson"
	"github.com/eugene-fedorenko/easyjson/jlexer"
	"github.com/eugene-fedorenko/easyjson/jwriter"
)

//easyjson:json
type NestedMarshaler struct {
	Value easyjson.MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}

