package os

import "io/fs"

func DirPerm() fs.FileMode { return 0o755 }

func FilePerm() fs.FileMode { return 0o644 }
