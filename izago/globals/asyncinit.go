package globals

import "sync"

// AsyncInit is used over the project to perform given tasks once.
var AsyncInit sync.Once
