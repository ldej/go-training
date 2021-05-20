package main

import "fmt"

type Example struct {
	Value  string
	PValue *string
	Data   map[string]string
}

func (e Example) SetValue(a string) {
	fmt.Printf("e address: %p\n", &e)
	e.Value = a
}

func (e *Example) SetValuePtr(a string) {
	fmt.Printf("e address: %p\n", e)
	e.Value = a
}

func (e Example) SetPValue(a *string) {
	fmt.Printf("e address: %p\n", &e)
	fmt.Printf("e.PValue address: %p\n", e.PValue)
	e.PValue = a
}

func (e *Example) SetPValuePtr(a *string) {
	fmt.Printf("e address: %p\n", e)
	fmt.Printf("e.PValue address: %p\n", e.PValue)
	e.PValue = a
}

func (e Example) SetData(a map[string]string) {
	fmt.Printf("e address: %p\n", &e)
	fmt.Printf("e.Data address before SetData: %p\n", e.Data)
	e.Data = a
	fmt.Printf("e.Data address after SetData: %p\n", e.Data)
}

func (e *Example) SetDataPtr(a map[string]string) {
	fmt.Printf("e address: %p\n", e)
	fmt.Printf("e.Data address before SetDataPtr: %p\n", e.Data)
	e.Data = a
	fmt.Printf("e.Data address after SetDataPtr: %p\n", e.Data)
}

func (e Example) SetDataKeyValue(key string, value string) {
	fmt.Printf("e address: %p\n", &e)
	fmt.Printf("e.Data address: %p\n", e.Data)
	e.Data[key] = value
}

func (e *Example) SetDataKeyValuePtr(key string, value string) {
	fmt.Printf("e address: %p\n", e)
	fmt.Printf("e.Data address: %p\n", e.Data)
	e.Data[key] = value
}

func main() {
	e := Example{Data: map[string]string{}}

	fmt.Printf("The adress of e is: %p\n", &e)

	e.SetValue("1")
	fmt.Println("Expected 1:", e.Value)
	// e.SetValue is a value receiver, which copies the object, so the value is lost

	e.SetValuePtr("2")
	fmt.Println("Expected 2:", e.Value)
	// e.SetValuePtr is a pointer receiver, which points to the object, so the value is present

	e.SetData(map[string]string{"3": "4"})
	fmt.Println("Expected 3, 4:", e.Data)
	// e.SetData is a value receiver, which copies the object, so the data is lost

	e.SetDataPtr(map[string]string{"5": "6"})
	fmt.Println("Expected 5, 6:", e.Data)
	// e.SetDataPtr is a pointer receiver, which points to the object, so the data is present

	e.SetDataKeyValue("7", "8")
	fmt.Println("Expected 7, 8:", e.Data)
	// e.SetDataKeyValue is a value receiver, which copies the object. However, the object contains a map[string]string,
	// which is a pointer itself, to the copy of the object contains a copy of the pointer, so the data is NOT lost
	// We demonstrate this by printing the address of the pointer

	e.SetDataKeyValuePtr("9", "10")
	fmt.Println("Expected 9, 10:", e.Data)
	// e.SetDataKeyValuePtr is a pointer receiver, which points to the object, so the data is present

	a := "11"
	e.SetPValue(&a)
	fmt.Println("Expected 11:", e.PValue)
	// e.SetPValue is a value receiver, which copies the object, so the value is lost

	b := "12"
	e.SetPValuePtr(&b)
	fmt.Println("Expected 12:", e.PValue)
	// e.SetPValuePtr is a pointer receiver, which points to the object, so the value is present
}
