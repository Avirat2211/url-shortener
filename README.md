# URL Shortener
A URL shortener build with Go (Gin Framework), Redis and PostgreSQL for the backend, along with a React+TypeScript frontend

## Features
- Shorten long URL's quickly
- Redirect short URL's to their original desinations
- Frontend built using React and TypeScript

## Tech Stack
### Backend:
- Go(Gin Framework)
- Redis
- PostgreSQL
### Frontend:
- React
- TypeScript
- Axios

## API Endpoints
1. Create Short URL

   POST `/create-short-url`

   Request Body:
   ```json
   {
     "long_url":"https://abcs.efg",
     "user_id":"123424525"
   }
   ```
   Response:
   ```json
   {
     "message":"short url created successfully"
     "short_url":"http://localhost:9808/aoijf341312",
   }
   ```

2. Retrieve Original URL

    GET `/:shortUrl`


## Setup Instructions
0. Clone the repository
   
   ```bash
   git clone https://github.com/Avirat2211/url-shortener.git
   cd url-shortener
   ```
### Backend Setup
1. Install Go and Redis
   
3. Install Dependencies:
   
   ```bash
   cd backend/
   go mod tidy
   ```
5. Setup .env file
   
   Create .env file in backend with:
   ```env
   ALLOWED_ORIGINS=
   PORT= 
   host     = 
   port     = 
   user     = 
   password = 
   dbname   = 
   Addr=    
   Password= 
   DB=
   ```
7. Run the server
   
    ```zsh
    go run main.go
    ```

### Frontend Setup
1. Navigate to frontend Directory
   
   ```bash
   cd frontend
   ```
3. Install Dependencies
   
   ```bash
   npm install
   ```
5. Setup .env file
   
   Create .env file in frontend with:
   
   ```env
   VITE_BACKEND_URL=http://localhost:9808
   ```
7. Start the frontend
   
   ```bash
   npm run dev
   ```
   
