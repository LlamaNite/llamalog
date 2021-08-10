package llamalog

const (
	SeqBlack   = "\u001b[30m"
	SeqRed     = "\u001b[31m"
	SeqGreen   = "\u001b[32m"
	SeqYellow  = "\u001b[33m"
	SeqBlue    = "\u001b[34m"
	SeqMagenta = "\u001b[35m"
	SeqCyan    = "\u001b[36m"
	SeqWhite   = "\u001b[37m"
	SeqReset   = "\u001b[0m"
)

func Black(text string) string {
	return SeqBlack + text + SeqReset
}

func Red(text string) string {
	return SeqRed + text + SeqReset
}

func Green(text string) string {
	return SeqGreen + text + SeqReset
}

func Yellow(text string) string {
	return SeqYellow + text + SeqReset
}

func Blue(text string) string {
	return SeqBlue + text + SeqReset
}

func Magenta(text string) string {
	return SeqMagenta + text + SeqReset
}

func Cyan(text string) string {
	return SeqCyan + text + SeqReset
}

func White(text string) string {
	return SeqWhite + text + SeqReset
}
