# YouTube Video Fetcher

## 📌 Project Overview
This project is a **YouTube Video Fetcher API** that continuously fetches the latest videos from YouTube based on a predefined search query. It stores video metadata in a PostgreSQL database and provides APIs to retrieve videos in a paginated format or search videos by title and description.

## 🚀 Features
- **Continuous Fetching**: Fetches the latest videos from YouTube every 10 seconds.
- **Database Storage**: Stores video metadata in PostgreSQL with proper indexing.
- **Paginated API**: Retrieves stored videos sorted in descending order of publish date.
- **Search API**: Searches videos by title and description.
- **Dockerized**: Can be deployed using Docker for easy setup.

## 🛠️ Tech Stack
- **Golang**: Backend API development
- **PostgreSQL**: Database for storing video metadata
- **YouTube Data API v3**: Fetching video details
- **Gin**: Web framework for handling API requests
- **GORM**: ORM for database interactions
- **Docker**: Containerization for easy deployment



## ⚙️ Setup and Installation
### 1️⃣ Clone the Repository
```sh
git clone [https://github.com/yourusername/youtube-video-fetcher.git](https://github.com/sushilkumar20/fampay-test.git)
cd fam
```

### 2️⃣ Set Up Environment Variables
Create a `.env` file in config directory and add:
```ini
YOUTUBE_API_KEYS=your_api_key_1  # Use multiple keys for quota management
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_db_user
POSTGRES_PASSWORD=your_db_password
POSTGRES_DATABASE=youtube_videos
PROBE_DURATION_IN_SECONDS=10  # Interval in seconds for fetching videos
YOUTUBE_BASE_URL=youtube_base_url
YOUTUBE_API_KEY=youtube_api_key
```

### 3️⃣ Run the Application
#### **Using Go Directly**
```sh
go mod tidy
go run main.go
```

#### **Using Docker**
```sh
docker build -t youtube-fetcher .
docker run -p 8080:8080 fam
```

#### **Using Docker Compose (with PostgreSQL)**
```sh
docker-compose up --build
```

## 📌 API Endpoints
### 1️⃣ **Get Paginated Videos**
- **Endpoint**: `GET /api/v1/videos?size=1&lastFetchedTime=2025-03-04T10:30:00Z`
- **lastFetchedTime**: published_at time of last video of previous page 
- **Description**: Returns a paginated list of videos sorted by published datetime.
- **Response**:
```json
{
  "videos": [
    {
      "title": "Cricket Match Highlights",
      "description": "Exciting match highlights...",
      "publishedAt": "2025-03-05T08:12:04Z",
      "thumbnails": "https://img.youtube.com/vi/abc123/default.jpg"
    }
  ]
}
```

### 2️⃣ **Search Videos**
- **Endpoint**: `GET /api/v1/videos/search?q=cricket&size=1&lastFetchedTime=2025-03-04T10:30:00Z`
- **Description**: Searches videos by title or description.
- **Response**:
```json
{
  "videos": [
    {
      "title": "How to play cricket",
      "description": "A tutorial on cricket basics...",
      "publishedAt": "2025-03-04T10:00:00Z"
    }
  ]
}
```

## 🏗️ Database Schema
```sql
CREATE TABLE videos (
    id SERIAL PRIMARY KEY,
    video_id TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    thumbnails TEXT
);

CREATE INDEX idx_published_at ON videos (published_at DESC);
CREATE INDEX idx_title_description ON videos USING GIN (to_tsvector('english', title || ' ' || description));
```

## 🛠️ Future Enhancements
- Improve **search API** with fuzzy matching for better results.
- Add **Redis caching** to optimize frequent queries.

