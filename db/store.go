package db

import (
	"encoding/binary"
	"os"

	"github.com/spf13/viper"
)

// Store is a store
type Store struct {
	Path string
}

// GetStore gets a store
func GetStore() Store {
	return Store{
		Path: viper.GetString("db"),
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int {
	return int(binary.BigEndian.Uint64(v))
}
