# UrbanKart

Initialize strucutre "User" with :Mobile,Address,Total and Trolley of type Cart
Initialize strucutre "List" with :U of type Array of User
Initialize structure "Product" with : Price,Itemno,Productname
Initialize structure "Cart" with : Kart which is array of Products

P1,P2,P3 are three products you initialize globally 

Function WhichUser - To find customer id
Function ScanUser - To get details of new user to be added
Function AddUser - Creates new user in the List
Function DispProduct - Displays the 3 products
Function GetQuan - Enters the quantity for adding products
Function Placeorder - Choose products for the respective user being called by customer ID
Function ChangeOrder - Add or remove products w.r.t customer id
Function CancelOrder - The user has his whole trolley = 0
Function DisplayOrder - Displays the users order


All these functions mostly work on customer ID, so that is very important as every other function works on customer ID
Note that I keep calling users from the list of users with customer id which is default index and further work on that
