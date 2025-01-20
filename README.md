# Swipe: A Command-Line HTTP Client

**Swipe** is a lightweight, fast, and Go-based command-line HTTP client inspired by cURL. It allows you to perform HTTP requests with ease and provides additional functionality for query parameters and response management.

---

## Features
- Specify HTTP methods with `-X` (e.g., GET, POST, DELETE).
- Add custom headers using `-H`.
- Include request data with `-d` for POST, PUT, or PATCH requests.
- Download responses to a file using `-o`.
- Add query parameters conveniently with `-q`.

---

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/swipe.git
   ```
2. Navigate to the project directory:
   ```sh
   cd swipe
   ```
3. Build the project:
   ```sh
   go build -o swipe
   ```
4. Add Swipe to your PATH (optional):
   ```sh
   export PATH=$PATH:/path/to/swipe
   ```

---

## Usage

### General Syntax
```sh
swipe <URL> [flags]
```

### Flags
- `-X`: Specify the HTTP method (e.g., GET, POST, PUT, DELETE).
- `-H`: Add headers in the format `"Header-Name: value"`. Repeatable for multiple headers.
- `-d`: Provide request data (e.g., JSON body) for POST/PUT requests.
- `-o`: Download the response body to a file. Specify the filename.
- `-q`: Add query parameters in the format `key=value`. Repeatable for multiple parameters.

---

## Examples

### **GET Request with Query Parameters**
Fetch a specific post from an API:
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -q id=1
```

### **POST Request with Data and Headers**
Create a new post with JSON data:
```sh
swipe https://jsonplaceholder.typicode.com/posts -X POST -H "Content-Type: application/json" -d '{"title":"foo","body":"bar","userId":1}'
```

### **Download Response to a File**
Save the response to `response.json`:
```sh
swipe https://jsonplaceholder.typicode.com/posts/1 -X GET -o response.json
```

### **Add Multiple Query Parameters**
Get with multiple filters:
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -q userId=1 -q title=foo
```

### **s**
Get with multiple Headers:
```sh
swipe https://example.com/soap-endpoint -X POST\
-H "Content-Type: text/xml" 
-H "SOAPAction: \"http://example.com/action\"" 
-d '<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <Example>
    </Example>
  </soap12:Body>
</soap12:Envelope>'
```
