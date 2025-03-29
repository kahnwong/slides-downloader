package spider

import (
	"fmt"
	"os"
)

func AppendLineToFile(filename string, data string) {
	// open file
	f, err := os.OpenFile(fmt.Sprintf("data/%s.txt", filename), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write to file
	if _, err = f.WriteString(fmt.Sprintf("%s\n", data)); err != nil {
		panic(err)
	}
}
