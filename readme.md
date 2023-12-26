# simple mock server

This is a simple example app with the following configuration:

## Configuration

```
base_url: ""
port: "8080" 

routes:

  # Example route
  - path: "/hello"
    method: "GET"   
    response:
      code: 200
      body: '{"message": "Hello, World!"}'

  - path: "/goodbye" 
    method: "POST"
    response: 
      code: 200
      body: '{"message": "Goodbye, World!"}'
```

## Usage

This app exposes two simple API routes:

- `GET /hello` - Returns a JSON response with a "Hello, World!" message
- `POST /goodbye` - Returns a JSON response with a "Goodbye, World!" message 


The app will start on port 8080. You can then access the API routes via HTTP requests, for example:

```
curl http://localhost:8080/hello
```

```
curl -X POST http://localhost:8080/goodbye
```

## License

This app is released under the MIT License. See `LICENSE` for details.