## unique request counter API Application

This document outlines how to run the Unique request counter API application locally or within a Docker container.

### Prerequisites

* **Go:** Install Go (version 1.22 or higher recommended) from https://golang.org/dl/.
* **Docker (optional):** Install Docker (https://www.docker.com/) if you plan to run the application in a container.

### 1. Running the Application Locally

#### 1.1 Clone the Repository

```bash
  git clone git@github.com:Durga-chikkala/unique-request-counter.git
  cd unique-request-counter
```
#### 1.2 Get all the dependencies by using:
```
    go get -v -d ./...
```


#### The application will start on port 8080. Access the endpoint at:


```
http://localhost:8080/api/verve/accept?id=<ID>&endpoint=<optional-endpoint>
```

Replace <ID> with a valid identifier and <optional-endpoint> with a custom endpoint URL (optional).

#### 1.4 Sending Requests for Testing
Use curl or any REST client tool to test the application.
Basic request without endpoint parameter:



```
curl "http://localhost:8080/api/verve/accept?id=123"
```

#### 1.5 Request with endpoint parameter:

```
curl "http://localhost:8080/api/verve/accept?id=123&endpoint=[http://example.com/receive](http://example.com/receive)"
```


#### 1.6 Checking Logs
The application logs the unique request counts every minute in the requests.log file. To verify logging functionality, execute:

```
tail -f requests.log
```


### 2. Running the Application with Docker
#### 2.1 Build the Docker Image
In the project root directory, run:

```
docker build -t request_counter .
```

#### 2.2 Run the Docker Container
Run the container and map port 8000 of the container to port 8000 on your host machine:


```
docker run -p 8080:8080 request_counter
```

#### 2.3 Sending Requests
With the container running, send test requests using curl or a REST client:

```
curl "http://localhost:8080/api/verve/accept?id=123"
curl "http://localhost:8080/api/verve/accept?id=123&endpoint=[http://example.com/receive](http://example.com/receive)"
```

#### 2.4 Checking Logs in Docker
View logs from the running container using:

```
docker logs <container_id>
```

Replace <container_id> with the actual ID of your container, which you can find with:

```
docker ps
```
