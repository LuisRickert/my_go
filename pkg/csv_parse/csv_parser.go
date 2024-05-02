package csv_parse

import (
	"regexp"
	"strings"
	"unicode"
)

// define csv struct
type csv_row struct {
	raw_data  string
	separator string
	size      int
	values    []string
	index     int
}
type csv_object struct {
	column_names []string
	values       []csv_row
	row_size     int
}


func clean_string(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, in)
}

func Csv_factory(in_sting string) csv_object {
	separator := ","

	var emptyValues = []csv_row{}
	result := csv_object{values: emptyValues}

	fill_empty_values := regexp.MustCompile(`,,`)

	for i, line := range strings.Split(strings.TrimSuffix(in_sting, "\n"), "\n") {
		var clean_values []string
		if i == 0 {
			// assume first row are column names
			for _, name := range strings.Split(line, separator) {
				clean_values = append(clean_values, clean_string(name))
			}

			result.column_names = clean_values
			result.row_size = len(clean_values)
		} else {
			// parse the row
			filled_row := fill_empty_values.ReplaceAll([]byte(line), []byte(",none,"))

			// split row at separator
			var row = csv_row{}
			for _, val := range strings.Split(strings.TrimSuffix(string(filled_row), separator), separator) {

				clean_values = append(clean_values, clean_string(string(val)))
				row = csv_row{index: i, raw_data: line, separator: separator, values: clean_values, size: len(clean_values)}
			}
			result.values = append(result.values, row)

		}

	}
	return result
}

func Get_column(column_name string, data csv_object) []string {

	var result []string
	// find index of column
	idx := -1
	for i, val := range data.column_names {
		if column_name == val {
			idx = i
			break
		}
	}
	if idx == -1 {
		panic("Column name: \"" + column_name + "\" not in row")
	}
	// get values from rows
	for _, row := range data.values {
		result = append(result, row.values[idx])
	}
	return result
}

