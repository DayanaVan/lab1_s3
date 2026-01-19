package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printHelp() {
	fmt.Println("Usage: ./l1 [FILE] [COMMAND] [ARGUMENTS]")
}

func main() {
	argc := len(os.Args)
	argv := os.Args

	if argc < 3 {
		printHelp()
		return
	}

	file, err := os.Open(argv[1])
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	switch argv[2][0] {
	case 'A':
		arr := *newArray()
		readArray(scanner, &arr)
		file.Close()
		switch argv[2][1:] {
		case "ADD":
			if argc != 4 {
				printHelp()
				return
			}
			arr.add(argv[3])
			fmt.Println(arr.size)
			fmt.Println(arr)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeArray(writer, &arr)
			file.Close()
		case "INSERT":
			if argc != 5 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			arr.insert(argv[3], index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeArray(writer, &arr)
			file.Close()
		case "GET":
			if argc != 4 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[3])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			fmt.Println(arr.get(index))
		case "REMOVE":
			if argc != 4 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[3])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			arr.remove(index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeArray(writer, &arr)
			file.Close()
		case "CHANGE":
			if argc != 5 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[3])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			arr.change(index, argv[4])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeArray(writer, &arr)
			file.Close()
		case "SIZE":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Println(arr.size)
		case "PRINT":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Println(arr.print())
		default:
			fmt.Println("Неизвестная операция")
		}
	case 'F':
		var fl ForwardList
		readForwardList(scanner, &fl)
		file.Close()
		switch argv[2][1:] {
		case "ADDTAIL":
			if argc != 4 {
				printHelp()
				return
			}
			fl.addTail(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "ADDHEAD":
			if argc != 4 {
				printHelp()
				return
			}
			fl.addHead(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "INSERT":
			if argc != 5 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			fl.insert(argv[3], index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "REMOVETAIL":
			if argc != 3 {
				printHelp()
				return
			}
			fl.removeTail()

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "REMOVEHEAD":
			if argc != 3 {
				printHelp()
				return
			}
			fl.removeHead()

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "REMOVEINDEX":
			if argc != 4 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[3])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			fl.remove(index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "PRINT":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Printf("С головы: %s\nС хвоста %s\n", fl.printFromHead(), fl.printFromTail())
		case "REMOVE":
			if argc != 5 {
				printHelp()
				return
			}
			num, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Номер вхождения не число")
				return
			}
			fl.removeKey(argv[3], num)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeForwardList(writer, &fl)
			file.Close()
		case "FIND":
			if argc != 5 {
				printHelp()
				return
			}
			num, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Номер вхождения не число")
				return
			}
			if fl.find(argv[3], num) != nil {
				fmt.Println("Найдено")
			} else {
				fmt.Println("Не найдено")
			}
		}
	case 'L':
		var fl List
		readList(scanner, &fl)
		file.Close()
		switch argv[2][1:] {
		case "ADDTAIL":
			if argc != 4 {
				printHelp()
				return
			}
			fl.addTail(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "ADDHEAD":
			if argc != 4 {
				printHelp()
				return
			}
			fl.addHead(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "INSERT":
			if argc != 5 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			fl.insert(argv[3], index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "REMOVETAIL":
			if argc != 3 {
				printHelp()
				return
			}
			fl.removeTail()

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "REMOVEHEAD":
			if argc != 3 {
				printHelp()
				return
			}
			fl.removeHead()

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "REMOVEINDEX":
			if argc != 4 {
				printHelp()
				return
			}
			index, err := strconv.Atoi(argv[3])
			if err != nil {
				fmt.Println("Индекс не число")
				return
			}
			fl.remove(index)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "PRINT":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Printf("С головы: %s\nС хвоста %s\n", fl.printFromHead(), fl.printFromTail())
		case "REMOVE":
			if argc != 5 {
				printHelp()
				return
			}
			num, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Номер вхождения не число")
				return
			}
			fl.removeKey(argv[3], num)

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeList(writer, &fl)
			file.Close()
		case "FIND":
			if argc != 5 {
				printHelp()
				return
			}
			num, err := strconv.Atoi(argv[4])
			if err != nil {
				fmt.Println("Номер вхождения не число")
				return
			}
			if fl.find(argv[3], num) != nil {
				fmt.Println("Найдено")
			} else {
				fmt.Println("Не найдено")
			}
		}
	case 'Q':
		var q Queue
		readQueue(scanner, &q)
		file.Close()
		switch argv[2][1:] {
		case "PUSH":
			if argc != 4 {
				printHelp()
				return
			}
			q.push(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeQueue(writer, &q)
			file.Close()
		case "POP":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Println(q.pop())

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeQueue(writer, &q)
			file.Close()
		case "PRINT":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Printf("С начала %s\n", q.print())
		}
	case 'S':
		var s Stack
		readStack(scanner, &s)
		file.Close()
		switch argv[2][1:] {
		case "PUSH":
			if argc != 4 {
				printHelp()
				return
			}
			s.push(argv[3])

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeStack(writer, &s)
			file.Close()
		case "POP":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Println(s.pop())

			file, err := os.OpenFile(argv[1], os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Не удалось открыть файл")
				return
			}
			writer := bufio.NewWriter(file)
			writeStack(writer, &s)
			file.Close()
		case "PRINT":
			if argc != 3 {
				printHelp()
				return
			}
			fmt.Printf("С конца %s\n", s.print())
		}
	default:
		fmt.Println("Неизвестная структура")
	}
}
