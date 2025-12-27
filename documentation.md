This is a temporary documentation file to document your features.
Don't overwrite this file, only append to it.

In here, after solving an issue successfully, **(IF AND ONLY IF IT IS MENTIONED IN THE NOTES)** , document everything you've added/changed.

Format of documentation: </br>

{your github id} {issue number and name}: </br>
- List of all features you've implemented, any particular format regarding them etc. Screenshots are allowed and appreciated.
- also - zero emojis. or em dashes. 
- Begin writing documentation from after the heading.

# DOCUMENTATION

### {LooninS} {#58 Time to mine}

- Added chat.go, changed main.go to call chat.go
- chat.go has 3 functions:
```go
connectToEchoServer() // connects to the server, sends the username, printout the server message
getUsername() // gets the username from the user/stdin and returns it as a string
getTimestamp() // returns the current timestamp as a string in the format of "02/01/2006 03:04:05 PM" 
```
- go.mod was updated to include the module github.com/gorilla/websocket
- tested the code locally 
