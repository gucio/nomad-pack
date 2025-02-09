// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flag

import (
	"os"
	"strconv"

	"github.com/posener/complete"
)

// -- Float64Var and float64Value
type Float64Var struct {
	Name       string
	Aliases    []string
	Usage      string
	Default    float64
	Hidden     bool
	EnvVar     string
	Target     *float64
	Completion complete.Predictor
}

type Float64VarP struct {
	*Float64Var
	Shorthand string
}

func (f *Set) Float64Var(i *Float64Var) {
	f.Float64VarP(&Float64VarP{
		Float64Var: i,
		Shorthand:  "",
	})
}

func (f *Set) Float64VarP(i *Float64VarP) {
	initial := i.Default
	if v, exist := os.LookupEnv(i.EnvVar); exist {
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			initial = i
		}
	}

	def := ""
	if i.Default != 0 {
		def = strconv.FormatFloat(i.Default, 'e', -1, 64)
	}

	f.VarFlagP(&VarFlagP{
		VarFlag: &VarFlag{
			Name:       i.Name,
			Aliases:    i.Aliases,
			Usage:      i.Usage,
			Default:    def,
			EnvVar:     i.EnvVar,
			Value:      newFloat64Value(initial, i.Target, i.Hidden),
			Completion: i.Completion,
		},
		Shorthand: "",
	})
}

type float64Value struct {
	hidden bool
	target *float64
}

func newFloat64Value(def float64, target *float64, hidden bool) *float64Value {
	*target = def
	return &float64Value{
		hidden: hidden,
		target: target,
	}
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f.target = v
	return nil
}

func (f *float64Value) Get() interface{} { return *f.target }
func (f *float64Value) String() string   { return strconv.FormatFloat(*f.target, 'g', -1, 64) }
func (f *float64Value) Example() string  { return "float" }
func (f *float64Value) Hidden() bool     { return f.hidden }
func (f *float64Value) Type() string     { return "float64" }
