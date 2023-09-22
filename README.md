# gamer-stats-api
The (wannabe) game stats api. It allows you to get easily the stats for different games.

## Current support games:
- ??

If you want a new game to be supported, please open an issue. 

## API documentation:
TODO

## Contributor guide:
Do you want to contribute to this project? Great! This is how you can contribute:

- Contribute to the codebase
- Report bugs
- Write new documentation
- Suggest new games

If you want to contribute to the codebase, follow those steps to install the project locally:
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

### Open a new Pull Request:
To open a new pull request and contribute to the codebase, you'll need first to fork the repository. Follow the instructions to fork and clone the forked repository locally:

https://docs.github.com/en/get-started/quickstart/fork-a-repo?tool=cli

After you have cloned the forked repository locally, you can start to work on the codebase. When you are done, you can open a new pull request.

### I'm super noob but I want to contribute to the codebase:
If you are a beginner and want to learn how to develop, you can still help. Look for issues [beginner friendly](https://github.com/dct-tournaments/gamer-stats-api/labels/beginner%20friendly) and try to solve them. If you have any questions, you can ask question in the issues and someone will try to help  you. 

This codebase is written in Golang, so if you never developed in the language, try the [tour of go](https://tour.golang.org/welcome/1) to learn the basics.

Also, if you want to learn how to use git, you can follow this [tutorial](https://learngitbranching.js.org/).
