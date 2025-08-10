# Splitexpense Backend

This is the backend for the Splitexpense application.

## Prerequisites

Before running the application, you need to have a PostgreSQL database running.

The application uses the following database connection string by default:
`postgres://user:password@localhost/splitexpense?sslmode=disable`

You can override this by creating a `.env` file in the root of the project with the following content:
```
DATABASE_URL=your_postgres_connection_string
JWT_SECRET=your_jwt_secret
```

## Running the application

To start the service, run the following command in the root of the project:

```bash
go run main.go
```

The server will start on port `8080` by default.

## API Endpoints

You can find example cURL requests for the `signup` and `login` endpoints in the `requests.txt` file.# split-expense
