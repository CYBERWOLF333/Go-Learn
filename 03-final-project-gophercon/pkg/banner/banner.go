package banner

import "fmt"

// ANSI Color constants
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Dim     = "\033[2m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// PrintBanner outputs the cool cyber ASCII logo
func PrintBanner(version string) {
	fmt.Printf(`%s%s
   ____             _               ____                          
  / ___| ___  _ __ | |__   ___ _ __|  _ \ ___  ___ ___  _ __  
 | |  _ / _ \| '_ \| '_ \ / _ \ '__| |_) / _ \/ __/ _ \| '_ \ 
 | |_| | (_) | |_) | | | |  __/ |  |  _ <  __/ (_| (_) | | | |
  \____|\___/| .__/|_| |_|\___|_|  |_| \_\___|\___\___/|_| |_|
             |_|                                              
%s%s       >>> Passive OSINT & Attack Surface Discovery Engine <<<%s
%s                v%s | Built in Go by CYBERWOLF333%s
%s=================================================================%s

`, Cyan, Bold, Yellow, Bold, Reset, Dim, version, Reset, Blue, Reset)
}

// Colorize helpers
func ColorRed(s string) string    { return Red + s + Reset }
func ColorGreen(s string) string  { return Green + s + Reset }
func ColorYellow(s string) string { return Yellow + s + Reset }
func ColorCyan(s string) string   { return Cyan + s + Reset }
func ColorBold(s string) string   { return Bold + s + Reset }
