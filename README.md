# nano-migrate

An env-aware, forward-only, postgres migrator for solo devs and tiny teams.

- :computer: A simple CLI tool
- :zap: recognizes `.env` files and environment variables for easy substitution into sql migration files
- :arrow_up: no messing around with down migrations: roll with the punches and roll forward.
- automatically searches for 

## Installation
(todo)


## Quickstart
1. Create a `.nanopl` file with the following entries:
```json
{
  "nn:mg": {
    "dbconn": "postgresql://user:pass:5432@host/db?schema=x",
    "migsdir": "./migrations",
    "idprefix": "nnnn",
    "substitutions": {
      "SCHEMA": "clients",
      "R1": "role1"
    }
  }
}
```

2. Run the following commands
```sh
nnmg init
nnmg create baseline schema
nnmg migrate
```
