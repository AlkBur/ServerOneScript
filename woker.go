package ServerOneScript

import (
	"fmt"
)

type Worker struct {
	ScriptChannel chan string
	quit          chan bool
}

func NewWorker() Worker {
	w := Worker{
		ScriptChannel: make(chan string),
		quit:          make(chan bool)}
	w.Start()
	return w
}

func (w Worker) Start() {
	go func() {
		for {
			select {
			case str := <-w.ScriptChannel:
				fmt.Println(str)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) RunScript(str string) {
	w.ScriptChannel <- str
}

func (w Worker) Stop() {
	w.quit <- true
}
