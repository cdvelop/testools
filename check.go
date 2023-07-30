package testools

import (
	"fmt"
	"log"
)

func CheckTest(prueba string, expected, response interface{}, a ...any) {

	if expected != response {
		fmt.Println("=>PRUEBA: ", prueba)
		fmt.Printf("=>SE ESPERABA:[%v]\n=>SE OBTUVO:[%v]\n", expected, response)
		fmt.Println(a...)
		log.Fatal()
	}
}
