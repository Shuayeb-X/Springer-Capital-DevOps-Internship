
## Week 1: Application Development & Continuous Integration (CI) Pipeline


**Project:** Deploy Pipeline with Canary and Automatic Rollback on Nomad\
**Organization:** Springer Capital\
**Phase:** Week 1 – Application & CI Pipeline\
**Operating System:** Kali Linux\
**Application Language:** Go\
**Containerization:** Docker\
**CI Platform:** GitHub Actions

---

# 1. Introduction

This repository documents the practical implementation of Week 1 of my Springer Capital DevOps Internship assignment.

The objective of the first week was to develop a small application, containerize it using Docker, and automate the application testing, Docker image building, and image publishing process using GitHub Actions.

Instead of manually building and publishing Docker images every time the application source code changes, I implemented a Continuous Integration (CI) pipeline that automatically performs these operations whenever new code is pushed to the `main` branch.

The application was developed using Go and exposes two HTTP endpoints:

- `/` – Returns application information.
- `/health` – Returns the application's health status.

The application was containerized using a multi-stage Dockerfile and published to Docker Hub and GitHub Container Registry (GHCR).

GitHub Actions was configured to perform the following operations automatically:

1. Checkout the source code.
2. Set up the Go environment.
3. Execute automated unit tests.
4. Run Go static analysis.
5. Authenticate with GitHub Container Registry.
6. Build the Docker image.
7. Push the Docker image to GHCR.

This implementation establishes the application and CI foundation for the subsequent Nomad deployment, canary release, and automatic rollback phases.

# 2. Week 1 Objectives

The assigned Week 1 requirements were:

| No. | Assigned Requirement                 | Implementation         |
| --- | ------------------------------------ | ---------------------- |
| 01  | Develop a small demo application     | Go HTTP application    |
| 02  | Implement a `/health` endpoint       | HTTP health endpoint   |
| 03  | Create a Dockerfile                  | Multi-stage Dockerfile |
| 04  | Create GitHub Actions workflow       | YAML workflow          |
| 05  | Build Docker image on push to `main` | Automated CI           |
| 06  | Push image to lab container registry | GHCR publishing        |

All six assigned objectives were completed.

Additional work included automated Go tests, static analysis, Docker Hub publishing, and local container verification.

---

# 3. Technology Stack and Tools

The following technologies were used throughout Week 1.

| Technology     | Version / Tool                | Purpose                 |
| -------------- | ----------------------------- | ----------------------- |
| Kali Linux     | Rolling 2026.3                | Development environment |
| Go             | Version specified in `go.mod` | Backend application     |
| Git            | Installed on Kali Linux       | Version control         |
| GitHub         | Cloud repository              | Source code hosting     |
| Docker         | 29.8.0                        | Containerization        |
| Docker Hub     | Public registry               | Image storage           |
| GHCR           | GitHub Container Registry     | Image storage           |
| GitHub Actions | Workflow automation           | Continuous Integration  |
| YAML           | Workflow configuration        | CI pipeline definition  |
| JSON           | HTTP response format          | API responses           |

## 3.1 Development Environment

All application development and local testing were performed on Kali Linux.

The project directory was:

```bash
/home/zet/springer-capital/springer-demo-app
```

To enter the project directory:

```bash
cd ~/springer-capital/springer-demo-app
```

**Command explanation:**

- `cd` – Changes the current working directory.
- `~` – Represents the current user's home directory.
- `springer-capital/springer-demo-app` – Project directory.

To inspect the project files:

```bash
ls -la
```

The `-l` option displays detailed file information, while `-a` includes hidden files and directories such as `.git` and `.github`.

# 4. Understanding the Project Architecture

Before implementation, it is important to understand how the components interact.

The application development and CI pipeline follow this architecture:

![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Go%20CI_CD%20Pipeline%20Infographic.png)

### How the architecture works

**Step 1 – Application Development**

The developer writes and modifies the Go application.

**Step 2 – Version Control**

The source code is tracked using Git.

**Step 3 – Code Push**

The developer commits and pushes changes to the `main` branch.

**Step 4 – CI Trigger**

GitHub Actions detects the push event and starts the configured workflow.

**Step 5 – Code Validation**

The workflow executes unit tests and Go static analysis.

**Step 6 – Docker Image Creation**

If the validation steps succeed, Docker builds the application image.

