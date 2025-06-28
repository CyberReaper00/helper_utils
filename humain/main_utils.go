package humain

import (
    "os"
    "fmt"
    "log"
    "bufio"
    "strings"
    "strconv"
)

func TrueType(data string) any {

	if i, err := strconv.Atoi(data); err == nil { return i }
	if b, err := strconv.ParseBool(data); err == nil { return b }
	if f, err := strconv.ParseFloat(data, 64); err == nil { return f }
	return data
}

func Input(msg string, val ...any) any {
    fmt.Printf(msg, val...)

    reader 	:= bufio.NewReader(os.Stdin)
    data, _	:= reader.ReadString('\n')
    data1	:= strings.TrimSpace(data)
	input	:= TrueType(data1)
    return input
}

const top = "╭────────────────────────────────────────╮"
const mid = "├────────────────────────────────────────┤"
const bot = "╰────────────────────────────────────────╯"
func InputMenu(msgs ...string) []any {

	if len(msgs) < 1 { log.Fatalln("Error: No arguments were provided for the menu") }

	fmt.Println("\033[H\033[2J")
	final_slice := make( []any, len(msgs) )
	for i, msg := range msgs {
		if i == 0 {
			fmt.Println(top)
		} else if i > 0 && i < len(msgs) {
			fmt.Println(mid)
		}
		final_slice[i] = Input("│ %s: ", msg)
	}

	fmt.Println(bot)
	return final_slice
}

func ExitMsg() {
    xyz := Input("\n\033[1;42m Press Enter to exit... \033[0m\n")
    if xyz != "" { ExitMsg() }
    return
}

func Err(msg string, err error, val ...any) {
    if err != nil { log.Fatalf( fmt.Sprintf(msg, val...) )}
}

func PrettyErr(msg string, err error, fg bool, val ...any) {
    var scheme int
    if fg { scheme = 3
    } else { scheme = 4 }

    if err != nil {
		log.Fatalf(fmt.Sprintf("\033[1;%d1m", scheme) +
		fmt.Sprintf(msg, val...) +
		"\033[0m\n")
	}
}

func PrettyMsg(msg string, color string, fg bool, val ...any) {
    var scheme int
    var color_code string

    if fg { scheme = 3
    } else { scheme = 4 }

    if !strings.Contains(color, "#") {
		switch color {
			case "black": 	color_code = fmt.Sprintf("\033[1;%d0m", scheme)
			case "red": 	color_code = fmt.Sprintf("\033[1;%d1m", scheme)
			case "green": 	color_code = fmt.Sprintf("\033[1;%d2m", scheme)
			case "orange": 	color_code = fmt.Sprintf("\033[1;%d3m", scheme)
			case "blue": 	color_code = fmt.Sprintf("\033[1;%d4m", scheme)
			case "purple": 	color_code = fmt.Sprintf("\033[1;%d5m", scheme)
			case "teal": 	color_code = fmt.Sprintf("\033[1;%d6m", scheme)
			case "beige": 	color_code = fmt.Sprintf("\033[1;%d7m", scheme)
			default: log.Fatalln(fmt.Sprintf("\033[1;41 Invalid color '%s' given\033[0m", color))
		}

    } else {
		parts := strings.Split(color, "#")
		if len(parts) < 2 { log.Fatalln("Invalid input was given") }
		hex := parts[len(parts) - 1]

		if len(hex) != 6 { log.Fatalln("Invalid hex code was entered") }

		r64, errR := strconv.ParseInt(hex[0:2], 16, 0)
		g64, errG := strconv.ParseInt(hex[2:4], 16, 0)
		b64, errB := strconv.ParseInt(hex[4:6], 16, 0)

		if errR != nil || errG != nil || errB != nil {
			log.Fatalln("" +
			"\033[1;32mParsing Error:\n" +
			"Unknown value was encountered when parsing hex code\033[0m")
		}

		r := int(r64)
		g := int(g64)
		b := int(b64)

		color_code = fmt.Sprintf("\033[%d8;2;%d;%d;%dm", scheme, r, g, b)
    }

    fmt.Printf(color_code+msg+"\033[0m", val...)
}

func IntSliceContains(slice []int, target int) bool {
    for _, num := range slice {
		if num == target { return true }
    }
    return false
}

func StrSliceContains(slice []string, target string) bool {
    for _, str := range slice {
		if str == target { return true }
    }
    return false
}
