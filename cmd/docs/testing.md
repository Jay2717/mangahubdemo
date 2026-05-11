# MangaHub — Current Features & Testing Guide

---

# 1. REST API Server

## Chức năng

* cung cấp API HTTP
* xử lý:

  * login
  * register
  * manga APIs
  * health check
  * reading list

---

# Run

```powershell
go run cmd/api-server/main.go
```

---

# Test

## Ping

```powershell
curl http://localhost:8080/ping -UseBasicParsing
```

Expected:

```json
{"message":"pong"}
```

---

# 2. JWT Authentication

## Chức năng

* đăng ký tài khoản
* login
* tạo JWT token
* bảo vệ API bằng middleware

---

# Register

```powershell
Invoke-WebRequest -Uri http://localhost:8080/register `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"username":"liem","password":"123456"}'
```

---

# Login

```powershell
Invoke-WebRequest -Uri http://localhost:8080/login `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"username":"liem","password":"123456"}'
```

Expected:

```json
{
  "token":"..."
}
```

---

# Save Token

```powershell
$token="YOUR_TOKEN"
```

---

# Test Protected API

```powershell
Invoke-WebRequest -Uri http://localhost:8080/manga `
-Headers @{"Authorization"="Bearer $token"}
```

---

# 3. Manga API

## Chức năng

* lấy danh sách manga
* thêm manga mới

---

# Get Manga List

```powershell
Invoke-WebRequest -Uri http://localhost:8080/manga `
-Headers @{"Authorization"="Bearer $token"} |
Select-Object -Expand Content
```

---

# Create Manga

```powershell
Invoke-WebRequest -Uri http://localhost:8080/manga `
-Method POST `
-Headers @{
"Content-Type"="application/json"
"Authorization"="Bearer $token"
} `
-Body '{"id":"chainsaw-man","title":"Chainsaw Man","author":"Tatsuki Fujimoto"}'
```

---

# 4. SQLite Database

## Chức năng

lưu:

* users
* manga
* reading progress
* reading list

---

# Database File

```text
mangahub.db
```

---

# Test

Mở bằng SQLite extension trong VSCode.

Check tables:

* users
* manga
* reading_progress
* reading_list

---

# 5. TCP Reading Progress Sync

## Chức năng

client gửi tiến trình đọc manga qua TCP socket.

Server:

* nhận dữ liệu
* lưu vào database

---

# Run TCP Server

```powershell
go run cmd/tcp-server/main.go
```

---

# Test TCP

```powershell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 7070)

$stream = $client.GetStream()

$message = '{"username":"liem","manga_id":"blue-box","chapter":30}' + "`n"

$data = [System.Text.Encoding]::UTF8.GetBytes($message)

$stream.Write($data, 0, $data.Length)

$stream.Close()

$client.Close()
```

---

# Expected Terminal Output

```text
=== Reading Progress ===
User: liem
Manga: blue-box
Chapter: 30
progress saved
```

---

# Verify Database

Check:

```text
reading_progress
```

table.

---

# 6. UDP Notification Service

## Chức năng

nhận notification realtime qua UDP.

Ví dụ:

* new chapter release
* manga update

---

# Run UDP Server

```powershell
go run cmd/udp-server/main.go
```

---

# Test UDP

```powershell
$udpClient = New-Object System.Net.Sockets.UdpClient

$message = "New chapter: Blue Box 190"

$data = [System.Text.Encoding]::UTF8.GetBytes($message)

$udpClient.Send($data, $data.Length, "localhost", 6060)

$udpClient.Close()
```

---

# Expected Output

```text
=== UDP Notification ===
Message: New chapter: Blue Box 190
```

---

# 7. WebSocket Chat Server

## Chức năng

chat realtime bằng WebSocket.

---

# Run Chat Server

```powershell
go run cmd/chat-server/main.go
```

---

# Install WebSocket Client

```powershell
npm install -g wscat
```

---

# Connect

```powershell
wscat -c ws://localhost:9093/chat
```

---

# Send Message

```text
hello mangahub
```

---

# Expected

Server broadcast lại message.

---

# 8. gRPC Manga Service

## Chức năng

gRPC service trả danh sách manga.

---

# Run gRPC Server

```powershell
go run cmd/grpc-server/main.go
```

---

# Test gRPC Client

```powershell
go run cmd/grpc-client/main.go
```

---

# Expected

```text
=== Manga List ===
blue-box Blue Box Kouji Miura
oshi-no-koi Oshi no Ko Aka Akasaka
```

---

# 9. Health Check API

## Chức năng

kiểm tra server còn hoạt động không.

---

# Test

```powershell
curl http://localhost:8080/health -UseBasicParsing
```

Expected:

```json
{"status":"ok"}
```

---

# 10. NGINX Reverse Proxy

## Chức năng

nginx làm gateway trước API servers.

---

# Start NGINX

```powershell
docker start mangahub-nginx
```

---

# Test Proxy

```powershell
curl http://localhost/api/ping -UseBasicParsing
```

Expected:

```json
{"message":"pong"}
```

---

# 11. Load Balancing

## Chức năng

nginx chia request cho:

* 8081
* 8082
* 8083

---

# Run 3 API Servers

## Terminal 1

```powershell
$env:PORT="8081"
go run cmd/api-server/main.go
```

---

## Terminal 2

```powershell
$env:PORT="8082"
go run cmd/api-server/main.go
```

---

## Terminal 3

```powershell
$env:PORT="8083"
go run cmd/api-server/main.go
```

---

# Test Load Balancing

```powershell
for ($i=0; $i -lt 10; $i++) {
  Invoke-WebRequest -Uri http://localhost/api/manga `
  -Headers @{"Authorization"="Bearer $token"} |
  Select-Object -Expand Content
}
```

---

# Expected

Response sẽ đổi:

```json
"server_port":"8081"
```

↓

```json
"server_port":"8082"
```

↓

```json
"server_port":"8083"
```

---

# 12. Reading List Feature

## Chức năng

user thêm manga vào library cá nhân.

---

# Add To Reading List

```powershell
Invoke-WebRequest -Uri http://localhost:8080/reading-list `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"username":"liem","manga_id":"blue-box"}'
```

---

# Get Reading List

```powershell
curl "http://localhost:8080/reading-list?username=liem" -UseBasicParsing
```

---

# Expected

```json
[
  {
    "id":1,
    "username":"liem",
    "manga_id":"blue-box"
  }
]
```

---

# 13. Docker Support

## Chức năng

containerized nginx service.

---

# Check Containers

```powershell
docker ps
```

Expected:

```text
mangahub-nginx
```