**Step 7 – Image Publishing**

The resulting image is pushed to GitHub Container Registry.

This process reduces manual work and ensures that application code is tested before an image is published.

---

# 5. Repository Structure

The project repository contains the following important files and directories:

```text
springer-demo-app/
│
├── .github/
│   └── workflows/
│       └── docker-publish.yml
│
├── .git/
│
├── dockerfile
├── .gitignore
├── go.mod
├── main.go
├── main_test.go
├── README.md
└── springer-demo-app
```

### Explanation of each file

| File                 | Responsibility                             |
| -------------------- | ------------------------------------------ |
| `.github/`           | GitHub-specific configuration              |
| `.github/workflows/` | GitHub Actions workflow files              |
| `docker-publish.yml` | CI pipeline configuration                  |
| `.git/`              | Git version-control metadata               |
| `dockerfile`         | Instructions for building the Docker image |
| `.gitignore`         | Files excluded from Git tracking           |
| `go.mod`             | Go module configuration                    |
| `main.go`            | Main application source code               |
| `main_test.go`       | Automated application tests                |
| `README.md`          | Project documentation                      |
| `springer-demo-app`  | Locally compiled executable                |

The compiled executable is a local build artifact. The Dockerfile creates its own executable inside the builder stage, so the local binary is not required for the Docker build.

# 6. Step-by-Step Application Development

## 6.1 Why did we need a demo application?

The main objective was to create a simple application that could be used to demonstrate the complete CI/CD workflow.

A small Go HTTP application was selected because:

- Go can produce lightweight Linux executables.
- It is suitable for backend and cloud-native applications.
- It works well with Docker.
- It has a built-in HTTP package.
- It supports automated testing through its standard testing package.

The application needed to provide a health endpoint because future deployment platforms need a way to determine whether the application is responding correctly.

## 6.2 Initialize the Go project

The project uses Go modules.

The module configuration is stored in:

```text
go.mod
```

To inspect it:

```bash
cat go.mod
```

The Go module identifies the project and manages its dependencies.

## 6.3 Create the application source file

The main application source code is stored in:

```text
main.go
```

The application was implemented using Go's standard libraries.

The main packages used are:

```go
import (
    "encoding/json"
    "log"
    "net/http"
)
```

### Package explanation

**`encoding/json`**

Used to encode Go data structures into JSON responses.

**`log`**

Used to print application startup messages and errors.

**`net/http`**

Provides HTTP server functionality, request handling, and response writing.

## 6.4 Complete application source code

The application source code:

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(map[string]string{
        "application": "Springer Capital Demo App",
        "status":      "running",
    })
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
    })
}

func newHandler() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/health", healthHandler)

    return mux
}

func main() {
    log.Println("Springer Demo App starting on port 8080")

    err := http.ListenAndServe(":8080", newHandler())

    if err != nil {
        log.Fatal(err)
    }
}
```

## 6.5 Understanding the application code

### A. Root handler

```go
func rootHandler(w http.ResponseWriter, r *http.Request)
```

This function handles requests to the root endpoint.

The parameters are:

- `w` – Writes the HTTP response.
- `r` – Contains the incoming HTTP request.

The route validation:

```go
if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
}
```

Checks whether the requested path is `/`.

If the requested path is different, the application returns HTTP 404.

The response header:

```go
w.Header().Set("Content-Type", "application/json")
```

Specifies that the response content is JSON.

The JSON encoder:

```go
json.NewEncoder(w).Encode(...)
```

Encodes the application information into a JSON response.

### B. Health handler

```go
func healthHandler(w http.ResponseWriter, r *http.Request)
```

This function handles health-check requests.

First, it verifies the HTTP method:

```go
if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

Only GET requests are accepted.

If another method is used, the application returns HTTP 405.

For a valid request:

```go
w.WriteHeader(http.StatusOK)
```

The application returns HTTP 200.

The JSON response is:

```json
{
  "status": "healthy"
}
```

This endpoint is particularly important because deployment platforms can periodically request it to determine application health.

### C. HTTP router

```go
func newHandler() http.Handler
```

Creates an HTTP request multiplexer.

```go
mux := http.NewServeMux()
```

Initializes a new router.

```go
mux.HandleFunc("/", rootHandler)
mux.HandleFunc("/health", healthHandler)
```

Registers the application's two HTTP handlers.

Returning the handler separately makes it easier to test the application without starting a real network server.

