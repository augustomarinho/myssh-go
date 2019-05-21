package ini

import "fmt"

type File struct {
}

func (init *File) Read() {
	fmt.Println("Reading file")
}
