//go:build debug
// +build debug

package xgb

const debug = true

func debugf(format string, args ...any) { logf(msg, args...) }