### D. Main function

```go
func main()
```

This is the application's entry point.

```go
http.ListenAndServe(":8080", newHandler())
```

Starts the HTTP server on port 8080.

The server listens on all available network interfaces because the address is `:8080`, rather than `127.0.0.1:8080`.

# 7. API Endpoint Documentation

The application exposes two primary endpoints.

## 7.1 Root Endpoint

**Endpoint:**

```http
GET /
```

**Purpose:**

Returns basic information about the running application.

**Request:**

```bash
curl http://localhost:8080/
```
![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/running-terminal.png)


**Expected response:**

```json
{
  "application": "Springer Capital Demo App",
  "status": "running"
}
```
![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/running-web.png)

**Expected HTTP status:**

```text
200 OK
```

## 7.2 Health Endpoint

**Endpoint:**

```http
GET /health
```

**Purpose:**

Provides an HTTP health-check endpoint for monitoring application availability.

**Request:**

```bash
curl http://localhost:8080/health
```
![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/endpoint-terminal.png)

**Expected response:**

```json
{
  "status": "healthy"
}
```
```bash
http://localhost:8080/health
```
![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/endpoint-web.png)

**Expected HTTP status:**

```text
200 OK
```

This endpoint will later be used by Nomad and its service-health monitoring configuration.

## 7.3 Unsupported HTTP Methods

The health endpoint only accepts GET requests.

For example:

```bash
curl -X POST http://localhost:8080/health
```

Expected result:

```text
405 Method Not Allowed
```

## 7.4 Unknown Routes

A request to an unregistered route should return:

```text
404 Not Found
```
For example:

```bash
curl http://localhost:8080/unknown
```
![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/invalid-terminal.png)


```bash
http://localhost:8080/unknown
```

![image alt](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/inavlid%20endpoint.png)

---

# 8. Application Testing and Code Quality

Writing application code is not enough. We also need to verify that it behaves as expected.

For this reason, automated unit tests were implemented.

## 8.1 Why unit testing?

Unit testing verifies individual functions and application behavior without requiring the complete deployment environment.

In this project, we wanted to verify:

- The root endpoint returns the expected response.
- The health endpoint returns a healthy status.
- Unsupported HTTP methods are rejected.
- Unknown routes return the correct HTTP response.

## 8.2 Test file

All application tests are stored in:

```text
main_test.go
```

Go provides a built-in testing package:

```go
import "testing"
```

The project contains four test cases.

| Test                                | Expected behavior                             |
| ----------------------------------- | --------------------------------------------- |
| `TestRootHandler`                   | Root endpoint returns application information |
| `TestHealthHandler`                 | Health endpoint returns healthy status        |
| `TestHealthHandlerMethodNotAllowed` | Non-GET request returns 405                   |
| `TestUnknownRoute`                  | Unknown path returns 404                      |

## 8.3 Execute automated tests

Command:

```bash
go test -v ./...
```

### Command breakdown

- `go` – Executes the Go toolchain.
- `test` – Runs automated tests.
- `-v` – Enables verbose output, displaying individual test results.
- `./...` – Runs tests in the current module and its subdirectories.

A successful execution should show each test with a PASS result.

The four test cases passed during implementation.

## 8.4 Format the source code

Command:

```bash
gofmt -w .
```

Explanation:

- `gofmt` formats Go source code.
- `-w` writes the formatted result back to the files.
- `.` targets the current directory.

Consistent formatting makes Go code easier to maintain and review.

## 8.5 Build the application

Command:

```bash
go build ./...
```

This checks whether the application can be compiled successfully.

A successful build confirms that there are no compilation errors in the packages being built.

## 8.6 Run static analysis

Command:

```bash
go vet ./...
```

`go vet` examines Go code for suspicious constructs and common programming mistakes.

It complements unit tests by checking issues that might not be caught by the test cases.

## 8.7 Why were these checks added to CI?

Running these commands manually is useful, but developers may forget to execute them before pushing code.

By including them in GitHub Actions, every push to `main` automatically triggers the same validation process.

# 9 Dockerizing Springer Demo Go Application



# Dockerizing Springer Demo Go Application
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/Go%20Application%20Containerization%20Workflow.png)

## 9.1. Creating the Dockerfile

First, I created a file named `dockerfile` in the project root directory.

### Dockerfile Code

