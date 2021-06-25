package models

import (
	"errors"
	"fmt"
	"strings"
)

type Todo struct {
	Items []string
}

func (t *Todo) AddItem(item string) {
	t.Items = append(t.Items, item)
}

func (t *Todo) DelItem(idx int) error {
	if idx < 0 || idx >= len(t.Items) {
		return errors.New("Array index out of boundary")
	}
	t.Items = append(t.Items[:idx], t.Items[idx+1:]...)
	return nil
}

func (t *Todo) String() string {
	if len(t.Items) == 0 {
		return "No todos!"
	}

	out := make([]string, len(t.Items))

	for i, item := range t.Items {
		out[i] = fmt.Sprintf("[%d] %s", i+1, item)
	}

	return strings.Join(out, "\n")
}
