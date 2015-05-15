
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
			shift(14),		/* * */
			shift(15),		/* body */
			shift(16),		/* id */
			shift(17),		/* . */
			shift(18),		/* : */
			shift(19),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			shift(14),		/* * */
			shift(15),		/* body */
			shift(16),		/* id */
			shift(17),		/* . */
			shift(18),		/* : */
			shift(19),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			shift(21),		/* { */
			nil,		/* } */
			shift(14),		/* * */
			shift(15),		/* body */
			shift(16),		/* id */
			shift(17),		/* . */
			shift(18),		/* : */
			shift(19),		/* # */
			shift(23),		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(6),		/* :, reduce: Selector */
			reduce(6),		/* #, reduce: Selector */
			reduce(6),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(7),		/* :, reduce: Selector */
			reduce(7),		/* #, reduce: Selector */
			reduce(7),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(8),		/* :, reduce: Selector */
			reduce(8),		/* #, reduce: Selector */
			reduce(8),		/* >, reduce: Selector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(9),		/* :, reduce: SimpleSelector */
			reduce(9),		/* #, reduce: SimpleSelector */
			reduce(9),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(10),		/* :, reduce: SimpleSelector */
			reduce(10),		/* #, reduce: SimpleSelector */
			reduce(10),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(11),		/* :, reduce: SimpleSelector */
			reduce(11),		/* #, reduce: SimpleSelector */
			reduce(11),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(12),		/* :, reduce: SimpleSelector */
			reduce(12),		/* #, reduce: SimpleSelector */
			reduce(12),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(13),		/* {, reduce: SimpleSelector */
			nil,		/* } */
			reduce(13),		/* *, reduce: SimpleSelector */
			reduce(13),		/* body, reduce: SimpleSelector */
			reduce(13),		/* id, reduce: SimpleSelector */
			reduce(13),		/* ., reduce: SimpleSelector */
			reduce(13),		/* :, reduce: SimpleSelector */
			reduce(13),		/* #, reduce: SimpleSelector */
			reduce(13),		/* >, reduce: SimpleSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			reduce(14),		/* :, reduce: FixedSelector */
			reduce(14),		/* #, reduce: FixedSelector */
			reduce(14),		/* >, reduce: FixedSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(15),		/* {, reduce: FixedSelector */
			nil,		/* } */
			reduce(15),		/* *, reduce: FixedSelector */
			reduce(15),		/* body, reduce: FixedSelector */
			reduce(15),		/* id, reduce: FixedSelector */
			reduce(15),		/* ., reduce: FixedSelector */
			reduce(15),		/* :, reduce: FixedSelector */
			reduce(15),		/* #, reduce: FixedSelector */
			reduce(15),		/* >, reduce: FixedSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(16),		/* {, reduce: NameSelector */
			nil,		/* } */
			reduce(16),		/* *, reduce: NameSelector */
			reduce(16),		/* body, reduce: NameSelector */
			reduce(16),		/* id, reduce: NameSelector */
			reduce(16),		/* ., reduce: NameSelector */
			reduce(16),		/* :, reduce: NameSelector */
			reduce(16),		/* #, reduce: NameSelector */
			reduce(16),		/* >, reduce: NameSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			shift(24),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S18
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			shift(26),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S20
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			nil,		/* * */
			nil,		/* body */
			shift(28),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(21),		/* {, reduce: ChildAnySelector */
			nil,		/* } */
			reduce(21),		/* *, reduce: ChildAnySelector */
			reduce(21),		/* body, reduce: ChildAnySelector */
			reduce(21),		/* id, reduce: ChildAnySelector */
			reduce(21),		/* ., reduce: ChildAnySelector */
			reduce(21),		/* :, reduce: ChildAnySelector */
			reduce(21),		/* #, reduce: ChildAnySelector */
			reduce(21),		/* >, reduce: ChildAnySelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			shift(14),		/* * */
			shift(15),		/* body */
			shift(16),		/* id */
			shift(17),		/* . */
			shift(18),		/* : */
			shift(19),		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(17),		/* {, reduce: ClassSelector */
			nil,		/* } */
			reduce(17),		/* *, reduce: ClassSelector */
			reduce(17),		/* body, reduce: ClassSelector */
			reduce(17),		/* id, reduce: ClassSelector */
			reduce(17),		/* ., reduce: ClassSelector */
			reduce(17),		/* :, reduce: ClassSelector */
			reduce(17),		/* #, reduce: ClassSelector */
			reduce(17),		/* >, reduce: ClassSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(18),		/* {, reduce: PClassSelector */
			nil,		/* } */
			reduce(18),		/* *, reduce: PClassSelector */
			reduce(18),		/* body, reduce: PClassSelector */
			reduce(18),		/* id, reduce: PClassSelector */
			reduce(18),		/* ., reduce: PClassSelector */
			reduce(18),		/* :, reduce: PClassSelector */
			reduce(18),		/* #, reduce: PClassSelector */
			reduce(18),		/* >, reduce: PClassSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(19),		/* {, reduce: IdSelector */
			nil,		/* } */
			reduce(19),		/* *, reduce: IdSelector */
			reduce(19),		/* body, reduce: IdSelector */
			reduce(19),		/* id, reduce: IdSelector */
			reduce(19),		/* ., reduce: IdSelector */
			reduce(19),		/* :, reduce: IdSelector */
			reduce(19),		/* #, reduce: IdSelector */
			reduce(19),		/* >, reduce: IdSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			shift(30),		/* } */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
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
			nil,		/* id */
			nil,		/* . */
			shift(31),		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			reduce(20),		/* {, reduce: ChildSelector */
			nil,		/* } */
			reduce(20),		/* *, reduce: ChildSelector */
			reduce(20),		/* body, reduce: ChildSelector */
			reduce(20),		/* id, reduce: ChildSelector */
			reduce(20),		/* ., reduce: ChildSelector */
			reduce(20),		/* :, reduce: ChildSelector */
			reduce(20),		/* #, reduce: ChildSelector */
			reduce(20),		/* >, reduce: ChildSelector */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S30
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
			reduce(5),		/* :, reduce: Rule */
			reduce(5),		/* #, reduce: Rule */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
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
			shift(32),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			shift(34),		/* int */
			shift(35),		/* float */
			nil,		/* | */
			nil,		/* ; */
			
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(22),		/* |, reduce: Value */
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			shift(37),		/* | */
			reduce(26),		/* ;, reduce: Values */
			
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
			nil,		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(23),		/* |, reduce: Value */
			reduce(23),		/* ;, reduce: Value */
			
		},

	},
	actionRow{ // S35
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			reduce(24),		/* |, reduce: Value */
			reduce(24),		/* ;, reduce: Value */
			
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			shift(38),		/* ; */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			nil,		/* } */
			nil,		/* * */
			nil,		/* body */
			shift(32),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			shift(34),		/* int */
			shift(35),		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			reduce(28),		/* }, reduce: Properties */
			nil,		/* * */
			nil,		/* body */
			shift(28),		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S39
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
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			reduce(25),		/* ;, reduce: Values */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* { */
			reduce(27),		/* }, reduce: Properties */
			nil,		/* * */
			nil,		/* body */
			nil,		/* id */
			nil,		/* . */
			nil,		/* : */
			nil,		/* # */
			nil,		/* > */
			nil,		/* int */
			nil,		/* float */
			nil,		/* | */
			nil,		/* ; */
			
		},

	},
	
}

