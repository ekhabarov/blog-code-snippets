package service

import "github.com/google/wire"

var Provider = wire.NewSet(
	// here we binds concrete type *Service satisfies a dependency of type Xlogic.
	New, wire.Bind(new(Xlogic), new(*Service)),
)
