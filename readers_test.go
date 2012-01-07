package euler

import (
	"testing"
)

func TestTableIntReadFile(t *testing.T) {
	file := "testdata/prob067.txt"
	table, err := TableIntReadFile(file)
	if err != nil {
		t.Error(file, " ", err)
	} else if len(table) != 100 {
		t.Error(file, " length wrong ", len(table))
	} else {
		for i, r := range table {
			if len(r) != i+1 {
				t.Error(file, " row length wrong ", i, len(r))
				break
			}
		}
	}
}
