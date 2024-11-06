package logService

// Define a custom type for colors
type color int

// Enum for console colors
const (
	reset color = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

// Color codes mapping
var colorCodes = map[color]string{
	reset:   "\033[0m",
	red:     "\033[31m",
	green:   "\033[32m",
	yellow:  "\033[33m",
	blue:    "\033[34m",
	magenta: "\033[35m",
	cyan:    "\033[36m",
	white:   "\033[37m",
}
