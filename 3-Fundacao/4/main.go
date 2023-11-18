package main

import "fmt"

type UID = int

var userName string = "dev.sidnei"
var uid UID = 1000

const fullName = "Sidnei Simmon"

func main() {

	fmt.Printf("A variavel userName é do tipo: %T e tem o valor: %v \n", userName, userName)
	fmt.Printf("A constante fullName é do tipo: %T e tem o valor: %v \n", fullName, fullName)
	fmt.Printf("O tipo %T tem o valor de %v \n", uid, uid)

}
