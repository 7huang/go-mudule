package testmod

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello test module sucess, %s", name)
}

func TestGit(){
	fmt.Println("TEST GIT")
}
