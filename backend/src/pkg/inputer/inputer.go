package inputer

import (
	"bufio"
	"fmt"
	"os"
)

func Inputer(l int, inputType string) (choice interface{}, err error) {
	//fmt.Println("(0 - выйти)")

	var formatString string
	reader := bufio.NewReader(os.Stdin)
	_ = reader
	for {
		switch inputType {
		case "int":
			formatString = "%d"
			var choice int

			_, err = fmt.Scanf(formatString, &choice)
			if err != nil {
				fmt.Println("ошибка ввода: %w", err)
				continue
			}
			//if choice == 0 {
			//	return choice, fmt.Errorf("выход")
			//}
			if inputType == "int" && l != 0 && (choice > l || choice < -1) {
				fmt.Println("не входит в рамки")
				continue
			}
			return choice, err
		case "str":
			formatString = "%s"
			var choice string

			_, err = fmt.Scanf(formatString, &choice)
			if err != nil {
				fmt.Println("ошибка ввода: %w", err)
				continue
			}

			return choice, err
		}
		//studio_id := studios[_].Id
		//fmt.Println(choice)

	}
}
