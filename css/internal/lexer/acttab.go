
package lexer

import(
	"fmt"
	"github.com/boombuler/termui/css/internal/token"
)

type ActionTable [NumStates] ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
 	ActionRow{ // S0
		Accept: 10,
 		Ignore: "",
 	},
 	ActionRow{ // S1
		Accept: 2,
 		Ignore: "",
 	},
 	ActionRow{ // S2
		Accept: 4,
 		Ignore: "",
 	},
 	ActionRow{ // S3
		Accept: 8,
 		Ignore: "",
 	},
 	ActionRow{ // S4
		Accept: 12,
 		Ignore: "",
 	},
 	ActionRow{ // S5
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S6
		Accept: 7,
 		Ignore: "",
 	},
 	ActionRow{ // S7
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S8
		Accept: 10,
 		Ignore: "",
 	},
 	ActionRow{ // S9
		Accept: 6,
 		Ignore: "",
 	},
 	ActionRow{ // S10
		Accept: 14,
 		Ignore: "",
 	},
 	ActionRow{ // S11
		Accept: 13,
 		Ignore: "",
 	},
 	ActionRow{ // S12
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S13
		Accept: 16,
 		Ignore: "",
 	},
 	ActionRow{ // S14
		Accept: 15,
 		Ignore: "",
 	},
 	ActionRow{ // S15
		Accept: 17,
 		Ignore: "",
 	},
 	ActionRow{ // S16
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S17
		Accept: 11,
 		Ignore: "",
 	},
 	ActionRow{ // S18
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S19
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S20
		Accept: 10,
 		Ignore: "",
 	},
 	ActionRow{ // S21
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S22
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S23
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S24
		Accept: -1,
 		Ignore: "!blockComment",
 	},
 	ActionRow{ // S25
		Accept: 9,
 		Ignore: "",
 	},
 		
}
