package main

import "testing"

func TestInsertElement_atStart(t *testing.T) {

	list := LinkedList{}

	list.Insert("Jeet")

	if list.len != 1 {
		t.Error("Element add failed", list.len)
		//t.Fail()
	}

}

func TestInsertElement_three(t *testing.T) {

	list := LinkedList{}

	list.Insert("Jeet")
	list.Insert("kashi")
	list.Insert("yash")

	if list.len != 3 {
		t.Error("Element add failed", list.len)
		//t.Fail()
	}

}
