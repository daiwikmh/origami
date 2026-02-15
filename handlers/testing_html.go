package handlers

const testingHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ORIGAMI API - CYBERPUNK TESTER</title>
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Orbitron:wght@400;700;900&display=swap');

        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: 'Orbitron', monospace;
            background: #000000;
            background-image:
                repeating-linear-gradient(0deg, transparent, transparent 2px, #0a0a0a 2px, #0a0a0a 4px),
                linear-gradient(180deg, #000000 0%, #0a0514 50%, #000000 100%);
            color: #00ff9f;
            line-height: 1.6;
            min-height: 100vh;
        }

        .scan-line {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 2px;
            background: linear-gradient(90deg, transparent, #00ff9f, transparent);
            animation: scan 4s linear infinite;
            pointer-events: none;
            z-index: 9999;
        }

        @keyframes scan {
            0% { transform: translateY(0); }
            100% { transform: translateY(100vh); }
        }

        .container { max-width: 1400px; margin: 0 auto; padding: 20px; position: relative; }

        header {
            background: linear-gradient(135deg, #ff006e 0%, #8338ec 50%, #3a86ff 100%);
            padding: 30px 20px;
            text-align: center;
            border: 2px solid #ff006e;
            box-shadow: 0 0 30px #ff006e, inset 0 0 30px rgba(255,0,110,0.3);
            margin-bottom: 30px;
            position: relative;
            overflow: hidden;
        }

        header::before {
            content: '';
            position: absolute;
            top: -50%;
            left: -50%;
            width: 200%;
            height: 200%;
            background: repeating-linear-gradient(
                45deg,
                transparent,
                transparent 10px,
                rgba(255,0,110,0.1) 10px,
                rgba(255,0,110,0.1) 20px
            );
            animation: headerAnim 20s linear infinite;
        }

        @keyframes headerAnim {
            0% { transform: translate(0, 0); }
            100% { transform: translate(50px, 50px); }
        }

        h1 {
            font-size: 3em;
            margin-bottom: 10px;
            text-shadow: 0 0 10px #ff006e, 0 0 20px #ff006e, 0 0 30px #ff006e;
            font-weight: 900;
            letter-spacing: 3px;
            position: relative;
            z-index: 1;
        }

        .subtitle {
            font-size: 1em;
            text-shadow: 0 0 5px #00ff9f;
            position: relative;
            z-index: 1;
        }

        .grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; }

        .card {
            background: rgba(10,5,20,0.9);
            border: 2px solid #8338ec;
            border-radius: 0px;
            padding: 25px;
            box-shadow: 0 0 20px rgba(131,56,236,0.5), inset 0 0 20px rgba(131,56,236,0.1);
            position: relative;
        }

        .card::before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            height: 2px;
            background: linear-gradient(90deg, transparent, #ff006e, transparent);
        }

        .card h2 {
            margin-bottom: 20px;
            color: #ff006e;
            text-shadow: 0 0 10px #ff006e;
            border-bottom: 2px solid #8338ec;
            padding-bottom: 10px;
            font-weight: 900;
        }

        .btn {
            background: linear-gradient(135deg, #ff006e 0%, #8338ec 100%);
            color: #ffffff;
            border: 2px solid #ff006e;
            padding: 12px 24px;
            cursor: pointer;
            font-size: 1em;
            font-family: 'Orbitron', monospace;
            font-weight: 700;
            transition: all 0.3s;
            box-shadow: 0 0 10px #ff006e;
            text-transform: uppercase;
            letter-spacing: 2px;
        }

        .btn:hover {
            box-shadow: 0 0 20px #ff006e, 0 0 30px #ff006e;
            transform: translateY(-2px);
            background: linear-gradient(135deg, #8338ec 0%, #ff006e 100%);
        }

        input, textarea {
            width: 100%;
            padding: 12px;
            margin: 8px 0;
            border: 2px solid #3a86ff;
            background: rgba(0,0,0,0.8);
            color: #00ff9f;
            font-size: 1em;
            font-family: 'Orbitron', monospace;
            box-shadow: inset 0 0 10px rgba(58,134,255,0.3);
        }

        input:focus, textarea:focus {
            outline: none;
            border-color: #ff006e;
            box-shadow: 0 0 15px #ff006e, inset 0 0 10px rgba(255,0,110,0.3);
        }

        .endpoint-list { list-style: none; max-height: 400px; overflow-y: auto; }

        .endpoint-list::-webkit-scrollbar { width: 10px; }
        .endpoint-list::-webkit-scrollbar-track { background: #000; }
        .endpoint-list::-webkit-scrollbar-thumb { background: #8338ec; box-shadow: 0 0 10px #8338ec; }

        .endpoint-item {
            background: rgba(0,0,0,0.8);
            padding: 15px;
            margin: 10px 0;
            cursor: pointer;
            border-left: 4px solid transparent;
            border-right: 2px solid #3a86ff;
            transition: all 0.3s;
        }

        .endpoint-item:hover {
            border-left-color: #ff006e;
            box-shadow: 0 0 15px rgba(255,0,110,0.5);
            transform: translateX(5px);
        }

        .endpoint-item.selected {
            border-left-color: #00ff9f;
            background: rgba(131,56,236,0.2);
            box-shadow: 0 0 20px rgba(0,255,159,0.5);
        }

        .method {
            display: inline-block;
            padding: 4px 12px;
            font-weight: bold;
            font-size: 0.9em;
            margin-right: 10px;
            box-shadow: 0 0 10px currentColor;
        }

        .method.get { background: #00ff9f; color: #000; }
        .method.post { background: #ff006e; color: #fff; }

        .nav {
            display: flex;
            gap: 15px;
            justify-content: center;
            margin: 20px 0;
        }

        .nav a {
            color: #00ff9f;
            text-decoration: none;
            padding: 10px 20px;
            background: rgba(0,0,0,0.8);
            border: 2px solid #3a86ff;
            transition: all 0.3s;
            text-transform: uppercase;
            font-weight: 700;
            letter-spacing: 1px;
        }

        .nav a:hover {
            background: rgba(131,56,236,0.3);
            box-shadow: 0 0 15px #3a86ff;
            border-color: #ff006e;
        }

        #response {
            background: #000;
            border: 2px solid #00ff9f;
            padding: 15px;
            margin-top: 20px;
            min-height: 300px;
            white-space: pre-wrap;
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
            overflow-x: auto;
            box-shadow: inset 0 0 20px rgba(0,255,159,0.2);
        }

        .status-code {
            display: inline-block;
            padding: 6px 12px;
            font-weight: bold;
            margin-bottom: 10px;
            box-shadow: 0 0 10px currentColor;
        }

        .status-200 { background: #00ff9f; color: #000; }
        .status-400 { background: #ffbe0b; color: #000; }
        .status-401 { background: #ff006e; color: #fff; }
        .status-429 { background: #ff9e00; color: #000; }
        .status-500 { background: #ff006e; color: #fff; }

        .info-box {
            background: rgba(58,134,255,0.1);
            padding: 15px;
            margin-bottom: 15px;
            border-left: 4px solid #3a86ff;
            box-shadow: 0 0 15px rgba(58,134,255,0.3);
        }

        .glitch {
            position: relative;
        }

        .glitch::before,
        .glitch::after {
            content: attr(data-text);
            position: absolute;
            top: 0;
            left: 0;
            opacity: 0.8;
        }

        .glitch::before {
            animation: glitch 0.3s infinite;
            color: #ff006e;
            z-index: -1;
        }

        .glitch::after {
            animation: glitch 0.3s infinite reverse;
            color: #3a86ff;
            z-index: -2;
        }

        @keyframes glitch {
            0% { transform: translate(0); }
            20% { transform: translate(-2px, 2px); }
            40% { transform: translate(-2px, -2px); }
            60% { transform: translate(2px, 2px); }
            80% { transform: translate(2px, -2px); }
            100% { transform: translate(0); }
        }
    </style>
</head>
<body>
    <div class="scan-line"></div>
    <div class="container">
        <header>
            <h1 class="glitch" data-text="⚡ ORIGAMI API TESTER ⚡">⚡ ORIGAMI API TESTER ⚡</h1>
            <p class="subtitle">[ CYBERPUNK INTERFACE v2.0 ]</p>
        </header>

        <div class="nav">
            <a href="/dashboard">◢ DASHBOARD</a>
            <a href="/test">◢ API TESTER</a>
        </div>

        <div class="info-box">
            <strong>⚠ SYSTEM NOTICE:</strong> Enter your API key, select target endpoint, execute request.
        </div>

        <div class="grid">
            <div class="card">
                <h2>⚙ REQUEST CONFIG</h2>
                <label><strong>API KEY:</strong></label>
                <input type="text" id="apiKey" placeholder="og_your_api_key_here">

                <label style="display: block; margin-top: 20px;"><strong>SELECT ENDPOINT:</strong></label>
                <ul class="endpoint-list" id="endpoints"></ul>

                <div id="marketIdInput" style="display: none; margin-top: 20px;">
                    <label><strong>MARKET ID:</strong></label>
                    <input type="text" id="marketId" placeholder="0x...">
                </div>

                <div id="limitInput" style="display: none; margin-top: 20px;">
                    <label><strong>LIMIT:</strong></label>
                    <input type="number" id="limit" value="10" min="1" max="50">
                </div>

                <div id="addressInput" style="display: none; margin-top: 20px;">
                    <label><strong>WALLET ADDRESS:</strong></label>
                    <input type="text" id="walletAddress" placeholder="0x...">
                </div>

                <button class="btn" style="margin-top: 20px; width: 100%;" onclick="sendRequest()">⚡ EXECUTE REQUEST</button>
            </div>

            <div class="card">
                <h2>◢ RESPONSE DATA</h2>
                <div id="statusCode"></div>
                <div id="response">[ AWAITING INPUT ]</div>
            </div>
        </div>
    </div>

    <script>
        const endpoints = [
            { path: '/origami/markets', method: 'GET', desc: 'Get all markets', requiresMarketId: false, requiresLimit: false, requiresAddress: false },
            { path: '/origami/markets/summary', method: 'GET', desc: 'Market summary', requiresMarketId: false, requiresLimit: false, requiresAddress: false },
            { path: '/origami/markets/:id/liquidity', method: 'GET', desc: 'Market liquidity', requiresMarketId: true, requiresLimit: false, requiresAddress: false },
            { path: '/origami/markets/:id/analytics', method: 'GET', desc: 'Market analytics', requiresMarketId: true, requiresLimit: false, requiresAddress: false },
            { path: '/origami/markets/:id/volatility', method: 'GET', desc: 'Market volatility', requiresMarketId: true, requiresLimit: false, requiresAddress: false },
            { path: '/origami/markets/:id/depth', method: 'GET', desc: 'Orderbook depth', requiresMarketId: true, requiresLimit: false, requiresAddress: false },
            { path: '/origami/signals/trending', method: 'GET', desc: 'Trending markets', requiresMarketId: false, requiresLimit: true, requiresAddress: false },
            { path: '/origami/signals/hot', method: 'GET', desc: 'Hot markets', requiresMarketId: false, requiresLimit: true, requiresAddress: false },
            { path: '/origami/signals/volatile', method: 'GET', desc: 'Volatile markets', requiresMarketId: false, requiresLimit: true, requiresAddress: false },
            { path: '/origami/signals/volume', method: 'GET', desc: 'Volume leaders', requiresMarketId: false, requiresLimit: true, requiresAddress: false },
            { path: '/origami/nft/verify/:address', method: 'GET', desc: '🎯 NFT Ownership Check', requiresMarketId: false, requiresLimit: false, requiresAddress: true },
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
                    '<div style="font-size: 0.9em; color: #3a86ff; margin-top: 5px;">' + ep.path + '</div>' +
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
            document.getElementById('addressInput').style.display = selectedEndpoint.requiresAddress ? 'block' : 'none';
        }

        async function sendRequest() {
            if (!selectedEndpoint) {
                alert('⚠ SELECT AN ENDPOINT FIRST');
                return;
            }

            const apiKey = document.getElementById('apiKey').value.trim();
            if (!apiKey) {
                alert('⚠ API KEY REQUIRED');
                return;
            }

            let url = selectedEndpoint.path;

            if (selectedEndpoint.requiresMarketId) {
                const marketId = document.getElementById('marketId').value.trim();
                if (!marketId) {
                    alert('⚠ MARKET ID REQUIRED');
                    return;
                }
                url = url.replace(':id', marketId);
            }

            if (selectedEndpoint.requiresAddress) {
                const address = document.getElementById('walletAddress').value.trim();
                if (!address) {
                    alert('⚠ WALLET ADDRESS REQUIRED');
                    return;
                }
                url = url.replace(':address', address);
            }

            if (selectedEndpoint.requiresLimit) {
                const limit = document.getElementById('limit').value;
                url += '?limit=' + limit;
            }

            const responseDiv = document.getElementById('response');
            const statusDiv = document.getElementById('statusCode');
            responseDiv.textContent = '[ TRANSMITTING REQUEST... ]';
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
                statusDiv.innerHTML = '<span class="status-code ' + statusClass + '">STATUS: ' + res.status + ' ' + res.statusText + '</span>';

                const data = await res.json();
                responseDiv.textContent = JSON.stringify(data, null, 2);
            } catch (err) {
                statusDiv.innerHTML = '<span class="status-code status-500">ERROR</span>';
                responseDiv.textContent = '[ SYSTEM ERROR: ' + err.message + ' ]';
            }
        }

        renderEndpoints();
    </script>
</body>
</html>
`
