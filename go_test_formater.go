package golang_unit_test_formatter

import (
	"fmt"
)

type TestNum int

const (
	success = "\u2713"
	failed  = "\u2717"
)

// Выводит форматированный заголовок теста.
func HeaderTestCase(id TestNum, comment string, input any) {
	fmt.Printf("\t Test %d:\t%s %v\n", id, comment, input)
}

// Возвращает форатированную строку положительного результата.
func StrCorrect(value any) string {
	return fmt.Sprintf("\033[32m\t\t%s\tGot correct: %v\033[0m", success, value)
}

// Возвращает форматированную строку отрицательного результата.
func StrWrong(value any) string {
	return fmt.Sprintf("\033[31m\t\t%s\tGot wrong: %v\033[0m", failed, value)
}

// Выводит результат в форматированном виде.
func ShowResult(result any, ok bool) {
	if ok {
		fmt.Printf("%s\n", StrCorrect(result))
	}
	if !ok {

		fmt.Printf("%s\n", StrWrong(result))
	}
}
