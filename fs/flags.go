package fs

import "os"

type flag struct{ flag int }

// Flags to open file.
var (
	FlagRead            = flag{os.O_RDONLY}
	FlagReadCreate      = flag{os.O_RDONLY | os.O_CREATE}
	FlagWrite           = flag{os.O_WRONLY | os.O_CREATE | os.O_TRUNC}
	FlagWriteAppend     = flag{os.O_WRONLY | os.O_CREATE | os.O_APPEND}
	FlagReadWriteAppend = flag{os.O_RDWR | os.O_CREATE | os.O_APPEND}
)