```dockerfile
# Stage 1: Build the Go application
FROM golang:1.26-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy project files into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o exe_binary .

# Stage 2: Create the lightweight runtime image
FROM alpine:3.22

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/exe_binary .

# Expose application port
EXPOSE 8080

# Start the application
CMD ["./exe_binary"]
```

### Dockerfile Explanation

| Instruction | Explanation |
|---|---|
| `FROM golang:1.26-alpine AS builder` | Uses Go Alpine as the build environment. |
| `WORKDIR /app` | Sets `/app` as the working directory. |
| `COPY . .` | Copies project files into the builder stage. |
| `RUN CGO_ENABLED=0 GOOS=linux go build -o exe_binary .` | Compiles the Go application into a Linux executable binary. |
| `FROM alpine:3.22` | Creates a separate lightweight runtime stage. |
| `COPY --from=builder /app/exe_binary .` | Copies only the compiled binary from the builder stage. |
| `EXPOSE 8080` | Documents the application's listening port. |
| `CMD ["./exe_binary"]` | Starts the application when the container runs. |

### Why Did I Use a Multi-Stage Build?

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/Docker%20and%20GHCR.png)

**Builder Stage:**
- Uses the Go compiler.
- Compiles the Go source code.
- Generates the executable binary.

**Runtime Stage:**
- Uses a lightweight Alpine Linux image.
- Copies only the compiled executable.
- Does not include the Go compiler or source code.

**Benefits:**
- Reduces the final Docker image size.
- Keeps the runtime environment lightweight.
- Separates build dependencies from runtime dependencies.
- Improves deployment efficiency.
---

## 9.2. Building the Docker Image

After creating the Dockerfile, I built the Docker image from the project root directory.

### Build Command

```bash
docker build -t springer-demo-app:v1 -f dockerfile .
```

### Command Explanation

| Command | Explanation |
|---|---|
| `docker build` | Builds a Docker image. |
| `-t` | Assigns a name and tag to the image. |
| `springer-demo-app` | Docker image name. |
| `v1` | Image version tag. |
| `-f dockerfile` | Specifies the Dockerfile name. |
| `.` | Specifies the current directory as the build context. |

### Verify the Docker Image

```bash
docker images
```

This command lists all locally available Docker images.

---

## 9.3. Running the Docker Container

After building the image, I created and started a Docker container to test the application.

### Run Command

```bash
docker run -d \
  --name springer-demo \
  -p 8080:8080 \
  springer-demo-app:v1
```

### Command Explanation

| Option | Explanation |
|---|---|
| `docker run` | Creates and starts a container. |
| `-d` | Runs the container in detached mode. |
| `--name springer-demo` | Assigns a name to the container. |
| `-p 8080:8080` | Maps host port 8080 to container port 8080. |
| `springer-demo-app:v1` | Specifies the Docker image. |

### Verify the Running Container

```bash
docker ps
```

To view all containers, including stopped containers:

```bash
docker ps -a
```

### Check Container Logs

```bash
docker logs springer-demo
```

---

## 9.4. Testing the Application

After starting the container, I tested the application using `curl`.

### Test Application Endpoint

```bash
curl http://localhost:8080/
```

### Test Health Endpoint

```bash
curl http://localhost:8080/health
```

### Expected Health Response

```json
{
  "status": "healthy"
}
```

These tests helped verify that the application was accessible through the mapped host port.

---

## 9.5. Docker Hub Authentication

After successfully testing the application locally, I authenticated with Docker Hub.

Docker Hub is a cloud-based registry used to store and distribute Docker images.

### Login Command

```bash
docker login
```

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/docker%20login%2Cpush.png)

Alternatively, specify the Docker Hub username:

```bash
docker login -u shuayeb009
```

I authenticated using my Docker Hub credentials.

**Security Note:**

- Never store Docker Hub passwords or access tokens inside a Dockerfile.
- Never commit credentials to GitHub.
- Use secure authentication methods such as Docker access tokens.

---

## 9.6. Tagging the Docker Image

Before publishing the image, I created a Docker Hub-compatible tag for the locally built image.

### Tag Command

```bash
docker tag springer-demo-app:v1 \
  shuayeb009/springer-demo-app:v1
```

### Command Explanation

