package urfave

import "github.com/urfave/cli/v3"

// bool

type BoolFlag struct {
	cli.BoolFlag
}

func Bool(f *cli.BoolFlag) *BoolFlag {
	return &BoolFlag{*f}
}

func (f *BoolFlag) V() bool {
	return f.Get().(bool) //nolint:forcetypeassert
}

// string

type StringFlag struct {
	cli.StringFlag
}

func String(f *cli.StringFlag) *StringFlag {
	return &StringFlag{*f}
}

func (f *StringFlag) V() string {
	return f.Get().(string) //nolint:forcetypeassert
}

// int

type IntFlag struct {
	cli.IntFlag
}

func Int(f *cli.IntFlag) *IntFlag {
	return &IntFlag{*f}
}

func (f *IntFlag) V() int {
	return f.Get().(int) //nolint:forcetypeassert
}

// string slice

type StringSliceFlag struct {
	cli.StringSliceFlag
}

func StringSlice(f *cli.StringSliceFlag) *StringSliceFlag {
	return &StringSliceFlag{*f}
}

func (f *StringSliceFlag) V() []string {
	return f.Get().([]string) //nolint:forcetypeassert
}
