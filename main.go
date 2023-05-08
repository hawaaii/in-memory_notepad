package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var notes []string
  var limit int
	var exit bool
	fmt.Println("Enter the maximum number of notes:")
  fmt.Scan(&limit)
	for {
		command, data, position, err := userPrompt("Enter a command and data: ")
		if err != nil {
			fmt.Println(err)
    } else {
			switch command {
			case "create":
				createNote(&notes, data, limit)
			case "update":
				updateNote(&notes, data, position)
			case "delete":
				deleteNote(&notes, position)
			case "list":
				listNotes(&notes)
			case "clear":
				clearNotes(&notes)
			case "exit":
				fmt.Print("[info] bye!\n")
				exit = true
			default:
				fmt.Println("[Error] Unknown command")
			}
		}
		if exit {break}
	}
}

func userPrompt(prompt string) (string, string, int, error) {
	fmt.Println(prompt)
	var position, data, strErr string
	var positionInt int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	command := input[0]

	if command == "update" || command == "delete" {
        if len(input) > 1 { position, data = input[1], strings.Join(input[2:], " ") } // check err
	} else {
		data = strings.Join(input[1:], " ")
		return command, data, 0, nil
	}
	positionInt, err := strconv.Atoi(position)
	if err != nil {
		strErr = "[Error] Invalid position: " + position
	}
	if data == "" && command == "update" {
		strErr = "[Error] Missing note argument"
	}
	if position == "" {
		strErr = "[Error] Missing position argument"
	}
	
    if strErr == "" {
		return command, data, positionInt, nil
	}
	return command, data, positionInt, errors.New(strErr)
}

func createNote(notes *[]string, data string, limit int) {
	if data == "" {
		fmt.Println("[Error] Missing note argument")
	} else if len(*notes) >= limit {
		fmt.Println("[Error] Notepad is full")
	} else {
		*notes = append(*notes, data)
		fmt.Println("[OK] The note was successfully created")
	}
}

func listNotes(notes *[]string) {
	if len(*notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for noteIdx, note := range *notes {
			fmt.Printf("[Info] %v: %v\n", noteIdx+1, note)
		}
	}
}

func clearNotes(notes *[]string) {
	*notes = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

func updateNote(notes *[]string, data string, position int) {
    notesCopy := *notes
	if len(notesCopy) == 0 {
		fmt.Println("[Error] There is nothing to update")
    }else if (position-1) > len(notesCopy) {
        fmt.Printf("[Error] Position %v is out of the boundaries [1, %v]\n",position,len(notesCopy)+1)
	} else {
		notesCopy[position-1] = data
		fmt.Println("[OK] The note at position", position, "was successfully updated")
	}
}

func deleteNote(notes *[]string, position int) {
	notesCopy := *notes
	if len(notesCopy) <= position {
		fmt.Println("[Error] There is nothing to delete")
    }else if (position-1) > len(notesCopy) {
        fmt.Printf("[Error] Position %v is out of the boundaries [1, %v]\n",position,len(notesCopy)+1)
	} else {
		*notes = append(notesCopy[:(position-1)], notesCopy[position:]...)
		fmt.Println("[OK] The note at position", position, "was successfully deleted")
	}
}
