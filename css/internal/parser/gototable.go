
/*
*/
package parser

const numNTSymbols = 15
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // CSS
		2, // Rules
		4, // Rule
		5, // Selector
		6, // SimpleSelector
		12, // FixedSelector
		9, // NameSelector
		10, // ClassSelector
		11, // IdSelector
		7, // ChildSelector
		8, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // CSS
		18, // Rules
		4, // Rule
		5, // Selector
		6, // SimpleSelector
		12, // FixedSelector
		9, // NameSelector
		10, // ClassSelector
		11, // IdSelector
		7, // ChildSelector
		8, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		20, // SimpleSelector
		12, // FixedSelector
		9, // NameSelector
		10, // ClassSelector
		11, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		24, // Properties
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		26, // SimpleSelector
		12, // FixedSelector
		9, // NameSelector
		10, // ClassSelector
		11, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		30, // Value
		33, // Values
		-1, // Properties
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		30, // Value
		36, // Values
		-1, // Properties
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		37, // Properties
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // CSS
		-1, // Rules
		-1, // Rule
		-1, // Selector
		-1, // SimpleSelector
		-1, // FixedSelector
		-1, // NameSelector
		-1, // ClassSelector
		-1, // IdSelector
		-1, // ChildSelector
		-1, // ChildAnySelector
		-1, // Value
		-1, // Values
		-1, // Properties
		

	},
	
}