| Part | Explanation |
|---|---|
| `docker tag` | Creates another tag for an existing image. |
| `springer-demo-app:v1` | Existing local image. |
| `shuayeb009` | Docker Hub username. |
| `springer-demo-app` | Docker Hub repository name. |
| `v1` | Image version tag. |

### Why Is Tagging Required?

Docker Hub uses the following image naming convention:

```text
USERNAME/REPOSITORY:TAG
```

Therefore, I tagged the local image with my Docker Hub username and repository name before publishing it.

The final image name became:

```text
shuayeb009/springer-demo-app:v1
```

### Verify the Tagged Image

```bash
docker images
```

---

## 9.7. Pushing the Image to Docker Hub

After tagging the image, I uploaded it to my Docker Hub repository.

### Push Command

```bash
docker push shuayeb009/springer-demo-app:v1
```
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/docker%20hub-push-ui.png)
### Command Explanation

| Command | Explanation |
|---|---|
| `docker push` | Uploads a Docker image to a registry. |
| `shuayeb009` | Docker Hub username. |
| `springer-demo-app` | Target repository. |
| `v1` | Image version tag. |

Docker uploaded the image layers to Docker Hub.

After the push operation completed successfully, the image became available in my Docker Hub repository.

---

## 9.8. Verifying the Published Image

Finally, I verified that the Docker image was successfully published to Docker Hub.

### Pull the Published Image

To verify that the image is available remotely, I can pull it from Docker Hub.

```bash
docker pull shuayeb009/springer-demo-app:v1
```

### Verify the Downloaded Image

```bash
docker images
```

### Run the Published Image

```bash
docker run -d \
  --name springer-demo-test \
  -p 8081:8080 \
  shuayeb009/springer-demo-app:v1
```

Test the application:

```bash
curl http://localhost:8081/
```

Health check:

```bash
curl http://localhost:8081/health
```

Using port `8081` avoids conflicts if the original container is still running on port `8080`.

### Docker Hub Repository

Published Docker image:

```text
shuayeb009/springer-demo-app:v1
```

Repository URL:

https://hub.docker.com/r/shuayeb009/springer-demo-app

---

## 9.9. Complete Docker Command Reference

All the major commands used in this Docker workflow are listed below.

| Step | Command | Purpose |
|---|---|---|
| 1 | `docker build -t springer-demo-app:v1 -f dockerfile .` | Build Docker image |
| 2 | `docker images` | List Docker images |
| 3 | `docker run -d --name springer-demo -p 8080:8080 springer-demo-app:v1` | Run container |
| 4 | `docker ps` | Check running containers |
| 5 | `docker logs springer-demo` | Check container logs |
| 6 | `curl http://localhost:8080/` | Test application |
| 7 | `docker login` | Authenticate with Docker Hub |
| 8 | `docker tag springer-demo-app:v1 shuayeb009/springer-demo-app:v1` | Tag image |
| 9 | `docker push shuayeb009/springer-demo-app:v1` | Push image |
| 10 | `docker pull shuayeb009/springer-demo-app:v1` | Pull published image |

---
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/docker%20login%2Cpush.png))

### Final Docker Image

```text
shuayeb009/springer-demo-app:v1
```

## 9.10. Pushing the Docker Image to Docker Hub

After authentication and tagging, I pushed the Docker image to Docker Hub.

### Push Command

```bash
docker push shuayeb009/springer-demo-app:v1
```

### Command Explanation

| Part | Meaning |
|---|---|
| `docker push` | Uploads a Docker image to a registry. |
| `shuayeb009` | Docker Hub username. |
| `springer-demo-app` | Docker Hub repository name. |
| `v1` | Image version tag. |

Docker uploads the image layers to Docker Hub.

After successful execution, the image becomes available in the remote repository.

---

## 9.11. Verifying the Published Image

After pushing the image, I verified its availability in Docker Hub.

### Docker Hub Repository

[View Springer Demo Docker Image](https://hub.docker.com/r/shuayeb009/springer-demo-app/tags)

### Pull Command

To verify that the published image is accessible from Docker Hub, I pulled it back to my local machine.

```bash
docker pull shuayeb009/springer-demo-app:v1
```

### Verify Local Images

```bash
docker image ls
```

This command lists locally available Docker images and confirms that the published image has been downloaded successfully.

### Inspect Docker Image

```bash
docker image inspect springer-demo-app:v1
```

This command displays detailed information about the Docker image, including its configuration, image ID, architecture, and layers.

---

## 9.12. Complete Docker Workflow

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/Go%20Application%20Containerization%20Workflow.png)

