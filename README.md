# Elasticsearch + Kibana + Go API

A comprehensive demonstration project showcasing Elasticsearch integration with a Go backend API using Gin framework, complete with Kibana visualization and Docker Compose setup.

## 📋 Project Overview

This project provides a practical implementation of:
- **Elasticsearch**: Distributed search and analytics engine 
- **Kibana**: Visualization and monitoring platform 
- **Go API**: RESTful API built with Gin framework
- **Docker Compose**: Multi-container orchestration for Elasticsearch cluster and Kibana

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    Gin Go API (:8000)                    │
├─────────────────────────────────────────────────────────┤
│  GET /          - Home endpoint                         │
│  GET /init      - Initialize sample book data           │
│  GET /search    - Search books by title                 │
│  POST /insert   - Insert new documents                  │
└─────────────────────────────────────────────────────────┘
         │                              │
         └──────────────────────────────┘
                      │
        ┌─────────────┴─────────────┐
        │                           │
   ┌────▼──────┐            ┌──────▼────┐
   │ es01:9200 │ ◄─────────►│ es02:9200 │
   └───────────┘            └───────────┘
        │                           │
        └─────────────┬─────────────┘
                      │
              ┌───────▼────────┐
              │ Kibana :5601   │
              └────────────────┘
```

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose installed
- Go 1.16+ (if running without Docker)
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd doc_Kibana
   ```

2. **Start Elasticsearch and Kibana:**
   ```bash
   docker-compose up -d
   ```
   
   This will start:
   - 2 Elasticsearch nodes (es01, es02) on port 9200
   - Kibana on port 5601

3. **Install Go dependencies:**
   ```bash
   go mod download
   ```

4. **Run the Go API server:**
   ```bash
   go run main.go
   ```
   
   The API will be available at `http://localhost:8000`

## 📚 API Endpoints

### 1. Home
```
GET /
```
Base endpoint for health check.

### 2. Initialize Sample Data
```
GET /init
```
Generates and inserts sample book data into Elasticsearch. Uses `gofakeit` library to generate realistic book information with random genres.

**Genres:** Fantasy, Science, Mystery, Historical, Romance, Horror, Biography, Adventure

### 3. Search Books
```
GET /search?q=<query>
```
Search books by title using Elasticsearch full-text search.

**Parameters:**
- `q` (required): Search query string

**Example:**
```bash
curl "http://localhost:8000/search?q=adventure"
```

### 4. Insert Document
```
POST /insert?index=<index_name>
Content-Type: application/json

{
  "title": "Book Title",
  "author": "Author Name",
  "genre": "Fantasy"
}
```
Insert a new document into the specified Elasticsearch index.

**Parameters:**
- `index` (required): Target Elasticsearch index name

**Example:**
```bash
curl -X POST "http://localhost:8000/insert?index=books" \
  -H "Content-Type: application/json" \
  -d '{"title":"The Hobbit","author":"J.R.R. Tolkien","genre":"Fantasy"}'
```

## 🛠️ Configuration

### Elasticsearch Connection
Configured in `config/elastic.go`:
- **Address:** `http://localhost:9200`
- **Nodes:** 2 cluster nodes (es01, es02)

### Server
- **Port:** 8000
- **Framework:** Gin

## 📁 Project Structure

```
.
├── main.go              # Entry point
├── docker-compose.yml   # Docker services configuration
├── config/
│   └── elastic.go       # Elasticsearch client initialization
├── handlers/
│   ├── book.go          # Sample data generation
│   ├── search.go        # Search handler
│   ├── insert.go        # Insert handler
│   └── home.go          # Home handler
├── routes/
│   └── routes.go        # Route registration
├── go.mod               # Go module definition
└── README.md            # This file
```

## 🔍 Kibana Access

Access Kibana at: `http://localhost:5601`

**Default Credentials:** No authentication required in this setup

### Common Tasks in Kibana:
1. Create index pattern: Management → Index Patterns → Create Index Pattern
2. Explore data: Discover tab
3. Create visualizations: Visualize tab
4. Build dashboards: Dashboard tab

## 🧹 Cleanup

Stop all containers:
```bash
docker-compose down
```

Remove volumes (data):
```bash
docker-compose down -v
```

## 📦 Dependencies

### Go Packages
- `github.com/gin-gonic/gin` - HTTP framework
- `github.com/elastic/go-elasticsearch/v8` - Elasticsearch client
- `github.com/brianvoe/gofakeit/v7` - Fake data generation

## 🐛 Troubleshooting

### Connection Refused (9200)
- Ensure Elasticsearch container is running: `docker-compose ps`
- Wait for Elasticsearch to fully initialize (takes ~30 seconds)
- Check logs: `docker-compose logs es01`

### Kibana Not Accessible (5601)
- Verify Kibana container is running
- Check Elasticsearch connectivity: `docker-compose logs kibana`
- Ensure port 5601 is not in use

### API Server Won't Start
- Install dependencies: `go mod download`
- Check if port 8000 is already in use
- Verify Elasticsearch is accessible at localhost:9200
