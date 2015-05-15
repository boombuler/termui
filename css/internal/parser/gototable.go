
/*
*/
package parser

const numNTSymbols = 21
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // CSS
		3, // Rules
		5, // Rule
		6, // SelectorWS
		7, // Selector
		8, // SimpleSelector
		11, // AttrSelectors
		12, // AttrSelector
		10, // NamedSelector
		9, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // CSS
		19, // Rules
		5, // Rule
		6, // SelectorWS
		7, // Selector
		8, // SimpleSelector
		11, // AttrSelectors
		12, // AttrSelector
		10, // NamedSelector
		9, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // CSS
		20, // Rules
		5, // Rule
		6, // SelectorWS
		7, // Selector
		8, // SimpleSelector
		11, // AttrSelectors
		12, // AttrSelector
		10, // NamedSelector
		9, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		22, // Gt
		-1, // Semicolon
		-1, // Pipe
		21, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		26, // AttrSelectors
		12, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		27, // AttrSelectors
		12, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		32, // Properties
		35, // Propertyname
		34, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		37, // SimpleSelector
		11, // AttrSelectors
		12, // AttrSelector
		10, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		38, // SimpleSelector
		11, // AttrSelectors
		12, // AttrSelector
		10, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		39, // CBClose
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		43, // Properties
		35, // Propertyname
		34, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		46, // Value
		49, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S46
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		53, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S47
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S48
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S49
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		55, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S50
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S51
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S52
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S53
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		46, // Value
		58, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S54
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S55
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S56
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S57
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S58
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	gotoRow{ // S59
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // SelectorWS
		-1, // Selector
		-1, // SimpleSelector
		-1, // AttrSelectors
		-1, // AttrSelector
		-1, // NamedSelector
		-1, // ChildSelector
		-1, // Value
		-1, // Values
		-1, // Properties
		-1, // Propertyname
		-1, // Property
		-1, // Gt
		-1, // Semicolon
		-1, // Pipe
		-1, // CBOpen
		-1, // CBClose
		

	},
	
}
