package utility

import "flag"

func FlagParse() (string, string) {
	mode := flag.String("mode", "version", "migration")
	path := flag.String("c", "", "config file")
	flag.Parse()

	switch *mode {
	case "up":
		return "up", *path
	case "down":
		return "down", *path
	case "reset":
		return "reset", *path
	case "version":
		return "version", *path
	}
	return "", ""
}
