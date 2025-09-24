package hustr

import (
    "strconv"
)

func IntSliceToStr(input []int64, type_ int) []string {
    out := make([]string, len(input))
    for i, val := range input { out[i] = strconv.FormatInt(val, type_) }
    return out
}

func FloatSliceToStr(input []float64, type_ int) []string {
    out := make([]string, len(input))
    for i, val := range input { out[i] = strconv.FormatFloat(val, 'f', -1, type_) }
    return out
}

func PrettyHttp(response *http.Response) string {

    dump, err := httputil.DumpResponse(response, true)
    if err != nil { err_("Error dumping response: %v", err) }
    input := string(dump)

    new_input := []string{}
    scanner := bufio.NewScanner(strings.NewReader(input))
    for scanner.Scan() {
	line := scanner.Text()
	if line != "" {
	    new_input = append(new_input, line)
	}
    }
    input = strings.Join(new_input, "\n")

    input = strings.ReplaceAll(input, "\n", "\n    ]")
    input = strings.ReplaceAll(input, ": ", " [\n        ")
    input = strings.ReplaceAll(input, ", ", ",\n        ")
    input = strings.ReplaceAll(input, "]", "], ")
    input = strings.ReplaceAll(input, "3e", "\n}\n3e")

    lines := strings.SplitN(input, "\n", 2)
    if len(lines) > 1 {
	lines[0] 	= fmt.Sprintf("%s {", strings.TrimSpace(lines[0]))
	new_line2 	:= strings.Replace(lines[1], "], ", "", 1)
	lines[1] 	= fmt.Sprintf(".   %s", strings.TrimSpace(new_line2))
	lines[len(lines)-1] = fmt.Sprintf("%s\n}", strings.TrimSpace(lines[len(lines)-1]))

	input 		= strings.Join(lines, "\n")
    }
    return input
}
