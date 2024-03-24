//go:build goexperiment.staticlockranking

package chantype

type lockRank int

// See: /src/runtime/lockrank_on.go
type lockRankStruct struct {
	rank lockRank
	pad  int
}
