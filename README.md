# MUZZ Application

![logo](resources/muzz.png)

Implementation of a simple application that fulfils the requirements of the MUZZ backend exercise.  
Discussion of **architectural and design choices** were left to the end of the document.

See API section for routes. They are _as per_ the spec, with the `discover` endpoint enriched with the following query params (all optional):

```txt
- min_age 
- max_age
- gender 
- sort_distance
- sort_attractiveness
```

- [Introduction](#introduction)
- [How To Run](#how-to-run)
  - [Integration Tests](#integration-tests)
- [API](#api)
- [Code Structure](#code-structure)
- [Data Model](#data-model)
- [Architecture and Design Choices](#architecture-and-design-choices)
  - [Attractiveness Calculation](#attractiveness-calculation)

## Introduction

This is a very simple Go app, made using Go 1.22.4
Main external libs used:

- [Gin](https://github.com/gin-gonic/gin) for routing, request context and middleware.
- [pgx](https://github.com/jackc/pgx) for PostgreSQL connection and driver.

App was developed on a Linux environment, with Docker version **24.0.5** and Docker Compose version **2.20.3**.

## How To Run

The application is fully containerized, so assuming Docker and Docker Compose are installed, you can run the app with
the following command:

```bash
# start the app in detached mode
docker-compose up -d app

# tail logs with
docker-compose logs -f app

# stop the app with
docker-compose down
```
By default the configurations in the `docker-compose.yml` file will start the app on port `80`.
That way you can access the app on `http://localhost/`, as per the spec.

**NOTE** Since docker compose has no health check, it might be the case that the DB is not yet accepting connections when you boot 
up the app. In that case, check if the pod is up and if not re-run:

```bash
docker ps 

docker-compose up -d app
```

Alternatively, you can run the app locally, as long a the instance of PostgresSQL is running:

```bash
# boot up db
docker-compose up -d postgres

# run the app on port 8080
make run

# OR

# run the app on port 80
make run-80
```

As most OSes don't allow you to run apps on the 80 port without admin privileges, so we kept both options.  
The provided makefile can start the app, removing any other apps running on the same port.

**Note:** The Postgres data base must be PostGIS ready with the extension installed. The provided docker-compose file
will take care of that.

### Integration Tests

Integration tests are available and can be run with the following command:

```bash
docker compose run e2e-test
```

These tests are located in the test folder and are a simple Python script that performs a sequence of requests to the API.
You don't need Python installed, as the script is run in a container, but they are quite useful to understand the API/see it working.

## API

- **Health Check**
    - **Method:** GET
    - **Endpoint:** `/health`
    - **Handler:** `handler.Health`
    - **Description:** Endpoint for health status monitoring.

---

- **Create User**
    - **Method:** POST
    - **Endpoint:** `/user/create`
    - **Handler:** `handler.CreateUser`
    - **Description:** Endpoint for creating a new user.
    - **Response:** `201 Created`
    - **Request Body:** None
    - **Response Body:**
    ```json
    {
        "result": {
            "id": 7,
            "email": "james@gmail.com",
            "password": "7q5Cs^1RF3#U",
            "name": "Elizabeth Davis",
            "gender": "F",
            "age": 58
        }
    }
    ```

---

- **Login**
    - **Method:** POST
    - **Endpoint:** `/login`
    - **Handler:** `handler.Login`
    - **Description:** Endpoint for user authentication and login.
    - **Response:** `201 Created`
    - **Request Body:** 
    ```json
    {
        "email": "email@gmail.com",
        "password": "7q5Cs^1RF3#U"
    }
    ```
    - **Response Body:**
    ```json
    {
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk3ODM4MzgsImlhdCI6MTcxOTY5NzQzOCwic3ViIjoiMSJ9.5l9piRKs_brYRYgxbADRFu4HrZmUU0gP2_Y08vG67Sc"
    }
    ```

---

- **Discover**
    - **Method:** GET
    - **Endpoint:** `/discover`
    - **EndpointParams:** `/discover?min_age=18&max_age=60&gender=F&sort_distance=true&sort_attractiveness=true`
    - **Authenticated:** Yes
    - **Handler:** `handler.Discover`
    - **Description:** Endpoint to discover potential matches. Query params are all optional.
    - **Response:** `200 OK`
    - **Response Body:**
    ```json
  {
    "results": [
      {
        "id": 3,
        "name": "Jessica Williams",
        "gender": "F",
        "age": 44,
        "distance_from_me": 10525.32
      },
      {
        "id": 4,
        "name": "Lauren Jones",
        "gender": "F",
        "age": 22,
        "distance_from_me": 11815.06
      }
    ]
  }
    
    ```

- **Swipe**
    - **Method:** POST
    - **Endpoint:** `/swipe`
    - **Authenticated:** Yes
    - **Handler:** `handler.Swipe`
    - **Description:** Endpoint for swiping left or right on matches.
    - **Response:** `201 Created`
    - **Error Response:** `400 Bad Request` | `409 Conflict`
    - **Request Body:**
      ```json
      {
      "user_id": 2,
      "preference": "YES"
      }
      ```
    - **Response Body:**
      ```json
      {
        "results": {
            "matched": false
        }
      }
      ```
---
<br/>

**Authentication**

Authenticated endpoints require a valid JWT token in the `Authorization` header. 
The token is generated by the `/login`.
Token must be a **Bearer** token.  
e.g. `Bearer <token>` (mind the space between `Bearer` and the token)

![auth](resources/token.png)

## Code Structure

```plaintext
.
├── api
│   └── router.go
├── config
│   └── config.go
├── db
│   └── db.go
└── pkg
    ├── cryptography
    ├── dao
    ├── handler
    ├── middleware
    ├── types
    └── workers
main.go
```

The best **entry point** would be the `router.go` file, which is responsible for setting up the mapping
of the routes and the handlers.

The `config.go` file is responsible for setting up the configuration of the app, 
while the `db.go` sets up the connection pool to the database.

## Data Model

![tables](resources/tables.png)

Foreign keys were not used due to increase overhead and complexity. Also, the app is simple enough to not require the 
extra check in terms of data consistency.

Location data is stored with the PostGIS extension, which allows for easy and efficient querying of the data.
This location data is indexed by default, and allows us to use distance operators and functions.

## Architecture and Design Choices

The app was designed to be **simple and easy to understand**. 
I considered going framework-less but the Go's HTTP STD lib makes things such as routing, query params, middlewares and request contexts
a bit more involved than necessary, specially compared with Gin.

For DB access, I decided not to go with an ORM, and instead to raw SQL via PGX. It helped to keep things simple and
it works for such a small application. Plus, PostGIS functionality may not be fully supported by most ORMs.

I chose **Postgres** as the database solution given that it is the system with which I am most familiar.
I also knew I could use PostGIS extension for location data.

The app is containerized, which makes it easy to run and deploy. The provided Dockerfile is multi-stage.

### Attractiveness Calculation

The attractiveness calculation is done as a counter the database, and it is just the number of right swipes a given user has.
Ideally, this would be a more complex algorithm, based on dislikes as well, and we could build a data pipeline of swipe (and other) events that calculate
a given person's attractiveness in relation to the user. This could also involve more data like preferences, location, hobbies, etc...

When we boot up the application we also **create a worker goroutine** that will re-calculate the attractiveness of each user.
It receives the swipe events via a channel, and increments the `likes` value in the DB.

### Assumptions

- Swipes are unique.
- Swipes are irreversible.
- Matches are irreversible.
- Swipes can only be made on valid users.
- Cannot delete users.
- User emails are unique.
- Users can be everywhere on Earth.

- We don't need paginated results
- We don't want bi-directional sorts i.e. since we don't have pagination we can treat a sort by distance result set with a reversed iterator as a sort by DESC
- We don't want to shard DB and/or have users distributed by DBs by geography
- We don't need to log out users -- the JWT expiration date suffices.
- For a small volume of data we don't need indexes on `age`. (`gender` has low cardinality so an index may be unadvised even for high data volumes).
- Cache would be over-engineering ATM
- We don't need localized error messages.

### Things I would have done with more time

- Much more unit tests. I decided to focus on e2e/integration tests as they are more valuable for this exercise.
- Leverage a caching solution using Redis. No needed for such a small app, but it would be a good show.
- Swagger API documentation. It would be nice to have a Swagger UI to test the API. This readme does a good job, but Swagger is more interactive.
- A more feature rich authentication system. We could fetch the user from DB at each request (a Redis cache would help) with roles and permission levels.
- More logs and log context. Generating a requestID and enriching logs with it + request context like user id would be ideal.
- Metrics: http request times, status codes, etc... 
