# API Tester

`api_tester` is a simple and concurrent HTTP load-testing tool written in Go.  
It allows you to send multiple HTTP requests to a target endpoint with configurable concurrency and payloads, while measuring response statistics.

---

## **Features**
- Send **GET, POST, PUT** (or any HTTP method) requests.
- **Custom payload** support for POST/PUT.
- **Concurrent requests** with configurable worker count.
- Collects:
  - Total requests
  - Errors
  - Average response time
  - Status code distribution
- Easy-to-use CLI.

---

## **Installation**
Clone the repository:

```bash
git clone https://github.com/your-username/api_tester.git
cd api_tester
````

Run directly with Go:

```bash
go run ./cmd/api_tester -url http://localhost:8080 -n 10 -c 2 -method GET
```

---

## **Usage**

### **CLI Flags**

* `-url` : Target endpoint URL (default: `http://localhost:8080`)
* `-method` : HTTP method (default: `GET`)
* `-data` : JSON payload for POST/PUT requests (default: `""`)
* `-n` : Total number of requests to send (default: `10`)
* `-c` : Number of concurrent workers (default: `2`)

---

### **Examples**

#### **GET Requests**

```bash
go run ./cmd/api_tester -url http://localhost:8080/api -n 20 -c 5 -method GET
```

#### **POST Requests with Data**

```bash
go run ./cmd/api_tester -url http://localhost:8080/api -n 20 -c 5 -method POST -data '{"name":"Rudranka"}'
```

---

## **Example Output**

```bash
config:  {http://localhost:8080/api GET  10 2}
Running 10 requests with concurrency=2

--- Test Summary ---

--- Test Results ---
Total Requests: 10
Errors: 0
Average Response Time: 524ms
Status Codes: map[200:10]
```

---

## **Project Structure**

```
api_tester/
│── cmd/
│   └── api_tester/
│       └── main.go        # Entry point
│
│── internal/
│   ├── config/            # CLI flag parsing
│   ├── httpclient/        # HTTP client logic
│   ├── runner/            # Main execution logic
│   └── stats/             # Stats aggregation and reporting
│
└── go.mod                 # Go module
```

---

## **Planned Enhancements**

* Min/Max/Median/95th percentile response time.
* Requests per second (RPS).
* Output results as JSON/CSV for further analysis.

---

## **License**

MIT License.

```

---

### **Next Step**
Would you like me to **add min/max/median response time + RPS** to your `stats.Print()` function and update the README accordingly?
```
