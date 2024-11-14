package lab6

import (
	"fmt"
)

type Fox struct {
	Name  string
	Color string
	Age   int
}

func (f *Fox) Speak() string {
	return "Ring-ding-ding-ding-dingeringeding!"
}

func (f *Fox) GetDescription() string {
	return fmt.Sprintf("This is %s, a %d-year-old %s fox.", f.Name, f.Age, f.Color)
}

func (f *Fox) GetView() string {
	return ` 
         \\ // 
        ( o.o )
         > ^ <
      `
}

func RunLab6() {
	fox := Fox{
		Name:  "Foxy",
		Color: "red",
		Age:   3,
	}

	fmt.Println(fox.Speak())
	fmt.Println(fox.GetDescription())
	fmt.Println(fox.GetView())
}
