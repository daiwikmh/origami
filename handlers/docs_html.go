package handlers

import "github.com/gin-gonic/gin"

const docsHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Origami API - Developer Documentation</title>
    <style>
        @import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;700&family=Inter:wght@400;600;700&display=swap');

        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: 'Inter', sans-serif;
            background: #000000;
            color: #E1C4E9;
            line-height: 1.6;
        }

        .container { max-width: 1200px; margin: 0 auto; padding: 40px 20px; }

        header {
            background: #070709;
            padding: 40px 20px;
            text-align: center;
            border-bottom: 2px solid #E1C4E9;
            margin-bottom: 40px;
        }

        h1 {
            font-size: 2.5em;
            margin-bottom: 10px;
            color: #E1C4E9;
            font-weight: 700;
        }

        .subtitle {
            font-size: 1.1em;
            opacity: 0.8;
        }

        .nav {
            display: flex;
            gap: 20px;
            justify-content: center;
            margin-bottom: 40px;
            padding: 20px;
            background: #070709;
            border-radius: 8px;
        }

        .nav a {
            color: #E1C4E9;
            text-decoration: none;
            padding: 10px 20px;
            background: #232323;
            border-radius: 4px;
            transition: all 0.3s;
            font-weight: 600;
        }

        .nav a:hover {
            background: #E1C4E9;
            color: #000000;
        }

        .section {
            background: #070709;
            border: 1px solid #232323;
            border-radius: 8px;
            padding: 30px;
            margin-bottom: 30px;
        }

        h2 {
            color: #E1C4E9;
            margin-bottom: 20px;
            font-size: 1.8em;
            border-bottom: 2px solid #232323;
            padding-bottom: 10px;
        }

        h3 {
            color: #E1C4E9;
            margin: 20px 0 15px 0;
            font-size: 1.3em;
        }

        .endpoint {
            background: #000000;
            border-left: 4px solid #E1C4E9;
            padding: 20px;
            margin: 20px 0;
            border-radius: 4px;
        }

        .method {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 4px;
            font-weight: bold;
            font-size: 0.9em;
            margin-right: 10px;
            font-family: 'JetBrains Mono', monospace;
        }

        .method.get { background: #E1C4E9; color: #000000; }
        .method.post { background: #232323; color: #E1C4E9; }

        .path {
            font-family: 'JetBrains Mono', monospace;
            color: #E1C4E9;
            font-size: 1.1em;
        }

        pre {
            background: #000000;
            border: 1px solid #232323;
            border-radius: 4px;
            padding: 15px;
            overflow-x: auto;
            margin: 15px 0;
            font-family: 'JetBrains Mono', monospace;
            font-size: 0.9em;
        }

        code {
            font-family: 'JetBrains Mono', monospace;
            background: #232323;
            padding: 2px 6px;
            border-radius: 3px;
            font-size: 0.9em;
        }

        .param {
            margin: 10px 0;
            padding: 10px;
            background: #232323;
            border-radius: 4px;
        }

        .param-name {
            color: #E1C4E9;
            font-weight: 600;
            font-family: 'JetBrains Mono', monospace;
        }

        .required {
            color: #E1C4E9;
            font-size: 0.8em;
            background: #000000;
            padding: 2px 6px;
            border-radius: 3px;
            margin-left: 10px;
        }

        .response-example {
            background: #000000;
            border: 1px solid #232323;
            padding: 15px;
            border-radius: 4px;
            margin-top: 10px;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            margin: 20px 0;
        }

        th, td {
            padding: 12px;
            text-align: left;
            border-bottom: 1px solid #232323;
        }

        th {
            background: #232323;
            color: #E1C4E9;
            font-weight: 600;
        }

        .toc {
            background: #232323;
            padding: 20px;
            border-radius: 8px;
            margin-bottom: 30px;
        }

        .toc a {
            color: #E1C4E9;
            text-decoration: none;
            display: block;
            padding: 8px 0;
            border-left: 3px solid transparent;
            padding-left: 15px;
            transition: all 0.3s;
        }

        .toc a:hover {
            border-left-color: #E1C4E9;
            padding-left: 20px;
        }

        .note {
            background: #232323;
            border-left: 4px solid #E1C4E9;
            padding: 15px;
            margin: 15px 0;
            border-radius: 4px;
        }

        .example-badge {
            background: #E1C4E9;
            color: #000000;
            padding: 4px 10px;
            border-radius: 4px;
            font-size: 0.85em;
            font-weight: 600;
            margin-left: 10px;
        }
    </style>
</head>
<body>
    <header>
        <h1>Origami API Documentation</h1>
        <p class="subtitle">Developer Guide for Injective Intelligence API</p>
    </header>

    <div class="container">
        <div class="nav">
            <a href="/dashboard">Dashboard</a>
            <a href="/test">API Tester</a>
            <a href="/docs">Documentation</a>
        </div>

        <div class="toc">
            <h3>Table of Contents</h3>
            <a href="#auth">Authentication</a>
            <a href="#rate-limits">Rate Limits</a>
            <a href="#markets">Market Endpoints</a>
            <a href="#signals">Signal Endpoints</a>
            <a href="#nft">NFT Verification</a>
            <a href="#errors">Error Handling</a>
            <a href="#examples">Code Examples</a>
        </div>

        <div class="section" id="auth">
            <h2>🔐 Authentication</h2>
            <p>All API requests require authentication using Bearer tokens in the Authorization header.</p>

            <div class="endpoint">
                <h3>Authorization Header</h3>
                <pre>Authorization: Bearer YOUR_API_KEY</pre>
                <div class="note">
                    <strong>Note:</strong> Get your API key from the <a href="/dashboard" style="color: #E1C4E9;">Dashboard</a>
                </div>
            </div>

            <h3>Example Request</h3>
            <pre>
curl -H "Authorization: Bearer og_your_api_key_here" \\
     https://api.origami.com/origami/markets</pre>

            <h3>Authentication Errors</h3>
            <table>
                <tr>
                    <th>Status</th>
                    <th>Error</th>
                    <th>Description</th>
                </tr>
                <tr>
                    <td>401</td>
                    <td><code>missing api key</code></td>
                    <td>No Authorization header provided</td>
                </tr>
                <tr>
                    <td>401</td>
                    <td><code>invalid api key</code></td>
                    <td>API key is invalid or revoked</td>
                </tr>
            </table>
        </div>

        <div class="section" id="rate-limits">
            <h2>⏱️ Rate Limits</h2>
            <p>Rate limits are enforced per API key on a rolling 1-minute window.</p>

            <table>
                <tr>
                    <th>Plan</th>
                    <th>Requests/Minute</th>
                    <th>Status Code</th>
                </tr>
                <tr>
                    <td>Default</td>
                    <td>100</td>
                    <td>429 if exceeded</td>
                </tr>
                <tr>
                    <td>Custom</td>
                    <td>Configurable</td>
                    <td>429 if exceeded</td>
                </tr>
            </table>

            <h3>Rate Limit Response</h3>
            <pre>
{
  "error": "rate limit exceeded",
  "rate_limit": 100,
  "window": "1 minute"
}</pre>
        </div>

        <div class="section" id="markets">
            <h2>📊 Market Endpoints</h2>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets</span>
                <h3>Get All Markets</h3>
                <p>Returns all available spot markets from Injective.</p>
                <pre>
curl -H "Authorization: Bearer YOUR_API_KEY" \\
     https://api.origami.com/origami/markets</pre>
                <div class="response-example">
                    <strong>Response:</strong>
                    <pre>
{
  "markets": [
    {
      "marketId": "0x...",
      "ticker": "INJ/USDT",
      "baseDenom": "inj",
      "quoteDenom": "peggy0x...",
      "makerFeeRate": "0",
      "takerFeeRate": "0.001"
    }
  ]
}</pre>
                </div>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets/summary</span>
                <h3>Get Market Summary</h3>
                <p>Returns simplified market data.</p>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets/:id/analytics</span>
                <h3>Get Market Analytics</h3>
                <p>Returns comprehensive analytics for a specific market.</p>
                <div class="param">
                    <span class="param-name">:id</span>
                    <span class="required">REQUIRED</span>
                    <p>Market ID (e.g., 0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa)</p>
                </div>
                <pre>
curl -H "Authorization: Bearer YOUR_API_KEY" \\
     https://api.origami.com/origami/markets/MARKET_ID/analytics</pre>
                <div class="response-example">
                    <strong>Response:</strong>
                    <pre>
{
  "market_id": "0x...",
  "current_price": 17.65,
  "volume_24h": 2450000.00,
  "price_change_24h": 0.85,
  "price_change_24h_pct": 5.06,
  "volatility": 0.12,
  "liquidity_score": 125840.52,
  "trending_score": 45.3,
  "orderbook_depth": { ... }
}</pre>
                </div>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets/:id/liquidity</span>
                <h3>Get Market Liquidity</h3>
                <p>Returns liquidity metrics and orderbook depth.</p>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets/:id/volatility</span>
                <h3>Get Market Volatility</h3>
                <p>Returns volatility indicator calculated from recent price history.</p>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/markets/:id/depth</span>
                <h3>Get Orderbook Depth</h3>
                <p>Returns multi-level orderbook depth metrics.</p>
            </div>
        </div>

        <div class="section" id="signals">
            <h2>📈 Signal Endpoints</h2>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/signals/trending?limit=10</span>
                <h3>Get Trending Markets</h3>
                <p>Returns markets ranked by trending score (volume + volatility + price change).</p>
                <div class="param">
                    <span class="param-name">limit</span>
                    <span>OPTIONAL</span>
                    <p>Number of results (1-50, default: 10)</p>
                </div>
                <pre>
curl -H "Authorization: Bearer YOUR_API_KEY" \\
     "https://api.origami.com/origami/signals/trending?limit=5"</pre>
                <div class="response-example">
                    <strong>Response:</strong>
                    <pre>
{
  "markets": [
    {
      "market_id": "0x...",
      "symbol": "INJ/USDT",
      "score": 45.3,
      "volume_24h": 2450000,
      "volatility": 0.12,
      "price_change_pct": 5.06
    }
  ],
  "count": 5
}</pre>
                </div>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/signals/hot?limit=10</span>
                <h3>Get Hot Markets</h3>
                <p>Returns markets with highest trending scores.</p>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/signals/volatile?limit=10</span>
                <h3>Get Most Volatile Markets</h3>
                <p>Returns markets sorted by volatility.</p>
            </div>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/signals/volume?limit=10</span>
                <h3>Get Volume Leaders</h3>
                <p>Returns markets sorted by 24h trading volume.</p>
            </div>
        </div>

        <div class="section" id="nft">
            <h2>🎯 NFT Verification</h2>

            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/origami/nft/verify/:address</span>
                <span class="example-badge">NEW</span>
                <h3>Verify NFT Ownership</h3>
                <p>Check if a wallet address holds the specified NFT (0x816070929010A3D202D8A6B89f92BeE33B7e8769).</p>
                <div class="param">
                    <span class="param-name">:address</span>
                    <span class="required">REQUIRED</span>
                    <p>Wallet address to check (e.g., 0x0d79347CB8748FC6bbB1d425b99a4f44e63826c9)</p>
                </div>
                <pre>
curl -H "Authorization: Bearer YOUR_API_KEY" \\
     https://api.origami.com/origami/nft/verify/0x0d79347CB8748FC6bbB1d425b99a4f44e63826c9</pre>
                <div class="response-example">
                    <strong>Response:</strong>
                    <pre>
{
  "address": "0x0d79347CB8748FC6bbB1d425b99a4f44e63826c9",
  "contract_address": "0x816070929010A3D202D8A6B89f92BeE33B7e8769",
  "has_nft": true,
  "status": "YES"
}</pre>
                </div>
            </div>

            <div class="endpoint">
                <span class="method post">POST</span>
                <span class="path">/origami/nft/verify/batch</span>
                <span class="example-badge">NEW</span>
                <h3>Batch Verify NFT Ownership</h3>
                <p>Check multiple addresses at once (max 50 addresses).</p>
                <pre>
curl -X POST \\
     -H "Authorization: Bearer YOUR_API_KEY" \\
     -H "Content-Type: application/json" \\
     -d '{"addresses": ["0x...", "0x..."]}' \\
     https://api.origami.com/origami/nft/verify/batch</pre>
                <div class="response-example">
                    <strong>Response:</strong>
                    <pre>
{
  "results": [
    {
      "address": "0x...",
      "contract_address": "0x816070929010A3D202D8A6B89f92BeE33B7e8769",
      "has_nft": true,
      "status": "YES"
    }
  ],
  "count": 2
}</pre>
                </div>
            </div>
        </div>

        <div class="section" id="errors">
            <h2>🚨 Error Handling</h2>
            <table>
                <tr>
                    <th>Status Code</th>
                    <th>Error Type</th>
                    <th>Description</th>
                </tr>
                <tr>
                    <td>400</td>
                    <td>Bad Request</td>
                    <td>Invalid parameters or missing required fields</td>
                </tr>
                <tr>
                    <td>401</td>
                    <td>Unauthorized</td>
                    <td>Missing or invalid API key</td>
                </tr>
                <tr>
                    <td>404</td>
                    <td>Not Found</td>
                    <td>Resource not found (e.g., invalid market ID)</td>
                </tr>
                <tr>
                    <td>429</td>
                    <td>Rate Limit Exceeded</td>
                    <td>Too many requests in the time window</td>
                </tr>
                <tr>
                    <td>500</td>
                    <td>Internal Server Error</td>
                    <td>Server error - contact support if persistent</td>
                </tr>
            </table>
        </div>

        <div class="section" id="examples">
            <h2>💻 Code Examples</h2>

            <h3>JavaScript/Node.js</h3>
            <pre>
const API_KEY = 'og_your_api_key';
const BASE_URL = 'https://api.origami.com';

async function getTrendingMarkets() {
  const response = await fetch(BASE_URL + '/origami/signals/trending?limit=5', {
    headers: {
      'Authorization': 'Bearer ' + API_KEY
    }
  });
  const data = await response.json();
  return data.markets;
}

getTrendingMarkets().then(markets => console.log(markets));</pre>

            <h3>Python</h3>
            <pre>
import requests

API_KEY = 'og_your_api_key'
BASE_URL = 'https://api.origami.com'

headers = {
    'Authorization': f'Bearer {API_KEY}'
}

response = requests.get(f'{BASE_URL}/origami/signals/trending?limit=5', headers=headers)
markets = response.json()['markets']
print(markets)</pre>

            <h3>cURL</h3>
            <pre>
curl -H "Authorization: Bearer og_your_api_key" \\
     https://api.origami.com/origami/signals/trending?limit=5</pre>

            <h3>React Hook</h3>
            <pre>
import { useState, useEffect } from 'react';

function useOrigamiAPI(endpoint) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('https://api.origami.com' + endpoint, {
      headers: {
        'Authorization': 'Bearer ' + process.env.REACT_APP_API_KEY
      }
    })
    .then(res => res.json())
    .then(data => {
      setData(data);
      setLoading(false);
    });
  }, [endpoint]);

  return { data, loading };
}

// Usage
function TrendingMarkets() {
  const { data, loading } = useOrigamiAPI('/origami/signals/trending?limit=5');

  if (loading) return <div>Loading...</div>;
  return <div>{data.markets.map(m => <div key={m.market_id}>{m.symbol}</div>)}</div>;
}</pre>
        </div>

        <div class="section">
            <h2>📚 Additional Resources</h2>
            <ul style="line-height: 2;">
                <li><a href="/test" style="color: #E1C4E9;">Interactive API Tester</a></li>
                <li><a href="/dashboard" style="color: #E1C4E9;">API Key Dashboard</a></li>
                <li><a href="https://github.com/yourusername/origami" style="color: #E1C4E9;">GitHub Repository</a></li>
                <li><a href="mailto:support@origami.com" style="color: #E1C4E9;">Contact Support</a></li>
            </ul>
        </div>
    </div>
</body>
</html>
`

// ServeDocs serves the developer documentation
func ServeDocs(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, docsHTML)
}
