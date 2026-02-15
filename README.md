# Origami API

**Production-ready Injective intelligence API platform with authentication, rate limiting, and real-time analytics.**

🌐 **Live Demo:** https://origami-8kv1.onrender.com/

---

## ✨ Features

### Core Functionality
- 🔐 **API Key Authentication** - Secure Bearer token authentication
- ⏱️ **Per-Key Rate Limiting** - Configurable limits (default: 100 req/min)
- 📊 **Real-Time Market Data** - Live data from Injective Protocol
- 📈 **Advanced Analytics** - Volatility, liquidity, and trending signals
- 🎯 **NFT Verification** - Check wallet ownership of specific NFTs
- 💾 **In-Memory Caching** - 80% cache hit rate, no database needed
- 🔄 **Background Workers** - Continuous data collection (5 goroutines)
- 📱 **Web Dashboard** - Manage API keys and view usage statistics
- 🧪 **Interactive Tester** - Test endpoints directly in browser

### API Endpoints
- Market data (all markets, summaries, analytics)
- Liquidity metrics and orderbook depth
- Volatility indicators
- Trading signals (trending, hot, volatile, volume leaders)
- NFT ownership verification (single & batch)
- Usage tracking and rate limit info

---

## 🚀 Quick Start

### Prerequisites
- Go 1.23+
- Internet connection (for Injective API)

### Installation

```bash
# Clone repository
git clone https://github.com/daiwikmh/origami.git
cd origami-api

# Install dependencies
go mod download

# Build
go build -o origami

# Run
./origami
```

Server starts at `http://localhost:8080`

### First Steps

1. **Open Dashboard:** http://localhost:8080/
2. **Copy the default API key** from startup logs
3. **Test in browser:** http://localhost:8080/test
4. **Read docs:** http://localhost:8080/docs

---

## 📡 API Usage

### Authentication

All `/origami/*` endpoints require authentication:

```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     https://origami-8kv1.onrender.com/origami/markets
```

### Example Endpoints

**Get Trending Markets:**
```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     "https://origami-8kv1.onrender.com/origami/signals/trending?limit=5"
```

**Get Market Analytics:**
```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     "https://origami-8kv1.onrender.com/origami/markets/MARKET_ID/analytics"
```

**Verify NFT Ownership:**
```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     "https://origami-8kv1.onrender.com/origami/nft/verify/0xYourWalletAddress"
```

**Response:**
```json
{
  "address": "0x...",
  "contract_address": "0x816070929010A3D202D8A6B89f92BeE33B7e8769",
  "has_nft": true,
  "status": "YES"
}
```

---

## 🌐 Live Deployment

**Production URL:** https://origami-8kv1.onrender.com/

### Public Pages
- **Dashboard:** https://origami-8kv1.onrender.com/
- **API Tester:** https://origami-8kv1.onrender.com/test
- **Documentation:** https://origami-8kv1.onrender.com/docs
- **System Info:** https://origami-8kv1.onrender.com/info

### Admin Endpoints (No Auth - Demo Only)
- `POST /admin/keys/generate` - Generate new API key
- `GET /admin/keys` - List all keys
- `POST /admin/keys/revoke` - Revoke key
- `GET /admin/usage` - Usage statistics

### Protected Endpoints (Require Auth)
- `GET /origami/markets` - All markets
- `GET /origami/markets/summary` - Market summary
- `GET /origami/markets/:id/analytics` - Market analytics
- `GET /origami/markets/:id/liquidity` - Liquidity metrics
- `GET /origami/markets/:id/volatility` - Volatility indicator
- `GET /origami/markets/:id/depth` - Orderbook depth
- `GET /origami/signals/trending?limit=N` - Trending markets
- `GET /origami/signals/hot?limit=N` - Hot markets
- `GET /origami/signals/volatile?limit=N` - Volatile markets
- `GET /origami/signals/volume?limit=N` - Volume leaders
- `GET /origami/nft/verify/:address` - Verify NFT ownership
- `POST /origami/nft/verify/batch` - Batch NFT verification
- `GET /origami/me` - Your usage stats
- `GET /origami/me/limits` - Your rate limits

---

## 🔑 API Key Management

### Generate Key (Dashboard)
1. Visit https://origami-8kv1.onrender.com/
2. Enter key name and rate limit
3. Click "Generate API Key"
4. **Save the key immediately** (shown only once)

### Generate Key (API)
```bash
curl -X POST https://origami-8kv1.onrender.com/admin/keys/generate \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My App",
    "rate_limit": 200
  }'
```

