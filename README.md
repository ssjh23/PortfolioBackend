# PortfolioBackend

## Overview
This is a REST API built in Golang Gin. It is for submission for 50.012 Lab 2, but its original intent was for me to build a backend for my portfolio website [here](https://seansoojunhao.me/). I wanted to build an editor for it so I don't have to keep going into the Frontend code and change it manually, while also having another project for Backend. It is built using the `Controller-Service-Repo` architecture, and the entry point is on `cmd/main.go` 

## Routes
1. `GET /about/get_profile_pic` <br><br>
This route gets a picture (`reduce.png`) and returns it as an `application/octet-stream`. This is to fulfil challenge `Have a route in your application that returns a content type that is not plaintext`

2. `GET /project/get_projects` <br><br>
This route returns projects from the hardcoded array in the repository folder. It takes in 3 Query params: <br>
    - `sortBy: Title or ""`. If `sortBy` uses `Title`, the projects returned will be in alphabetical order by title. Otherwise if empty, will remain unordered<br>
    - `count: 0 - Number of projects`. if specified, count will return the exact number of projects, starting from the beginning of the array list. If the count exceeds the number of projects, an error with status `400` will be thrown. If `count = 0`, all the projects that can be returned will be returned <br>
    - `offset: 0 - (Number of projects-1)`. if specified, the return value for the projects will start from the index specified by `offset`. if `offset` equals or is more than the number of projects,an error with status `400` will be thrown. Returned is the requested `Projects`<br><br>
    The way the queries are considered are in the order: `sortBy -> offset -> count`<br>
3. `POST /project/add_project` <br><br>
This route adds a project using `form-data` with values `title, description, techstack`. Note that `techstack` is a csv. If any of the fields are blank on request, an error with status `400` will be thrown. Returned is the updated `Projects` This is to fulfil `File upload in a POST request, using multipart/form-data`

4. `DEL /project/delete_project`<br><br>
This route deletes a project with an `Authorization` header. Here the auth header is hardcoded to be my student id `1005263`. If the `Authorization` header is missing or the value is invalid, an error status `401` with an `Unauthorized` message is thrown`. It uses one query param, `title` which corresponds to the title of the project. If the `title` is not valid, an error status `400` is thrown. The `title` has to match an existing project. Returned is the updated `Projects`. This is to fulfil `Some form of authorization through inspecting the request headers`



  
 

## Build using docker
1. Be in the Root directory where `docker-compose.yml` and `Dockerfile` is
2. Run `docker build -t rest_api:1.0 . ` to build the container
3. Run `docker run -p 8000:3001 rest_api:1.0  ` This exposes the docker container to port `8000` but the input is connected to port `3001` where the backend is running
