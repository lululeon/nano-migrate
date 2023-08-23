#! /bin/bash

echo "🗲  making sure you have go installed..."
if ! command -v go >/dev/null 2>&1
then
  echo "  ❌ Uh-oh... You don't seem to have go installed."
  echo "Please see here: https://go.dev/doc/install for instructions. After you've installed, you can run this setup again! :-)"
  exit
else
  echo "  ✔️  Go is available!"
fi

echo "🗲  making sure you have docker installed..."
if ! command -v docker >/dev/null 2>&1
then
    echo "  ⚠️  Warning: Docker isn't running (or isn't installed?)."
    echo "  You don't need docker if you've already got a postgres instance that you want to use for development and testing."
    echo "  Otherwise, you can go here: https://docs.docker.com/get-docker/ for install instructions. If you're sure it's installed, try the command 'sudo service docker start'."
else
  echo "  ✔️  docker is available!"
fi

echo "🗲  creating 'nn_datavol' folder in your home directory..."
mkdir -p $HOME/nn_datavol/migrations

echo "🗲  creating standard '.env' file for a quick start..."
cat << EOF > ./.env
PG_PORT=5432
DB_DATAPATH=$HOME/nn_datavol/nnmg
PG_URL=postgres://postgres:localpass@localhost:5432/localdb
MIGPATH=$HOME/nn_datavol/migrations
EOF

echo "🗲  installing gotestsum for readable test reports..."
go install gotest.tools/gotestsum@latest

echo "✨  Done!"
