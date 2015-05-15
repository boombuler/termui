
package parser

import css "github.com/boombuler/termui/css"
import internal "github.com/boombuler/termui/css/internal"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : CSS	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CSS : Rules	<< X[0], nil >>`,
		Id: "CSS",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CSS : error	<<  >>`,
		Id: "CSS",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rule Rules	<< append(X[1].(internal.Rules), X[0].(internal.Rule)), nil >>`,
		Id: "Rules",
		NTType: 2,
		Index: 3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[1].(internal.Rules), X[0].(internal.Rule)), nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< internal.Rules{ X[0].(internal.Rule) }, nil >>`,
		Id: "Rules",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.Rules{ X[0].(internal.Rule) }, nil
		},
	},
	ProdTabEntry{
		String: `Rule : Selector "{" Properties "}"	<< internal.Rule{ X[0].(css.Selector), X[2].(internal.PropertyValues)}, nil >>`,
		Id: "Rule",
		NTType: 3,
		Index: 5,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.Rule{ X[0].(css.Selector), X[2].(internal.PropertyValues)}, nil
		},
	},
	ProdTabEntry{
		String: `Selector : SimpleSelector	<< X[0], nil >>`,
		Id: "Selector",
		NTType: 4,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Selector : ChildSelector	<<  >>`,
		Id: "Selector",
		NTType: 4,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Selector : ChildAnySelector	<<  >>`,
		Id: "Selector",
		NTType: 4,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : NameSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 5,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : ClassSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 5,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : PClassSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 5,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : IdSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : FixedSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 5,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FixedSelector : "*"	<< css.AnySelector, nil >>`,
		Id: "FixedSelector",
		NTType: 6,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.AnySelector, nil
		},
	},
	ProdTabEntry{
		String: `FixedSelector : "body"	<< css.BodySelector, nil >>`,
		Id: "FixedSelector",
		NTType: 6,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.BodySelector, nil
		},
	},
	ProdTabEntry{
		String: `NameSelector : id	<< css.NameSelector(str(X[0])), nil >>`,
		Id: "NameSelector",
		NTType: 7,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.NameSelector(str(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `ClassSelector : "." id	<< css.ClassSelector(str(X[1])), nil >>`,
		Id: "ClassSelector",
		NTType: 8,
		Index: 17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.ClassSelector(str(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `PClassSelector : ":" id	<< css.PseudoClassSelector(str(X[1])), nil >>`,
		Id: "PClassSelector",
		NTType: 9,
		Index: 18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.PseudoClassSelector(str(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `IdSelector : "#" id	<< css.IdSelector(str(X[0])), nil >>`,
		Id: "IdSelector",
		NTType: 10,
		Index: 19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.IdSelector(str(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `ChildSelector : Selector ">" SimpleSelector	<< css.ParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil >>`,
		Id: "ChildSelector",
		NTType: 11,
		Index: 20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.ParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil
		},
	},
	ProdTabEntry{
		String: `ChildAnySelector : Selector SimpleSelector	<<  >>`,
		Id: "ChildAnySelector",
		NTType: 12,
		Index: 21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : id	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 13,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Value : int	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 13,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Value : float	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 13,
		Index: 24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Values : Value "|" Values	<< append(X[2].([]string), X[0].(string)), nil >>`,
		Id: "Values",
		NTType: 14,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[2].([]string), X[0].(string)), nil
		},
	},
	ProdTabEntry{
		String: `Values : Value	<< []string{X[0].(string)}, nil >>`,
		Id: "Values",
		NTType: 14,
		Index: 26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []string{X[0].(string)}, nil
		},
	},
	ProdTabEntry{
		String: `Properties : id ":" Values ";" Properties	<< append(X[4].(internal.PropertyValues), internal.PropertyValue { str(X[0]), X[2] }), nil >>`,
		Id: "Properties",
		NTType: 15,
		Index: 27,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[4].(internal.PropertyValues), internal.PropertyValue { str(X[0]), X[2] }), nil
		},
	},
	ProdTabEntry{
		String: `Properties : id ":" Values ";"	<< internal.PropertyValues{ internal.PropertyValue { str(X[0]), X[2] } }, nil >>`,
		Id: "Properties",
		NTType: 15,
		Index: 28,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.PropertyValues{ internal.PropertyValue { str(X[0]), X[2] } }, nil
		},
	},
	
}
