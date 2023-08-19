#! /bin/bash

echo "🗲 creating 'nn_datavol' folder in your home directory..."
mkdir -p $HOME/nn_datavol/migrations

echo "🗲 creating standard '.env' file for a quick start..."
cat << EOF > ./.env
PG_PORT=5432
DB_DATAPATH=$HOME/nn_datavol/nnmg
PG_URL=postgres://postgres:localpass@localhost:5432/localdb
MIGPATH=$HOME/nn_datavol/migrations
EOF

echo "✨ Done!"
