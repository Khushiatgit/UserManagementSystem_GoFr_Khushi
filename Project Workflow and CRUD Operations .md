**User Management System** 👥 is a project made to efficiently manage users in a system, in a workplace, a social media app or users of a particular service.
It is created using **GoFR** 🚀 and have used **MongoDB** cluster for storing of the data and user details.

To keep it simple, a user model has following properties.
![image](https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/c12398ec-bc37-4365-9d98-7d18f6118012)

**Lets go through the following Postman Collection to dive deep into the functionalities** 🎯 

<img width="382" alt="User Management System Collection" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/80e2c7f8-d5d1-4376-836e-bce5a180b557">



Now making use of the **CreateUserEndpoint** using HTTP **Post** Method, lets create 5 users in the system. 

**USER 1**
 <img width="1057" alt="Added User 1" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/704b4989-da32-4aac-9e72-cb8d04f67b67">
**USER 2**
<img width="1057" alt="Added User 2" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/c10264af-840d-47ea-b40e-530d31ce78ed">
**USER 3**
<img width="1057" alt="Added User 3" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/8ee0f1b9-2e1d-40eb-af4a-8e6e6b6f7830">
**USER 4**
<img width="1057" alt="Added User 4" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/33243677-0f94-4cbb-80ef-f671b0517e0b">
**USER 5**
<img width="1057" alt="Added User 5" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/1ad3f275-b651-44fa-ac58-b98fe8ed09ea">

Using the **GetUserEndpoint**, we'll fetch the details of all users present in the system using HTTP **Get** Method

<img width="1057" alt="Get ALL users" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/82f87ff8-1af6-44ba-a59e-67cd7ba6b841">

Let us update the details of USER 2 and USER 5 with details of updation sent in the request body to **UpdateUserEndpoint** using HTTP **Put** request,

<img width="1060" alt="Screenshot 2023-12-17 at 6 44 11 PM" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/379404a9-07b7-4bb2-9367-4ce24dcc0cb2">

<img width="1060" alt="Screenshot 2023-12-17 at 6 44 27 PM" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/c7532e9b-00de-4e33-be73-f60a47461ebc">


the **details have been updated**, let us check using GET request,

<img width="1057" alt="Get After Updations" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/190b0387-c78b-4371-88f2-7de2b73c1681">

Now, finally let us make use of **DeleteUserendpoint** to **delete User 4** from the system, using DElete method.

<img width="1057" alt="Delete User 4" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/d8da429d-bda8-478f-b4a0-f29b8c9b14c9">

Check using GET if the **user 4 has been deleted  ?**

<img width="1057" alt="Get After Deleting User 4" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/0a478713-f562-40e6-899f-ac817a2981bd">

We are almost done through the Project Workflow, finally we can see if the data is being correctly stored in the MongoDB collection, users :

<img width="1440" alt="UMS_MongoDb" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/0ca4a3f3-c942-4525-84cd-0ccff4ac177a">


Thank you, if you have followed till here 😇






