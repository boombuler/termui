
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			shift(24),		/* + */
			shift(25),		/* > */
			nil,		/* ; */
			nil,		/* | */
			shift(26),		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(27),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			shift(30),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
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
			shift(31),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			shift(32),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			shift(33),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(35),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			shift(38),		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			shift(40),		/* ws */
			nil,		/* error */
			shift(13),		/* # */
			shift(14),		/* id */
			shift(15),		/* : */
			shift(16),		/* . */
			shift(17),		/* * */
			shift(18),		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			reduce(37),		/* ws, reduce: Plus */
			nil,		/* error */
			reduce(37),		/* #, reduce: Plus */
			reduce(37),		/* id, reduce: Plus */
			reduce(37),		/* :, reduce: Plus */
			reduce(37),		/* ., reduce: Plus */
			reduce(37),		/* *, reduce: Plus */
			reduce(37),		/* body, reduce: Plus */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			reduce(39),		/* ws, reduce: Gt */
			nil,		/* error */
			reduce(39),		/* #, reduce: Gt */
			reduce(39),		/* id, reduce: Gt */
			reduce(39),		/* :, reduce: Gt */
			reduce(39),		/* ., reduce: Gt */
			reduce(39),		/* *, reduce: Gt */
			reduce(39),		/* body, reduce: Gt */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(45),		/* ws, reduce: CBOpen */
			nil,		/* error */
			nil,		/* # */
			reduce(45),		/* id, reduce: CBOpen */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			reduce(7),		/* +, reduce: SelectorWS */
			reduce(7),		/* >, reduce: SelectorWS */
			nil,		/* ; */
			nil,		/* | */
			reduce(7),		/* {, reduce: SelectorWS */
			nil,		/* } */
			
		},

	},
	actionRow{ // S28
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(44),		/* ws, reduce: CBOpen */
			nil,		/* error */
			nil,		/* # */
			reduce(44),		/* id, reduce: CBOpen */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			shift(44),		/* } */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(45),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			shift(46),		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(35),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(31),		/* }, reduce: Properties */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(48),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(49),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			shift(51),		/* int */
			shift(52),		/* float */
			nil,		/* + */
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
			reduce(38),		/* ws, reduce: Gt */
			nil,		/* error */
			reduce(38),		/* #, reduce: Gt */
			reduce(38),		/* id, reduce: Gt */
			reduce(38),		/* :, reduce: Gt */
			reduce(38),		/* ., reduce: Gt */
			reduce(38),		/* *, reduce: Gt */
			reduce(38),		/* body, reduce: Gt */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
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
			nil,		/* $ */
			reduce(36),		/* ws, reduce: Plus */
			nil,		/* error */
			reduce(36),		/* #, reduce: Plus */
			reduce(36),		/* id, reduce: Plus */
			reduce(36),		/* :, reduce: Plus */
			reduce(36),		/* ., reduce: Plus */
			reduce(36),		/* *, reduce: Plus */
			reduce(36),		/* body, reduce: Plus */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			reduce(23),		/* ws, reduce: ChildSelector */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
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
			reduce(6),		/* $, reduce: Rule */
			shift(54),		/* ws */
			nil,		/* error */
			reduce(6),		/* #, reduce: Rule */
			reduce(6),		/* id, reduce: Rule */
			reduce(6),		/* :, reduce: Rule */
			reduce(6),		/* ., reduce: Rule */
			reduce(6),		/* *, reduce: Rule */
			reduce(6),		/* body, reduce: Rule */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(47),		/* $, reduce: CBClose */
			reduce(47),		/* ws, reduce: CBClose */
			nil,		/* error */
			reduce(47),		/* #, reduce: CBClose */
			reduce(47),		/* id, reduce: CBClose */
			reduce(47),		/* :, reduce: CBClose */
			reduce(47),		/* ., reduce: CBClose */
			reduce(47),		/* *, reduce: CBClose */
			reduce(47),		/* body, reduce: CBClose */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			shift(55),		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(34),		/* ws, reduce: Propertyname */
			nil,		/* error */
			nil,		/* # */
			reduce(34),		/* id, reduce: Propertyname */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(34),		/* int, reduce: Propertyname */
			reduce(34),		/* float, reduce: Propertyname */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S47
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
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(30),		/* }, reduce: Properties */
			
		},

	},
	actionRow{ // S48
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
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S49
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
			nil,		/* + */
			nil,		/* > */
			reduce(25),		/* ;, reduce: Value */
			reduce(25),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(56),		/* ws */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			reduce(29),		/* ;, reduce: Values */
			shift(58),		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S51
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
			nil,		/* + */
			nil,		/* > */
			reduce(26),		/* ;, reduce: Value */
			reduce(26),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(27),		/* ws, reduce: Value */
			nil,		/* error */
			nil,		/* # */
			nil,		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			reduce(27),		/* ;, reduce: Value */
			reduce(27),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S53
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
			nil,		/* + */
			nil,		/* > */
			shift(60),		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(46),		/* $, reduce: CBClose */
			reduce(46),		/* ws, reduce: CBClose */
			nil,		/* error */
			reduce(46),		/* #, reduce: CBClose */
			reduce(46),		/* id, reduce: CBClose */
			reduce(46),		/* :, reduce: CBClose */
			reduce(46),		/* ., reduce: CBClose */
			reduce(46),		/* *, reduce: CBClose */
			reduce(46),		/* body, reduce: CBClose */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
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
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S56
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
			nil,		/* + */
			nil,		/* > */
			reduce(24),		/* ;, reduce: Value */
			reduce(24),		/* |, reduce: Value */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(61),		/* ws */
			nil,		/* error */
			nil,		/* # */
			shift(49),		/* id */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			shift(51),		/* int */
			shift(52),		/* float */
			nil,		/* + */
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
			reduce(43),		/* ws, reduce: Pipe */
			nil,		/* error */
			nil,		/* # */
			reduce(43),		/* id, reduce: Pipe */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(43),		/* int, reduce: Pipe */
			reduce(43),		/* float, reduce: Pipe */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
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
			shift(63),		/* ws */
			nil,		/* error */
			nil,		/* # */
			reduce(35),		/* id, reduce: Property */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(35),		/* }, reduce: Property */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* ws, reduce: Semicolon */
			nil,		/* error */
			nil,		/* # */
			reduce(41),		/* id, reduce: Semicolon */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(41),		/* }, reduce: Semicolon */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(42),		/* ws, reduce: Pipe */
			nil,		/* error */
			nil,		/* # */
			reduce(42),		/* id, reduce: Pipe */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			reduce(42),		/* int, reduce: Pipe */
			reduce(42),		/* float, reduce: Pipe */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S62
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
			nil,		/* + */
			nil,		/* > */
			reduce(28),		/* ;, reduce: Values */
			nil,		/* | */
			nil,		/* { */
			nil,		/* } */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* ws, reduce: Semicolon */
			nil,		/* error */
			nil,		/* # */
			reduce(40),		/* id, reduce: Semicolon */
			nil,		/* : */
			nil,		/* . */
			nil,		/* * */
			nil,		/* body */
			nil,		/* int */
			nil,		/* float */
			nil,		/* + */
			nil,		/* > */
			nil,		/* ; */
			nil,		/* | */
			nil,		/* { */
			reduce(40),		/* }, reduce: Semicolon */
			
		},

	},
	
}

