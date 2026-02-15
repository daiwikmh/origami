# 🚀 Origami API - Deployment Guide

Complete guide for deploying Origami API to production.

---

## 📋 Table of Contents

1. [Prerequisites](#prerequisites)
2. [Local Development](#local-development)
3. [Building for Production](#building-for-production)
4. [Deployment Options](#deployment-options)
5. [Environment Configuration](#environment-configuration)
6. [Monitoring & Logging](#monitoring--logging)
7. [Security Best Practices](#security-best-practices)
8. [Scaling](#scaling)
9. [Troubleshooting](#troubleshooting)

---

## ✅ Prerequisites

### Required Software
- **Go** 1.23+ ([Download](https://golang.org/dl/))
- **Git** for version control
- **systemd** (for Linux deployments)
- **nginx** or **Apache** (for reverse proxy)
- **SSL Certificate** (Let's Encrypt recommended)

### Recommended
- **Docker** (optional, for containerized deployment)
- **Monitoring tools** (Prometheus, Grafana)
- **Log aggregation** (ELK Stack, Loki)

---

## 💻 Local Development

### 1. Clone Repository

```bash
git clone https://github.com/yourusername/origami-api.git
cd origami-api
```

### 2. Install Dependencies

```bash
go mod download
go mod verify
```

### 3. Build & Run Locally

```bash
# Build
go build -o origami

# Run
./origami
```

Server starts on `http://localhost:8080`

### 4. Access Endpoints
- **Dashboard**: http://localhost:8080/dashboard
- **API Tester**: http://localhost:8080/test
- **Documentation**: http://localhost:8080/docs

---

## 🏗️ Building for Production

### Optimized Build

```bash
# Build with optimizations
go build -ldflags="-s -w" -o origami

# Verify binary
./origami --version

# Check binary size
ls -lh origami
```

### Cross-Platform Builds

```bash
# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o origami-linux-amd64

# Linux (ARM64 - for AWS Graviton, Raspberry Pi)
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o origami-linux-arm64

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o origami-darwin-amd64

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o origami-darwin-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o origami-windows-amd64.exe
```

---

## 🌐 Deployment Options

### Option 1: Traditional Server (VPS/Dedicated)

#### Step 1: Setup Server

```bash
# SSH into your server
ssh user@your-server.com

# Create application directory
sudo mkdir -p /opt/origami
sudo chown $USER:$USER /opt/origami
```

#### Step 2: Upload Binary

```bash
# From local machine
scp origami-linux-amd64 user@your-server.com:/opt/origami/origami

# On server, make executable
chmod +x /opt/origami/origami
```

#### Step 3: Create systemd Service

```bash
sudo nano /etc/systemd/system/origami.service
```

Add the following configuration:

```ini
[Unit]
Description=Origami API Service
After=network.target
Wants=network-online.target

[Service]
Type=simple
User=origami
Group=origami
WorkingDirectory=/opt/origami
ExecStart=/opt/origami/origami
Restart=always
RestartSec=10
StandardOutput=append:/var/log/origami/access.log
StandardError=append:/var/log/origami/error.log

# Security settings
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/log/origami

# Resource limits
LimitNOFILE=65536
LimitNPROC=4096

[Install]
WantedBy=multi-user.target
```

#### Step 4: Create User & Directories

```bash
# Create system user
sudo useradd -r -s /bin/false origami

# Create log directory
sudo mkdir -p /var/log/origami
sudo chown origami:origami /var/log/origami

# Set permissions
sudo chown -R origami:origami /opt/origami
```

#### Step 5: Start Service

```bash
# Reload systemd
sudo systemctl daemon-reload

# Enable service (start on boot)
sudo systemctl enable origami

# Start service
sudo systemctl start origami

# Check status
sudo systemctl status origami

# View logs
sudo journalctl -u origami -f
```

---

### Option 2: Docker Deployment

#### Create Dockerfile

```dockerfile
# /opt/origami/Dockerfile

FROM golang:1.23-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o origami .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/origami .

# Create non-root user
RUN addgroup -g 1000 origami && \
    adduser -D -u 1000 -G origami origami && \
    chown -R origami:origami /app

USER origami

EXPOSE 8080

CMD ["./origami"]
```

#### Build & Run Docker Container

```bash
# Build image
docker build -t origami-api:latest .

# Run container
docker run -d \
  --name origami-api \
  --restart unless-stopped \
  -p 8080:8080 \
  -v /var/log/origami:/var/log/origami \
  origami-api:latest

# View logs
docker logs -f origami-api

# Stop container
docker stop origami-api

# Remove container
docker rm origami-api
```

#### Docker Compose

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  origami-api:
    build: .
    container_name: origami-api
    restart: unless-stopped
    ports:
      - "8080:8080"
    volumes:
      - ./logs:/var/log/origami
    environment:
      - GIN_MODE=release
    healthcheck:
      test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:8080/"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

Run with Docker Compose:

```bash
# Start
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

---

### Option 3: Cloud Platforms

#### AWS EC2

1. **Launch EC2 Instance** (Amazon Linux 2 or Ubuntu)
2. **Configure Security Group**: Allow inbound TCP port 8080 (or 443 for HTTPS)
3. **SSH into instance** and follow "Traditional Server" steps above
4. **Optional**: Use AWS ELB (Elastic Load Balancer) for HTTPS termination

#### AWS ECS/Fargate

```bash
# Tag image for ECR
docker tag origami-api:latest 123456789.dkr.ecr.us-east-1.amazonaws.com/origami-api:latest

# Push to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 123456789.dkr.ecr.us-east-1.amazonaws.com
docker push 123456789.dkr.ecr.us-east-1.amazonaws.com/origami-api:latest

# Create ECS task definition and service through AWS Console or CLI
```

#### Google Cloud Run

```bash
# Build and push to Google Container Registry
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/origami-api

# Deploy to Cloud Run
gcloud run deploy origami-api \
  --image gcr.io/YOUR_PROJECT_ID/origami-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

#### DigitalOcean App Platform

1. Connect GitHub repository
2. Select Dockerfile or Buildpack
3. Configure environment variables
4. Deploy with one click

#### Heroku

```bash
# Login to Heroku
heroku login

# Create app
heroku create your-origami-api

# Deploy
git push heroku main

# View logs
heroku logs --tail
```

---

## ⚙️ Environment Configuration

### Environment Variables

Create `/opt/origami/.env`:

```env
# Server Configuration
PORT=8080
GIN_MODE=release

# API Configuration
API_BASE_URL=https://api.yourorigami.com
CORS_ORIGINS=https://yourapp.com,https://app.yourorigami.com

# Rate Limiting
DEFAULT_RATE_LIMIT=100

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### Configuration File

Create `/opt/origami/config.yaml`:

```yaml
server:
  port: 8080
  read_timeout: 15s
  write_timeout: 15s
  idle_timeout: 60s

api:
  base_url: "https://api.yourorigami.com"
  default_rate_limit: 100

logging:
  level: "info"
  format: "json"
  output: "/var/log/origami/app.log"

security:
  cors_origins:
    - "https://yourapp.com"
    - "https://app.yourorigami.com"
```

---

## 🔒 Reverse Proxy (nginx)

### Install nginx

```bash
sudo apt update
sudo apt install nginx
```

### Configure nginx

Create `/etc/nginx/sites-available/origami`:

```nginx
upstream origami_backend {
    server 127.0.0.1:8080;
    keepalive 64;
}

server {
    listen 80;
    server_name api.yourorigami.com;

    # Redirect HTTP to HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name api.yourorigami.com;

    # SSL Configuration
    ssl_certificate /etc/letsencrypt/live/api.yourorigami.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.yourorigami.com/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;

    # Security Headers
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;

    # Logging
    access_log /var/log/nginx/origami-access.log;
    error_log /var/log/nginx/origami-error.log;

    # Gzip Compression
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml;

    location / {
        proxy_pass http://origami_backend;
        proxy_http_version 1.1;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Connection "";

        # Timeouts
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;

        # Buffer settings
        proxy_buffering on;
        proxy_buffer_size 4k;
        proxy_buffers 8 4k;
        proxy_busy_buffers_size 8k;
    }

    # Health check endpoint (no auth required)
    location /health {
        proxy_pass http://origami_backend/;
        access_log off;
    }
}
```

### Enable Site

```bash
# Create symlink
sudo ln -s /etc/nginx/sites-available/origami /etc/nginx/sites-enabled/

# Test configuration
sudo nginx -t

# Reload nginx
sudo systemctl reload nginx
```

### SSL Certificate (Let's Encrypt)

```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx

# Obtain certificate
sudo certbot --nginx -d api.yourorigami.com

# Test renewal
sudo certbot renew --dry-run

# Auto-renewal is configured by default
```

---

## 📊 Monitoring & Logging

### System Monitoring

```bash
# View logs
sudo journalctl -u origami -f

# Check resource usage
htop

# Check disk space
df -h

# Check memory usage
free -m
```

### Application Metrics

Add health check endpoint in `main.go`:

```go
r.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "healthy",
        "uptime": time.Since(startTime).String(),
    })
})
```

### Log Rotation

Create `/etc/logrotate.d/origami`:

```
/var/log/origami/*.log {
    daily
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 origami origami
    sharedscripts
    postrotate
        systemctl reload origami > /dev/null 2>&1 || true
    endscript
}
```

---

## 🔐 Security Best Practices

### 1. Firewall Configuration

```bash
# Enable UFW
sudo ufw enable

# Allow SSH
sudo ufw allow 22/tcp

# Allow HTTP/HTTPS (if using nginx)
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# If exposing API directly (NOT recommended for production)
# sudo ufw allow 8080/tcp

# Check status
sudo ufw status
```

### 2. Secure API Keys

- **Never commit API keys** to version control
- **Rotate keys regularly** (monthly recommended)
- **Use environment variables** or secret managers (AWS Secrets Manager, HashiCorp Vault)
- **Implement key expiration** in production

### 3. Rate Limiting

- Already implemented per-key rate limiting
- Consider adding IP-based rate limiting in nginx:

```nginx
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;

location /origami {
    limit_req zone=api_limit burst=20 nodelay;
    # ... rest of proxy config
}
```

### 4. CORS Configuration

Update in code to whitelist specific origins:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://yourapp.com"},
    AllowMethods:     []string{"GET", "POST"},
    AllowHeaders:     []string{"Authorization", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))
```

---

## 📈 Scaling

### Vertical Scaling (Single Server)

1. **Upgrade server resources** (CPU, RAM)
2. **Optimize Go application**:
   ```bash
   # Increase GOMAXPROCS
   export GOMAXPROCS=8
   ```

### Horizontal Scaling (Multiple Servers)

1. **Deploy multiple instances** behind a load balancer
2. **Use sticky sessions** for API key rate limiting
3. **Implement Redis** for distributed rate limiting (future enhancement)

#### AWS Application Load Balancer

```bash
# Create target group
aws elbv2 create-target-group \
  --name origami-targets \
  --protocol HTTP \
  --port 8080 \
  --vpc-id vpc-xxxxx

# Create load balancer
aws elbv2 create-load-balancer \
  --name origami-lb \
  --subnets subnet-xxxxx subnet-yyyyy \
  --security-groups sg-xxxxx

# Register targets
aws elbv2 register-targets \
  --target-group-arn arn:aws:elasticloadbalancing:... \
  --targets Id=i-xxxxx Id=i-yyyyy
```

---

## 🔧 Troubleshooting

### Common Issues

#### 1. Port Already in Use

```bash
# Check what's using port 8080
sudo lsof -i :8080

# Kill process
sudo kill -9 PID
```

#### 2. Permission Denied

```bash
# Fix ownership
sudo chown -R origami:origami /opt/origami
sudo chmod +x /opt/origami/origami
```

#### 3. Service Won't Start

```bash
# Check logs
sudo journalctl -u origami -n 100 --no-pager

# Check systemd status
sudo systemctl status origami

# Restart service
sudo systemctl restart origami
```

#### 4. High Memory Usage

```bash
# Check memory
free -m

# Monitor Go memory
# Add to main.go:
go func() {
    http.ListenAndServe("localhost:6060", nil)
}()

# Then use pprof:
go tool pprof http://localhost:6060/debug/pprof/heap
```

#### 5. Slow API Responses

```bash
# Check system load
uptime

# Check disk I/O
iostat -x 1

# Check network
netstat -s
```

---

## ✅ Pre-Deployment Checklist

- [ ] Build optimized binary (`-ldflags="-s -w"`)
- [ ] Configure systemd service
- [ ] Setup nginx reverse proxy
- [ ] Obtain SSL certificate
- [ ] Configure firewall (UFW)
- [ ] Setup log rotation
- [ ] Configure monitoring/alerts
- [ ] Test health check endpoint
- [ ] Verify rate limiting works
- [ ] Test API with production keys
- [ ] Document deployment process
- [ ] Setup automated backups (if using database)
- [ ] Configure CORS for production domains
- [ ] Review security settings
- [ ] Plan rollback strategy

---

## 📞 Support

- **Issues**: https://github.com/yourusername/origami/issues
- **Email**: support@yourorigami.com
- **Docs**: https://api.yourorigami.com/docs

---

**Deployment Complete! 🎉**
