
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
		String: `CSS : ws Rules	<< X[1], nil >>`,
		Id: "CSS",
		NTType: 1,
		Index: 1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `CSS : Rules	<< X[0], nil >>`,
		Id: "CSS",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CSS : error	<<  >>`,
		Id: "CSS",
		NTType: 1,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rule Rules	<< append(X[1].(internal.Rules), X[0].(internal.Rule)), nil >>`,
		Id: "Rules",
		NTType: 2,
		Index: 4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[1].(internal.Rules), X[0].(internal.Rule)), nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< internal.Rules{X[0].(internal.Rule) }, nil >>`,
		Id: "Rules",
		NTType: 2,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.Rules{X[0].(internal.Rule) }, nil
		},
	},
	ProdTabEntry{
		String: `Rule : SelectorWS CBOpen Properties CBClose	<< internal.Rule{X[0].(css.Selector), X[2].(internal.PropertyValues)}, nil >>`,
		Id: "Rule",
		NTType: 3,
		Index: 6,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.Rule{X[0].(css.Selector), X[2].(internal.PropertyValues)}, nil
		},
	},
	ProdTabEntry{
		String: `SelectorWS : Selector ws	<< X[0], nil >>`,
		Id: "SelectorWS",
		NTType: 4,
		Index: 7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Selector : SimpleSelector	<< X[0], nil >>`,
		Id: "Selector",
		NTType: 5,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Selector : ChildSelector	<< X[0], nil >>`,
		Id: "Selector",
		NTType: 5,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : NamedSelector	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 6,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : AttrSelectors	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 6,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SimpleSelector : NamedSelector AttrSelectors	<< X[0], nil >>`,
		Id: "SimpleSelector",
		NTType: 6,
		Index: 12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AttrSelectors : AttrSelector AttrSelectors	<< append(X[1].(css.AndSelector), X[0].(css.Selector)), nil >>`,
		Id: "AttrSelectors",
		NTType: 7,
		Index: 13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[1].(css.AndSelector), X[0].(css.Selector)), nil
		},
	},
	ProdTabEntry{
		String: `AttrSelectors : AttrSelector	<< css.AndSelector{X[0].(css.Selector)}, nil >>`,
		Id: "AttrSelectors",
		NTType: 7,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.AndSelector{X[0].(css.Selector)}, nil
		},
	},
	ProdTabEntry{
		String: `AttrSelector : "#" id	<< css.IDSelector(str(X[1])), nil >>`,
		Id: "AttrSelector",
		NTType: 8,
		Index: 15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.IDSelector(str(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `AttrSelector : ":" id	<< css.PseudoClassSelector(str(X[1])), nil >>`,
		Id: "AttrSelector",
		NTType: 8,
		Index: 16,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.PseudoClassSelector(str(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `AttrSelector : "." id	<< css.ClassSelector(str(X[1])), nil >>`,
		Id: "AttrSelector",
		NTType: 8,
		Index: 17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.ClassSelector(str(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `NamedSelector : id	<< css.NameSelector(str(X[0])), nil >>`,
		Id: "NamedSelector",
		NTType: 9,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.NameSelector(str(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `NamedSelector : "*"	<< css.AnySelector, nil >>`,
		Id: "NamedSelector",
		NTType: 9,
		Index: 19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.AnySelector, nil
		},
	},
	ProdTabEntry{
		String: `NamedSelector : "body"	<< css.BodySelector, nil >>`,
		Id: "NamedSelector",
		NTType: 9,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.BodySelector, nil
		},
	},
	ProdTabEntry{
		String: `ChildSelector : SelectorWS Gt SimpleSelector	<< css.ParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil >>`,
		Id: "ChildSelector",
		NTType: 10,
		Index: 21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.ParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil
		},
	},
	ProdTabEntry{
		String: `ChildSelector : Selector ws SimpleSelector	<< css.AnyParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil >>`,
		Id: "ChildSelector",
		NTType: 10,
		Index: 22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.AnyParentSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil
		},
	},
	ProdTabEntry{
		String: `ChildSelector : SelectorWS Plus SimpleSelector	<< css.AfterSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil >>`,
		Id: "ChildSelector",
		NTType: 10,
		Index: 23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return css.AfterSelector{X[0].(css.Selector), X[2].(css.Selector)}, nil
		},
	},
	ProdTabEntry{
		String: `Value : Value ws	<< X[0], nil >>`,
		Id: "Value",
		NTType: 11,
		Index: 24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : id	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 11,
		Index: 25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Value : int	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 11,
		Index: 26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Value : float	<< str(X[0]), nil >>`,
		Id: "Value",
		NTType: 11,
		Index: 27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Values : Value Values	<< append([]string{X[0].(string)}, X[1].([]string)...), nil >>`,
		Id: "Values",
		NTType: 12,
		Index: 28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append([]string{X[0].(string)}, X[1].([]string)...), nil
		},
	},
	ProdTabEntry{
		String: `Values : Value	<< []string{X[0].(string)}, nil >>`,
		Id: "Values",
		NTType: 12,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []string{X[0].(string)}, nil
		},
	},
	ProdTabEntry{
		String: `Properties : Property Properties	<< append(X[1].(internal.PropertyValues), X[0].(internal.PropertyValue)), nil >>`,
		Id: "Properties",
		NTType: 13,
		Index: 30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[1].(internal.PropertyValues), X[0].(internal.PropertyValue)), nil
		},
	},
	ProdTabEntry{
		String: `Properties : Property	<< internal.PropertyValues{X[0].(internal.PropertyValue)}, nil >>`,
		Id: "Properties",
		NTType: 13,
		Index: 31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.PropertyValues{X[0].(internal.PropertyValue)}, nil
		},
	},
	ProdTabEntry{
		String: `Propertyname : Propertyname ws	<< X[0], nil >>`,
		Id: "Propertyname",
		NTType: 14,
		Index: 32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Propertyname : id ws ":"	<< str(X[0]), nil >>`,
		Id: "Propertyname",
		NTType: 14,
		Index: 33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Propertyname : id ":"	<< str(X[0]), nil >>`,
		Id: "Propertyname",
		NTType: 14,
		Index: 34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return str(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Property : Propertyname Values Semicolon	<< internal.PropertyValue { X[0].(string), X[1] }, nil >>`,
		Id: "Property",
		NTType: 15,
		Index: 35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return internal.PropertyValue { X[0].(string), X[1] }, nil
		},
	},
	ProdTabEntry{
		String: `Plus : Plus ws	<<  >>`,
		Id: "Plus",
		NTType: 16,
		Index: 36,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Plus : "+"	<<  >>`,
		Id: "Plus",
		NTType: 16,
		Index: 37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Gt : Gt ws	<<  >>`,
		Id: "Gt",
		NTType: 17,
		Index: 38,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Gt : ">"	<<  >>`,
		Id: "Gt",
		NTType: 17,
		Index: 39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Semicolon : Semicolon ws	<<  >>`,
		Id: "Semicolon",
		NTType: 18,
		Index: 40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Semicolon : ";"	<<  >>`,
		Id: "Semicolon",
		NTType: 18,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CBOpen : CBOpen ws	<<  >>`,
		Id: "CBOpen",
		NTType: 19,
		Index: 42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CBOpen : "{"	<<  >>`,
		Id: "CBOpen",
		NTType: 19,
		Index: 43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CBClose : CBClose ws	<<  >>`,
		Id: "CBClose",
		NTType: 20,
		Index: 44,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CBClose : "}"	<<  >>`,
		Id: "CBClose",
		NTType: 20,
		Index: 45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
