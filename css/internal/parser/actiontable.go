
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
			shift(3),		/* error */
			nil,		/* { */
			nil,		/* } */
			shift(13),		/* * */
			shift(14),		/* body */
			shift(15),		/* id */
			shift(16),		/* . */
			shift(17),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: CSS */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S3
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: CSS */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Rules */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			shift(13),		/* * */
			shift(14),		/* body */
			shift(15),		/* id */
			shift(16),		/* . */
			shift(17),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			shift(19),		/* { */
			nil,		/* } */
			shift(13),		/* * */
			shift(14),		/* body */
			shift(15),		/* id */
			shift(16),		/* . */
			shift(17),		/* # */
			shift(21),		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(6),		/* {, reduce: Selector */
			nil,		/* } */
			reduce(6),		/* *, reduce: Selector */
			reduce(6),		/* body, reduce: Selector */
			reduce(6),		/* id, reduce: Selector */
			reduce(6),		/* ., reduce: Selector */
			reduce(6),		/* #, reduce: Selector */
			reduce(6),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(7),		/* {, reduce: Selector */
			nil,		/* } */
			reduce(7),		/* *, reduce: Selector */
			reduce(7),		/* body, reduce: Selector */
			reduce(7),		/* id, reduce: Selector */
			reduce(7),		/* ., reduce: Selector */
			reduce(7),		/* #, reduce: Selector */
			reduce(7),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(8),		/* {, reduce: Selector */
			nil,		/* } */
			reduce(8),		/* *, reduce: Selector */
			reduce(8),		/* body, reduce: Selector */
			reduce(8),		/* id, reduce: Selector */
			reduce(8),		/* ., reduce: Selector */
			reduce(8),		/* #, reduce: Selector */
			reduce(8),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(9),		/* {, reduce: SimpleSelector */
			nil,		/* } */
			reduce(9),		/* *, reduce: SimpleSelector */
			reduce(9),		/* body, reduce: SimpleSelector */
			reduce(9),		/* id, reduce: SimpleSelector */
			reduce(9),		/* ., reduce: SimpleSelector */
			reduce(9),		/* #, reduce: SimpleSelector */
			reduce(9),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(10),		/* {, reduce: SimpleSelector */
			nil,		/* } */
			reduce(10),		/* *, reduce: SimpleSelector */
			reduce(10),		/* body, reduce: SimpleSelector */
			reduce(10),		/* id, reduce: SimpleSelector */
			reduce(10),		/* ., reduce: SimpleSelector */
			reduce(10),		/* #, reduce: SimpleSelector */
			reduce(10),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(11),		/* {, reduce: SimpleSelector */
			nil,		/* } */
			reduce(11),		/* *, reduce: SimpleSelector */
			reduce(11),		/* body, reduce: SimpleSelector */
			reduce(11),		/* id, reduce: SimpleSelector */
			reduce(11),		/* ., reduce: SimpleSelector */
			reduce(11),		/* #, reduce: SimpleSelector */
			reduce(11),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(12),		/* {, reduce: SimpleSelector */
			nil,		/* } */
			reduce(12),		/* *, reduce: SimpleSelector */
			reduce(12),		/* body, reduce: SimpleSelector */
			reduce(12),		/* id, reduce: SimpleSelector */
			reduce(12),		/* ., reduce: SimpleSelector */
			reduce(12),		/* #, reduce: SimpleSelector */
			reduce(12),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(13),		/* {, reduce: FixedSelector */
			nil,		/* } */
			reduce(13),		/* *, reduce: FixedSelector */
			reduce(13),		/* body, reduce: FixedSelector */
			reduce(13),		/* id, reduce: FixedSelector */
			reduce(13),		/* ., reduce: FixedSelector */
			reduce(13),		/* #, reduce: FixedSelector */
			reduce(13),		/* >, reduce: FixedSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(14),		/* {, reduce: FixedSelector */
			nil,		/* } */
			reduce(14),		/* *, reduce: FixedSelector */
			reduce(14),		/* body, reduce: FixedSelector */
			reduce(14),		/* id, reduce: FixedSelector */
			reduce(14),		/* ., reduce: FixedSelector */
			reduce(14),		/* #, reduce: FixedSelector */
			reduce(14),		/* >, reduce: FixedSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(15),		/* {, reduce: NameSelector */
			nil,		/* } */
			reduce(15),		/* *, reduce: NameSelector */
			reduce(15),		/* body, reduce: NameSelector */
			reduce(15),		/* id, reduce: NameSelector */
			reduce(15),		/* ., reduce: NameSelector */
			reduce(15),		/* #, reduce: NameSelector */
			reduce(15),		/* >, reduce: NameSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(22),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(23),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Rules */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(25),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(19),		/* {, reduce: ChildAnySelector */
			nil,		/* } */
			reduce(19),		/* *, reduce: ChildAnySelector */
			reduce(19),		/* body, reduce: ChildAnySelector */
			reduce(19),		/* id, reduce: ChildAnySelector */
			reduce(19),		/* ., reduce: ChildAnySelector */
			reduce(19),		/* #, reduce: ChildAnySelector */
			reduce(19),		/* >, reduce: ChildAnySelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			shift(13),		/* * */
			shift(14),		/* body */
			shift(15),		/* id */
			shift(16),		/* . */
			shift(17),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(16),		/* {, reduce: ClassSelector */
			nil,		/* } */
			reduce(16),		/* *, reduce: ClassSelector */
			reduce(16),		/* body, reduce: ClassSelector */
			reduce(16),		/* id, reduce: ClassSelector */
			reduce(16),		/* ., reduce: ClassSelector */
			reduce(16),		/* #, reduce: ClassSelector */
			reduce(16),		/* >, reduce: ClassSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(17),		/* {, reduce: IdSelector */
			nil,		/* } */
			reduce(17),		/* *, reduce: IdSelector */
			reduce(17),		/* body, reduce: IdSelector */
			reduce(17),		/* id, reduce: IdSelector */
			reduce(17),		/* ., reduce: IdSelector */
			reduce(17),		/* #, reduce: IdSelector */
			reduce(17),		/* >, reduce: IdSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			shift(27),		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			shift(28),		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(18),		/* {, reduce: ChildSelector */
			nil,		/* } */
			reduce(18),		/* *, reduce: ChildSelector */
			reduce(18),		/* body, reduce: ChildSelector */
			reduce(18),		/* id, reduce: ChildSelector */
			reduce(18),		/* ., reduce: ChildSelector */
			reduce(18),		/* #, reduce: ChildSelector */
			reduce(18),		/* >, reduce: ChildSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Rule */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			reduce(5),		/* *, reduce: Rule */
			reduce(5),		/* body, reduce: Rule */
			reduce(5),		/* id, reduce: Rule */
			reduce(5),		/* ., reduce: Rule */
			reduce(5),		/* #, reduce: Rule */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(29),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			shift(31),		/* int */
			shift(32),		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(20),		/* |, reduce: Value */
			nil,		/* : */
			reduce(20),		/* ;, reduce: Value */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			shift(34),		/* | */
			nil,		/* : */
			reduce(24),		/* ;, reduce: Values */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(21),		/* |, reduce: Value */
			nil,		/* : */
			reduce(21),		/* ;, reduce: Value */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(22),		/* |, reduce: Value */
			nil,		/* : */
			reduce(22),		/* ;, reduce: Value */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			shift(35),		/* ; */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(29),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			shift(31),		/* int */
			shift(32),		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			reduce(26),		/* }, reduce: Properties */
			nil,		/* * */
			nil,		/* body */
			shift(25),		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			reduce(23),		/* ;, reduce: Values */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			reduce(25),		/* }, reduce: Properties */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* : */
			nil,		/* ; */
			
		},

	},
	
}

