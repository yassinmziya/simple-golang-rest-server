# simple-golang-rest-server

## Using Golang Encoder/Decoder instead of Marshal/Unmarshal to handle JSON

In a Go REST server, using `encoding/json`'s `Decoder`/`Encoder` instead of `json.Marshal` and `json.Unmarshal` has several advantages that improve code readability, performance, and functionality. Here's why you might prefer `Decoder` and `Encoder`:

### 1. **Streaming Support**
- **`Decoder` and `Encoder`:** Work directly with streams (e.g., `io.Reader` and `io.Writer`), making them ideal for processing large JSON payloads incrementally without loading the entire payload into memory.
- **`Marshal` and `Unmarshal`:** Require the entire JSON data to be in memory as a `[]byte`. This can lead to high memory usage for large payloads.

### 2. **Performance**
- **`Decoder` and `Encoder`:** Avoid intermediate `[]byte` allocations since they operate directly on streams.
- **`Marshal` and `Unmarshal`:** Involve additional steps to convert data to/from `[]byte`, which can be less efficient.

### 3. **Ease of Use with HTTP**
- **`Decoder` and `Encoder`:** Directly integrate with HTTP request and response bodies (`http.Request.Body` and `http.ResponseWriter`), simplifying the code.
  ```go
  // Example using Decoder
  err := json.NewDecoder(req.Body).Decode(&data)
  if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  ```
- **`Marshal` and `Unmarshal`:** Require manual handling of request/response bodies, involving more boilerplate code:
  ```go
  body, err := io.ReadAll(req.Body)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  err = json.Unmarshal(body, &data)
  ```

### 4. **Error Handling**
- **`Decoder`:** Provides more granular error handling with features like `DisallowUnknownFields`, which ensures unknown fields in JSON are detected.
  ```go
  decoder := json.NewDecoder(req.Body)
  decoder.DisallowUnknownFields() // Prevent unknown fields
  err := decoder.Decode(&data)
  ```
- **`Unmarshal`:** Doesn't natively support such checks; you'd need custom validation.

### 5. **Readability and Maintenance**
- **`Decoder` and `Encoder`:** Result in cleaner and more concise code for handling JSON directly with streams.
- **`Marshal` and `Unmarshal`:** Add complexity when dealing with streams, especially for HTTP use cases.

### When to Use `Marshal`/`Unmarshal`
- When you're working with JSON data already loaded into memory as `[]byte`.
- For small, simple payloads where the overhead of stream processing isn't a concern.

### Conclusion
For most REST server implementations, especially when dealing with HTTP request/response bodies, `json.Decoder` and `json.Encoder` are more efficient and cleaner than `json.Marshal` and `json.Unmarshal`.
