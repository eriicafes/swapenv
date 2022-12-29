package fs

import "os"

type flag struct {
	flag   int
	create bool
}

// Flags to open file.
var (
	FlagRead            = flag{flag: os.O_RDONLY}
	FlagReadCreate      = flag{flag: os.O_RDONLY | os.O_CREATE, create: true}
	FlagWrite           = flag{flag: os.O_WRONLY | os.O_CREATE | os.O_TRUNC, create: true}
	FlagWriteAppend     = flag{flag: os.O_WRONLY | os.O_CREATE | os.O_APPEND, create: true}
	FlagReadWriteAppend = flag{flag: os.O_RDWR | os.O_CREATE | os.O_APPEND, create: true}
)
