# nano-migrate

An env-aware, forward-only, postgres migrator for solo devs and tiny teams.

- :computer: A simple CLI tool
- :zap: recognizes `.env` files and environment variables for easy substitution into sql migration files
- :arrow_up: no messing around with down migrations: roll with the punches and roll forward.

## Installation

(todo)

## Quickstart

Supply these environment variables

```
# database connection string
PG_URL

# path to the folder that holds your migration files
MIGPATH
```

Then run the following command to initialise the migrator

```sh
nnmg init
```

To create a brand new migration, simply use `nnmg create` followed by a description of the schema change. e.g.:

```txt
nnmg create add foo column to things table
```

:point_up: This will create a new file in your migrations folder, named `nnnn-add-foo-column-to-things.sql`, or similar. It will be empty, because nano-migrate doesn't opine on how you write your migrations and it doesn't support down migrations anyway, so really... enter whatever sql you'd like to run! :-)

:warning: **Do not use transactions**, as your migration will be wrapped in a postgres transaction.

To run all pending migrations that have not yet been applied to your database:

```sh
nnmg migrate
```

That's it!

---

## Contributing

Here's our tiny [contributing](./CONTRIBUTING.md) guide. We especially welcome contributors who are learning go as well as people who want the experience of contributing to a project! :-)