---

## 913. Important Docker Commands Summary

| Command | Purpose |
|---|---|
| `docker version` | Checks Docker version. |
| `docker info` | Displays Docker engine information. |
| `docker build -t springer-demo-app:v1 -f dockerfile .` | Builds the Docker image. |
| `docker images` | Lists local Docker images. |
| `docker image ls` | Lists local Docker images. |
| `docker run -d --name springer-demo -p 8080:8080 springer-demo-app:v1` | Creates and runs a container. |
| `docker ps` | Lists running containers. |
| `docker ps -a` | Lists all containers. |
| `docker logs springer-demo` | Displays container logs. |
| `docker login` | Authenticates with Docker Hub. |
| `docker tag springer-demo-app:v1 shuayeb009/springer-demo-app:v1` | Tags the image for Docker Hub. |
| `docker push shuayeb009/springer-demo-app:v1` | Pushes the image to Docker Hub. |
| `docker pull shuayeb009/springer-demo-app:v1` | Downloads the published image. |
| `docker image inspect springer-demo-app:v1` | Displays image configuration. |

---


### Final Docker Image

```text
shuayeb009/springer-demo-app:v1
```

### Docker Hub Repository

https://hub.docker.com/r/shuayeb009/springer-demo-app/tags


# 10 Publishing Docker Image to GitHub Container Registry (GHCR)

## 10.1. GHCR Authentication

First, I created a GitHub Personal Access Token (Classic) with the following permissions:

- `write:packages`
- `read:packages`

Then, I authenticated with GitHub Container Registry using Docker.

### Login Command

```bash
echo "$GHCR_TOKEN" | docker login ghcr.io -u Shuayeb-X --password-stdin
```

**Command Explanation:**

| Part | Explanation |
|---|---|
| `docker login` | Authenticates with a container registry. |
| `ghcr.io` | GitHub Container Registry address. |
| `-u Shuayeb-X` | GitHub username. |
| `--password-stdin` | Reads the access token securely from standard input. |

**Expected Output:**

```text
Login Succeeded
```

## 10.2. Building the Docker Image

I built the Springer Demo application image using the Dockerfile.

```bash
docker build -t springer-demo-app:v1 -f dockerfile .
```

**Explanation:**

- `docker build`: Builds a Docker image.
- `-t`: Assigns an image name and tag.
- `springer-demo-app:v1`: Local image name and version.
- `-f dockerfile`: Specifies the Dockerfile.
- `.`: Uses the current directory as the build context.

## 10.3. Tagging the Image for GHCR

After building the image, I tagged it with the GHCR repository name.

```bash
docker tag springer-demo-app:v1 ghcr.io/shuayeb-x/springer-demo-app:v1
```
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/GHCR%20login%20and%20push-2.png)

**Explanation:**

The GHCR image naming format is:

```text
ghcr.io/USERNAME/IMAGE_NAME:TAG
```

- `ghcr.io`: Registry address.
- `shuayeb-x`: GitHub username in lowercase.
- `springer-demo-app`: Image repository name.
- `v1`: Image version.

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/GHCR%20login%20and%20push-1.png)


## 10.4. Pushing the Image to GHCR

After authentication and tagging, I pushed the image to GitHub Container Registry.

```bash
docker push ghcr.io/shuayeb-x/springer-demo-app:v1
```
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/GHCR%20login%20and%20push-2.png)

Docker uploaded the image layers to GHCR.

**Published Image:**

```text
ghcr.io/shuayeb-x/springer-demo-app:v1
```

## 10.5. Pulling and Verifying the Image

I verified the published image by pulling it from GHCR.

```bash
docker pull ghcr.io/shuayeb-x/springer-demo-app:v1
```

Then, I inspected the image:

```bash
docker image inspect ghcr.io/shuayeb-x/springer-demo-app:v1
```

I verified the following properties:

- Image ID
- Repository tags
- Architecture (`amd64`)
- Operating system (`linux`)
- Exposed port (`8080/tcp`)
- Working directory (`/app`)
- Startup command (`./exe_binary`)

## 10.6. Making the GHCR Package Public

I changed the package visibility to Public from GitHub Container Registry package settings.

This allows Nomad and other Docker environments to pull the image without registry authentication.

**GHCR Package:**

