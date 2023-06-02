package exit

import "os"

func main() { // want "there is os.Exit in the main function"
	os.Exit(1)
}
