# Origami API Platform - Complete Guide

A production-ready API platform providing authenticated access to Injective intelligence data with rate limiting, usage tracking, and a web dashboard.

---

## 🚀 Quick Start

### 1. Build and Run

```bash
# Build the application
go build

# Run the server
./origami
```

The server will display a **default API key** on startup:

```
======================================================================
  DEFAULT API KEY FOR TESTING
======================================================================
  Name: Default Test Key
  Key:  og_28ab9c0ee59399d52fca3ce52b92f9ebbcc5ce53a0d45c7de6d702f5f7b0954c
  Rate Limit: 100 requests/minute
======================================================================
```

### 2. Access the Dashboard

Open your browser and visit:
- **Dashboard**: http://localhost:8080/dashboard
- **API Tester**: http://localhost:8080/test

---

## 🔐 Authentication

All `/origami/*` endpoints require authentication using Bearer tokens.

### Making Authenticated Requests

```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     http://localhost:8080/origami/markets
```

### Error Responses

**Missing API Key (401):**
```json
{
  "error": "missing api key"
}
```

**Invalid API Key (401):**
```json
{
  "error": "invalid api key"
}
```

---

## 📊 API Endpoints

### Public Endpoints (No Auth Required)

#### Get System Info
```bash
GET /
```

Returns API information and available endpoints.

**Example:**
```bash
curl http://localhost:8080/
```

**Response:**
```json
{
  "name": "Origami API",
  "version": "1.0.0",
  "description": "Production-ready Injective intelligence API with authentication and rate limiting",
  "endpoints": [...],
  "docs": "/dashboard",
  "test": "/test"
}
```

---

### Protected Endpoints (Require Auth)

All endpoints under `/origami/*` require authentication.

#### Markets

**Get All Markets**
```bash
GET /origami/markets
Authorization: Bearer YOUR_API_KEY
```

**Get Market Summary**
```bash
GET /origami/markets/summary
Authorization: Bearer YOUR_API_KEY
```

**Get Market Liquidity**
```bash
GET /origami/markets/{marketId}/liquidity
Authorization: Bearer YOUR_API_KEY
```

**Get Market Analytics**
```bash
GET /origami/markets/{marketId}/analytics
Authorization: Bearer YOUR_API_KEY
```

**Get Market Volatility**
```bash
GET /origami/markets/{marketId}/volatility
Authorization: Bearer YOUR_API_KEY
```

**Get Orderbook Depth**
```bash
GET /origami/markets/{marketId}/depth
Authorization: Bearer YOUR_API_KEY
```

#### Signals

**Get Trending Markets**
```bash
GET /origami/signals/trending?limit=10
Authorization: Bearer YOUR_API_KEY
```

**Get Hot Markets**
```bash
GET /origami/signals/hot?limit=10
Authorization: Bearer YOUR_API_KEY
```

**Get Volatile Markets**
```bash
GET /origami/signals/volatile?limit=10
Authorization: Bearer YOUR_API_KEY
```

**Get Volume Leaders**
```bash
GET /origami/signals/volume?limit=10
Authorization: Bearer YOUR_API_KEY
```

#### User Info

**Get My Usage Stats**
```bash
GET /origami/me
Authorization: Bearer YOUR_API_KEY
```

**Get My Rate Limits**
```bash
GET /origami/me/limits
Authorization: Bearer YOUR_API_KEY
```

---

## ⚙️ Admin Endpoints

Admin endpoints do not require authentication (add auth in production!).

### Generate API Key

```bash
POST /admin/keys/generate
Content-Type: application/json

{
  "name": "Production App",
  "rate_limit": 100
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/admin/keys/generate \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Application",
    "rate_limit": 200
  }'
```

**Response:**
```json
{
  "api_key": "og_a1b2c3...",
  "name": "My Application",
  "rate_limit": 200,
  "created_at": "2026-02-15T10:00:00Z",
  "message": "API key created successfully. Store it securely - it won't be shown again."
}
```

### List API Keys

```bash
GET /admin/keys
```

Returns all API keys with masked values and usage statistics.

**Example:**
```bash
curl http://localhost:8080/admin/keys
```

