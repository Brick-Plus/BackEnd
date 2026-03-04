package main

type Article struct {
	ID    int
	Name  string
	Price int // en centimes
}

var articles = []Article{
	{ID: 1, Name: "Remy", Price: 1000},
	{ID: 2, Name: "lego", Price: 1200},
	{ID: 3, Name: "brick", Price: 200},
	{ID: 4, Name: "teste", Price: 700},
	{ID: 5, Name: "Hugo qui coute chere", Price: 1250},
}
