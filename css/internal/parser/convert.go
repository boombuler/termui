package parser

import (
	"github.com/boombuler/termui/css/internal/token"
)

func str(at Attrib) string {
	tkn, ok := at.(*token.Token)
	if !ok {
		return ""
	}
	return string(tkn.Lit)
}
