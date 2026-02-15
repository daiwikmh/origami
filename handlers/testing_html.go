package handlers

const testingHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Origami API - Endpoint Tester</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif; background: #0f172a; color: #e2e8f0; line-height: 1.6; }
        .container { max-width: 1400px; margin: 0 auto; padding: 20px; }
        header { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 40px 20px; text-align: center; border-radius: 10px; margin-bottom: 30px; }
        h1 { font-size: 2.5em; margin-bottom: 10px; }
        .subtitle { opacity: 0.9; font-size: 1.1em; }
        .grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; }
        .card { background: #1e293b; border-radius: 10px; padding: 25px; box-shadow: 0 4px 6px rgba(0,0,0,0.3); }
        .card h2 { margin-bottom: 20px; color: #a78bfa; border-bottom: 2px solid #334155; padding-bottom: 10px; }
        .btn { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; border: none; padding: 12px 24px; border-radius: 6px; cursor: pointer; font-size: 1em; transition: transform 0.2s; }
        .btn:hover { transform: translateY(-2px); }
        .btn-secondary { background: #334155; }
        input, select, textarea { width: 100%; padding: 12px; margin: 8px 0; border: 2px solid #334155; border-radius: 6px; background: #0f172a; color: #e2e8f0; font-size: 1em; font-family: inherit; }
        input:focus, select:focus, textarea:focus { outline: none; border-color: #667eea; }
        textarea { font-family: 'Courier New', monospace; min-height: 200px; }
        .endpoint-list { list-style: none; }
        .endpoint-item { background: #0f172a; padding: 15px; margin: 10px 0; border-radius: 6px; cursor: pointer; transition: border-left 0.2s; border-left: 4px solid transparent; }
        .endpoint-item:hover { border-left-color: #667eea; }
        .endpoint-item.selected { border-left-color: #a78bfa; background: #1e293b; }
        .method { display: inline-block; padding: 4px 8px; border-radius: 4px; font-weight: bold; font-size: 0.9em; margin-right: 10px; }
        .method.get { background: #10b981; color: white; }
        .method.post { background: #3b82f6; color: white; }
        .nav { display: flex; gap: 15px; justify-content: center; margin: 20px 0; }
        .nav a { color: #a78bfa; text-decoration: none; padding: 10px 20px; border-radius: 6px; background: #1e293b; transition: background 0.2s; }
        .nav a:hover { background: #334155; }
        #response { background: #0f172a; border: 2px solid #334155; border-radius: 6px; padding: 15px; margin-top: 20px; min-height: 300px; white-space: pre-wrap; font-family: 'Courier New', monospace; font-size: 0.9em; overflow-x: auto; }
        .status-code { display: inline-block; padding: 6px 12px; border-radius: 4px; font-weight: bold; margin-bottom: 10px; }
        .status-200 { background: #10b981; color: white; }
        .status-400 { background: #f59e0b; color: white; }
        .status-401 { background: #ef4444; color: white; }
        .status-429 { background: #f59e0b; color: white; }
        .status-500 { background: #ef4444; color: white; }
        .info-box { background: #1e3a8a; padding: 15px; border-radius: 6px; margin-bottom: 15px; border-left: 4px solid #3b82f6; }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>🧪 API Endpoint Tester</h1>
            <p class="subtitle">Test Origami API endpoints with your API key</p>
        </header>

        <div class="nav">
            <a href="/dashboard">Dashboard</a>
            <a href="/test">API Tester</a>
        </div>

        <div class="info-box">
            <strong>ℹ️ How to use:</strong> Enter your API key, select an endpoint from the list below, and click "Send Request" to test the API.
        </div>

        <div class="grid">
            <div class="card">
                <h2>🛠️ Request Configuration</h2>
                <label><strong>API Key:</strong></label>
                <input type="text" id="apiKey" placeholder="og_your_api_key_here">

                <label style="display: block; margin-top: 20px;"><strong>Select Endpoint:</strong></label>
                <ul class="endpoint-list" id="endpoints"></ul>

                <div id="marketIdInput" style="display: none; margin-top: 20px;">
                    <label><strong>Market ID:</strong></label>
                    <input type="text" id="marketId" placeholder="0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa">
                </div>

                <div id="limitInput" style="display: none; margin-top: 20px;">
                    <label><strong>Limit:</strong></label>
                    <input type="number" id="limit" value="10" min="1" max="50">
                </div>

                <button class="btn" style="margin-top: 20px; width: 100%;" onclick="sendRequest()">🚀 Send Request</button>
            </div>

            <div class="card">
                <h2>📡 Response</h2>
                <div id="statusCode"></div>
                <div id="response">No request sent yet...</div>
            </div>
        </div>
    </div>

    <script>
        const endpoints = [
            { path: '/origami/markets', method: 'GET', desc: 'Get all markets', requiresMarketId: false, requiresLimit: false },
            { path: '/origami/markets/summary', method: 'GET', desc: 'Market summary', requiresMarketId: false, requiresLimit: false },
            { path: '/origami/markets/:id/liquidity', method: 'GET', desc: 'Market liquidity', requiresMarketId: true, requiresLimit: false },
            { path: '/origami/markets/:id/analytics', method: 'GET', desc: 'Market analytics', requiresMarketId: true, requiresLimit: false },
            { path: '/origami/markets/:id/volatility', method: 'GET', desc: 'Market volatility', requiresMarketId: true, requiresLimit: false },
            { path: '/origami/markets/:id/depth', method: 'GET', desc: 'Orderbook depth', requiresMarketId: true, requiresLimit: false },
            { path: '/origami/signals/trending', method: 'GET', desc: 'Trending markets', requiresMarketId: false, requiresLimit: true },
            { path: '/origami/signals/hot', method: 'GET', desc: 'Hot markets', requiresMarketId: false, requiresLimit: true },
            { path: '/origami/signals/volatile', method: 'GET', desc: 'Volatile markets', requiresMarketId: false, requiresLimit: true },
            { path: '/origami/signals/volume', method: 'GET', desc: 'Volume leaders', requiresMarketId: false, requiresLimit: true },
        ];

        let selectedEndpoint = null;

        function renderEndpoints() {
            const list = document.getElementById('endpoints');
            const items = [];

            for (let i = 0; i < endpoints.length; i++) {
                const ep = endpoints[i];
                const methodClass = ep.method.toLowerCase();
                const item = '<li class="endpoint-item" onclick="selectEndpoint(' + i + ')">' +
                    '<span class="method ' + methodClass + '">' + ep.method + '</span>' +
                    '<strong>' + ep.desc + '</strong>' +
                    '<div style="font-size: 0.9em; color: #94a3b8; margin-top: 5px;">' + ep.path + '</div>' +
                    '</li>';
                items.push(item);
            }

            list.innerHTML = items.join('');
        }

        function selectEndpoint(idx) {
            selectedEndpoint = endpoints[idx];

            const items = document.querySelectorAll('.endpoint-item');
            for (let i = 0; i < items.length; i++) {
                if (i === idx) {
                    items[i].classList.add('selected');
                } else {
                    items[i].classList.remove('selected');
                }
            }

            document.getElementById('marketIdInput').style.display = selectedEndpoint.requiresMarketId ? 'block' : 'none';
            document.getElementById('limitInput').style.display = selectedEndpoint.requiresLimit ? 'block' : 'none';
        }

        async function sendRequest() {
            if (!selectedEndpoint) {
                alert('Please select an endpoint first');
                return;
            }

            const apiKey = document.getElementById('apiKey').value.trim();
            if (!apiKey) {
                alert('Please enter your API key');
                return;
            }

            let url = selectedEndpoint.path;

            if (selectedEndpoint.requiresMarketId) {
                const marketId = document.getElementById('marketId').value.trim();
                if (!marketId) {
                    alert('Please enter a Market ID');
                    return;
                }
                url = url.replace(':id', marketId);
            }

            if (selectedEndpoint.requiresLimit) {
                const limit = document.getElementById('limit').value;
                url += '?limit=' + limit;
            }

            const responseDiv = document.getElementById('response');
            const statusDiv = document.getElementById('statusCode');
            responseDiv.textContent = 'Sending request...';
            statusDiv.innerHTML = '';

            try {
                const res = await fetch(url, {
                    method: selectedEndpoint.method,
                    headers: {
                        'Authorization': 'Bearer ' + apiKey,
                        'Accept': 'application/json'
                    }
                });

                const statusClass = 'status-' + Math.floor(res.status / 100) + '00';
                statusDiv.innerHTML = '<span class="status-code ' + statusClass + '">Status: ' + res.status + ' ' + res.statusText + '</span>';

                const data = await res.json();
                responseDiv.textContent = JSON.stringify(data, null, 2);
            } catch (err) {
                statusDiv.innerHTML = '<span class="status-code status-500">Error</span>';
                responseDiv.textContent = 'Error: ' + err.message;
            }
        }

        renderEndpoints();
    </script>
</body>
</html>
`
