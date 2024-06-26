package main

type Vacancy struct {
	name   string
	salary int
	url    string
	reqs   []int
}

var Vacancies []Vacancy = []Vacancy{
	{
		name:   "Example",
		salary: 50000,
		url:    "https://api.hh.ru/vacancies/",
		reqs:   []int{1, 0, 0, 0},
	},
	{
		name:   "Разработчик",
		salary: 100000,
		url:    "https://gauctr.ru/job/dorozhnyj-rabochij-ot-120-000-rublej-na-ruki/",
		reqs:   []int{1, 1, 0, 0},
	},
	{
		name:   "Администратор",
		salary: 100000,
		url:    "https://api.superjob.ru/2.0/vacancies/",
		reqs:   []int{2, 1, 1, 1},
	},
}
