
**User Management System (UMS) - GoFR** 🚀🚀


<img width="1155" alt="Screenshot 2023-12-17 at 3 17 24 AM" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/f1e35811-2845-4d8e-bc12-b0db0c9890e5">



 **Overview** 👥

The User Management System (UMS) is a simple RESTful API implemented in Gofr, designed for managing user-related operations. This project provides basic functionality to create, retrieve, update, and delete user records in a MongoDB database. The system exposes endpoints to interact with user data, making it a foundational component for applications requiring user management capabilities.

**Sequence Diagram**

<img width="840" alt="Screenshot 2023-12-17 at 7 22 36 PM" src="https://github.com/Khushiatgit/UserManagementSystem_GoFR_Khushi/assets/83766368/a1c2e673-fe01-44c9-b795-25ed280358b5">




**Features :**

**Create User**: Add a new user to the system by providing user details such as username and email.

**Get User:** Fetch a list of all users stored in the system.

**Update User:** Modify user information, including updating the username and email.

**Delete User:** Remove a user from the system based on their ID.

**Technologies Used :**

**Gofr**: A lightweight and fast web framework for Go (Golang), facilitating the creation of RESTful APIs.

**MongoDB:** A NoSQL database used for storing user data.

**Getting Started :**

To set up and run the User Management System:

_**Clone the repository:**_

git clone <repository_url>
cd user-management-system-gofr

_**Install dependencies:**_

go mod download
go get gofr.dev


_**Configure MongoDB:**_

Replace the MongoDB connection string in main.go with your own connection details.

_**Run the application:**_

go run main.go


**Access the API:**

The API will be accessible at http://localhost:8000.

_API Endpoints_

**Create User:**

POST /users


**Get User(s):**

GET /users - Get all users


**Update User:**

PUT /users/{id}


**Delete User:**

DELETE /users/{id}

__**Usage :**__

**Create User:**

Send a POST request to /users with user details in the request body.

**Get User :**

To retrieve all users, send a GET request to /users.

**Update User:**

Send a PUT request to /users/{id} with updated user information in the body.


**Delete User:**

Send a DELETE request to /users/{id} to remove a user from the system.

**Please refer Project Workflow and CRUD Operations.md for complete understanding of the flow and internals of the project, it also has the snapshots for testing of endpoints using Postman🎯 !**
