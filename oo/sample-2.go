package main

import(
	"fmt"
)

type Human struct {
	Name string
	w Walk
}

func(p *Human) SetWalk(w Walk){ p.w = w}
func(p Human) Walk(){p.w.Start()}

type Walk interface{
	Start()
}

type Slow struct {}

type Fast struct {}

func(s Slow) Start(){
	fmt.Println("Run slow")
}

func(f Fast) Start(){
	fmt.Println("Run fast")
}

func main()  {
	h := new(Human)
	h.SetWalk(new(Slow))
	h.Walk()
	h.SetWalk(new(Fast))
	h.Walk()
}