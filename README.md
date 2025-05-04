# Book CRUD Application

A RESTful API application built with Go that provides CRUD (Create, Read, Update, Delete) operations for managing books.

## Project Structure

```
.
├── cmd/            # Application entry points
├── config/         # Configuration files
├── controllers/    # HTTP request handlers
├── middleware/     # HTTP middleware components
├── models/         # Data models
├── repository/     # Data access layer
├── routes/         # Route definitions
├── services/       # Business logic layer
└── utils/          # Utility functions
```

## Prerequisites

- Go 1.16 or higher
- Git

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd book_crud
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

## API Endpoints

The API provides the following endpoints:

- `GET /books` - Get all books
- `GET /books/{id}` - Get a specific book
- `POST /books` - Create a new book
- `PUT /books/{id}` - Update a book
- `DELETE /books/{id}` - Delete a book

## Development

To start development:

1. Make sure you have Go installed and configured
2. Install project dependencies
3. Run the application in development mode

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
