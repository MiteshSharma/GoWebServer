# Go HTTP Server

We are running an HTTP server at a specified port to accept requests written in go using its net/http package. We have a build.go file which helps with building our binary for deployment purpose. We can directly build project by going to cmd package and running main.go.

We divided the responsibility of serving a request to different modules for clarity. All our requests go through middlewares which helps in logging, measuring metrics etc. A request is handled by the api package whose main responsibility is to validate request parameters and forming right response. Once a request is validated by api package, it is passed to app package to handle all business logic. App package requests repository package for fetching any data from database or cache. We have defined other packages like logger for centralized logging, metrics for managing metrics of requests, settings to keep static or dynamic config data.

    HTTP Request --> Hits Server --> Go through middleware --> API --> App --> Repository

### API

Api package is used to receive an incoming request, validate the request for any bad input parameters. Generate a proper response after running our business logic.

### App

App package's main responsibility is to execute business logic. This is the heart of our server, as it takes to request and process it to return the desired output. App package can call repository package to fetch or store data in database or cache.

### Repository

Repository package is a wrapper on database and cache, so no other package can directly access the database. This package handle all create, update, fetch and delete operation on database tables or cache.

We have divided our logic in different modules like user, auth, common, bi event handling with all core logic moved to core folder. Core logic includes:

### BI

Business intelligence package is to take events and passing them to other modules for handling. This uses an event dispatcher package which helps make handling events async.

### Metrics

Metrics package helps abstract out handling of metrics. Other packages send data to metrics package and it makes sure this data is sent to Prometheus for visualization.

### Logger

Logger package helps abstract out handling of the logger. This package helps store logs in console or file.


### Bus

Bus package helps send messages and receive responses from different modules and provide separation of responsibility by not letting them interfere with each other.

### MQ

MQ package provide support to send message to messaging queue. We have implemented SQS for this case but others can be added.


### Event Dispatcher

Event dispatcher provides support to dispatch event using a bufferent queue and go routine. This makes handling messages async.