package todo

import (
	"time"
	"errors"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []item

func (t Todos) Add(task string) {
	items := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Now(),
	}

	t = append(t, items)
}	

func (t Todos) Complete(index int) error {
	if index < 0 || index > len(t) {
		return errors.New("Invalid index")	
	}

	t[index].Done = true
	t[index].CompletedAt = time.Now() 

	return nil
}	

func (t Todos) Delete(index int) error {
	if index < 0 || index > len(t) {
		return errors.New("Invalid index")	
	}

	/*
	* [5, 6, 4, 6, 7]
	* inclusive (index is 4: include from index 0, up to given index, but not including index) 
	* exclusive (from given index, include everything from that point to the end)
	*/
	t = append(t[:index-1], t[index:]...)
	
	return nil
}

func Load() {
	
}
