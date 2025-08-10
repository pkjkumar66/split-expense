graph TD
    A[User] -->|HTTP Request| B(API Gateway);

    subgraph "Authentication"
        B --> C{/api/v1/auth};
        C -->|/signup| D[Auth Handler: SignUp];
        C -->|/login| E[Auth Handler: Login];
    end

    subgraph "User Management"
        B --> F{/api/v1/users};
        F -->|/me (GET)| G[User Handler: GetMe];
        F -->|/me (PUT)| H[User Handler: UpdateMe];
    end

    D --> I[Auth Service];
    E --> I;
    G --> J[User Service];
    H --> J;

    I --> K[(Database)];
    J --> K;
```

### Application Flow Explanation

This application is a Go-based expense-splitting service with a RESTful API. Here's a breakdown of its architecture and flow:

1.  **API Gateway**: All incoming HTTP requests are routed through an API gateway, which directs traffic to the appropriate handlers.

2.  **Authentication**:
    *   **Sign-Up**: New users can create an account by sending a `POST` request to `/api/v1/auth/signup`. The `AuthHandler` processes the request, and the `AuthService` handles the business logic of creating a new user in the database.
    *   **Login**: Registered users can log in by sending a `POST` request to `/api/v1/auth/login`. The `AuthHandler` validates the credentials, and the `AuthService` returns a JWT for session management.

3.  **User Management**:
    *   **Get Profile**: Authenticated users can retrieve their profile information with a `GET` request to `/api/v1/users/me`.
    *   **Update Profile**: Authenticated users can update their profile with a `PUT` request to `/api/v1/users/me`.

4.  **Services and Database**:
    *   The handlers delegate business logic to the corresponding services (`AuthService`, `UserService`).
    *   The services interact with the database to perform CRUD operations.

The remaining handlers for groups, expenses, and balances are not yet implemented, so those features are not functional.