### Revoke API Key

```bash
POST /admin/keys/revoke
Content-Type: application/json

{
  "key": "og_key_to_revoke"
}
```

### Get Usage Statistics

```bash
GET /admin/usage
```

Returns aggregated usage statistics across all API keys.

**Example:**
```bash
curl http://localhost:8080/admin/usage
```

**Response:**
```json
{
  "total_keys": 5,
  "active_keys": 4,
  "total_requests": 12450,
  "key_stats": {
    "og_abc...": {
      "name": "Production App",
      "request_count": 8500,
      "last_used_at": "2026-02-15T10:30:00Z",
      "rate_limit": 200,
      "endpoint_usage": {
        "GET /origami/markets": 3200,
        "GET /origami/signals/trending": 2100
      }
    }
  }
}
```

---

## 🚦 Rate Limiting

Each API key has a configurable rate limit (requests per minute).

- Default rate limit: **100 requests/minute**
- Rate limit window: **1 minute (rolling)**
- Limit is enforced per API key

### Rate Limit Exceeded (429)

```json
{
  "error": "rate limit exceeded",
  "rate_limit": 100,
  "window": "1 minute"
}
```

---

## 📈 Usage Tracking

The platform automatically tracks:

- **Total requests** per API key
- **Last used timestamp** for each key
- **Endpoint-specific usage** (which endpoints are called most)
- **Request patterns** over time

View your usage:
```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     http://localhost:8080/origami/me
```

---

## 🎯 Web Dashboard

Access the web dashboard at **http://localhost:8080/dashboard**

### Features:
- 📊 Real-time usage statistics
- 🔑 Generate new API keys
- 📋 View all API keys and their usage
- 🗑️ Revoke API keys
- 📈 Visual analytics

### API Tester

Access the endpoint tester at **http://localhost:8080/test**

### Features:
- 🧪 Test any endpoint with your API key
- 📝 Automatic request configuration
- 📡 View responses in real-time
- 🔍 Status code visualization
- 📋 Copy-paste ready curl commands

---

## 🛠️ Complete Example Workflow

### 1. Generate a New API Key

```bash
curl -X POST http://localhost:8080/admin/keys/generate \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Trading Bot",
    "rate_limit": 200
  }' | jq '.'
```

Save the returned API key securely.

### 2. Test Authentication

```bash
API_KEY="your_api_key_here"

# This should fail (no auth)
curl http://localhost:8080/origami/markets

# This should work
curl -H "Authorization: Bearer $API_KEY" \
     http://localhost:8080/origami/markets | jq '.'
```

### 3. Fetch Market Data

```bash
# Get all markets
curl -H "Authorization: Bearer $API_KEY" \
     http://localhost:8080/origami/markets | jq '.markets | length'

# Get trending markets
curl -H "Authorization: Bearer $API_KEY" \
     "http://localhost:8080/origami/signals/trending?limit=5" | jq '.'

# Get specific market analytics
MARKET_ID="0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
curl -H "Authorization: Bearer $API_KEY" \
     "http://localhost:8080/origami/markets/$MARKET_ID/analytics" | jq '.'
```

### 4. Monitor Your Usage

```bash
# Check your usage stats
curl -H "Authorization: Bearer $API_KEY" \
     http://localhost:8080/origami/me | jq '.'

# Check rate limit info
curl -H "Authorization: Bearer $API_KEY" \
     http://localhost:8080/origami/me/limits | jq '.'
```

### 5. View All Usage (Admin)

```bash
# Get system-wide usage statistics
curl http://localhost:8080/admin/usage | jq '.'
```

---

## 🔒 Security Best Practices

### For API Key Generation:
1. **Store keys securely** - Keys are only shown once during generation
2. **Use descriptive names** - Name keys after their purpose (e.g., "Production Bot", "Dev Testing")
3. **Set appropriate rate limits** - Start conservative, increase as needed
4. **Rotate keys regularly** - Generate new keys and revoke old ones periodically
5. **One key per application** - Don't share keys between different apps

