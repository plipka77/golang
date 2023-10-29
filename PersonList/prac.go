package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	personList := new(PersonLinkedList)

	fmt.Println("Welcome to Person Linked List Creator")
	for true {
		fmt.Printf("Choose from one of the following options:\n1) Append\n2) Insert\n3) Delete\n4) Print \n5) Get Size\n6) Clear List\n7) Exit\n")
		input, _ := reader.ReadBytes('\n')

		switch input[0] {
		case '1':
			person := createPersonNode(reader)
			personList.append(person)
		case '2':
			index, valid := getIndex(reader)
			if !valid {
				continue
			}
			person := createPersonNode(reader)
			personList.insert(person, index)
		case '3':
			index, valid := getIndex(reader)
			if !valid {
				continue
			}
			personList.delete(index)
		case '4':
			personList.print()
		case '5':
			fmt.Printf("The current list size is: %d\n", personList.getSize())
		case '6':
			personList.clear()
		case '7':
			return
		default:
			fmt.Println("Invalid option selection")
		}
	}
}

func createPersonNode(reader *bufio.Reader) PersonNode {
	fmt.Println("Please provide a name:")
	name, _ := reader.ReadBytes('\n')

	fmt.Println("Please provide an id:")
	id, _ := reader.ReadBytes('\n')

	result := new(PersonNode)
	result.id = string(id[:len(id)-1])
	result.name = string(name[:len(name)-1])
	return *result
}

func getIndex(reader *bufio.Reader) (int, bool) {
	fmt.Println("Please provide an index:")
	indexBytes, _ := reader.ReadBytes('\n')
	index, err := strconv.Atoi(string(indexBytes[:len(indexBytes)-1]))
	if err != nil {
		fmt.Println("Please enter a valid integer...")
		return 0, false
	}

	return index, true
}