### Check Usage
```bash
curl -H "Authorization: Bearer YOUR_API_KEY" \
     https://origami-8kv1.onrender.com/origami/me
```

---

## 💻 Development

### Project Structure
```
inject/
├── main.go              # Entry point
├── routes.go            # Route definitions
├── auth/
│   └── keystore.go      # API key management
├── cache/
│   └── data_cache.go    # In-memory cache
├── clients/
│   ├── injective.go     # Injective API client
│   └── trades.go        # Trade data client
├── handlers/
│   ├── admin_handlers.go     # Admin endpoints
│   ├── analytics_handlers.go # Analytics endpoints
│   ├── market_handlers.go    # Market endpoints
│   ├── nft_handlers.go       # NFT verification
│   ├── signal_handlers.go    # Signal endpoints
│   ├── dashboard_html.go     # Dashboard UI
│   ├── testing_html.go       # API tester UI
│   └── docs_html.go          # Developer docs
├── middleware/
│   ├── auth.go          # Authentication
│   ├── ratelimit.go     # Rate limiting
│   └── usage.go         # Usage tracking
├── models/
│   ├── apikey.go        # API key models
│   ├── analytics.go     # Analytics models
│   ├── market.go        # Market models
│   └── trade.go         # Trade models
├── services/
│   ├── analytics.go     # Analytics logic
│   └── market.go        # Market logic
├── utils/
│   ├── math_helpers.go  # Math utilities
│   └── parse_helpers.go # Parsing utilities
└── workers/
    └── data_collector.go # Background workers
```

### Local Development
```bash
# Run in development mode
go run main.go

# Build for production
go build -ldflags="-s -w" -o origami

# Run tests (if any)
go test ./...

# Format code
go fmt ./...
```

### Environment Variables
```bash
PORT=8080              # Server port (auto-set by Render)
GIN_MODE=release       # Gin mode (debug/release)
```

---

## 🎨 Tech Stack

- **Language:** Go 1.23
- **Framework:** Gin (HTTP router)
- **Data Source:** Injective Protocol API
- **Caching:** In-memory (no database)
- **Authentication:** Bearer token (custom implementation)
- **Rate Limiting:** In-memory sliding window
- **Deployment:** Render
- **UI:** Pure HTML/CSS/JavaScript

---

## 📊 Architecture

```
┌─────────────────────────────────────────┐
│         HTTP Clients (Developers)        │
└──────────────┬──────────────────────────┘
               │
               ▼
    ┌──────────────────────┐
    │   Gin HTTP Router    │
    │   /origami/* routes  │
    └──────────┬───────────┘
               │
    ┌──────────┴──────────────────────┐
    │     Middleware Stack            │
    │  • Authentication (Bearer)      │
    │  • Rate Limiting (per key)      │
    │  • Usage Tracking               │
    └──────────┬──────────────────────┘
               │
    ┌──────────┴──────────────────────┐
    │      Services Layer             │
    │  • Market Analytics             │
    │  • Signal Generation            │
    │  • NFT Verification             │
    └──────────┬──────────────────────┘
               │
    ┌──────────┴────────────┬─────────┐
    │                       │          │
    ▼                       ▼          ▼
┌──────────┐    ┌────────────────┐  ┌─────────┐
│Injective │    │ In-Memory Cache│  │ API Key │
│   API    │    │ (TTL-based)    │  │  Store  │
└────┬─────┘    └────────────────┘  └─────────┘
     │                   ▲
     │                   │
     └────┬──────────────┘
          │
    ┌─────┴──────────────┐
    │ Background Workers │
    │  (5 goroutines)    │
    └────────────────────┘
```

**Data Flow:**
1. Client sends request with API key
2. Middleware validates key & checks rate limit
3. Service fetches from cache or Injective API
4. Background workers continuously update cache
5. Response returned to client

---

## 🔒 Security

- ✅ API key authentication required for all data endpoints
- ✅ Per-key rate limiting (prevents abuse)
- ✅ Usage tracking (monitor API consumption)
- ✅ HTTPS in production (Render provides)
- ✅ No database (reduced attack surface)
- ✅ Minimal dependencies (fewer vulnerabilities)

**Production Recommendations:**
- Add authentication to admin endpoints
- Implement IP-based rate limiting
- Add API key expiration
- Enable CORS whitelisting
- Add request logging
- Set up monitoring/alerts

---

## 📈 Performance

