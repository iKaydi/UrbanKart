# UrbanKart
A simple ordering application where you can register, order products and modify or cancel accordingly.
There also is an admin that will assign a representative to users according to place

Note all of this is in one main file
Need to understand how to do multiple files

# Shopping
The code consists of 4 structures. Namely : User,Product,List and Cart.
List and Cart are an array of User and Product respectively.
The whole code runs around using customer ID to call the respective customer from the list of customers
The ID's are the index from the list of users starting from 0 till 'length-1'.
Every operation that is done first asks for user-ID and then proceeds to carry out the task.

# Delivery

This part consists of 2 strucutres : Representative and DeliveryOrg
Delivery Org is the array of representatives
Indexing follows same rule as Product

Note : 99 for switching to Admin
