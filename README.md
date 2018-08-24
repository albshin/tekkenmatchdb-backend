# TekkenMatchDB (Backend)

## Why?
I am using this as a learning experience for constructing a production ready API. Unlike some other projects, I am invested in Tekken as a player so this should motivate me more to finish the project.

## Dependencies
Following the footsteps of many other Gophers, I have decided not to use a framework and to keep dependencies to a minimum. This also helps me learn about concepts at a lower level.
* [chi](https://github.com/go-chi/chi) (Routing)
* [sqlx](https://github.com/jmoiron/sqlx) (Database accessibility)
* [pq](https://github.com/lib/pq) (Postgres driver)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) (Input validation. Possiby replaced with stdlib in the future)
    * govalidator (imported by ozzo)
* [squirrel](https://github.com/Masterminds/squirrel) (Query builder)

## Some planned features
* Matchlist
    * Sorting (by date, event, etc.)
    * Filtering (by character, player, event, rank, etc.)
* Player Page
    * Matches by player
    * Matchup
        * Worst matchups
        * Stats vs another player

## Getting Started
Certain values are hardcoded for now and will be replaced with environment variables.

1. Create a table named "matches" into your PostgreSQL database.
2. Insert both the country.sql and schema.sql into the "matches" database. (country.sql first)
3. Run ```go run main.go```