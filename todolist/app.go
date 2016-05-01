package todolist

import (
	"fmt"
	"regexp"
	"strconv"
)

type App struct {
	TodoStore Store
}

func NewApp() *App {
	app := &App{TodoStore: NewFileStore()}
	app.TodoStore.Load()
	return app
}

func (a *App) AddTodo(input string) {
	parser := &Parser{}
	todo := parser.ParseNewTodo(input)

	a.TodoStore.Add(todo)
	a.TodoStore.Save()
	fmt.Println("Todo added.")
}

func (a *App) DeleteTodo(input string) {
	id := a.getId(input)
	if id != -1 {
		a.TodoStore.Delete(id)
		a.TodoStore.Save()
		fmt.Println("Todo deleted.")
	} else {
		fmt.Println("Could not find id.")
	}
}

func (a *App) ListTodos(input string) {
	//filtered := NewFilter(a.TodoStore.Todos()).filter()
	grouped := a.getGroups(input)

	formatter := NewFormatter(grouped)
	formatter.Print()
}

func (a *App) getId(input string) int {

	re, _ := regexp.Compile("\\d+")
	if re.MatchString(input) {
		id, _ := strconv.Atoi(re.FindString(input))
		return id
	} else {
		return -1
	}
}

func (a *App) getGroups(input string) *GroupedTodos {
	grouper := &Grouper{}
	contextRegex, _ := regexp.Compile(`by c.*$`)
	projectRegex, _ := regexp.Compile(`by p.*$`)

	var grouped *GroupedTodos

	if contextRegex.MatchString(input) {
		grouped = grouper.GroupByContext(a.TodoStore.Todos())
	} else if projectRegex.MatchString(input) {
		grouped = grouper.GroupByContext(a.TodoStore.Todos())
	} else {
		grouped = grouper.GroupByNothing(a.TodoStore.Todos())
	}
	return grouped
}
