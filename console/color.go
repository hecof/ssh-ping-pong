package console

// Red adds red color to the msg
func Red(msg string) string {
	return "\033[0;31m" + msg + "\033[0m"
}

// RedLn adds red color to the msg
func RedLn(msg string) string {
	return "\033[0;31m" + msg + "\033[0m\n"
}

// Green adds green color to the msg
func Green(msg string) string {
	return "\033[0;32m" + msg + "\033[0m"
}

// GreenLn adds green color to the msg and a new line at the end
func GreenLn(msg string) string {
	return "\033[0;32m" + msg + "\033[0m\n"
}

// Cyan adds cyan color to the msg
func Cyan(msg string) string {
	return "\033[0;36m" + msg + "\033[0m"
}

// CyanLn adds cyan color to the msg and a new line at the end
func CyanLn(msg string) string {
	return "\033[0;36m" + msg + "\033[0m\n"
}

// Yellow adds yellow color to the msg
func Yellow(msg string) string {
	return "\033[0;33m" + msg + "\033[0m"
}

