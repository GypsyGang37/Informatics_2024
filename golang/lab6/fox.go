package lab6

import (
	"fmt"
)

// Fox представляет лисицу с определенными характеристиками.
type Fox struct {
	Name  string
	Color string
	Age   int
}

// Speak возвращает строку, представляющую звук, который издает лисица.
func (f *Fox) Speak() string {
	return "Ring-ding-ding-ding-dingeringeding!"
}

// GetDescription возвращает описание лисицы.
func (f *Fox) GetDescription() string {
	return fmt.Sprintf("This is %s, a %d-year-old %s fox.", f.Name, f.Age, f.Color)
}

// GetView возвращает ASCII-изображение лисицы.
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
