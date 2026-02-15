# 🚀 Origami API - dApp Integration Guide

Complete guide for integrating Origami API into your decentralized application (dApp).

---

## 📋 Table of Contents

1. [Quick Start](#quick-start)
2. [Authentication Setup](#authentication-setup)
3. [Market Data Integration](#market-data-integration)
4. [NFT Verification](#nft-verification)
5. [React Integration](#react-integration)
6. [Vue.js Integration](#vuejs-integration)
7. [Web3 Integration](#web3-integration)
8. [Error Handling](#error-handling)
9. [Rate Limiting](#rate-limiting)
10. [Production Checklist](#production-checklist)

---

## 🎯 Quick Start

### Step 1: Get Your API Key

Visit the dashboard to generate an API key:
```
http://localhost:8080/dashboard
```

### Step 2: Test Your Key

```javascript
const API_KEY = 'og_your_api_key_here';
const BASE_URL = 'http://localhost:8080';

// Test authentication
fetch(`${BASE_URL}/origami/markets`, {
  headers: {
    'Authorization': `Bearer ${API_KEY}`
  }
})
.then(res => res.json())
.then(data => console.log('Markets:', data))
.catch(err => console.error('Error:', err));
```

---

## 🔐 Authentication Setup

### Basic Setup

```javascript
// config.js
export const API_CONFIG = {
  baseURL: 'http://localhost:8080',
  apiKey: process.env.REACT_APP_ORIGAMI_API_KEY,
  endpoints: {
    markets: '/origami/markets',
    trending: '/origami/signals/trending',
    nftVerify: '/origami/nft/verify',
  }
};
```

### API Client Class

```javascript
// api/origamiClient.js
class OrigamiAPI {
  constructor(apiKey, baseURL = 'http://localhost:8080') {
    this.apiKey = apiKey;
    this.baseURL = baseURL;
  }

  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`;
    const headers = {
      'Authorization': `Bearer ${this.apiKey}`,
      'Content-Type': 'application/json',
      ...options.headers
    };

    try {
      const response = await fetch(url, {
        ...options,
        headers
      });

      if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'API request failed');
      }

      return await response.json();
    } catch (error) {
      console.error('Origami API Error:', error);
      throw error;
    }
  }

  // Market endpoints
  async getMarkets() {
    return this.request('/origami/markets');
  }

  async getMarketSummary() {
    return this.request('/origami/markets/summary');
  }

  async getMarketAnalytics(marketId) {
    return this.request(`/origami/markets/${marketId}/analytics`);
  }

  async getMarketLiquidity(marketId) {
    return this.request(`/origami/markets/${marketId}/liquidity`);
  }

  async getMarketVolatility(marketId) {
    return this.request(`/origami/markets/${marketId}/volatility`);
  }

  async getOrderbookDepth(marketId) {
    return this.request(`/origami/markets/${marketId}/depth`);
  }

  // Signal endpoints
  async getTrending(limit = 10) {
    return this.request(`/origami/signals/trending?limit=${limit}`);
  }

  async getHotMarkets(limit = 10) {
    return this.request(`/origami/signals/hot?limit=${limit}`);
  }

  async getVolatileMarkets(limit = 10) {
    return this.request(`/origami/signals/volatile?limit=${limit}`);
  }

  async getVolumeLeaders(limit = 10) {
    return this.request(`/origami/signals/volume?limit=${limit}`);
  }

  // NFT verification
  async verifyNFT(address) {
    return this.request(`/origami/nft/verify/${address}`);
  }

  async verifyNFTBatch(addresses) {
    return this.request('/origami/nft/verify/batch', {
      method: 'POST',
      body: JSON.stringify({ addresses })
    });
  }

  // User endpoints
  async getMyUsage() {
    return this.request('/origami/me');
  }

  async getMyLimits() {
    return this.request('/origami/me/limits');
  }
}

export default OrigamiAPI;
```

---

## 📊 Market Data Integration

### Display Trending Markets

```javascript
import OrigamiAPI from './api/origamiClient';

const api = new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY);

async function displayTrendingMarkets() {
  try {
    const response = await api.getTrending(5);
    const markets = response.markets;

    markets.forEach(market => {
      console.log(`${market.symbol}: Score ${market.score.toFixed(2)}`);
      console.log(`  Volume: $${market.volume_24h.toLocaleString()}`);
      console.log(`  Volatility: ${(market.volatility * 100).toFixed(2)}%`);
      console.log(`  Price Change: ${market.price_change_pct.toFixed(2)}%`);
    });
  } catch (error) {
    console.error('Failed to fetch trending markets:', error);
  }
}
```

### Live Market Data Component (React)

```jsx
import React, { useState, useEffect } from 'react';
import OrigamiAPI from '../api/origamiClient';

const TrendingMarkets = () => {
  const [markets, setMarkets] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const api = new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY);

  useEffect(() => {
    const fetchMarkets = async () => {
      try {
        const data = await api.getTrending(10);
        setMarkets(data.markets || []);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchMarkets();
    const interval = setInterval(fetchMarkets, 30000); // Refresh every 30s

    return () => clearInterval(interval);
  }, []);

  if (loading) return <div>Loading trending markets...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="trending-markets">
      <h2>🔥 Trending Markets</h2>
      {markets.map(market => (
        <div key={market.market_id} className="market-card">
          <h3>{market.symbol}</h3>
          <div className="metrics">
            <span>Score: {market.score.toFixed(2)}</span>
            <span>Volume: ${market.volume_24h.toLocaleString()}</span>
            <span className={market.price_change_pct >= 0 ? 'positive' : 'negative'}>
              {market.price_change_pct >= 0 ? '+' : ''}{market.price_change_pct.toFixed(2)}%
            </span>
          </div>
        </div>
      ))}
    </div>
  );
};

export default TrendingMarkets;
```

---

## 🎯 NFT Verification

### Check if User Holds NFT

```javascript
import OrigamiAPI from './api/origamiClient';

const api = new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY);

async function checkNFTOwnership(walletAddress) {
  try {
    const result = await api.verifyNFT(walletAddress);

    console.log(`Address: ${result.address}`);
    console.log(`Has NFT: ${result.status}`); // "YES" or "NO"
    console.log(`Contract: ${result.contract_address}`);

    return result.has_nft; // boolean
  } catch (error) {
    console.error('NFT verification failed:', error);
    return false;
  }
}

// Usage
const userAddress = '0x0d79347CB8748FC6bbB1d425b99a4f44e63826c9';
const hasNFT = await checkNFTOwnership(userAddress);

if (hasNFT) {
  console.log('✅ User holds the required NFT!');
} else {
  console.log('❌ User does not hold the NFT');
}
```

### Batch NFT Verification

```javascript
async function checkMultipleAddresses(addresses) {
  try {
    const result = await api.verifyNFTBatch(addresses);

    result.results.forEach(item => {
      console.log(`${item.address}: ${item.status}`);
    });

    return result.results;
  } catch (error) {
    console.error('Batch verification failed:', error);
    return [];
  }
}

// Usage
const addresses = [
  '0x0d79347CB8748FC6bbB1d425b99a4f44e63826c9',
  '0x1234567890123456789012345678901234567890',
  '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd'
];

const results = await checkMultipleAddresses(addresses);
```

### NFT-Gated Access Component

```jsx
import React, { useState, useEffect } from 'react';
import { useWallet } from '@solana/wallet-adapter-react'; // or your Web3 provider
import OrigamiAPI from '../api/origamiClient';

const NFTGatedContent = ({ children }) => {
  const { publicKey } = useWallet();
  const [hasAccess, setHasAccess] = useState(false);
  const [loading, setLoading] = useState(true);

  const api = new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY);

  useEffect(() => {
    const verifyAccess = async () => {
      if (!publicKey) {
        setHasAccess(false);
        setLoading(false);
        return;
      }

      try {
        const result = await api.verifyNFT(publicKey.toString());
        setHasAccess(result.has_nft);
      } catch (error) {
        console.error('Access verification failed:', error);
        setHasAccess(false);
      } finally {
        setLoading(false);
      }
    };

    verifyAccess();
  }, [publicKey]);

  if (loading) {
    return <div>Verifying access...</div>;
  }

  if (!publicKey) {
    return <div>Please connect your wallet</div>;
  }

  if (!hasAccess) {
    return (
      <div className="access-denied">
        <h2>🔒 Access Restricted</h2>
        <p>You need to hold the required NFT to access this content.</p>
        <p>Contract: 0x816070929010A3D202D8A6B89f92BeE33B7e8769</p>
      </div>
    );
  }

  return <>{children}</>;
};

export default NFTGatedContent;
```

---

## ⚛️ React Integration

### Complete React Hook

```javascript
// hooks/useOrigamiAPI.js
import { useState, useEffect, useCallback } from 'react';
import OrigamiAPI from '../api/origamiClient';

export const useOrigamiAPI = () => {
  const [api] = useState(() => new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY));
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const request = useCallback(async (fn) => {
    setLoading(true);
    setError(null);

    try {
      const result = await fn(api);
      return result;
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  }, [api]);

  return {
    api,
    loading,
    error,
    request
  };
};

// Usage in component
import { useOrigamiAPI } from '../hooks/useOrigamiAPI';

function MarketDashboard() {
  const { api, loading, error, request } = useOrigamiAPI();
  const [markets, setMarkets] = useState([]);

  useEffect(() => {
    request(api => api.getTrending(10))
      .then(data => setMarkets(data.markets))
      .catch(console.error);
  }, [request]);

  return (
    <div>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error}</p>}
      {markets.map(m => <div key={m.market_id}>{m.symbol}</div>)}
    </div>
  );
}
```

---

## 🎨 Vue.js Integration

### Vue Composable

```javascript
// composables/useOrigami.js
import { ref } from 'vue';
import OrigamiAPI from '../api/origamiClient';

export function useOrigami() {
  const api = new OrigamiAPI(import.meta.env.VITE_ORIGAMI_API_KEY);
  const loading = ref(false);
  const error = ref(null);

  async function fetchData(fn) {
    loading.value = true;
    error.value = null;

    try {
      const result = await fn(api);
      return result;
    } catch (err) {
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  }

  return {
    api,
    loading,
    error,
    fetchData
  };
}
```

### Vue Component Example

```vue
<template>
  <div class="markets">
    <h2>Trending Markets</h2>
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">Error: {{ error }}</div>
    <div v-else>
      <div v-for="market in markets" :key="market.market_id" class="market-card">
        <h3>{{ market.symbol }}</h3>
        <p>Score: {{ market.score.toFixed(2) }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useOrigami } from '../composables/useOrigami';

const { api, loading, error, fetchData } = useOrigami();
const markets = ref([]);

onMounted(async () => {
  const data = await fetchData(api => api.getTrending(10));
  markets.value = data.markets;
});
</script>
```

---

## 🌐 Web3 Integration

### Combine with MetaMask/Web3

```javascript
import OrigamiAPI from './api/origamiClient';
import { ethers } from 'ethers';

class Web3OrigamiIntegration {
  constructor(apiKey) {
    this.origami = new OrigamiAPI(apiKey);
    this.provider = null;
    this.signer = null;
  }

  async connectWallet() {
    if (!window.ethereum) {
      throw new Error('MetaMask not installed');
    }

    this.provider = new ethers.providers.Web3Provider(window.ethereum);
    await this.provider.send('eth_requestAccounts', []);
    this.signer = this.provider.getSigner();

    const address = await this.signer.getAddress();
    return address;
  }

  async verifyCurrentUser() {
    const address = await this.signer.getAddress();
    const nftResult = await this.origami.verifyNFT(address);
    return nftResult.has_nft;
  }

  async getMarketsForUser() {
    const hasNFT = await this.verifyCurrentUser();

    if (hasNFT) {
      // Premium user - get all analytics
      return await this.origami.getMarkets();
    } else {
      // Regular user - get summary only
      return await this.origami.getMarketSummary();
    }
  }
}

// Usage
const integration = new Web3OrigamiIntegration(process.env.REACT_APP_ORIGAMI_API_KEY);
await integration.connectWallet();
const markets = await integration.getMarketsForUser();
```

---

## 🚨 Error Handling

### Comprehensive Error Handler

```javascript
class OrigamiAPIError extends Error {
  constructor(message, statusCode, details) {
    super(message);
    this.name = 'OrigamiAPIError';
    this.statusCode = statusCode;
    this.details = details;
  }
}

async function handleOrigamiRequest(requestFn) {
  try {
    return await requestFn();
  } catch (error) {
    if (error.message.includes('missing api key')) {
      throw new OrigamiAPIError('API key not provided', 401, error);
    } else if (error.message.includes('invalid api key')) {
      throw new OrigamiAPIError('Invalid API key', 401, error);
    } else if (error.message.includes('rate limit exceeded')) {
      throw new OrigamiAPIError('Rate limit exceeded - please wait', 429, error);
    } else if (error.message.includes('Market not found')) {
      throw new OrigamiAPIError('Market not found', 404, error);
    } else {
      throw new OrigamiAPIError('Unknown API error', 500, error);
    }
  }
}

// Usage
try {
  const data = await handleOrigamiRequest(() => api.getTrending());
  console.log(data);
} catch (error) {
  if (error instanceof OrigamiAPIError) {
    console.error(`Error ${error.statusCode}: ${error.message}`);

    if (error.statusCode === 429) {
      // Rate limited - retry after delay
      setTimeout(() => retryRequest(), 60000);
    }
  }
}
```

---

## ⏱️ Rate Limiting

### Rate Limit Helper

```javascript
class RateLimiter {
  constructor(api) {
    this.api = api;
    this.queue = [];
    this.processing = false;
  }

  async execute(fn) {
    return new Promise((resolve, reject) => {
      this.queue.push({ fn, resolve, reject });
      this.processQueue();
    });
  }

  async processQueue() {
    if (this.processing || this.queue.length === 0) return;

    this.processing = true;
    const { fn, resolve, reject } = this.queue.shift();

    try {
      const result = await fn();
      resolve(result);
    } catch (error) {
      if (error.message.includes('rate limit exceeded')) {
        // Re-queue the request
        this.queue.unshift({ fn, resolve, reject });
        // Wait 60 seconds before processing more
        await new Promise(r => setTimeout(r, 60000));
      } else {
        reject(error);
      }
    } finally {
      this.processing = false;
      // Process next item after a small delay
      setTimeout(() => this.processQueue(), 100);
    }
  }
}

// Usage
const api = new OrigamiAPI(API_KEY);
const limiter = new RateLimiter(api);

// These will be queued and executed with rate limiting
const result1 = await limiter.execute(() => api.getTrending());
const result2 = await limiter.execute(() => api.getHotMarkets());
```

---

## ✅ Production Checklist

### Environment Variables

```env
# .env
REACT_APP_ORIGAMI_API_KEY=og_your_production_key_here
REACT_APP_ORIGAMI_API_URL=https://api.yourorigami.com
```

### Best Practices

1. **Never expose API keys in client code**
   - Use environment variables
   - Use a backend proxy if possible

2. **Implement caching**
   ```javascript
   const cache = new Map();
   const CACHE_TTL = 30000; // 30 seconds

   async function getCachedData(key, fetchFn) {
     const cached = cache.get(key);
     if (cached && Date.now() - cached.timestamp < CACHE_TTL) {
       return cached.data;
     }

     const data = await fetchFn();
     cache.set(key, { data, timestamp: Date.now() });
     return data;
   }
   ```

3. **Handle offline scenarios**
4. **Monitor rate limits**
5. **Implement retry logic**
6. **Log errors properly**

---

## 🎉 Complete Example dApp

```jsx
import React, { useState, useEffect } from 'react';
import { useWallet } from '@solana/wallet-adapter-react';
import OrigamiAPI from './api/origamiClient';

function InjectiveDashboard() {
  const { publicKey } = useWallet();
  const [api] = useState(() => new OrigamiAPI(process.env.REACT_APP_ORIGAMI_API_KEY));
  const [hasNFT, setHasNFT] = useState(false);
  const [trending, setTrending] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function loadData() {
      try {
        // Verify NFT ownership
        if (publicKey) {
          const nftCheck = await api.verifyNFT(publicKey.toString());
          setHasNFT(nftCheck.has_nft);
        }

        // Load trending markets
        const markets = await api.getTrending(5);
        setTrending(markets.markets);
      } catch (error) {
        console.error('Error loading data:', error);
      } finally {
        setLoading(false);
      }
    }

    loadData();
    const interval = setInterval(loadData, 30000);
    return () => clearInterval(interval);
  }, [publicKey, api]);

  if (loading) return <div>Loading...</div>;

  return (
    <div className="dashboard">
      <header>
        <h1>Injective Market Dashboard</h1>
        {hasNFT && <span className="badge">🎯 Premium Member</span>}
      </header>

      <section className="trending">
        <h2>🔥 Trending Markets</h2>
        {trending.map(market => (
          <div key={market.market_id} className="market-card">
            <h3>{market.symbol}</h3>
            <div className="stats">
              <span>Score: {market.score.toFixed(2)}</span>
              <span>Volume: ${market.volume_24h.toLocaleString()}</span>
              <span className={market.price_change_pct >= 0 ? 'up' : 'down'}>
                {market.price_change_pct >= 0 ? '↑' : '↓'}
                {Math.abs(market.price_change_pct).toFixed(2)}%
              </span>
            </div>
          </div>
        ))}
      </section>
    </div>
  );
}

export default InjectiveDashboard;
```

---

## 📚 Additional Resources

- [Origami API Documentation](./ORIGAMI_PLATFORM.md)
- [API Tester](http://localhost:8080/test)
- [Dashboard](http://localhost:8080/dashboard)

---

**Built with ❤️ for Web3 Developers**
