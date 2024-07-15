function up {
   docker compose up
}

function down {
    docker compose down
}

function stop {
    docker compose stop
}

function restart {
    docker compose restart
}

function create_network() {
    docker network create point
}

function run() {
    create_network
    docker compose up -d
}
