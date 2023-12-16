
**User Management System (UMS) - GoFR**

**Overview**

The User Management System (UMS) is a simple RESTful API implemented in Gofr, designed for managing user-related operations. This project provides basic functionality to create, retrieve, update, and delete user records in a MongoDB database. The system exposes endpoints to interact with user data, making it a foundational component for applications requiring user management capabilities.

**Features**

**Create User**: Add a new user to the system by providing user details such as username and email.

**Get User:** Fetch a list of all users stored in the system.

**Update User:** Modify user information, including updating the username and email.

**Delete User:** Remove a user from the system based on their ID.

**Technologies Used**

**Gofr**: A lightweight and fast web framework for Go (Golang), facilitating the creation of RESTful APIs.

**MongoDB:** A NoSQL database used for storing user data.

**Getting Started**

To set up and run the User Management System:

_**Clone the repository:**_

git clone <repository_url>
cd user-management-system-gofr

_**Install dependencies:**_

go mod download

_**Configure MongoDB:**_

Replace the MongoDB connection string in main.go with your own connection details.

Run the application:

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

_**Usage**_

**Create User:**

Send a POST request to /users with user details in the request body.

**Get User(s):**

To retrieve all users, send a GET request to /users.

**Update User:**

Send a PUT request to /users/{id} with updated user information in the body.


**Delete User:**

Send a DELETE request to /users/{id} to remove a user from the system.
