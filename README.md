# Tiny Site Backend Setup Guide

This README provides step-by-step instructions for setting up the Tiny Site Backend on your local development environment. The Tiny Site Backend is built using the Go programming language, PostgreSQL for the database, and Google authentication for authentication. Before you begin, ensure that you have the following prerequisites installed on your system:

1. Go 1.21 or higher
2. PostgreSQL 15 or higher
3. pgAdmin (optional but recommended for database management)

## Prerequisites

### Go 1.21

Make sure you have Go version 1.21 or higher installed on your system. You can download and install Go from the official website: [Go Downloads](https://golang.org/dl/)

### PostgreSQL

You will need PostgreSQL as the database for the Tiny Site Backend. You can download and install PostgreSQL from the official website: [PostgreSQL Downloads](https://www.postgresql.org/download/)

### pgAdmin (Optional)

pgAdmin is a popular database management tool for PostgreSQL. It is recommended to install pgAdmin for easier database management. You can download and install pgAdmin from the official website: [pgAdmin Downloads](https://www.pgadmin.org/download/)

## Setup Instructions

1. Clone the Tiny Site Backend repository from GitHub:

   ```bash
   git clone https://github.com/Real-Dev-Squad/tiny-site-backend.git
   ```

2. Navigate to the project directory:

   ```bash
   cd tiny-site-backend
   ```

3. Create a `.env` file for your local development environment. You can use the sample `.env` file provided in the `environments` directory as a template. Copy the `dev.env` file to `.env`:

   ```bash
   cp environments/dev.env .env
   ```

   Edit the `.env` file and configure the environment variables according to your local setup. Make sure to set the database connection details and Google OAuth credentials.

4. **Setting Up Google OAuth**

   - Go to the [Google Developers Console](https://console.developers.google.com/).

   - Click on "Select a Project" in the top navigation and create a new project if you don't have one.

   - In the left sidebar, click on "APIs & Services" and then "Credentials."

   - Click on "Create Credentials" and select "OAuth client ID."

   - Choose "Web application" as the application type.

   - In the "Authorized JavaScript origins" field, add `http://localhost:8000` since this is your local development environment.

   - In the "Authorized redirect URIs" field, add the URI where Google should redirect after authentication. For local development, use `http://localhost:8000/auth/google/callback`.

   - Click the "Create" button.

   - You will be provided with a client ID and client secret. Add these values to your `.env` file that you created earlier:

     ```
     GOOGLE_CLIENT_ID=your-client-id
     GOOGLE_CLIENT_SECRET=your-client-secret
     ```

   - Save the changes to your `.env` file.

5. Create the PostgreSQL database for the project. You can use pgAdmin or command-line tools to create the database. Update the `.env` file with the database connection details.

    ```
    DB_URL=postgres://username:password@host:port/database_name
    TEST_DB_URL=postgres://username:password@host:port/test_database_name
    ```

6. Install the Go project dependencies:

   ```bash
   go mod download
   ```

7. Build and run the Tiny Site Backend:

   ```bash
   go run main.go
   ```

- Use Air for hot reloading

    ```bash
    go get -u github.com/cosmtrek/air
    air
    ```

    The backend server should now be running at `http://localhost:8000s`.

8. Access the Tiny Site Backend API via a web browser or API client, and start developing your application.

## Usage

You can access the Tiny Site Backend API at `http://localhost:8000` once the server is running. Refer to the project's documentation or API endpoints for more information on how to interact with the backend.
