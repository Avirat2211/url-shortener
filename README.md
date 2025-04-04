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
0. Ensure you have Docker and Docker compose installed
1. Setup .env file
   
   Create .env file in backend with:
   ```env
   ALLOWED_ORIGINS=http://localhost:5173 
   PORT=9808
   POSTGRES_HOST=postgres
   POSTGRES_PORT=5432
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=securepassword
   POSTGRES_DB=mydb
   REDIS_PORT=6379
   REDIS_HOST=redis
   Addr=redis:6379
   DB=0
   ```
2. Build and Run the containers

   Use Docker Compose to start the services:
   ```bash
   docker-compose up --build
   ```

   To stop the container:
   ```bash
   docker-compose down
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
   
