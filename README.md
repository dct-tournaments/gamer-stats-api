# gamer-stats-api

## Contributor guide:

### Local instalation:

#### Requirements:

- Docker


#### Step 1:

Build the docker image with the command:

`docker build -f deployment/local/Dockerfile --tag gamer-stats-api .`

#### Step 2:

Run the image with the command:
`docker run -p 8080:8080 gamer-stats-api`

#### Step 3:

Check the health endpoint at `localhost:8080/health` 
