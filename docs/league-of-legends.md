# League of Legends Stats API 

## Overview
Provide the stats for a league of legends player.

## API Endpoints

### GET /stats/v0/league-of-legends?
#### Params
- `username` - The summoner name of the player
- `region` - The region of the player. Valid values are br1,eun1,euw1,jp1,kr,la1,la2,na1,oc1,ph2,ru,sg2,th2,tr1,tw2,vn2
- `start_time` - Epoch timestamp in seconds. The matchlist started storing timestamps on June 16th, 2021. Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.

#### Request
```
curl -X GET "http://localhost:8080/stats-api/v0/league-of-legends?username=kahnoel&region=euw1"
```

#### Response
```json
{
    "stats": {
        "kills": 53,
        "deaths": 119,
        "assists": 78
    },
    "query_executed_at": 1703425687
}
```
