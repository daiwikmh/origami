package handlers

const dashboardHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Origami API - Dashboard</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif; background: #0f172a; color: #e2e8f0; line-height: 1.6; }
        .container { max-width: 1200px; margin: 0 auto; padding: 20px; }
        header { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 40px 20px; text-align: center; border-radius: 10px; margin-bottom: 30px; }
        h1 { font-size: 2.5em; margin-bottom: 10px; }
        .subtitle { opacity: 0.9; font-size: 1.1em; }
        .card { background: #1e293b; border-radius: 10px; padding: 25px; margin-bottom: 20px; box-shadow: 0 4px 6px rgba(0,0,0,0.3); }
        .card h2 { margin-bottom: 20px; color: #a78bfa; border-bottom: 2px solid #334155; padding-bottom: 10px; }
        .btn { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; border: none; padding: 12px 24px; border-radius: 6px; cursor: pointer; font-size: 1em; transition: transform 0.2s; }
        .btn:hover { transform: translateY(-2px); }
        .btn:active { transform: translateY(0); }
        input, textarea { width: 100%; padding: 12px; margin: 8px 0; border: 2px solid #334155; border-radius: 6px; background: #0f172a; color: #e2e8f0; font-size: 1em; }
        input:focus, textarea:focus { outline: none; border-color: #667eea; }
        .key-list { list-style: none; }
        .key-item { background: #0f172a; padding: 15px; margin: 10px 0; border-radius: 6px; border-left: 4px solid #667eea; }
        .key-value { font-family: 'Courier New', monospace; background: #1e293b; padding: 8px; border-radius: 4px; display: inline-block; margin: 5px 0; }
        .stats-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0; }
        .stat-box { background: #0f172a; padding: 20px; border-radius: 6px; text-align: center; }
        .stat-value { font-size: 2em; color: #a78bfa; font-weight: bold; }
        .stat-label { color: #94a3b8; margin-top: 5px; }
        .success { color: #10b981; }
        .error { color: #ef4444; }
        .warning { color: #f59e0b; }
        .nav { display: flex; gap: 15px; justify-content: center; margin: 20px 0; }
        .nav a { color: #a78bfa; text-decoration: none; padding: 10px 20px; border-radius: 6px; background: #1e293b; transition: background 0.2s; }
        .nav a:hover { background: #334155; }
        #message { padding: 15px; border-radius: 6px; margin: 15px 0; display: none; }
        #message.success { background: #064e3b; border-left: 4px solid #10b981; }
        #message.error { background: #7f1d1d; border-left: 4px solid #ef4444; }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>🎯 Origami API</h1>
            <p class="subtitle">Production-Ready Injective Intelligence API Platform</p>
        </header>

        <div class="nav">
            <a href="/dashboard">Dashboard</a>
            <a href="/test">API Tester</a>
        </div>

        <div id="message"></div>

        <div class="card">
            <h2>📊 API Usage Statistics</h2>
            <div class="stats-grid" id="stats">
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
            <h2>🔑 Generate New API Key</h2>
            <form id="generateForm">
                <input type="text" id="keyName" placeholder="API Key Name (e.g., Production App)" required>
                <input type="number" id="rateLimit" placeholder="Rate Limit (requests/minute, default: 100)" min="1" max="10000">
                <button type="submit" class="btn">Generate API Key</button>
            </form>
            <div id="newKeyDisplay" style="display: none; margin-top: 20px; background: #0f172a; padding: 20px; border-radius: 6px;">
                <p class="success" style="margin-bottom: 10px;">✓ API Key Generated Successfully!</p>
                <p><strong>Name:</strong> <span id="newKeyName"></span></p>
                <p><strong>Key:</strong> <span class="key-value" id="newKeyValue"></span></p>
                <p><strong>Rate Limit:</strong> <span id="newKeyLimit"></span> requests/minute</p>
                <p class="warning" style="margin-top: 10px;">⚠️ Save this key now - it won't be shown again!</p>
            </div>
        </div>

        <div class="card">
            <h2>📋 Your API Keys</h2>
            <button class="btn" onclick="loadKeys()">Refresh Keys</button>
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
                    list.innerHTML = '<li class="key-item">No API keys yet</li>';
                    return;
                }

                const items = [];
                for (let i = 0; i < data.keys.length; i++) {
                    const key = data.keys[i];
                    const statusBadge = key.is_active ? '<span class="success">● Active</span>' : '<span class="error">● Inactive</span>';
                    const createdDate = new Date(key.created_at).toLocaleString();
                    const requestCount = (key.request_count || 0).toLocaleString();

                    const item = '<li class="key-item">' +
                        '<div><strong>' + key.name + '</strong> ' + statusBadge + '</div>' +
                        '<div>Key: <span class="key-value">' + key.key_preview + '</span></div>' +
                        '<div>Requests: ' + requestCount + ' | Rate Limit: ' + key.rate_limit + '/min</div>' +
                        '<div style="font-size: 0.9em; color: #94a3b8;">Created: ' + createdDate + '</div>' +
                        '</li>';
                    items.push(item);
                }

                list.innerHTML = items.join('');
            } catch (err) {
                showMessage('Error loading keys: ' + err.message, 'error');
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
                    showMessage(data.error || 'Error generating key', 'error');
                }
            } catch (err) {
                showMessage('Error: ' + err.message, 'error');
            }
        });

        loadStats();
        loadKeys();
        setInterval(loadStats, 5000);
    </script>
</body>
</html>
`
