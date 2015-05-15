
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* ws */
			shift(4),		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: CSS */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S4
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: CSS */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Rules */
			nil,		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			shift(23),		/* > */
			nil,		/* ; */
			nil,		/* | */
			shift(24),		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(25),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(8),		/* ws, reduce: Selector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(9),		/* ws, reduce: Selector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(10),		/* ws, reduce: SimpleSelector */
			nil,		/* error */
			shift(13),		/* # */
			nil,		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(11),		/* ws, reduce: SimpleSelector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(14),		/* ws, reduce: AttrSelectors */
			nil,		/* error */
			shift(13),		/* # */
			nil,		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(28),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(18),		/* ws, reduce: NamedSelector */
			nil,		/* error */
			reduce(18),		/* #, reduce: NamedSelector */
			nil,		/* id */
			reduce(18),		/* :, reduce: NamedSelector */
			reduce(18),		/* ., reduce: NamedSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(29),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(30),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(19),		/* ws, reduce: NamedSelector */
			nil,		/* error */
			reduce(19),		/* #, reduce: NamedSelector */
			nil,		/* id */
			reduce(19),		/* :, reduce: NamedSelector */
			reduce(19),		/* ., reduce: NamedSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(20),		/* ws, reduce: NamedSelector */
			nil,		/* error */
			reduce(20),		/* #, reduce: NamedSelector */
			nil,		/* id */
			reduce(20),		/* :, reduce: NamedSelector */
			reduce(20),		/* ., reduce: NamedSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: CSS */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Rules */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(31),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(33),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(36),		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(36),		/* ws, reduce: Gt */
			nil,		/* error */
			reduce(36),		/* #, reduce: Gt */
			reduce(36),		/* id, reduce: Gt */
			reduce(36),		/* :, reduce: Gt */
			reduce(36),		/* ., reduce: Gt */
			reduce(36),		/* *, reduce: Gt */
			reduce(36),		/* body, reduce: Gt */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(42),		/* ws, reduce: CBOpen */
			nil,		/* error */
			nil,		/* # */
			reduce(42),		/* id, reduce: CBOpen */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			reduce(7),		/* >, reduce: SelectorWS */
			nil,		/* ; */
			nil,		/* | */
			reduce(7),		/* {, reduce: SelectorWS */
			nil,		/* } */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(12),		/* ws, reduce: SimpleSelector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(13),		/* ws, reduce: AttrSelectors */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(15),		/* ws, reduce: AttrSelector */
			nil,		/* error */
			reduce(15),		/* #, reduce: AttrSelector */
			nil,		/* id */
			reduce(15),		/* :, reduce: AttrSelector */
			reduce(15),		/* ., reduce: AttrSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(16),		/* ws, reduce: AttrSelector */
			nil,		/* error */
			reduce(16),		/* #, reduce: AttrSelector */
			nil,		/* id */
			reduce(16),		/* :, reduce: AttrSelector */
			reduce(16),		/* ., reduce: AttrSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(17),		/* ws, reduce: AttrSelector */
			nil,		/* error */
			reduce(17),		/* #, reduce: AttrSelector */
			nil,		/* id */
			reduce(17),		/* :, reduce: AttrSelector */
			reduce(17),		/* ., reduce: AttrSelector */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* ws, reduce: CBOpen */
			nil,		/* error */
			nil,		/* # */
			reduce(41),		/* id, reduce: CBOpen */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			shift(40),		/* } */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(41),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			shift(42),		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(33),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(30),		/* }, reduce: Properties */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(44),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(45),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			shift(47),		/* int */
			shift(48),		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(35),		/* ws, reduce: Gt */
			nil,		/* error */
			reduce(35),		/* #, reduce: Gt */
			reduce(35),		/* id, reduce: Gt */
			reduce(35),		/* :, reduce: Gt */
			reduce(35),		/* ., reduce: Gt */
			reduce(35),		/* *, reduce: Gt */
			reduce(35),		/* body, reduce: Gt */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(21),		/* ws, reduce: ChildSelector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(22),		/* ws, reduce: ChildSelector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Rule */
			shift(50),		/* ws */
			nil,		/* error */
			reduce(6),		/* #, reduce: Rule */
			reduce(6),		/* id, reduce: Rule */
			reduce(6),		/* :, reduce: Rule */
			reduce(6),		/* ., reduce: Rule */
			reduce(6),		/* *, reduce: Rule */
			reduce(6),		/* body, reduce: Rule */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: CBClose */
			reduce(44),		/* ws, reduce: CBClose */
			nil,		/* error */
			reduce(44),		/* #, reduce: CBClose */
			reduce(44),		/* id, reduce: CBClose */
			reduce(44),		/* :, reduce: CBClose */
			reduce(44),		/* ., reduce: CBClose */
			reduce(44),		/* *, reduce: CBClose */
			reduce(44),		/* body, reduce: CBClose */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			shift(51),		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(33),		/* ws, reduce: Propertyname */
			nil,		/* error */
			nil,		/* # */
			reduce(33),		/* id, reduce: Propertyname */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(33),		/* int, reduce: Propertyname */
			reduce(33),		/* float, reduce: Propertyname */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(29),		/* }, reduce: Properties */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(31),		/* ws, reduce: Propertyname */
			nil,		/* error */
			nil,		/* # */
			reduce(31),		/* id, reduce: Propertyname */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(31),		/* int, reduce: Propertyname */
			reduce(31),		/* float, reduce: Propertyname */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(24),		/* ws, reduce: Value */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(24),		/* ;, reduce: Value */
			reduce(24),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(28),		/* ;, reduce: Values */
			shift(54),		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(25),		/* ws, reduce: Value */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(25),		/* ;, reduce: Value */
			reduce(25),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(26),		/* ws, reduce: Value */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(26),		/* ;, reduce: Value */
			reduce(26),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			shift(56),		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(43),		/* $, reduce: CBClose */
			reduce(43),		/* ws, reduce: CBClose */
			nil,		/* error */
			reduce(43),		/* #, reduce: CBClose */
			reduce(43),		/* id, reduce: CBClose */
			reduce(43),		/* :, reduce: CBClose */
			reduce(43),		/* ., reduce: CBClose */
			reduce(43),		/* *, reduce: CBClose */
			reduce(43),		/* body, reduce: CBClose */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(32),		/* ws, reduce: Propertyname */
			nil,		/* error */
			nil,		/* # */
			reduce(32),		/* id, reduce: Propertyname */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(32),		/* int, reduce: Propertyname */
			reduce(32),		/* float, reduce: Propertyname */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(23),		/* ws, reduce: Value */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(23),		/* ;, reduce: Value */
			reduce(23),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(57),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(45),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			shift(47),		/* int */
			shift(48),		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* ws, reduce: Pipe */
			nil,		/* error */
			nil,		/* # */
			reduce(40),		/* id, reduce: Pipe */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(40),		/* int, reduce: Pipe */
			reduce(40),		/* float, reduce: Pipe */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(59),		/* ws */
			nil,		/* error */
			nil,		/* # */
			reduce(34),		/* id, reduce: Property */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(34),		/* }, reduce: Property */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(38),		/* ws, reduce: Semicolon */
			nil,		/* error */
			nil,		/* # */
			reduce(38),		/* id, reduce: Semicolon */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(38),		/* }, reduce: Semicolon */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(39),		/* ws, reduce: Pipe */
			nil,		/* error */
			nil,		/* # */
			reduce(39),		/* id, reduce: Pipe */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(39),		/* int, reduce: Pipe */
			reduce(39),		/* float, reduce: Pipe */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			reduce(27),		/* ;, reduce: Values */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(37),		/* ws, reduce: Semicolon */
			nil,		/* error */
			nil,		/* # */
			reduce(37),		/* id, reduce: Semicolon */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(37),		/* }, reduce: Semicolon */
			
		},

	},
	
}

