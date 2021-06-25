package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Resource  string
	Action    string
	Arguments []string
}

func NewCommand(text string) *Command {
	if !strings.HasPrefix(strings.ToLower(text), "af ") {
		return nil
	}

	cols := strings.Split(text, " ")

	// Add dummy items to prevent index out of boundary
	if len(cols) < 4 {
		cols = append(cols, make([]string, 4-len(cols))...)
	}

	return &Command{
		Resource:  strings.ToLower(cols[1]),
		Action:    strings.ToLower(cols[2]),
		Arguments: cols[3:],
	}
}

func (c *Command) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %v", c.Resource, c.Action, c.Arguments)
}

const HELP_MESSAGE = `Afred Bot accepts following commands:
- af help
  Print this message

- af todo
  Print all your todos

- af todo add buy the milk
	Add a todo called "buy the milk"

- af todo del 1
	Delete todo item [1]`

const READ_DB_ERROR = "Cannot read from database"
const WRITE_DB_ERROR = "Cannot write to database"

func (c *Command) Execute(conf *Conf, msg *ReceivedMessage) (string, error) {
	if c == nil {
		return "Sorry I cannot understand", errors.New("Nil Command")
	}

	if c.Resource == "help" {
		return HELP_MESSAGE, nil
	} else if c.Resource == "todo" {
		if c.Action == "" {
			todo, err := ReadTODOs(conf, msg.UserID())
			if err != nil {
				return READ_DB_ERROR, err
			}
			return fmt.Sprint(todo), nil
		} else if c.Action == "add" {
			todo, err := ReadTODOs(conf, msg.UserID())
			if err != nil {
				return READ_DB_ERROR, err
			}
			item := strings.Join(c.Arguments, " ")
			todo.AddItem(item)
			if err := WriteTODOs(conf, msg.UserID(), todo); err != nil {
				return WRITE_DB_ERROR, err
			}
			return fmt.Sprintf("Added a new todo:\n%v", todo), nil
		} else if c.Action == "del" {
			todo, err := ReadTODOs(conf, msg.UserID())
			if err != nil {
				return READ_DB_ERROR, err
			}
			idx, err := strconv.Atoi(c.Arguments[0])
			if err != nil {
				return "Invalid number", err
			}
			if err := todo.DelItem(idx - 1); err != nil {
				return fmt.Sprintf("Failed to delete todo at [%d]", idx), err
			}
			if err := WriteTODOs(conf, msg.UserID(), todo); err != nil {
				return WRITE_DB_ERROR, err
			}
			return fmt.Sprintf("Deleted todo [%d]\n%v", idx, todo), nil
		}
	}

	return "Sorry, I cannot understand", errors.New("Unknown Command")
}
