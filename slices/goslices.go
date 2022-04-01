package main

import "fmt"

func main() {

	// var c [5]int
	// c[0] = 10
	// c[1] = 20
	// c[2] = 30
	// c[3] = 40
	// c[4] = 50

	//c as an array, array assignement is copy by value where slice assignment is by reference
	//c := [5]int{10, 20, 30, 40, 50}

	numbers := []int{1, 3}

	fmt.Printf("slice numbers - len : %v , cap : %v , address : %p \n", len(numbers), cap(numbers), &numbers)

	numbers = append(numbers, 4)
	fmt.Printf("slice numbers - len : %v , cap : %v , address : %p \n", len(numbers), cap(numbers), &numbers)

	//c as an slice
	c := []int{10, 20, 30, 40, 50}

	a := c
	b := c
	// b := c[0:3]
	// a := c[2:5]

	fmt.Printf("a type %T\n  b type %T\n c type %T\n", a, b, c)

	fmt.Println("a", a, "len", len(a), "cap", cap(a))

	fmt.Println("b", b, "len", len(b), "cap", cap(b))

	fmt.Println("c", c, "len", len(c), "cap", cap(c))

	fmt.Println("after changin array element")

	c[2] = 300

	fmt.Println("a", a, "len", len(a), "cap", cap(a))

	fmt.Println("b", b, "len", len(b), "cap", cap(b))

	fmt.Println("c", c, "len", len(c), "cap", cap(c))

	//after append

	b = append(b, 6)

	fmt.Println("b = append(b, 6)")
	fmt.Println("a", a, "len", len(a), "cap", cap(a))

	fmt.Println("b", b, "len", len(b), "cap", cap(b))

	fmt.Println("c", c, "len", len(c), "cap", cap(c))

	//memory allocation

	//case 1
	s := []int{}
	s = append(s, 1)
	s = append(s, 2)
	t := append(s, 3)
	u := append(s, 4)
	fmt.Println("s: ", s, "\nt: ", t, "\nu: ", u)
	fmt.Println("s", s, "len", len(s), "cap", cap(s), "address of s %p", &s[1])
	fmt.Println("t", t, "len", len(t), "cap", cap(t), "address of t %p, %p", &t[1], &t[2])
	fmt.Println("u", u, "len", len(u), "cap", cap(u), "address of u %p, %p", &u[1], &u[2])

	fmt.Println("changing value in existing slice")
	s[0] = 100
	fmt.Println("s: ", s, "\nt: ", t, "\nu: ", u)
	fmt.Println("s", s, "len", len(s), "cap", cap(s), "address of s %p", &s[1])
	fmt.Println("t", t, "len", len(t), "cap", cap(t), "address of t %p, %p", &t[1], &t[2])
	fmt.Println("u", u, "len", len(u), "cap", cap(u), "address of u %p, %p", &u[1], &u[2])

	//case 2
	s = append(s, 3)
	x := append(s, 4)
	fmt.Println("Intermediate x", x, "len", len(x), "cap", cap(x), "address of x %p, %p", &x[2], &x[3])
	y := append(s, 5)
	fmt.Println("s: ", s, "\nx: ", x, "\ny: ", y)
	fmt.Println("s", s, "len", len(s), "cap", cap(s), "address of s %p", &s[2])
	fmt.Println("x", x, "len", len(x), "cap", cap(x), "address of x %p, %p", &x[2], &x[3])
	fmt.Println("y", y, "len", len(y), "cap", cap(y), "address of y %p, %p", &y[2], &y[3])

	fmt.Println("After changing first element of s")
	//case 2
	s[0] = 1
	y[3] = 500
	fmt.Println("s: ", s, "\nx: ", x, "\ny: ", y)
	fmt.Println("s", s, "len", len(s), "cap", cap(s), "address of s %p", &s[2])
	fmt.Println("x", x, "len", len(x), "cap", cap(x), "address of x %p, %p", &x[2], &x[3])
	fmt.Println("y", y, "len", len(y), "cap", cap(y), "address of y %p, %p", &y[2], &y[3])

}
