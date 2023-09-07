# GraphQL POC implemented in GO

This Go application is a GraphQL TODO app that leverages GraphQL queries, mutations and subscriptions. It provides functionality to create tasks and users, mark tasks as completed, and query tasks for users. Users can subscrie to ne notified when they have overdue tasks.

## Features

- Create tasks with due dates and assign them to users.
- Mark tasks as completed.
- Query tasks for a specific user.
- Create users.
- Real-time notifications for overdue tasks using GraphQL subscriptions and an in-memory pub/sub system.

## Getting Started

1. Clone the repository to your local machine:

   ```shell
   git clone https://github.com/liviudnicoara/go-graphql-poc.git
   cd go-graphql-poc
   ```

2. Install the required dependencies:

    ```shell
    go mod tidy
    ```
3. Build and run the application:

    ```shell
    go run main.go
    ```

The application will start the GraphQL server, and you can access the GraphQL Playground at http://localhost:8080 in your web browser.

### Usage
GraphQL Playground
Access the GraphQL Playground by opening http://localhost:8080 in your web browser.

Use the Playground to send queries and mutations to the GraphQL server. The provided schema includes the following types:

User: Represents a user with an ID and a name.
Task: Represents a task with an ID, text, done flag, due date, and the user it's assigned to.

DateTime: A custom scalar type for representing date and time values.

Mutations
Use GraphQL mutations to create tasks, mark them as completed, and assign tasks to users.

Queries
Use GraphQL queries to retrieve tasks for a specific user and perform other data retrievals.

Subscriptions
The application includes GraphQL subscriptions for notifying users when they have overdue tasks.

Contributing
Contributions are welcome! If you'd like to contribute to this project or have suggestions for improvements, please open an issue or create a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.