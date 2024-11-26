# Hopp 🐇

Hash Ownership Ping Protocol (Hopp) implementation written in Go. 
## License
This software is licensed for personal, academic, and non-commercial use only.  
Commercial use requires prior authorization. See the [LICENSE](./LICENSE) file for details.

## Getting Started

#### 📄 Documentation needed.  
#### 🚫 Testing and DB not currently supported.


## 🛠️ Current Work in Progress
#### - Adding database support for integration testing.  
#### - Better documentation.  


## 🚀 Future Goals
#### - Better Docker functionality & Docker compose support
#### - Add full API documentation.  
#### - Implement a CLI tool and GUI  
#### - Authentication


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
