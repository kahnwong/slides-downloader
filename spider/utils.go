package spider

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func AppendLineToFile(filename string, data string) {
	// open file
	f, err := os.OpenFile(fmt.Sprintf("data/%s.txt", filename), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal().Msg("failed to close file")
		}
	}(f)

	// write to file
	if _, err = fmt.Fprintf(f, "%s\n", data); err != nil {
		panic(err)
	}
}
