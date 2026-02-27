# Emergency Backend

This is a simple Go backend that demonstrates JWT authentication and PostgreSQL database connection.

## Setup

1. **Environment variables**

   Create a `.env` file or set the following variables in your shell. The example below shows the MySQL-style variables you're currently using:

   ```env
   DB_USER=root
   DB_PASS=Laxman1103!@/
   DB_HOST=localhost:3306
   DB_NAME=app
   JWT_SECRET=your-very-secret-key
   PORT=8080
   ```

   (The code also allows you to switch to a `DATABASE_URL` style DSN if you prefer; adjust `db/db.go` accordingly.)

2. **Install dependencies**

   ```powershell
   go mod tidy
   ```

3. **Run the server**

   ```powershell
   go run main.go
   ```

## Endpoints

- `POST /login` - expects JSON with `username` and `password`, returns a JWT token.
- `GET /api/hello` - protected route, requires `Authorization: Bearer <token>` header.

## Database

This example currently connects to a **MySQL** server using the `go-sql-driver/mysql` driver. Ensure your database contains a `users` table such as:

```sql
CREATE TABLE users (
  id CHAR(36) PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL
);
```

Replace the password comparison logic with proper hashing (e.g. [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)).

If you switch back to PostgreSQL, you can change the import and DSN construction in `db/db.go` and update the environment variables accordingly.

## Notes

- This code uses `github.com/golang-jwt/jwt/v5` for token handling and `github.com/gorilla/mux` for routing.
- For real projects, handle errors and logging more robustly, and avoid storing plain text passwords.
