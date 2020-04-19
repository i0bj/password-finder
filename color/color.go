package color

import "fmt"

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[0;97m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
	//  Bold
	BBlack  = Color("\033[1;30m%s\033[0m") // Black
	BRed    = Color("\033[1;31m%s\033[0m") // Red
	BGreen  = Color("\033[1;32m%s\033[0m") // Green
	BYellow = Color("\033[1;33m%s\033[0m") // Yellow
	BBlue   = Color("\033[1;34m%s\033[0m") // Blue
	BPurple = Color("\033[1;35m%s\033[0m") // Purple
	BCyan   = Color("\033[1;36m%s\033[0m") //  Cyan
	BWhite  = Color("\033[1;37m%s\033[0m") //  White

	//  Underline
	UBlack  = Color("\033[4;30m%s\033[0m") //  Black
	URed    = Color("\033[4;31m%s\033[0m") //  Red
	UGreen  = Color("\033[4;32m%s\033[0m") //  Green
	UYellow = Color("\033[4;33m%s\033[0m") //  Yellow
	UBlue   = Color("\033[4;34m%s\033[0m") //  Blue
	UPurple = Color("\033[4;35m%s\033[0m") //  Purple
	UCyan   = Color("\033[4;36m%s\033[0m") //  Cyan
	UWhite  = Color("\033[4;37m%s\033[0m") //  White

	//  Background
	OnBlack  = Color("\033[40m%s\033[0m") //  Black
	OnRed    = Color("\033[41m%s\033[0m") //  Red
	OnGreen  = Color("\033[42m%s\033[0m") //  Green
	OnYellow = Color("\033[43m%s\033[0m") //  Yellow
	OnBlue   = Color("\033[44m%s\033[0m") //  Blue
	OnPurple = Color("\033[45m%s\033[0m") //  Purple
	OnCyan   = Color("\033[46m%s\033[0m") //  Cyan
	OnWhite  = Color("\033[47m%s\033[0m") //  White

	//  High Intensty
	IBlack  = Color("\033[0;90m%s\033[0m") //  Black
	IRed    = Color("\033[0;91m%s\033[0m") //  Red
	IGreen  = Color("\033[0;92m%s\033[0m") //  Green
	IYellow = Color("\033[0;93m%s\033[0m") //  Yellow
	IBlue   = Color("\033[0;94m%s\033[0m") //  Blue
	IPurple = Color("\033[0;95m%s\033[0m") //  Purple
	ICyan   = Color("\033[0;96m%s\033[0m") //  Cyan
	IWhite  = Color("\033[0;97m%s\033[0m") //  White

	//  Bold High Intensty
	BIBlack  = Color("\033[1;90m%s\033[0m") //  Black
	BIRed    = Color("\033[1;91m%s\033[0m") //  Red
	BIGreen  = Color("\033[1;92m%s\033[0m") //  Green
	BIYellow = Color("\033[1;93m%s\033[0m") //  Yellow
	BIBlue   = Color("\033[1;94m%s\033[0m") //  Blue
	BIPurple = Color("\033[1;95m%s\033[0m") //  Purple
	BICyan   = Color("\033[1;96m%s\033[0m") //  Cyan
	BIWhite  = Color("\033[1;97m%s\033[0m") //  White

	//  High Intensty backgrounds
	OnIBlack  = Color("\033[0;100m%s\033[0m") //  Black
	OnIRed    = Color("\033[0;101m%s\033[0m") //  Red
	OnIGreen  = Color("\033[0;102m%s\033[0m") //  Green
	OnIYellow = Color("\033[0;103m%s\033[0m") //  Yellow
	OnIBlue   = Color("\033[0;104m%s\033[0m") //  Blue
	OnIPurple = Color("\033[10;95m%s\033[0m") //  Purple
	OnICyan   = Color("\033[0;106m%s\033[0m") //  Cyan
	OnIWhite  = Color("\033[0;107m%s\033[0m") //  White

)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
