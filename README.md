# Dealbreaker RESTFUL API  
Dealbreaker API is the API for a my implementation of the famous Red Flag games. The API is created with [Go](https://go.dev/), [Go Echo](https://echo.labstack.com/) and [PostgreSQL](https://www.postgresql.org/).

## Dealbreaker Game 
To play the game, each player will be dealt a 3 card hand; two *perk* cards and one *dealbreaker*. Each set of cards represent the player's date. After each player has been dealt the cards, they will discuss to see which of the set of cards that were dealt would be their preferred date. 

### Usage 
1. Clone this repo git clone `https://github.com/EzlosSWM/dealbreaker.git`

2. Navigate to the directory `cd dealbreaker`

3. Download dependancies go mod download && go mod verify

4. Copy the example.env to .env
```bash 
$ cp example.env .env
```

5. Run 
    - `make all`

### Endpoints 
**Get**

*/healthcheck*

*"/api/v1/jokes"*
- Returns all jokes

*"/api/v1/jokes/type/:joke_type"* 
- Filters the jokes returned by the type of joke 

**Delete**

*"api/v1/card/:id"*
- Deleted the specified card by ID 

**Post**

*"api/v1/card"*
- Adds a new card to the list of cards 
```JSON
{
    "joke_type": "dealbreaker",
    "joke": "King/Queen of halitosis."
}
```

*"api/v1/batch"*
- Used to batch upload data into the database 

#### Joke Types 
- perk 
- dealbreaker

### Todo 
- [ ] Dockerize
- [ ] Deploy