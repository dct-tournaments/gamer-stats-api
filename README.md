# gamer-stats-api
The (wannabe) game stats api. It allows you to get easily the stats for different games.

## API documentation:

### v0.0.1

#### Supported games and events
- League of Legends
  - kills
  - deaths
  - assists

ps: if you want a new game or event, open a new issue and we will try to add it.

#### Response protocol:
The response protocol is the same for all supported the games. The response is a JSON object with the following structure:

```json
{  
    "stats": {
        "in_game_event_name": total_count_of_events_for_the_player,
        "in_game_event_name": total_count_of_events_for_the_player,
        "in_game_event_name": total_count_of_events_for_the_player
    },
    "query_executed_at": "timestamp of the query"
}
```

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
`docker run -e LEAGUE_OF_LEGENDS_API_KEY="your_lol_api_key" -p 8080:8080 gamer-stats-api`

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
