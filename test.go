package main

import "fmt"

type perm string

const (
	Read  perm = "read"
	Write perm = "write"
	Exec  perm = "execute"
)

var Admin = "admin"
var User = perm("user")

func checkPermission(p perm) {
	// check the permission
}

func main() {
	fmt.Println(Admin)
	fmt.Println(User)
	checkPermission(Read)
	checkPermission(Write)
	checkPermission(Exec)
	checkPermission(User)
	checkPermission(perm(Admin))

}
