# Teltech coding challenge


## How to run

This app is available on this URL https://challenge-hfpike7vfq-uc.a.run.app.  
Vue app is available to test this app on this URL https://cc-teltech-40c37.web.app/.

App can be started by navigating into the root of the project and running commands in terminal: 

```
    export ADDR=localhost:4000

    go run main.go
```

Or by running it in docker container

```
    docker build -t cc-teltech .

    docker run --name cc-teltech -p 4000:4000 -it cc-teltech
```

## How it works

The app has 4 routes `add, subtract, divide, multiply` all of which take `x` and `y` as parameters.  
Response is returned as `Result` in `json` format.
```
    type Result struct {
        Action string  `json:"action"`
        X      float64 `json:"x"`
        Y      float64 `json:"y"`
        Answer float64 `json:"answer"`
        Cached bool    `json:"cached"`
    }
```

### Cache

I've solved caching by creating middleware which checks every request made if it has been made in last 1 minute. Responses are stored in a `map[string]Cache`.  
`Cache` is a type I've created to hold information like request URL, result, and time when it was last used.

```
    type Cache struct {
        Url      string
        Result   models.Result
        LastUsed time.Time
    }
```

### Clearing cache 

To clear `Cache` from map I've used simple `for` loop to go trough all items and check when it was last used. Check is triggered by `Ticker` which is set to send tick every second. 

## Examples

##  Add
```
curl --request GET \
  --url https://challenge-hfpike7vfq-uc.a.run.app/add\?x\=5\&y\=5


# If tested locally
curl --request GET \
  --url http://localhost:4000/add\?x\=5\&y\=5

```
### Response 
```
    {"action":"add","x":5,"y":5,"answer":10,"cached":false}
```


## Subtract
```
curl --request GET \
  --url https://challenge-hfpike7vfq-uc.a.run.app/subtract\?x\=5\&y\=5


# If tested locally
curl --request GET \
  --url http://localhost:4000/subtract\?x\=5\&y\=5
```
Response
```
    {"action":"subtract","x":5,"y":5,"answer":0,"cached":false}
```

## Divide
```
curl --request GET \
  --url https://challenge-hfpike7vfq-uc.a.run.app/divide\?x\=5\&y\=5


# If tested locally
curl --request GET \
  --url http://localhost:4000/divide\?x\=5\&y\=5
```
Response
```
    {"action":"divide","x":5,"y":5,"answer":1,"cached":false}
```

## Multiply
```
curl --request GET \
  --url https://challenge-hfpike7vfq-uc.a.run.app/multiply\?x\=5\&y\=5

# If tested locally
curl --request GET \
  --url http://localhost:4000/multiply\?x\=5\&y\=5
```
Response 
```
    {"action":"multiply","x":5,"y":5,"answer":25,"cached":false}
```

## Cached response
Response 
```
    {"action":"multiply","x":5,"y":5,"answer":25,"cached":true}
```


## Deployment

With use of Docker I've deployed this App to Google Cloud Run.



## Tests

In tests I've checked if every route handler would return correct answer and if it would return correct responses when `x` or `y` parameter is missing.
