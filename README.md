# Help Desk Bot

### Setup

#### Prerequisites

1.  Go

    - Installation
      - MacOS Homebrew installation: `brew install go`
      - Windows and Debian-based Linux distros: [Official installation](https://go.dev/doc/install)
      - Arch-based Linux distros: `sudo pacman -S go`
    - `GOPATH` and `GOBIN` environment variable set up
      - Guide: [Go Wiki: Setting GOPATH](https://go.dev/wiki/SettingGOPATH)

2.  Docker installed

- Windows and MacOS: install [Docker Desktop](https://docs.docker.com/desktop/)
- Debian-based Linux distros installation guide [here](https://docs.docker.com/engine/install/debian/)
- Arch-based Linux distros installation:
  ```bash
  sudo pacman -S docker docker-compose
  ```

#### Running the server locally

1. Fork and clone the repository

   ```bash
   git clone https://github.com/<your_username>/chatbot-backend.git
   cd chatbot-backend
   ```

2. Create your own telegram bot for development: https://core.telegram.org/bots/tutorial#getting-ready

3. Create a `.env` file in the root directory of the project and add the following environment variables

   ```bash
   cat .env.example >> .env
   ```

   Set `TELEGRAM_TOKEN` in `.env` to the token you got from the previous step

4. Build the Docker image

   NOTE: If you have a local postgres instance running, you need to set `PGPORT` variable in the .env file to the port number **other than** the default 5432. (e.g. `PGPORT=5433`)

   ```bash
   docker compose --env-file .env build
   # or run `make dockerbuild` if you are using MacOS, Linux or WSL
   ```

5. Run the Docker container

   ```bash
    docker compose --env-file .env up
    # or run `make dockerup` if you are using MacOS, Linux or WSL
   ```

6. Run the migrations
   Keep the docker container from step 6 running and open a new terminal window

   ```bash
   docker exec -it chatbot-backend-server-1 sh
   ```

   You are now running the container's interactive shell. Run the following commands

   ```bash
   go run cmd/migratedb/main.go
   ```

#### Development Environment

1. Install development tools

   - golangci-Lint: https://golangci-lint.run/welcome/install/

2. Set up git hooks
   Assuming you are inside the project directory

   ```bash
   golangci-lint init
   ```