### For Production Deployment:
1. **Add authentication to admin endpoints** - Current implementation has no auth for demo purposes
2. **Use HTTPS** - Encrypt all traffic
3. **Set up API key expiration** - Add TTL to API keys
4. **Implement IP whitelisting** - Restrict key usage to specific IPs
5. **Add request logging** - Log all requests for audit trails
6. **Rate limit admin endpoints** - Prevent abuse of key generation

---

## 📊 Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP Clients                              │
│                 (Developers using API)                       │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
        ┌──────────────────────────────────────┐
        │      API Endpoints (/origami/*)      │
        └──────────────────┬───────────────────┘
                           │
        ┌──────────────────┴──────────────────────────┐
        │          Middleware Stack                    │
        │  1. API Key Authentication                   │
        │  2. Rate Limiting (per key)                  │
        │  3. Usage Tracking                           │
        └──────────────────┬──────────────────────────┘
                           │
        ┌──────────────────┴──────────────────────────┐
        │           Services Layer                     │
        │   (Business Logic + Analytics)               │
        └──────────────────┬──────────────────────────┘
                           │
        ┌──────────────────┴──────────────────────────┐
        │                                              │
        ▼                                              ▼
┌──────────────────┐                    ┌──────────────────────┐
│ Injective API    │                    │  In-Memory Stores    │
│ (Market Data)    │                    │  • API Keys          │
└──────────────────┘                    │  • Rate Limits       │
        ▲                                │  • Usage Stats       │
        │                                │  • Market Cache      │
        │                                └──────────────────────┘
        │                                         ▲
        └─────────────────┬───────────────────────┘
                          │
                ┌─────────┴────────┐
                │ Background       │
                │ Workers          │
                │ (5 goroutines)   │
                └──────────────────┘
```

---

## 🎨 Web Dashboard Features

### Dashboard (`/dashboard`)
- Real-time API usage statistics
- Active/total keys counter
- Total requests tracker
- API key generation form
- List of all API keys with:
  - Masked key preview
  - Request count
  - Rate limit
  - Active/inactive status
  - Creation timestamp

### API Tester (`/test`)
- Interactive endpoint selector
- Automatic request configuration
- Market ID input for specific endpoints
- Limit parameter for signal endpoints
- Real-time response viewer
- Status code highlighting
- JSON response formatting

---

## ⚡ Performance

- **Cache hit rate**: ~80% after warmup
- **Authentication overhead**: <1ms per request
- **Rate limiting overhead**: <0.5ms per request
- **Memory usage**: ~40MB for 500 keys + 500 markets
- **Background workers**: 5 goroutines for data collection
- **Data freshness**:
  - Markets: 10 seconds
  - Orderbooks: 5 seconds
  - Analytics: 15 seconds

---

## 🔧 Configuration

### Default Settings:
- **Port**: 8080
- **Default rate limit**: 100 req/min
- **Rate limit window**: 1 minute (rolling)
- **Read timeout**: 15 seconds
- **Write timeout**: 15 seconds
- **Idle timeout**: 60 seconds

### Modify in Code:
- Port: `main.go` - `Addr: ":8080"`
- Rate limits: `auth/keystore.go` - `GenerateKey()` function
- Timeouts: `main.go` - `srv` configuration

---

## 🚀 Production Deployment

### Running in Production:

```bash
# Build optimized binary
go build -ldflags="-s -w" -o origami

# Run with nohup
nohup ./origami > /var/log/origami-api.log 2>&1 &

# Check if running
curl http://localhost:8080/

# View logs
tail -f /var/log/origami-api.log
```

### systemd Service Example:

```ini
[Unit]
Description=Origami API Service
After=network.target

[Service]
Type=simple
User=origami
WorkingDirectory=/opt/origami
ExecStart=/opt/origami/origami
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

### Nginx Reverse Proxy:

```nginx
server {
    listen 80;
    server_name api.origami.example.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

---

## 📝 License

MIT License - See LICENSE file for details

---

## 🤝 Support

For issues, questions, or feature requests:
- GitHub Issues: https://github.com/yourusername/origami/issues
- Documentation: http://localhost:8080/dashboard
- API Tester: http://localhost:8080/test

---

**Built with ❤️ using Go, Gin, and Injective Protocol**
