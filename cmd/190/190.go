// JSON unmasrshal
// json to go website - https://mholt.github.io/json-to-go/
package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {

	s := `[{"Name":"John","Age":30},{"Name":"Jane","Age":25}]`
	bs := []byte(s)

	var employees []employee
	err := json.Unmarshal(bs, &employees)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employees)

	for _, emp := range employees {
		fmt.Println(emp.Name, emp.Age)
	}

}