[Springer Demo App - GHCR Package](https://github.com/users/shuayeb-x/packages/container/package/springer-demo-app)


## 10.7 GHCR Publishing Commands

| Command | Purpose |
|---|---|
| `echo "$GHCR_TOKEN" \| docker login ghcr.io -u Shuayeb-X --password-stdin` | Authenticates with GitHub Container Registry. |
| `docker build -t springer-demo-app:v1 -f dockerfile .` | Builds the Docker image. |
| `docker tag springer-demo-app:v1 ghcr.io/shuayeb-x/springer-demo-app:v1` | Tags the image for GHCR. |
| `docker push ghcr.io/shuayeb-x/springer-demo-app:v1` | Pushes the image to GHCR. |
| `docker pull ghcr.io/shuayeb-x/springer-demo-app:v1` | Downloads the image from GHCR. |
| `docker image inspect ghcr.io/shuayeb-x/springer-demo-app:v1` | Displays GHCR image configuration. |


# 11 GitHub Actions Workflow

First, I created a directory named `.github/workflows` in the project's root directory.

```bash
mkdir -p .github/workflows
```

Then, I created the workflow file:

```bash
nano .github/workflows/docker-publish.yml
```

I wrote the GitHub Actions workflow configuration in this file.

### **Workflow Code**

```yaml

name: Build and Push Docker Image to GHCR

on:
  push:
    branches:
      - main

permissions:
  contents: read
  packages: write

jobs:
  build-and-push:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout source code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Run Go tests
        run: go test -v ./...

      - name: Run Go vet
        run: go vet ./...

      - name: Login to GitHub Container Registry
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Build and push Docker image
        uses: docker/build-push-action@v6
        with:
          context: .
          file: dockerfile
          push: true
          tags: ghcr.io/shuayeb-x/springer-demo-app:v1

```

### **Workflow Code Explanation**

| Configuration | Purpose |
|---|---|
| `name` | Defines the workflow name. |
| `on: push` | Triggers the workflow when a Git push occurs. |
| `branches: main` | Runs only when changes are pushed to the main branch. |
| `permissions` | Provides repository read and package write permissions. |
| `runs-on` | Uses an Ubuntu runner. |
| `actions/checkout@v4` | Checks out the source code from the GitHub repository. |
| `actions/setup-go@v5` | Sets up the Go environment. |
| `go test -v ./...` | Runs automated Go tests. |
| `go vet ./...` | Performs static analysis of the Go code. |
| `docker/login-action@v3` | Logs in to GHCR. |
| `GITHUB_TOKEN` | Used for GitHub Actions authentication. |
| `docker/build-push-action@v6` | Builds and pushes the Docker image. |
| `push: true` | Enables pushing the image to the registry. |
| `tags` | Defines the GHCR image name and version. |

### **How Did I Run the GitHub Actions Workflow?**

After creating and saving the workflow file, I committed and pushed it to the Git repository.

```bash
git add .github/workflows/docker-publish.yml
git commit -m "Add GitHub Actions workflow for GHCR publishing"
git push origin main
```

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/worklow-yaml.png)

Then, I went to the **Actions** tab of the GitHub repository and verified the workflow execution.

During the first run, I encountered a `403 Forbidden` error due to GHCR permissions. Later, I granted the repository **Write** permission from GHCR Package Settings → Manage Actions access.

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/failed.png)

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/failed%20log.png)

### After that, the workflow executed successfully.
![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/workflow%20runs%20successfully.png)


### **Workflow Link**

[GitHub Actions – Workflow Runs](https://github.com/Shuayeb-X/devops-intern-final/actions?utm_source=chatgpt.com)

### Automated the trigger whenever any code push to the main branch

![Image Alt Text](https://github.com/Shuayeb-X/Springer-Capital-DevOps-Internship/blob/main/Project%201%3A%20Deploy%20pipeline%20with%20canary%20and%20automatic%20rollback%20on%20Nomad/Week%201%20%3AApplication%20Development%20%26%20Continuous%20Integration%20(CI)%20Pipeline/docs/Day%202/autmated%20success.png)

**Final Result:** I automated the entire CI/CD pipeline using GitHub Actions. Whenever code is pushed to the `main` branch, GitHub Actions automatically runs Go tests, Go vet, builds the Docker image, authenticates with GHCR, and pushes the image to GitHub Container Registry without manual intervention.

