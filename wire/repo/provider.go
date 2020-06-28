package repo

import "github.com/google/wire"

var Provider = wire.NewSet(
	// here we binds concrete type *DB satisfies a dependency of type Database.
	New, wire.Bind(new(Database), new(*DB)),
)
