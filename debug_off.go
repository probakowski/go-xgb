//go:build !debug
// +build !debug

package xgb

const debug = false

func debugf(format string, args ...any) {}
