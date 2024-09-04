package logic

import "strings"

type Task uint8

const (
	TermSig Task = iota
	LoadSig
	LogSig
	HelpSig
	NilSig
	// Scalability)))))))))))))))))))))))))))
)

func (t Task) String() string {
	return [...]string{
		"q",
		"load",
		"loglevel",
		"help",
		"",
	}[t]
}

func parseCommand(input string) (Task, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return NilSig, []string{} // wrong one
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "q":
		return TermSig, args
	case "load":
		if len(args) == 2 {
			return LoadSig, args
		}
		return NilSig, []string{} // wrong one
	case "loglevel":
		if len(args) == 1 {
			return LogSig, args
		}
		return NilSig, []string{} // wrong one
	case "help":
		return HelpSig, args
	default:
		return NilSig, []string{} // wrong one
	}
}
