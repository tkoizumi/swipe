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
### Parsing JSON for a specific field

To parse (single or multiple) JSON response for a specific field, use the `-E` option:
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -q id=1 -E title
```
### Parsing JSON for multiple fields

To parse (single or multiple) JSON response for multiple fields, use the `-E` option with a comma separated values of field names:
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -q id=1 -E 'title,id'
```
### Extract JSON and transform API response  

To extract and transform API responses effortlessly, use the `-P` flag, you can define structured mappings with flexible expressions.  
```sh
swipe https://jsonplaceholder.typicode.com/posts -X GET -P '{"#": .id, "name": .title}' -o test.json -q id=1
```
