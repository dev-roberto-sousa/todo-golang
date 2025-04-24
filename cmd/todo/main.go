package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dev-roberto-sousa/todo-golang"
)

func main() {
	const todoFileName = ".todo.json"

	l := &todo.List{}
	// É uma boa prática usar STDERR em cli's em GO
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")

		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
