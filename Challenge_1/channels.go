package main

type idol struct {
	name string
	age  int
}

func checkIdolName(idols [3]idol) [3]string {
	is_en_mem := make(chan string)
	go checkName(is_en_mem, idols)
	is_en := [3]string{}
	is_en[0] = <-is_en_mem
	is_en[1] = <-is_en_mem
	is_en[2] = <-is_en_mem
	return is_en

}

func checkName(is_en_mem chan<- string, idol [3]idol) {
	for _, e := range idol {
		if e.name == "Nerissa" {
			is_en_mem <- "Nerrissa"
		}

		if e.name == "Noel" {
			is_en_mem <- "Noel"
		}
		if e.name == "Mori" {
			is_en_mem <- "Mori"
		}

	}
}
