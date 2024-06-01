# GO Discord file server

### Store the images of your application in discord, why not?


## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

If you never created a discord bot follow [this tutorial](https://medium.com/@mssandeepkamath/building-a-simple-discord-bot-using-go-12bfca31ad5d)

1. Clone the repo:

   ```bash
   git clone https://github.com/diyst/go_discord_file_server.git
    ```

2. Install the dependencies
    ```bash
    go mod  tidy
    ```

3. Configure enviroment variables
    ```
    PORT=3080
    AUTH_TOKEN= The auth token generated in discord
    CHANNEL_ID= The discord channel id where the images will be stored
    ```

4. Run the application
    ```bash
    go run main.go
    ```go run main.go