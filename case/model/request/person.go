package request

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonGreeting struct {
	Person   Person `json:"person"`
	Greeting string `json:"greeting"`
}