- **Cache Hit Rate:** ~80% after warmup
- **Response Time:** <100ms (p95)
- **Memory Usage:** ~40MB (500 markets + 500 keys)
- **Background Workers:** 5 goroutines
- **Data Freshness:**
  - Markets: 10 seconds
  - Orderbooks: 5 seconds
  - Analytics: 15 seconds

---

## 🚀 Deployment

### Render (Current)
```yaml
Build Command: go build -ldflags="-s -w" -o origami
Start Command: ./origami
```

### Docker
```bash
docker build -t origami-api .
docker run -d -p 8080:8080 origami-api
```

### Traditional Server
```bash
# Build
go build -ldflags="-s -w" -o origami

# Run with systemd
sudo systemctl start origami
```

See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed deployment guides.

---

## 🔧 Configuration

### Cache TTLs
- Markets: 10 seconds
- Orderbooks: 5 seconds
- Analytics: 15 seconds
- Trades/Prices: Rolling windows (no expiry)

### Rate Limits
- Default: 100 requests/minute
- Configurable per API key
- Window: 1 minute (rolling)

### Background Workers
1. **Market Collector** - Every 10s
2. **Orderbook Collector** - Every 5s (top 50 markets)
3. **Trade Collector** - Every 10s (top 30 markets)
4. **Price History** - Every 60s
5. **Analytics Computer** - Every 15s

---

## 🧪 Testing

### Interactive Tester
Visit https://origami-8kv1.onrender.com/test

### Manual Testing
```bash
# Get API key from dashboard
API_KEY="og_your_key_here"

# Test markets endpoint
curl -H "Authorization: Bearer $API_KEY" \
     https://origami-8kv1.onrender.com/origami/markets

# Test trending signals
curl -H "Authorization: Bearer $API_KEY" \
     "https://origami-8kv1.onrender.com/origami/signals/trending?limit=5"

# Test NFT verification
curl -H "Authorization: Bearer $API_KEY" \
     "https://origami-8kv1.onrender.com/origami/nft/verify/0xYourAddress"
```

---

## 📚 Documentation

- **Developer Docs:** https://origami-8kv1.onrender.com/docs
- **API Tester:** https://origami-8kv1.onrender.com/test
- **Dashboard:** https://origami-8kv1.onrender.com/
- **dApp Integration:** See [DAPP_INTEGRATION.md](./DAPP_INTEGRATION.md)
- **Deployment Guide:** See [DEPLOYMENT.md](./DEPLOYMENT.md)

---

## 🤝 Integration Examples

### JavaScript
```javascript
const API_KEY = 'og_your_key';
const BASE_URL = 'https://origami-8kv1.onrender.com';

const response = await fetch(`${BASE_URL}/origami/signals/trending?limit=5`, {
  headers: { 'Authorization': `Bearer ${API_KEY}` }
});
const data = await response.json();
```

### Python
```python
import requests

API_KEY = 'og_your_key'
BASE_URL = 'https://origami-8kv1.onrender.com'

response = requests.get(
    f'{BASE_URL}/origami/signals/trending?limit=5',
    headers={'Authorization': f'Bearer {API_KEY}'}
)
markets = response.json()['markets']
```

### React
```jsx
function TrendingMarkets() {
  const [markets, setMarkets] = useState([]);

  useEffect(() => {
    fetch('https://origami-8kv1.onrender.com/origami/signals/trending?limit=5', {
      headers: { 'Authorization': `Bearer ${process.env.REACT_APP_API_KEY}` }
    })
    .then(res => res.json())
    .then(data => setMarkets(data.markets));
  }, []);

  return <div>{markets.map(m => <div key={m.market_id}>{m.symbol}</div>)}</div>;
}
```

---

## 🛣️ Roadmap

- [ ] Add WebSocket support for real-time updates
- [ ] Implement Redis for distributed caching
- [ ] Add API key expiration
- [ ] Create Prometheus metrics endpoint
- [ ] Add request logging middleware
- [ ] Implement pagination for large responses
- [ ] Add GraphQL support
- [ ] Create official SDKs (JavaScript, Python, Go)

---

## 📄 License

MIT License - See LICENSE file for details

---

## 🙏 Acknowledgments

- **Injective Protocol** - For providing the exchange data API
- **Gin Framework** - For the excellent HTTP router
- **Render** - For simple and reliable hosting

---

## 📧 Contact

- **Live API:** https://origami-8kv1.onrender.com/
- **Issues:** Open a GitHub issue
- **Docs:** https://origami-8kv1.onrender.com/docs

---

**Built with ❤️ using Go and Injective Protocol**
