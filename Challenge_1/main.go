package main

<<<<<<< HEAD
import "time"

type text struct {
	body string
	date time.Time
=======
func main() {
	sound_demon := idol{name: "Nerissa", age: 2500}
	knight := idol{name: "Noel", age: 25}
	reaper := idol{name: "Mori", age: 2000}

	idol_list := [3]idol{sound_demon, knight, reaper}
	test := checkIdolName(idol_list)
	for _, i := range test {
		print(i)
	}

>>>>>>> 80ad60e (Continued channels practice)
}
