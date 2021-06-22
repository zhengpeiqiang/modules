package structModel

type A struct {
	Name string
	Age  int
	Job  string
}

type B struct {
	num         string
	performance performance
	rang        int
}

type performance struct {
	cpu    string
	memory string
	disk   string
}

type C struct {
	height int
	width  int
	worth  float64
}
