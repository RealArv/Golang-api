package main

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {

}

func EncodeJson() {
	lcoCourses := []course{
		{"asd", 123, "asd", "asd", []string{"qwe", "qw"}},
	}
}
