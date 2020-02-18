package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	phone string
}
type Contactlist struct {
	Contacts []*Contact
}

func (contactlist *Contactlist) add (contact Contact)  {
	contactlist.Contacts = append(contactlist.Contacts, &contact)
}



func (contactlist Contactlist) printstr() {
	for _,val := range contactlist.Contacts {
		fmt.Println("ID: ", val.ID )
		fmt.Println("Name: ", val.Name)
		fmt.Println("Phone number: ", val.phone)
		fmt.Println("--------------------")
	}

}
func (contactlist *Contactlist) update (id int, clist Contact) {
	for _, value := range contactlist.Contacts {
		if id == value.ID {
			value.ID = clist.ID
			value.Name = clist.Name
			value.phone = clist.phone
		}
	}
}
func (contactlist Contactlist) delete (id int) {
	for index, value := range contactlist.Contacts {
		if id == value.ID {
			contactlist.Contacts = append(contactlist.Contacts[:index], contactlist.Contacts[index+1:]...)
		}
	}
}

func (contactList Contactlist) search(letter string) Contactlist {
	var arr Contactlist
	for _, value := range contactList.Contacts {

		if strings.Contains(strings.ToUpper(value.Name), strings.ToUpper(letter)) {
			arr.Contacts = append(arr.Contacts, value)
		}
	}
	return arr
}

func main () {

	var clist Contactlist

	c1 := Contact{
		ID:    23,
		Name:  "Otabek",
		phone: "+998998136254",
	}

	c2 := Contact{
		ID:    11,
		Name:  "Oybek",
		phone: "+998993917199",
	}

	c3 := Contact{
		ID:    8,
		Name:  "Zayniddin",
		phone: "+998978785055",
	}

	c4 := Contact{
		ID:    72,
		Name:  "Mashhurbek",
		phone: "+998991234567",
	}

	fmt.Println("ADD -----------------")
	clist.add(c1)
	clist.add(c2)
	clist.add(c3)

	fmt.Println("PRINT --------------- ")
	clist.printstr()

	fmt.Println("SEARCH -----------------")

	searchList := clist.search("q")
	searchList.printstr()


	fmt.Println("COntacts update now ----")
	clist.update(11, c4)

	clist.printstr()

	clist.delete(23)
	clist.printstr()
}