package main

import (
	"fmt"
	"time"
)

func main() {
	/* i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i) */
	//um loop for implementa a execução repetida de um código baseado em um contador ou uma variável de loop.
	//Ao contrário do que ocorre com outras linguagens de programação que têm vários constructos de looping
	//como o while, do etc., o go tem somente o loop for
	for j := 0; j < 10; j += 3 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
}
