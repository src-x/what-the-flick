```
 __     __  __  __  ______  ______     ______ __  __  ______       ______ __      __  ______  __  __    
/\ \  _ \ \/\ \_\ \/\  __ \/\__  _\   /\__  _/\ \_\ \/\  ___\     /\  ___/\ \    /\ \/\  ___\/\ \/ /    
\ \ \/ ".\ \ \  __ \ \  __ \/_/\ \/   \/_/\ \\ \  __ \ \  __\     \ \  __\ \ \___\ \ \ \ \___\ \  _"-.  
 \ \__/".~\_\ \_\ \_\ \_\ \_\ \ \_\      \ \_\\ \_\ \_\ \_____\    \ \_\  \ \_____\ \_\ \_____\ \_\ \_\ 
  \/_/   \/_/\/_/\/_/\/_/\/_/  \/_/       \/_/ \/_/\/_/\/_____/     \/_/   \/_____/\/_/\/_____/\/_/\/_/ 
```

## Demo
![Demo of what-the-flick](./demo.gif)

## Technologies
RediSearch, ClickHouse, gRPC, Gin, React, and Docker

## File Structure
```
what-the-flick
|-- clickhouse
|   |-- Dockerfile
|   |-- insert_movie_data.py
|   |-- insert_movie_data.sh
|   `-- test.new
|-- grpc-clickhouse
|   |-- recommend_client
|   |   `-- client.go
|   |-- recommend_server
|   |   `-- server.go
|   |-- recommendpb
|   |   |-- recommend.pb.go
|   |   `-- recommend.proto
|   |-- Dockerfile
|   `-- generate.sh
|-- grpc-redisearch
|   |-- redisearchpb
|   |   |-- redisearch.pb.go
|   |   `-- redisearch.proto
|   |-- search_client
|   |   `-- client.go
|   |-- search_server
|   |   `-- server.go
|   |-- Dockerfile
|   `-- generate.sh
|-- movie-website
|   |-- public
|   |   |-- favicon.ico
|   |   |-- index.html
|   |   `-- manifest.json
|   |-- src
|   |   |-- DownshiftAutoComplete.js
|   |   |-- DownshiftCheckbox.js
|   |   |-- index.js
|   |   `-- styles.css
|   |-- Dockerfile
|   `-- package.json
|-- redisearch
|   |-- Dockerfile
|   `-- suggestion_completion_data.py
|-- docker-compose.yml
`-- start.sh
```

## Requirements
* Docker
* TMDB API key

## Getting Started / Installation
```
1. git clone https://github.com/src-x/what-the-flick.git
```
```
2. sh start.sh
```
```
3. input your tmdb api key
```

## Pointers/References

Grpc and Gin tutorial by Tensor Programming:
* https://www.youtube.com/watch?v=Y92WWaZJl24
* https://github.com/tensor-programming/grpc_tutorial

Downshift Sandboxes from Kent C. Dodds:
* https://codesandbox.io/s/w6gyj30kn
* https://codesandbox.io/s/v6yn8j0l03

Dataset and API

* Dataset: https://grouplens.org/datasets/movielens/
* API: https://www.themoviedb.org/?language=en-US

## Future Improvements and Additions

* replace existing API server with gRPC-Web
* add python notebooks to demonstrate how to use redisearch and clickhouse
* organize and improve frontend react functionality 
* use typescript...
* insert larger movielens datasets for databases
* add more user-based similarities and improve buggy sql queries
* and much more