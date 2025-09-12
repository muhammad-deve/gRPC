# gRPC (Google Remote Procedure Call)

## 1. What is RPC?
**RPC (Remote Procedure Call)** is a concept that allows a program to execute a procedure (function) on another machine as if it were a local function call.

Examples of RPC frameworks:
- JSON-RPC
- XML-RPC
- gRPC
----

## 2. What is gRPC?
**gRPC** is an open-source, high-performance RPC framework developed by **Google** 
It is designed for efficient communication between services in **distributed systems** such as microservices.

### Key Features of gRPC
- **Language-agnostic** (supports many programming languages)
- **High performance** (built on top of HTTP/2)
- **Supports streaming** (client, server, and bidirectional)
- **Uses Protocol Buffers (Protobuf)** for efficient serialization
- **Auto-generates code** from `.proto` files
- **Secure communication** (supports TLS/SSL)

---

## 3. How gRPC Works
1. Define your service and messages in a `.proto` file.
2. The `.proto` file acts as a **contract** between services.
3. Use the `protoc` compiler to generate client and server code in your target language.
4. Implement the server logic and run the gRPC server.
5. The client calls methods as if they were local functions.

### Example `.proto` file
```proto
syntax = "proto3";

service UserService {
  rpc GetUser (GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
  string id = 1;
}

message GetUserResponse {
  string id = 1;
  string name = 2;
  string email = 3;
}
```

---

## 4. Types of gRPC Communication
gRPC supports **four communication types**:

1. **Unary RPC**
   - Client sends a single request → Server returns a single response.

2. **Server Streaming RPC**
   - Client sends a single request → Server sends multiple responses (stream).

3. **Client Streaming RPC**
   - Client sends multiple requests (stream) → Server sends a single response.

4. **Bidirectional Streaming RPC**
   - Both client and server send a stream of messages to each other.

---

## 5. gRPC vs REST
| Feature              | gRPC                              | REST                    |
|-----------------------|-----------------------------------|-------------------------|
| Protocol              | HTTP/2                           | HTTP/1.1                |
| Data format           | Protobuf (binary, compact)       | JSON (text, verbose)    |
| Performance           | Faster (binary, multiplexing)    | Slower (text, no multiplexing) |
| Streaming             | Supported (bi-directional)       | Limited (SSE, WebSocket) |
| Code generation       | Automatic from `.proto`          | Manual, boilerplate     |
| Browser support       | Limited (needs gRPC-Web)         | Native support          |

---

## 6. When to Use gRPC?
Use **gRPC** when:
- You have **microservices** that need to communicate efficiently.
- **Low latency and high throughput** are critical.
- You need **bi-directional streaming** (e.g., chat apps, IoT, real-time updates).
- You want **strict contracts** between services with auto-generated code.

Use **REST** when:
- You need direct **browser-to-server communication**.
- Simplicity and human-readable APIs are more important than performance.
- You are building **public APIs** widely consumed by different clients.

---

## 7. Pros and Cons of gRPC

### ✅ Pros
- High performance and efficiency (binary Protobuf + HTTP/2).
- Strongly typed contracts between services.
- Supports streaming and multiplexing.
- Auto code generation reduces boilerplate.

### ❌ Cons
- Not natively supported by browsers (requires gRPC-Web).
- Harder to debug compared to JSON APIs.
- Learning curve for Protobuf and gRPC tooling.

---

## 8. Summary
- **RPC** is a concept.  
- **gRPC** is a modern, fast, language-agnostic framework for RPC, created by Google.  
- It’s widely used in **microservices architectures** because of its **speed, reliability, and streaming capabilities**.  
- REST is still preferred for browser-facing APIs, but gRPC is the better choice for **service-to-service communication** in distributed systems. 
