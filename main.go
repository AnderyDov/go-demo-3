package main

import "fmt"

func main() {
	m := map[string]string{
		"PurpleSchool": "https://purpleschool.ru",
	}

	fmt.Print("!", m["kkk"])

loop:
	for {
		res := ""
		fmt.Printf(`Выберите пунк меню:
1. Посмотреть закладки
2. Добавить заклпдку
3. Удалить закладку
4. Выход
`)
		fmt.Scan(&res)
		switch res {
		case "1":
			showNotes(m)
		case "2":
			addNote(m)
		case "3":
			delNote(m)
		case "4":
			fmt.Println("exit")
			break loop
		default:
			fmt.Println("exit")
			break loop

		}
	}

}

func showNotes(mm map[string]string) {
	for key, el := range mm {
		fmt.Println(key, " - ", el)
	}
}

func addNote(mm map[string]string) {
	key := ""
	value := ""
	fmt.Println("Name new note: ")
	fmt.Scan(&key)
	fmt.Println("Value new nite: ")
	fmt.Scan(&value)
	mm[key] = value
	fmt.Println("Added new nite - ", mm[key], " - ", value)

}

func delNote(mm map[string]string) {
	key := ""
	fmt.Println("Name deleted note: ")
	fmt.Scan(&key)
	delete(mm, key)
	fmt.Println("Deleted note - ", mm[key])
}
