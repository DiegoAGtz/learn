package main

import (
	"fmt"
	"sort"
)

type user struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type porEdad []user
type porApellido []user

func (a porEdad) Len() int {return len(a)}
func (a porEdad) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a porEdad) Less(i, j int) bool {return a[i].Edad < a[j].Edad}

func (a porApellido) Len() int {return len(a)}
func (a porApellido) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a porApellido) Less(i, j int) bool {return a[i].Apellido < a[j].Apellido}

func main() {
	u1 := user{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := user{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := user{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

    for _, user := range users {
        fmt.Println(user.Nombre, user.Apellido, user.Edad)
        for _, dicho := range user.Dichos {
            fmt.Println("\t", dicho)
        }
    }

    fmt.Println("\nPor Edad------")
    sort.Sort(porEdad(users))
    for _, user := range users {
        fmt.Println(user.Nombre, user.Apellido, user.Edad)
        sort.Strings(user.Dichos)
        for _, dicho := range user.Dichos {
            fmt.Println("\t", dicho)
        }

    }

    fmt.Println("\nPor Apellido------")
    sort.Sort(porApellido(users))
    for _, user := range users {
        fmt.Println(user.Nombre, user.Apellido, user.Edad)
        for _, dicho := range user.Dichos {
            fmt.Println("\t", dicho)
        }
    }
}
