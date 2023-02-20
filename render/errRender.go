package render

import "fmt"

func ErrorCheck(err error) {
	if err != nil {
		fmt.Printf("Error parsing template %s \n", err)
	}

}
