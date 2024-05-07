package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Получаем текст, считанный из стандартного ввода
	inputText := scanner.Text()

	// Выводим текст на экран
	fmt.Println(inputText)

}
