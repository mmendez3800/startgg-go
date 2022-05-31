# StartGG API Wrapper (WIP)
This project implements a wrapper for the StartGG API using Go (Golang). For those that do not wish to go through the full implementation of the GraphQL language, this wrapper provides an easy and concise way of to make API calls for the StartGG platform.

## General Info
It is currently a work-in-progress and not finalized. Additional changes and updates are still needed such as implementation of other objects and the ability to perform calls for nested objects.

## Technologies
The programs in this project were run using the following:
* go1.15.6

## Setup
To use the wrapper, you will first want to install it into your working package or module. This can be done by running the below command in terminal:
```bash
go install github.com/mmendez3800/startgg-go
```

After doing this, you can then easily import the wrapper throughout your files using the below line of code:
```go
import "github.com/mmendez3800/startgg-go"
```

## Execution

### Authentication Token
In order to perform any call related to the Start GG API, the authentication token needs to be indicated. This is done by calling
the below method. If this is not done, none of the subsequent calls using the wrapper can be used. The string in this example would be replaced by the token provided by Start GG.

```go
startgg.SetAuthToken("STARTGG_TOKEN")
```

### Player
```go
import (
    "fmt"

    "github.com/mmendez3800/startgg-go"
)

// Indicate the authorization token to be used for API calls
startgg.SetAuthToken("STARTGG_TOKEN")

// Pull the results of the Player object by ID reference
idResult, err := startgg.PlayerByID(219999)
if err != nil {
    fmt.Printf("Error - %s", err)
}
fmt.Printf("Results - %+v", idResult)
```

### Tournament
```go
import (
    "fmt"

    "github.com/mmendez3800/startgg-go"
)

// Indicate the authorization token to be used for API calls
startgg.SetAuthToken("STARTGG_TOKEN")

// Pull the results of the Tournament object by ID reference
idResult, err := startgg.TournamentByID(326786)
if err != nil {
    fmt.Printf("Error - %s", err)
}
fmt.Printf("Results - %+v", idResult)

// Pull the results of the Tournament object by slug reference
slugResult, err := startgg.TournamentBySlug("tournament/genesis-8")
if err != nil {
    fmt.Printf("Error - %s", err)
}
fmt.Printf("Results - %+v", slugResult)
```

### User
```go
import (
    "fmt"

    "github.com/mmendez3800/startgg-go"
)

// Indicate the authorization token to be used for API calls
startgg.SetAuthToken("STARTGG_TOKEN")

// Pull the results of the User object by ID reference
idResult, err := startgg.UserByID(41259)
if err != nil {
    fmt.Printf("Error - %s", err)
}
fmt.Printf("Results - %+v", idResult)

// Pull the results of the User object by slug reference
slugResult, err := startgg.UserBySlug("user/fd58260e")
if err != nil {
    fmt.Printf("Error - %s", err)
}
fmt.Printf("Results - %+v", slugResult)
```
