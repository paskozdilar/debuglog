package debuglog

import "log"

func Print(v ...any) {
	if enabled {
		v = append([]any{"[debuglog] "}, v...)
		log.Print(v...)
	}
}

func Printf(format string, v ...any) {
	if enabled {
		format = "[debuglog] " + format
		log.Printf(format, v...)
	}
}

func Println(v ...any) {
	if enabled {
		v = append([]any{"[debuglog]"}, v...)
		log.Println(v...)
	}
}
