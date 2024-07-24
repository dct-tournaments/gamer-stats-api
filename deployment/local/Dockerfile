
FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/gamer-stats-api ./cmd/

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

RUN cd bin

ENV LEAGUE_OF_LEGENDS_API_KEY=""

# Run
CMD ["./bin/gamer-stats-api"]
