# Movie Management API

This is a simple API built with Go that allows you to manage a list of movies. You can create, read, update, and delete movies using HTTP requests. The movies are stored in memory, so they will be lost when the server restarts.

## Table of Contents

1. [API Endpoints](#api-endpoints)
   - [Get All Movies](#1-get-all-movies)
   - [Get a Specific Movie](#2-get-a-specific-movie)
   - [Create a New Movie](#3-create-a-new-movie)
   - [Update a Movie](#4-update-a-movie)
   - [Delete a Movie](#5-delete-a-movie)
2. [Running the Application](#running-the-application)
3. [License](#license)

---

## API Endpoints

### 1. **Get All Movies**

- **URL**: `/movies`
- **Method**: `GET`
- **Description**: Retrieves a list of all movies.
- **Response**:
  - Status Code: `200 OK`
  - Body: A JSON array of all movies.

**Example Request:**

```bash
GET http://localhost:4000/movies
[
  {
    "id": "1",
    "isbn": "1234567",
    "title": "The Rock",
    "director": {
      "firstName": "Joe",
      "lastName": "Jones"
    }
  },
  {
    "id": "2",
    "isbn": "098748765",
    "title": "The Musk",
    "director": {
      "firstName": "Eli",
      "lastName": "Job"
    }
  }
]

```

### 2. **Get a single movie**

- **URL**:`/movies/1`
- **Method**:`GET`
- **Description**: Retrieves a one movie.
- **Response**:

  - Status Code: `200 OK`
  - Body: A JSON object representing the movie.

**Example Request:**

```bash
GET http://localhost:4000/movies/1
  {
    "id": "1",
    "isbn": "1234567",
    "title": "The Rock",
    "director": {
      "firstName": "Joe",
      "lastName": "Jones"
  }
```

### 3. **Create a New Movie**

- **URL**: `/movies`
- **Method**: `POST`
- **Description**: Creates a new movie.
- **Response**:
  - Status Code: `201 OK`
  - Body: A JSON object representing the new movie.

**Example Request:**

```bash
POST http://localhost:4000/movies
Content-Type: application/json
{
  "isbn": "5678901",
  "title": "New Movie",
  "director": {
    "firstName": "Alice",
    "lastName": "Smith"
    }
}
```

### 4. **Update a Movie**

- **URL**:`/movies/1`
- **Method**:`UPDATE`
- **Description**: Updates an existing movie based on its ID.
- **Response**:

  - Status Code: `200 OK`
  - Body: A JSON object with updated movie details.

**Example Request:**

```bash
PUT http://localhost:4000/movies/1
Content-Type: application/json

{
  "isbn": "7654321",
  "title": "Updated Movie",
  "director": {
    "firstName": "Joe",
    "lastName": "Jones"
  }
}
```

### 5. **Delete a Movie**

- **URL**:`/movies/1`
- **Method**:`DELETE`
- **Description**: Deletes a movie based on its ID.
- **Response**:

  - Status Code: `200 OK`
  - Body: A JSON array of all movies.

**Example Request:**

```bash
DELETE http://localhost:4000/movies/1
[
  {
    "id": "1",
    "isbn": "1234567",
    "title": "The Rock",
    "director": {
      "firstName": "Joe",
      "lastName": "Jones"
    }
  },
  {
    "id": "2",
    "isbn": "098748765",
    "title": "The Musk",
    "director": {
      "firstName": "Eli",
      "lastName": "Job"
    }
  }
]
```
