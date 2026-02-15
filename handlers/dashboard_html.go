package handlers

const dashboardHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ORIGAMI API - COMMAND CENTER</title>
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Orbitron:wght:400;700;900&display=swap');

        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: 'Orbitron', monospace;
            background: #000000;
            background-image:
                repeating-linear-gradient(0deg, transparent, transparent 2px, #0a0a0a 2px, #0a0a0a 4px),
                linear-gradient(180deg, #000000 0%, #0a0514 50%, #000000 100%);
            color: #E1C4E9;
            line-height: 1.6;
        }

            0% { transform: translateY(0); }
            100% { transform: translateY(100vh); }
        }

        .container { max-width: 1200px; margin: 0 auto; padding: 20px; }

        header {
            background: linear-gradient(135deg, #E1C4E9 0%, #232323 50%, #232323 100%);
            padding: 30px 20px;
            text-align: center;
            border: 2px solid #E1C4E9;
            box-shadow: 0 0 30px #E1C4E9, inset 0 0 30px rgba(255,0,110,0.3);
            margin-bottom: 30px;
            position: relative;
            overflow: hidden;
        }

        h1 {
            font-size: 3em;
            margin-bottom: 10px;
            text-shadow: 0 0 10px #E1C4E9, 0 0 20px #E1C4E9;
            font-weight: 900;
            letter-spacing: 3px;
        }

        .subtitle {
            font-size: 1em;
            text-shadow: 0 0 5px #E1C4E9;
        }

        .card {
            background: rgba(10,5,20,0.9);
            border: 2px solid #232323;
            padding: 25px;
            margin-bottom: 20px;
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
            background: linear-gradient(90deg, transparent, #E1C4E9, transparent);
        }

        .card h2 {
            margin-bottom: 20px;
            color: #E1C4E9;
            text-shadow: 0 0 10px #E1C4E9;
            border-bottom: 2px solid #232323;
            padding-bottom: 10px;
            font-weight: 900;
        }

        .btn {
            background: linear-gradient(135deg, #E1C4E9 0%, #232323 100%);
            color: #ffffff;
            border: 2px solid #E1C4E9;
            padding: 12px 24px;
            cursor: pointer;
            font-size: 1em;
            font-family: 'Orbitron', monospace;
            font-weight: 700;
            transition: all 0.3s;
            box-shadow: 0 0 10px #E1C4E9;
            text-transform: uppercase;
            letter-spacing: 2px;
        }

        .btn:hover {
            box-shadow: 0 0 20px #E1C4E9, 0 0 30px #E1C4E9;
            transform: translateY(-2px);
        }

        input {
            width: 100%;
            padding: 12px;
            margin: 8px 0;
            border: 2px solid #232323;
            background: rgba(0,0,0,0.8);
            color: #E1C4E9;
            font-size: 1em;
            font-family: 'Orbitron', monospace;
            box-shadow: inset 0 0 10px rgba(58,134,255,0.3);
        }

        input:focus {
            outline: none;
            border-color: #E1C4E9;
            box-shadow: 0 0 15px #E1C4E9, inset 0 0 10px rgba(255,0,110,0.3);
        }

        .key-list { list-style: none; }

        .key-item {
            background: rgba(0,0,0,0.8);
            padding: 15px;
            margin: 10px 0;
            border-left: 4px solid #E1C4E9;
            box-shadow: 0 0 10px rgba(0,255,159,0.3);
        }

        .key-value {
            font-family: 'Courier New', monospace;
            background: #000;
            padding: 8px;
            display: inline-block;
            margin: 5px 0;
            border: 1px solid #232323;
            box-shadow: 0 0 5px #232323;
        }

        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 15px;
            margin: 20px 0;
        }

        .stat-box {
            background: rgba(0,0,0,0.8);
            padding: 20px;
            text-align: center;
            border: 2px solid #232323;
            box-shadow: 0 0 15px rgba(131,56,236,0.5);
        }

        .stat-value {
            font-size: 2em;
            color: #E1C4E9;
            font-weight: bold;
            text-shadow: 0 0 10px #E1C4E9;
        }

        .stat-label {
            color: #232323;
            margin-top: 5px;
            text-transform: uppercase;
            font-size: 0.8em;
        }

        .success { color: #E1C4E9; text-shadow: 0 0 5px #E1C4E9; }
        .error { color: #E1C4E9; text-shadow: 0 0 5px #E1C4E9; }
        .warning { color: #ffbe0b; text-shadow: 0 0 5px #ffbe0b; }

        .nav {
            display: flex;
            gap: 15px;
            justify-content: center;
            margin: 20px 0;
        }

        .nav a {
            color: #E1C4E9;
            text-decoration: none;
            padding: 10px 20px;
            background: rgba(0,0,0,0.8);
            border: 2px solid #232323;
            transition: all 0.3s;
            text-transform: uppercase;
            font-weight: 700;
        }

        .nav a:hover {
            background: rgba(131,56,236,0.3);
            box-shadow: 0 0 15px #232323;
            border-color: #E1C4E9;
        }

        #message {
            padding: 15px;
            margin: 15px 0;
            display: none;
            border-left: 4px solid;
        }

        #message.success {
            background: rgba(0,255,159,0.1);
            border-left-color: #E1C4E9;
            box-shadow: 0 0 10px rgba(0,255,159,0.5);
        }

        #message.error {
            background: rgba(255,0,110,0.1);
            border-left-color: #E1C4E9;
            box-shadow: 0 0 10px rgba(255,0,110,0.5);
        }

        #newKeyDisplay {
            background: rgba(0,255,159,0.1);
            border: 2px solid #E1C4E9;
            box-shadow: 0 0 20px rgba(0,255,159,0.3);
        }
    </style>
</head>
<body>
    
    <div class="container">
        <header>
            <h1>⚡ ORIGAMI COMMAND CENTER ⚡</h1>
            <p class="subtitle">[ API KEY MANAGEMENT SYSTEM ]</p>
        </header>

        <div class="nav">
            <a href="/">◢ DASHBOARD</a>
            <a href="/test">◢ API TESTER</a>
            <a href="/docs">◢ DOCUMENTATION</a>
        </div>

        <div id="message"></div>

        <div class="card">
            <h2>◢ SYSTEM STATISTICS</h2>
            <div class="stats-grid">
                <div class="stat-box">
                    <div class="stat-value" id="totalKeys">-</div>
                    <div class="stat-label">Total Keys</div>
                </div>
                <div class="stat-box">
                    <div class="stat-value" id="activeKeys">-</div>
                    <div class="stat-label">Active Keys</div>
                </div>
                <div class="stat-box">
                    <div class="stat-value" id="totalRequests">-</div>
                    <div class="stat-label">Total Requests</div>
                </div>
            </div>
        </div>

        <div class="card">
            <h2>⚙ GENERATE API KEY</h2>
            <form id="generateForm">
                <input type="text" id="keyName" placeholder="KEY NAME (e.g., Production App)" required>
                <input type="number" id="rateLimit" placeholder="RATE LIMIT (requests/minute, default: 100)" min="1" max="10000">
                <button type="submit" class="btn">⚡ GENERATE KEY</button>
            </form>
            <div id="newKeyDisplay" style="display: none; margin-top: 20px; padding: 20px;">
                <p class="success" style="margin-bottom: 10px;">✓ API KEY GENERATED</p>
                <p><strong>NAME:</strong> <span id="newKeyName"></span></p>
                <p><strong>KEY:</strong> <span class="key-value" id="newKeyValue"></span></p>
                <p><strong>RATE LIMIT:</strong> <span id="newKeyLimit"></span> req/min</p>
                <p class="warning" style="margin-top: 10px;">⚠ SAVE THIS KEY - IT WILL NOT BE SHOWN AGAIN</p>
            </div>
        </div>

        <div class="card">
            <h2>◢ ACTIVE API KEYS</h2>
            <button class="btn" onclick="loadKeys()">⟳ REFRESH</button>
            <ul class="key-list" id="keysList"></ul>
        </div>
    </div>

    <script>
        function showMessage(text, type) {
            const msg = document.getElementById('message');
            msg.textContent = text;
            msg.className = type;
            msg.style.display = 'block';
            setTimeout(function() { msg.style.display = 'none'; }, 5000);
        }

        function copyKey(id, btn) {
    const text = document.getElementById(id).innerText;

    navigator.clipboard.writeText(text).then(() => {
        const original = btn.innerText;
        btn.innerText = "Copied!";
        setTimeout(() => btn.innerText = original, 1500);
    }).catch(err => {
        console.error("Copy failed:", err);
    });
}

        async function loadStats() {
            try {
                const res = await fetch('/admin/usage');
                const data = await res.json();
                document.getElementById('totalKeys').textContent = data.total_keys || 0;
                document.getElementById('activeKeys').textContent = data.active_keys || 0;
                document.getElementById('totalRequests').textContent = (data.total_requests || 0).toLocaleString();
            } catch (err) {
                console.error('Error loading stats:', err);
            }
        }

        async function loadKeys() {
            try {
                const res = await fetch('/admin/keys');
                const data = await res.json();
                const list = document.getElementById('keysList');

                if (!data.keys || data.keys.length === 0) {
                    list.innerHTML = '<li class="key-item">[ NO API KEYS FOUND ]</li>';
                    return;
                }

                const items = [];
                for (let i = 0; i < data.keys.length; i++) {
                    const key = data.keys[i];
                    const statusBadge = key.is_active ? '<span class="success">● ACTIVE</span>' : '<span class="error">● INACTIVE</span>';
                    const createdDate = new Date(key.created_at).toLocaleString();
                    const requestCount = (key.request_count || 0).toLocaleString();

                    const item = '<li class="key-item">' +
                        '<div><strong>' + key.name + '</strong> ' + statusBadge + '</div>' +
const keyId = 'key-' + i;

const item = '<li class="key-item">' +
    '<div><strong>' + key.name + '</strong> ' + statusBadge + '</div>' +
    '<div>KEY: <span class="key-value">' + key.key_preview + '</span></div>' +
    '<code id="' + keyId + '" style="display:none;">' + key.key + '</code>' +
    '<div>' +
        '<button onclick="copyKey(\'' + keyId + '\', this)">COPY</button>' +
    '</div>' +
    '<div>REQUESTS: ' + requestCount + ' | RATE: ' + key.rate_limit + '/min</div>' +
    '<div style="font-size: 0.8em; color: #232323;">CREATED: ' + createdDate + '</div>' +
    '</li>';
                        '<div>REQUESTS: ' + requestCount + ' | RATE: ' + key.rate_limit + '/min</div>' +
                        '<div style="font-size: 0.8em; color: #232323;">CREATED: ' + createdDate + '</div>' +
                        '</li>';
                    items.push(item);
                }

                list.innerHTML = items.join('');
            } catch (err) {
                showMessage('[ ERROR LOADING KEYS: ' + err.message + ' ]', 'error');
            }
        }

        document.getElementById('generateForm').addEventListener('submit', async function(e) {
            e.preventDefault();
            const name = document.getElementById('keyName').value;
            const rateLimit = parseInt(document.getElementById('rateLimit').value) || 100;

            try {
                const res = await fetch('/admin/keys/generate', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ name: name, rate_limit: rateLimit })
                });

                const data = await res.json();

                if (res.ok) {
                    document.getElementById('newKeyName').textContent = data.name;
                    document.getElementById('newKeyValue').textContent = data.api_key;
                    document.getElementById('newKeyLimit').textContent = data.rate_limit;
                    document.getElementById('newKeyDisplay').style.display = 'block';
                    document.getElementById('generateForm').reset();
                    loadKeys();
                    loadStats();
                } else {
                    showMessage('[ ERROR: ' + (data.error || 'Unknown error') + ' ]', 'error');
                }
            } catch (err) {
                showMessage('[ SYSTEM ERROR: ' + err.message + ' ]', 'error');
            }
        });

        loadStats();
        loadKeys();
        setInterval(loadStats, 5000);
    </script>
</body>
</html>
`
