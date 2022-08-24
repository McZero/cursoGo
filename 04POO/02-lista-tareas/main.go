package main

import (
	"fmt"
)

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *TaskList) toPrint() {
	for i, task := range tl.tasks {
		fmt.Println(i, task.name)
	}
}

type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar curso de GO este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de HTML",
		desc:      "Completar curso de HTML esta semana",
		completed: true,
	}

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println(lista)

	t3 := Task{
		name:      "Curso de CSS",
		desc:      "Completar curso de CSS esta semana",
		completed: false,
	}

	lista.appendTask(&t3)

	fmt.Println(lista)

	lista.toPrint()

	lista.deleteTask(1)

	lista.toPrint()

	//t1.toPrint()
	//t2.toPrint()
}
