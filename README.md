# Hopp 🐇

Hash Ownership Ping Protocol (Hopp) implementation written in Go. 

## Getting Started

#### 📄 Documentation needed.  
#### 🚫 Testing and DB not currently supported.


## 🛠️ Current Work in Progress
#### - Adding database support for integration testing.  
#### - Better documentation.  
#### - Endpoint redirects


## 🚀 Future Goals
#### - Add full API documentation.  
#### - Plugin/API for CRM during Bid phase of ping post
#### - Implement a CLI tool and GUI  


## Makefile

### Run build make command with tests:  
```bash
make all
```

### Build the application:  
```bash
make build
```

### Run the application:  
```bash
make run
```

### Create DB container:  
```bash
make docker-run
```

### Shutdown DB container:  
```bash
make docker-down
```

### DB Integration tests:  
```bash
make itest
```

### Live reload the application:  
```bash
make watch
```

### Run the test suite:  
```bash
make test
```

### Clean up binary from the last build:  
```bash
make clean
```
