# Console Shop RESTful API

![Golang](https://img.shields.io/badge/Go-1.21-blue.svg)
![Gin](https://img.shields.io/badge/Gin-1.9.1-green.svg)

This is a small practice project aimed at developing a RESTful API using the Go programming language and the Gin framework. The project simulates a console shop, allowing you to perform various operations such as viewing the list of available consoles, adding new items to the existing list, and accessing specific consoles by their IDs.

## Features

- View the list of available consoles in the shop.
- Add a new console to the shop's inventory.
- Access a specific console's details by its unique ID.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.16 or higher installed on your machine. You can download it [here](https://golang.org/dl/).
- Gin 1.8.0 or higher installed. You can find the suitable version [here](https://github.com/gin-gonic/gin/releases).
- Basic knowledge of Go programming.
- Familiarity with RESTful API concepts.

## Getting Started

To get this project up and running on your local machine, follow these simple steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/ButcherWithSmile/console-shop.git
   ```

2. Change to the project directory:

   ```bash
   cd console-shop
   ```

3. Run the Go application:

   ```bash
   go run main.go
   ```

The API server will start, and you can access it at `http://localhost:8880`.

## API Endpoints

### List all consoles

- **URL:** `/consoles`
- **Method:** `GET`
- **Description:** Get a list of all available consoles in the shop.

### Add a new console

- **URL:** `/consoles`
- **Method:** `POST`
- **Description:** Add a new console to the shop's inventory.
- **Request Body:** JSON object containing console details (e.g., brand, model, price).

### Get a specific console

- **URL:** `/consoles/{id}`
- **Method:** `GET`
- **Description:** Retrieve the details of a specific console by its unique ID.

## Usage Example

Here's a simple example of how to use the API:

1. List all available consoles:

   ```bash
   curl -X GET http://localhost:8880/consoles
   ```

2. Add a new console:

   ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"id": "5", "brand": "Sony", "model": "PlayStation 4 Slim Console 1TB", "price": 189.99}' http://localhost:8880/consoles
   ```

3. Get the details of a specific console (replace `{id}` with the actual ID):

   ```bash
   curl -X GET http://localhost:8880/consoles/{id}
   ```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, please open an issue or create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
