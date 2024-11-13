// Package sysexits provides exit status codes similar to FreeBSD's sysexits.h
package sysexits

// Exit status codes
const (
	EX_OK          = 0  // Successful termination
	EX_USAGE       = 64 // Command line usage error
	EX_DATAERR     = 65 // Data format error
	EX_NOINPUT     = 66 // Cannot open input
	EX_NOUSER      = 67 // Addressee unknown
	EX_NOHOST      = 68 // Host name unknown
	EX_UNAVAILABLE = 69 // Service unavailable
	EX_SOFTWARE    = 70 // Internal software error
	EX_OSERR       = 71 // System error (e.g., can't fork)
	EX_OSFILE      = 72 // Critical OS file missing
	EX_CANTCREAT   = 73 // Can't create (user) output file
	EX_IOERR       = 74 // Input/output error
	EX_TEMPFAIL    = 75 // Temporary failure; user is invited to retry
	EX_PROTOCOL    = 76 // Remote error in protocol
	EX_NOPERM      = 77 // Permission denied
	EX_CONFIG      = 78 // Configuration error
)
