package testools

import (
	"fmt"
	"reflect"
)

func (r Request) CheckTest(expected, response interface{}, comment ...any) bool {

	if !reflect.DeepEqual(expected, response) {
		fmt.Println("=>PRUEBA: ", r.TestName)
		fmt.Printf("=>SE ESPERABA:[%v]\n\n=>SE OBTUVO:[%v]\n", expected, response)
		fmt.Println(comment...)
		return false
	}

	return true
}
