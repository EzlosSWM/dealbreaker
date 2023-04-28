# Dealbreaker RESTFUL API  
Dealbreaker API is the API for a my implementation of the famous Red Flag games. The API is created with [Go](https://go.dev/), [Go Echo](https://echo.labstack.com/) [PostgreSQL](https://www.postgresql.org/).

## Dealbreaker Game 
To play the game, each player will be dealt a 3 card hand; two *perk* cards and one *dealbreaker*. Each set of cards represent the player's date. After each player has been dealt the cards, they will discuss to see which of the set of cards that were dealt would be their preferred date. 


### Routes 

**Post**

`"/api/v1/perk"`

`"/api/v1/dealbreaker"`

**Get**

`/healthcheck`

`"/api/v1/jokes"` 

<!-- filters the jokes returned by the topic  -->
`"/api/v1/jokes/topic/:topic"`

<!-- filters the jokes returned by the type of joke  -->
`"/api/v1/jokes/type/:joke_type"` 

**Delete**

`"/jokes/:id"`

#### Joke Topic
- general 
- sexual 
- finance 
- domestic

#### Joke Types 
- perk 
- dealbreaker

### Todo 
- [] pagination
- [] custom error handling 
- [] delete all exisiting items from DB
