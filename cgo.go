//go:build gcc
// +build gcc

package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
)

// DriverName is the default driver name for SQLite.
const DriverName = "sqlite3"
