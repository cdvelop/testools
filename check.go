package testools

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckTest(prueba string, expected, response interface{}, t *testing.T, a ...any) {

	if !reflect.DeepEqual(expected, response) {
		fmt.Println(a...)
		// if expected != response {
		fmt.Println("=>PRUEBA: ", prueba)
		fmt.Printf("=>SE ESPERABA:[%v]\n\n=>SE OBTUVO:[%v]\n", expected, response)
		t.Fatal()
	}
}

func (r Request) CheckTest(expected, response interface{}, comment ...any) {
	CheckTest(r.TestName, expected, response, r.T, comment...)
}
