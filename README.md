# sakalli

If you're lazy like me and you don't want to spend some time implementing a websocket server in your backend to send notifications/page updates to your frontend client then this is the tool for you!

## Usage

## First of all let's run the server 

### Build and run
1. Build

  * `make build`
2. Run the server, obviously!
  * `./bin/sakalli`  runs on port 8080 by default
  * `./bin/sakalli --port 8000` will run on port 8000

### Or run via `go run`

  * `make run`
  * or `make run PORT=8000`

### or simply use docker 
  * `docker run -p 8080:8080 zedobaia/sakalli-docker:latest`

That's it for the server!
## let set up a simple javascript client
Let's assume you're running on `localhost:8080`
and `sakalli.js` is at `/static/sakalli.js`
 ```
 <script src="/static/sakalli.js"></script>
 script>
    const sakalli = new Sakalli("localhost:8080"); // or your server ip
    window.onload = function () {
      let connected = sakalli.connect("555"); // 555 is the user id, to isolate each user 
      if (connected) {
        sakalli.on('notification', data => {
          console.log(data);
        })
        sakalli.on('opened', function () {
          console.log("connection established!!");
        })
        sakalli.on('closed', () => {
          console.log("connection closed!!");
        })
      }
    };
  </script>
 ```
 Wow, that was easy!!
 ## Now let's notify the client from your backend, this is the last part I promise
 So this is basically just a POST request to sakalli service, it will take care of the rest.
 There are two endpoints
| Endpoint  |  JSON Format |  Description |
| ------------ | ------------ | ------------ |
| `http://localhost:8080/send/<user-id>/`  |  `{"page" : "string", "type" : "string", "data" : {"prop1" : "value1", "prop2" : "value2"}, //free form}`| To send an event to all connected clients for a specific user |
| `http://localhost:8080/broadcast/`  |  `{page" : "string", type" : "string", data" : {"prop1" : "value1", "prop2" : "value2"}, ids" : ["user1_id", "user2_id"]}` | To send an event to all connected clients for multiple users |


Example sending to all connected clients for user with id "555"
```
curl --location --request POST 'http://localhost:8080/send/555' \
--header 'Content-Type: application/json' \
--data '{
    "page" : "/books/",
    "type" : "book_created",
    "data" : {
        "object_type" : "Book",
        "id" : 15,
        "value" : null
    }
}'
```

Example broadcasting to all connected clients for users with ids "555" and "kkk"

```
curl --location --request POST 'http://localhost:8080/broadcast/' \
--header 'Content-Type: application/json' \
--data '{
    "page" : "",
    "type" : "broadcast",
    "data" : {
        "object_type" : "Book",
        "id" : 15,
        "value" : null
    },
    "ids" : ["555", "kkk"]
}'
```
