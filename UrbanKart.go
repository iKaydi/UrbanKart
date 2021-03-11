package main

import "fmt"

type User struct { //Struct User to store mobile and address
	Mobile  int
	Address string
	Trolley Cart
}

type List struct { //Struct List to store array of Users
	U []User
}

type Product struct {
	Price       int
	Itemno      int
	Productname string
}

type Cart struct {
	Kart  []Product
	Total int
}

var P1 Product = Product{15000, 1, "TV"}
var P2 Product = Product{25000, 2, "Laptop"}
var P3 Product = Product{7500, 3, "Phone"}

func WhichUSer() int {
	fmt.Printf("Which user number are you?")
	var i int
	fmt.Scanln(&i)
	return i
}

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

func DispProducts() {
	fmt.Printf("1.TV\n2.Laptop\n3.Phone\n")
}

func GetQuan() int {
	var i int
	fmt.Printf("Enter quantity")
	fmt.Scanln(&i)
	return i
}

func (L *List) GetProducts(quan int, cust_id int) {
	var ch, k int
	//L.U[cust_id].Trolley.Kart[cust_id] = Product{}
	//fmt.Printf("%v", Bag)
	var TempProduct Product
	for k = 0; k < quan; k++ {
		//L.U[cust_id].Trolley.Kart[k] = Product{}
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
			//L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			//fmt.Printf("%v", L.U[cust_id])
			//Bag[k].Productname = "TV"
			//TempProduct = Product{Bag[k].Itemno, Bag[k].Price, Bag[k].Productname}
			//fmt.Printf("%v", TempProduct)
			//L.U[cust_id].Trolley = make([]L.U[cust_id].Kart,0)
			//L.U[cust_id].Trolley.Kart[z].Price = 15000
			//L.U[cust_id].Trolley.Kart[z].Itemno = 1
			//L.U[cust_id].Trolley.Kart[z].Productname = "TV"
			L.U[cust_id].Trolley.Total = 15000 + L.U[cust_id].Trolley.Total
			//fmt.Printf("%v", TempProduct)
			//TempProduct = Product{L.U[cust_id].Trolley.Kart[z].Price, L.U[cust_id].Trolley.Kart[z].Itemno, L.U[cust_id].Trolley.Kart[z].Productname}
			//L.U[cust_id].Trolley.Kart[k] = append(L.U[cust_id].Trolley.Kart[k],TempProduct)
			//fmt.Printf("%v", TempProduct)
		}
		if ch == 2 {
			//Bag[k].Itemno = 2
			//Bag[k].Price = 25000
			//Bag[k].Productname = "Laptop"
			BIN := 2
			BP := 25000
			BN := "Laptop"
			TempProduct = Product{BP, BIN, BN}
			//L.U[cust_id].Trolley.Kart = append(Bag, TempProduct)
			L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			L.U[cust_id].Trolley.Total = 25000 + L.U[cust_id].Trolley.Total
			//fmt.Printf("%v", TempProduct)
			//TempProduct = Product{Bag[k].Itemno, Bag[k].Price, Bag[k].Productname}
			//fmt.Printf("%v", TempProduct)
			//L.U[cust_id].Trolley.Kart[z].Price = 25000
			//L.U[cust_id].Trolley.Kart[z].Itemno = 2
			//L.U[cust_id].Trolley.Kart[z].Productname = "Laptop"
			//TempProduct = Product{L.U[cust_id].Trolley.Kart[z].Price, L.U[cust_id].Trolley.Kart[z].Itemno, L.U[cust_id].Trolley.Kart[z].Productname}
			//L.U[cust_id].Trolley.Kart[k] = Product{L.U[cust_id].Trolley.Total + 25000, 2, "Laptop"}
			//fmt.Printf("%v", TempProduct)
		}
		if ch == 3 {
			//Bag[k].Itemno = 3
			//Bag[k].Price = 7500
			//Bag[k].Productname = "Phone"
			BIN := 3
			BP := 7500
			BN := "Phone"
			TempProduct = Product{BP, BIN, BN}
			//fmt.Printf("%v", TempProduct)
			//L.U[cust_id].Trolley.Kart = append(Bag, TempProduct)
			L.U[cust_id].Trolley.Kart = append(L.U[cust_id].Trolley.Kart, TempProduct)
			L.U[cust_id].Trolley.Total = 7500 + L.U[cust_id].Trolley.Total
			//	TempProduct = Product{Bag[k].Itemno, Bag[k].Price, Bag[k].Productname}
			//		fmt.Printf("%v", TempProduct)
			//L.U[cust_id].Trolley.Kart[z].Price = 7500
			//L.U[cust_id].Trolley.Kart[z].Itemno = 3
			//L.U[cust_id].Trolley.Kart[z].Productname = "Phone"
			//TempProduct = Product{L.U[cust_id].Trolley.Kart[z].Price, L.U[cust_id].Trolley.Kart[z].Itemno, L.U[cust_id].Trolley.Kart[z].Productname}
			//L.U[cust_id].Trolley.Kart[k] = Product{L.U[cust_id].Trolley.Total + 7500, 3, "Phone"}
			//fmt.Printf("%v", TempProduct)
		}
	}
	fmt.Printf("%v", L.U)
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
			user1 := User{n, s, cart1}
			list1.AddUsers(user1)
			fmt.Printf("%d %v \n", i, list1)
			i++
		}
		if ch == 2 {
			DispProducts()
		}
		if ch == 3 {
			cust_id = WhichUSer() //Customer number
			quan = GetQuan()      //Quantity
			list1.GetProducts(quan, cust_id)

			//fmt.Printf("%v", list1.U[cust_id]) //succesfull print
			//fmt.Println(quan)                  //succesful print
			//Prod()

		}
	}
}

//Array of products is your kart
//so get the total
