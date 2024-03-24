//go:build !goexperiment.staticlockranking

package chantype

// See: /src/runtime/lockrank_off.go
type lockRankStruct struct{}
