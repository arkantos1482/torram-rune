package types

const (
	// ModuleName defines the module name
	ModuleName = "rune"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rune"
)

var (
	ParamsKey = []byte("p_rune")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
