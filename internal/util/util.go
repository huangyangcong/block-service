package util

import "github.com/google/wire"

// ProviderSet is util providers.
var ProviderSet = wire.NewSet(NewEmailNotify)
