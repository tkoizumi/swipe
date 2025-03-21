# Swipe: A Command-Line HTTP Client

**Swipe** is a lightweight, fast, Go-based command-line HTTP client inspired by cURL and jq. It allows you to perform HTTP requests and provides the added functionality of transforming responses using jq syntax, all without requiring additional tools.

---

## Features
- Specify HTTP methods with `-X` (e.g., GET, POST, DELETE)
- Add custom headers using `-H`
- Include request data with `-d` for POST, PUT, or PATCH requests
- Download responses to a file using `-o`
- Add query parameters conveniently with `-q`
- Transform JSON responses on the fly using `-P` with JQ syntax

---

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/tkoizumi/swipe
   ```
2. Navigate to the project directory:
   ```sh
   cd swipe
   ```
3. Build the project:
   ```sh
   make
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
or
```sh
swipe [flags] <URL> 
```

### Flags
- `-X`: Specify the HTTP method (e.g., GET, POST, PUT, DELETE). By default, swipe uses GET if no method is specified.
- `-H`: Add headers in the format `"Header-Name: value"`. Repeatable for multiple headers.
- `-d`: Provide request data (e.g., JSON body) for POST/PUT requests.
- `-o`: Download the response body to a file. Specify the filename.
- `-q`: Add query parameters in the format `key=value`. Repeatable for multiple parameters.
- `-P`: Transform JSON response using JQ syntax.

---

## Examples
### **GET Request**
Fetch all posts from an API:
```sh
swipe https://jsonplaceholder.typicode.com/posts
```

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

### **SOAP Request with Multiple Headers**
When sending a SOAP request, you might need to include multiple headers for authentication or other purposes. You can add multiple headers using the -H flag for each one:
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

### **Using File References as Payloads**
Swipe allows you to reference files as payloads using the -d flag. This feature simplifies sending large JSON, XML, or other types of payloads without embedding them directly in the command.
```sh
swipe https://jsonplaceholder.typicode.com/posts -X POST -H "Content-Type: application/json" -d @data.json
```

### Including Response Headers
To perform a GET request and include the response headers in the output:
```sh
swipe -i https://jsonplaceholder.typicode.com/posts/1
```

### Basic Authentication

Swipe supports Basic Authentication using the `-u` and `-p` flags to specify a username and password.
```sh
swipe https://httpbin.org/basic-auth/user/passwd -u user -p passwd
```

### Using Custom API Key Header

Some APIs require authentication via a custom header. To send an API key in the request header, use the `-H` option:
```sh
swipe -H "X-API-Key: YOUR_API_KEY" https://api.example.com/data
```
### Using .pem and Private Key for SOAP API Access 

To authenticate and access SOAP APIs in the **swipe** project using a **.pem certificate** and a **private key**, use the following flags:
- `-E cert.pem` – Specifies the client certificate for authentication.
- `-Y private.key` – Provides the corresponding private key for secure communication.
```sh
swipe https://api.example.com/soap-endpoint -X POST -E cert.pem -Y private.key -d some.xml
```

### Extract JSON and transform API response (jq syntax) 

To extract and transform API responses effortlessly, use the `-P` flag, you can define structured mappings with flexible expressions. The syntax is the same as jq transformations.
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -P ".[] | {title, id}"
```
### Filtering Based on Conditions
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -P "[.[] | select(.userId == 1)]"
```
### Renaming Fields in the Output
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -P ".[] | {heading: .title, post_id: .id}"
```
