package main

import "fmt"

type User struct {
	nombre string
	email  string
	estado bool
}
type Student struct {
	user   User
	codigo string
}

type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	//Relación de uno a uno
	/*
		u1 := User{
			nombre: "Carlos",
			email:  "carlos@gmail.com",
			estado: true,
		}

		u2 := User{
			nombre: "Ernesto",
			email:  "ernesto@gmail.com",
			estado: true,
		}

		e1 := Student{
			user:   u1,
			codigo: "std-0001",
		}

		fmt.Println(u2)

		fmt.Println(e1)*/

	//Relacion de uno a muchos
	v1 := Video{titulo: "01-Introducción"}
	v2 := Video{titulo: "02-Instalación"}

	c1 := Curso{
		titulo: "Curso de GO",
		videos: []Video{v1, v2},
	}

	v1.curso = c1
	v2.curso = c1

	fmt.Println(v1.curso.titulo)

	for _, video := range v1.curso.videos {
		fmt.Println(video.titulo)
	}

}
