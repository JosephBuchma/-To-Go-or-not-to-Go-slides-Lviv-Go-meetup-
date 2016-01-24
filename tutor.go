package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// CustomTypes OMIT

type age uint8

func (a age) String() string {
	return fmt.Sprintf("%d years old", a)
}

type Gopher struct {
	Name string `json:"name"`
	Age  age    `json:"age"`
}

func (g Gopher) Whoami() string {
	return fmt.Sprintf("Hi, my name is %s and I'm %s", g.Name, g.Age)
}

func (g *Gopher) Rename(name string) {
	g.Name = name
}

// CustomTypes OMIT

type Heizenberg struct{}

// TypeSwitch OMIT

type BusyError struct {
	msg string
}

func (e BusyError) Error() string {
	return e.msg
}

type JessieFailError struct {
	msg    string
	repeat int
}

func (e JessieFailError) Error() string {
	return fmt.Sprintf("Jessie error: %s\n", strings.Repeat(e.msg+" ", e.repeat))
}

func (g *Heizenberg) MakeMeth() error {
	switch rand.Intn(3) {
	case 0:
		return BusyError{"I'm busy"}
	case 1:
		return JessieFailError{"B**ch!", rand.Intn(5) + 1}
	}
	return nil
}

// TypeSwitch OMIT

// TypeSwitchExample OMIT

func TypeSwitchExample() {
	H := Heizenberg{}
	for err := H.MakeMeth(); err != nil; err = H.MakeMeth() {
		fmt.Println(err)

		switch err := err.(type) { // HL
		case BusyError:
			fmt.Println("Waiting...")
			time.Sleep(500 * time.Millisecond)
		case JessieFailError:
			fmt.Printf("Teaching Jessie %d times\n", err.repeat)
			time.Sleep(time.Duration(err.repeat) * 100 * time.Millisecond)
		}
		fmt.Println("Trying again...")
	}

	fmt.Println("METH IS DONE!!!")
}

// TypeSwitchExample OMIT

// EmbeddingStructs OMIT

type MyLog struct {
	*log.Logger
	debug bool
}

func (ml MyLog) Debug(msg string) {
	if ml.debug {
		ml.Logger.Print(fmt.Sprintf("[DEBUG] %s", msg))
	}
}

func UseMyLog() {
	logger := MyLog{log.New(os.Stdout, "|logger| ", log.Ltime), true}

	logger.Print("Proxied method of log.Logger")
	logger.Debug("LoooooooooooooL")
}

// EmbeddingStructs OMIT

// CustomTypesShowcase OMIT
func CustomTypesShowcase() {
	gopher := &Gopher{Name: "Walter White", Age: 50}
	fmt.Println(gopher.Whoami())
	gopher.Rename("Heizenberg")
	fmt.Println(gopher.Whoami())
	fmt.Println(gopher.Age)
}

// CustomTypesShowcase OMIT

// EmptyInterfaces OMIT
func EmptyInterfacesDemo() {
	var foo interface{}
	foo = 3
	foo = "lol"
	foo = Gopher{"John Snow", 30}

	gop, ok := foo.(Gopher)
	fmt.Println(ok)
	fmt.Println(gop.Whoami())

	_, ok = foo.(string)
	fmt.Println(ok)
}

// EmptyInterfaces OMIT


func main() {
	rand.Seed(time.Now().UnixNano())
	// CustomTypesShowcaseM OMIT
	//CustomTypesShowcase()
	// CustomTypesShowcaseM OMIT
	// EmbeddingStructsM OMIT
	//UseMyLog()
	// EmbeddingStructsM OMIT
	// TypeSwitchExampleM OMIT
	//TypeSwitchExample()
	// TypeSwitchExampleM OMIT
	// EmptyInterfacesM OMIT
	//EmptyInterfacesDemo()
	// EmptyInterfacesM OMIT
}
