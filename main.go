package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readOperation() string {
	fmt.Println("Digite a operação (soma, subtrai, multiplica, divide, sair):")
	reader := bufio.NewReader(os.Stdin)
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func readNumbers() []int {
	fmt.Println("Digite os números separados por espaço:")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Split(strings.TrimSpace(line), " ")
	nums := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		nums[i] = num
	}
	return nums
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}

func main() {
	for {
		op := readOperation()
		if op == "sair" {
			break
		}
		nums := readNumbers()
		var result int
		switch op {
		case "soma":
			result = soma(nums...)
		case "subtrai":
			result = subtrai(nums...)
		case "multiplica":
			result = multiplica(nums...)
		case "divide":
			result = divide(nums...)
		default:
			fmt.Println("Operação desconhecida")
			continue
		}
		fmt.Println("Resultado:", result)
	}
}
