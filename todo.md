
### **1. Start with the HTTP Server Setup**
**Goal**: Create a running HTTP server to handle requests.
- **Files**: `internal/server/server.go`, `cmd/server/main.go`
- **Steps**:
    1. Write a `StartServer()` function in `server.go` to set up the server using `net/http`.
    2. Add a simple route like `/health` to test server availability.
    3. In `main.go`, call `StartServer()` to launch the server.

**Test**:
- Use `curl` or your browser to check `http://localhost:8080/health`. It should return `200 OK` or a simple message like `"Server is running"`.

---

### **2. Implement Note Struct and Basic Storage**
**Goal**: Define the `Note` struct and enable saving/loading notes to/from JSON files.
- **Files**: `internal/notes/note.go`, `internal/notes/storage.go`
- **Steps**:
    1. Define the `Note` struct with fields like `ID`, `Title`, `Content`, and `IsLocked`.
    2. Write `SaveNoteToFile` and `LoadNoteFromFile` functions in `storage.go`.
    3. Create test data (e.g., hardcoded notes) to save and load.

**Test**:
- Manually call `SaveNoteToFile` and `LoadNoteFromFile` in a temporary `main.go` setup.
- Verify JSON files are created, and data is read correctly.

---

### **3. Set Up In-Memory Note Management**
**Goal**: Manage notes in RAM for fast access.
- **Files**: `internal/notes/ram.go`
- **Steps**:
    1. Implement the `RAMStore` struct to manage notes in memory.
    2. Add functions:
        - `AddNoteToRAM`
        - `RemoveNoteFromRAM`
        - `GetNoteFromRAM`
    3. Integrate these functions with the storage layer (e.g., load notes into RAM on startup).

**Test**:
- Write a simple script in `main.go` to add, retrieve, and remove notes from RAM.
- Confirm RAM and file storage integration works smoothly.

---

### **4. Implement Client Connection Management**
**Goal**: Allow clients to connect and associate them with notes.
- **Files**: `internal/clients/client.go`, `internal/server/server.go`
- **Steps**:
    1. Define the `Client` struct in `client.go`.
    2. Add functions:
        - `NewClient` to initialize a client.
        - `DisconnectClient` to handle disconnections.
        - `SendChangeToClient` to simulate sending updates.
    3. Use WebSockets (e.g., `gorilla/websocket`) to handle client-server communication.

**Test**:
- Use a simple WebSocket client (browser or CLI tool) to connect and disconnect.
- Confirm connections are tracked, and messages can be sent/received.

---

### **5. Add Search Functionality**
**Goal**: Allow clients to search for notes by title or content.
- **Files**: `internal/server/server.go`, `internal/notes/ram.go`
- **Steps**:
    1. Add a `/search` route in `server.go`.
    2. Implement search logic in `ram.go` to filter notes based on a query.

**Test**:
- Use `curl` or Postman to test the `/search` endpoint.
- Return a list of matching notes and verify correct results.

---

### **6. Enable Reading Notes**
**Goal**: Serve full note content when requested.
- **Files**: `internal/server/server.go`, `internal/notes/ram.go`
- **Steps**:
    1. Add a `/note/{id}` route to fetch note details by ID.
    2. Retrieve the note from RAM and return it as JSON.

**Test**:
- Test the route using note IDs.
- Confirm correct notes are served.

---

### **7. Implement Editing Functionality**
**Goal**: Allow one client to edit a note and propagate changes.
- **Files**: `internal/server/server.go`, `internal/clients/client.go`, `internal/notes/ram.go`
- **Steps**:
    1. Add a `/note/{id}/edit` route to lock the note and set the client in edit mode.
    2. Handle edit requests using WebSockets for real-time updates.
    3. Stream changes to all clients viewing the note.

**Test**:
- Simulate editing with one client and viewing with another.
- Confirm changes propagate correctly.

---

### **8. Implement Queuing for Edit Access**
**Goal**: Allow multiple clients to queue for editing a locked note.
- **Files**: `internal/notes/ram.go`, `internal/server/server.go`
- **Steps**:
    1. Add a queue in `RAMStore` for clients waiting to edit a note.
    2. Modify the `/note/{id}/edit` route to handle queuing logic.

**Test**:
- Simulate multiple clients requesting to edit the same note.
- Confirm the queue is respected.

---

### **9. Finalize RAM Persistence Logic**
**Goal**: Ensure changes in RAM are periodically written to file.
- **Files**: `internal/notes/ram.go`, `internal/notes/storage.go`
- **Steps**:
    1. Write logic to save notes to storage whenever a client exits edit mode.
    2. Remove notes from RAM if no clients are connected.

**Test**:
- Verify that no data is lost during server restarts or client disconnects.

---

### **10. Conduct End-to-End Testing**
**Goal**: Test all features together in realistic scenarios.
- Use tools like Postman and WebSocket clients to simulate client-server interactions.
- Identify edge cases, like handling simultaneous edits or large notes.
