

# Overview
This is the backend server that will host clients and respond to their requests to load notes,
edit notes, and save notes. It will check which clients are connected to which notes as well
as monitor which notes are being edited.

Open notes will be stored in RAM for easy access to being updated, should the client choose to
edit a note. When a note is saved, it will be written to a JSON file.

## Clients
Clients will be web-based and have a simple text editor displayed for them on their page. The
client will be able to search for notes and request a note to read. Clients can read as many
notes as they would like because they will be read-only and will only receive server-side updates
to the notes as they come.

Clients can request to edit a note in which case the server will check if the note is already
being edited. If it is not being edited already, it will lock that note and will allow the client
to stream their changes to the server. The server will then continuously update the note in RAM
and stream the changes to each client that has that particular notes open for reading.

## Server
The server will accept and monitor connections from clients. When a client connects it will keep
that client in a running list of clients as well as monitor how long the client has been
connected. The server will receive GET requests from a client's search inquiry and the server will
respond back with a list of notes that fit the search criteria.

Once a note is chosen via GET request, the server will respond will the full contents of the note.
This will be considered "read mode". Any new changes will be sent to the clients who have the note
open. Each change will be a small change packet containing only the change and not the full note.

If a client requests to edit the note, the note will become locked to where no other client can
edit the note. If another attempts to request to go into "edit mode", they will be added to a
queue to edit the note once it becomes available again. Closing the connection will remove them
from the queue.

A client in edit mode will now be able to edit the note text in their text editor instead of
simply view it, with each edit made being sent up as a change packet to the server. The server
will make the change to the RAM note and echo out the change to all other clients. Every time a
client exits edit mode the RAM note will be written to storage to ensure that changes aren't
entirely lost when power shuts off abruptly. The RAM note will only be written one last time to
storage and then removed from RAM ONLY when there are no more clients connected to that note,
otherwise, it will stay in RAM for easy access.