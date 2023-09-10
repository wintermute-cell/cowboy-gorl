package gui

import (
	"cowboy-gorl/pkg/logging"
	"strconv"
	"strings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
parseStyleDefinition converts a string of style definitions into a map representation.

Input format: "property:value|property2:value2|...".

Note:
- If a value itself contains colons, the entire value after the first colon is
considered as the value. For example, "url:http:\/\/example.com" will be parsed
as {"url": "http:\/\/example.com"}.
- Extra spaces between properties, colons, and values are trimmed in the output.
*/
func parseStyleDef(style_definition string) map[string]any {
    if style_definition == "" {
        return make(map[string]any)
    }

	// remove potential leading/trailing pipe symbol
    style_definition = strings.Trim(style_definition, "|")

	pairs := strings.Split(style_definition, "|")
	style_map := make(map[string]any)

	// Define conversion functions
	converters := map[string]func(string) any {
        "color": func(value string) any {
            rgba := strings.Split(value, ",")
            if len(rgba) != 4 {
                logging.Warning("Invalid color format for value: %v", value)
                return nil
            }

            var result [4]uint8
            for i, v := range rgba {
                f, err := strconv.Atoi(v) // directly parse to int
                if err != nil || f < 0 || f > 255 { 
                    logging.Warning("Invalid color component in value: %v", value)
                    return nil
                }
                result[i] = uint8(f)
            }
            return rl.NewColor(result[0], result[1], result[2], result[3])
        },
        "background": func(value string) any {
            rgba := strings.Split(value, ",")
            if len(rgba) != 4 {
                logging.Warning("Invalid color format for value: %v", value)
                return nil
            }

            var result [4]uint8
            for i, v := range rgba {
                f, err := strconv.Atoi(v) // directly parse to int
                if err != nil || f < 0 || f > 255 { 
                    logging.Warning("Invalid color component in value: %v", value)
                    return nil
                }
                result[i] = uint8(f)
            }
            return rl.NewColor(result[0], result[1], result[2], result[3])
        },
		"font": func(value string) any {
			return value // already a string
		},
		"debug": func(value string) any {
            b, err := strconv.ParseBool(value)
            if err != nil{
                logging.Warning("Invalid bool value: %v", value)
            }
            return b
		},
		// add more converters here as needed
	}

	for _, pair := range pairs {
		if pair == "" {
			logging.Warning("Found empty pair in gui styledef!")
			continue
		}
		pv := strings.SplitN(pair, ":", 2)
		if pv[0] == "" || len(pv) != 2 || pv[1] == "" {
			logging.Warning("Found empty property or value in gui styledef: %v", pair)
			continue
		}

		// Trim spaces in case there are any around the property or value
		prop := strings.TrimSpace(pv[0])
		val := strings.TrimSpace(pv[1])

		// Convert to appropriate data type
		if converter, ok := converters[prop]; ok {
			style_map[prop] = converter(val)
		} else {
			logging.Warning("Unknown property in gui styledef: %v", prop)
		}
	}

	return style_map
}
