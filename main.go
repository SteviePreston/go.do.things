package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Learn Go")
	todos.add("Learn C")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}
