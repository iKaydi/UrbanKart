package main

import "fmt"

type User struct { //Struct User to store mobile and address With total and array of products
	Mobile  int
	Address string
	Trolley Cart
	Total   int
}

type List struct { //Struct List to store array of Users
	U []User
}

type Product struct { //Struct for the Product details
	Price       int
	Itemno      int
	Productname string
}

type Cart struct {
	Kart []Product //Struct for array of products
}

//Global declaration for ease of use
var P1 Product = Product{15000, 1, "TV"}
var P2 Product = Product{25000, 2, "Laptop"}
var P3 Product = Product{7500, 3, "Phone"}

//Find out customer ID
func WhichUSer() int {
	fmt.Printf("Which user number are you?")
	var i int
	fmt.Scanln(&i)
	return i
}

//Enter details of new user
func ScanUser() (int, string, Cart) {
	var n int
	var s string
	k := Cart{}
	fmt.Println("Enter Number")
	fmt.Scanln(&n)
	fmt.Println("Enter Address")
	fmt.Scanln(&s)
	return n, s, k
}

//Add User
func (L *List) AddUsers(user User) bool { //Func with reference to array of users with argument User type

	if L.U == nil { //If in the list of users, there is no one i.e. NULL
		L.U = make([]User, 0) //Make a new user
	}
	for _, v := range L.U {
		if v.Mobile == user.Mobile {
			return false
		}
	}
	L.U = append(L.U, user) //If there are already users, just append
	return true
}

//Display products
func DispProducts() {
	fmt.Printf("1.TV\n2.Laptop\n3.Phone\n")
}

//Quantity of Products to proceed with
func GetQuan() int {
	var i int
	fmt.Printf("Enter quantity")
	fmt.Scanln(&i)
	return i
}

//Placing order with customer id and quantity of products
func (L *List) PlaceOrder(quan int, cust_id int) {
	var ch, k int
	var TempProduct Product
	for k = 0; k < quan; k++ {
		DispProducts()
		fmt.Println("What would you like to choose?")
		fmt.Scanln(&ch)
		if ch == 1 {
			BIN := 1
			BP := 15000
			BN := "TV"
			TempProduct = Product{BP, BIN, BN}
			L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			fmt.Printf("%v", L.U[cust_id].Trolley)
			L.U[cust_id].Total = L.U[cust_id].Total + 15000
		}
		if ch == 2 {
			BIN := 2
			BP := 25000
			BN := "Laptop"
			TempProduct = Product{BP, BIN, BN}
			L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			L.U[cust_id].Total = L.U[cust_id].Total + 25000
		}
		if ch == 3 {
			BIN := 3
			BP := 7500
			BN := "Phone"
			TempProduct = Product{BP, BIN, BN}
			L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			L.U[cust_id].Total = 7500 + L.U[cust_id].Total
		}
	}
	fmt.Printf("%v", L.U)
}

//After entering user-id, user can add or remove items from trolley
func (L *List) ChangeOrder(cust_id int) {
	for true {
		fmt.Printf("What would you like to remove?\nPress 99 to exit\n")
		fmt.Printf("%v", L.U[cust_id].Trolley)
		var a int
		fmt.Scanln(&a)
		if a == 99 {
			break
		} else {
			L.U[cust_id].Total = L.U[cust_id].Total - L.U[cust_id].Trolley.Kart[a].Price
			L.U[cust_id].Trolley.Kart[a] = Product{0, 0, ""}
			fmt.Printf("%v", L.U[cust_id].Trolley)
		}
	}
}

//Creates a whole new cart for user
func (L *List) CancelOrder(cust_id int) {
	for k := range L.U {
		if k == cust_id {
			L.U[cust_id].Trolley = Cart{}
			L.U[cust_id].Total = 0
			fmt.Printf("%v", L.U)
		}
	}
}

//Display users cart
func (L *List) DisplayOrder(cust_id int) {
	for k := range L.U {
		if k == cust_id {
			fmt.Println("Products you have ordered are")
			fmt.Println(L.U[cust_id].Trolley)
			fmt.Println("Total :", L.U[cust_id].Total)
		}
	}
}

func main() {

	list1 := List{}
	var cart1 Cart
	var i, ch, n, quan, cust_id int
	var s string
	for true {
		fmt.Println("What would you like to do?\n1.Register\n2.Display Products\n3.Place Order\n4.Change Order\n5.Cancel Order\n6.Display Order\n")
		fmt.Scanln(&ch)
		if ch == 1 {
			n, s, cart1 = ScanUser()
			user1 := User{n, s, cart1, 0}
			list1.AddUsers(user1)
			fmt.Printf("%d %v \n", i, list1)
			i++
		}
		if ch == 2 {
			DispProducts()
		}
		if ch == 3 {
			cust_id = WhichUSer()
			quan = GetQuan()
			list1.PlaceOrder(quan, cust_id)
		}
		if ch == 4 {
			cust_id = WhichUSer()
			quan = GetQuan()
			fmt.Println("1.Add\n2.Remove")
			var a int
			fmt.Scanln(&a)
			if a == 1 {
				list1.PlaceOrder(quan, cust_id)
			}
			if a == 2 {
				list1.ChangeOrder(cust_id)
			}
		}
		if ch == 5 {
			cust_id = WhichUSer()
			list1.CancelOrder(cust_id)
		}

		if ch == 6 {
			cust_id = WhichUSer()
			list1.DisplayOrder(cust_id)
		}
	}
}
