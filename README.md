# Tiny Site Backend Setup Guide: A Beginner-Friendly Approach

Welcome to the Tiny Site Backend setup guide! This guide will walk you through setting up the project on your local machine, even if you're new to web development. Let's break it down into simple steps.

## What You'll Need

Before we start, make sure you have these tools installed:

1. Go (version 1.21 or newer) - This is the programming language we use.
2. PostgreSQL (version 15 or newer) - Our database system.
3. pgAdmin (optional) - A helpful tool for managing your database.

Don't have these? No worries! Here's where to get them:
- Go: https://golang.org/dl/
- PostgreSQL: https://www.postgresql.org/download/
- pgAdmin: https://www.pgadmin.org/download/

## Setting Up Your Project

1. **Get the Code**
   Open your terminal and run:
   ```
   git clone https://github.com/Real-Dev-Squad/tiny-site-backend.git
   cd tiny-site-backend
   ```

2. **Set Up Your Environment**
   Create a file named `.env` in the project folder. This file will hold important settings. Here's a template:

   ```
   JWT_SECRET=<your-secret-key> e.g. my-secret-key
   JWT_VALIDITY_IN_HOURS=<validity-in-hours> e.g. 72
   JWT_ISSUER=<issuer> e.g. tiny-site-backend

   PORT=<port> e.g. 8000
   DOMAIN=<domain> e.g. localhost
   AUTH_REDIRECT_URL=<auth-redirect-url> e.g. http://localhost:3000
   ALLOWED_ORIGINS=<allowed-origins> e.g. http://localhost:3000

   DB_URL=postgresql://<username>:<password>@<host>:<port>/<database_name>?sslmode=disable
   e.g. postgresql://postgres:postgres@localhost:5432/tiny-site?sslmode=disable

   TEST_DB_URL=postgresql://<username>:<password>@<host>:<port>/<test_database_name>?sslmode=disable
   e.g. postgresql://postgres:postgres@localhost:5432/test-tiny-site?sslmode=disable

   DB_MAX_OPEN_CONNECTIONS=<max-open-connections> e.g. 10

   GOOGLE_CLIENT_ID=<google-client-id>
   GOOGLE_CLIENT_SECRET=<google-client-secret>
   GOOGLE_REDIRECT_URL="http://<domain>:<port>/v1/auth/google/callback"
   e.g. http://localhost:8000/v1/auth/google/callback
   ```

   Replace the placeholder values with your actual settings.

3. **Google OAuth Setup**
   To allow Google sign-in:
   - Go to https://console.developers.google.com/
   - Create a new project
   - Set up OAuth credentials for a web application
   - Add `http://localhost:8000` as an authorized JavaScript origin
   - Add `http://localhost:8000/auth/google/callback` as a redirect URI
   - Copy the provided client ID and secret to your `.env` file

4. **Set Up Your Database**
   - Open pgAdmin
   - Create a new database called `tiny-site`
   - If you choose a different name, update the `DB_URL` in your `.env` file

5. **Prepare Your Database**
   Run this command to set up your database tables:
   ```
   migrate -path ./migrations -database "postgresql://username:password@host:port/database_name?sslmode=disable" up
   ```
   Replace `username`, `password`, `host`, `port`, and `database_name` with your actual database details.

6. **Install Project Dependencies**
   Run:
   ```
   go mod download
   ```

7. **Start Your Server**
   You have two options:

   a. Basic start:
      ```
      go run main.go
      ```

   b. For automatic reloading during development:
      ```
      go get -u github.com/cosmtrek/air
      air
      ```

Great job! Your server should now be running at `http://localhost:8000`.

## Next Steps

- Explore the API endpoints in your browser or with a tool like Postman.
- You can also use [Tiny Site Frontend]([https://](https://github.com/Real-Dev-Squad/tiny-site-frontend)) to try out the full application.
- Happy coding!


## Need Help?

If you run into any issues, feel free to ask for help in the #tiny-site channel on our [Discord server](https://discord.com/channels/673083527624916993/785579160264769586). We're here to help you succeed! 
