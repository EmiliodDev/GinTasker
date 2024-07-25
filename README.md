# GinTasker

GinTasker is a simple Todo List API built using Go, Gin, GORM, and Viper. This API allows users to manage their tasks efficiently with basic CRUD operations: **Create**, **Read**, **Update**, and **Delete**. It's designed as a lightweight and straightforward solution for task management.

## Features

- **Task Management**: Create, view, update, and delete tasks.
- **Built with Go**: Utilizes Go's robust features for building scalable and fast applications.
- **Frameworks and Libraries**: Incorporate Gin for routing, GORM for ORM, and Viper for configuration management.
- **No User Authorization**: Focused on core functionality without authentication.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Installation

To set up GinTasker on your local machine, follow these steps:

1. **Clone the repository**:
```bash
git clone https://github.com/EmiliodDev/GinTasker
cd GinTasker
```

2. **Install dependencies**:
```bash
go mod tidy
```

3. **Configure the application**:

    Create a `config.yaml` file:
```yaml
appname: GinTasker
port: "8080"

db:
  host: localhost
  port: "3306"
  user: youruser
  password: yourpassword
  name: yourdbname
```

4. **Run the application**:
```bash
make run
```

5. **Access the API**:

    The API will be available at `http://localhost:8080`.

## Usage

You can interact with the API using tools like [Postman](https://www.postman.com/) or [cURL](https://curl.se/). Here are some example requests:

- **List all tasks**:
    ```bash
    curl http://localhost:8080/api/v1/tasks
    ```

- **Get task by ID**:
    ```bash
    curl http://localhost:8080/api/v1/task/{id}
    ```

- **Create a new task**:
    ```bash
    curl -X POST http://localhost:8080/api/v1/create -H "Content-Type: application/json" -d '{"name": "New Task", "description": "Complete the assignment.", "completed": false}'
    ```

- **Update task**:
    ```bash
    curl -X PUT http://localhost:8080/api/v1/update/{id} -H "Content-Type: application/json" -d '{"name": "Updated Task", "description": "Update the assignment.", "completed": true}'
    ```

- **Delete task**:
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/delete/{id}
    ```

## API Endpoints

| Method |	Endpoint         |	Description          |
|--------|-------------------|-----------------------|
| GET	 |/api/v1/tasks      |Get all tasks          |
| GET	 |/api/v1/task/{id}  |Get a specific task    |
| POST	 |/api/v1/create	 |Create a new task      |
| PUT	 |/api/v1/update/{id}|Update a specific task |
| DELETE |/api/v1/delete/{id}|Delete a specific task |

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
