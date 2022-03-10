# Google Search Bot 
This repo defines a Discord Bot that can query Google, and track a user's search history

## How to build
In order to use this repository, you need to create a discord bot. This can be done by following the steps [here](https://www.writebots.com/discord-bot-token/). Make sure you also follow the instructions to add the bot to a server.

Once you have your token, clone this repo and rename the `.env.example` file to `.env`. Then, replace the value under `BOT_TOKEN` with your token. Once you have done this, you can build and execute the code like any other go project.

`go build -o <OUTPUT_FILE_NAME>`
`go run <OUTPUT_FILE_NAME>`

## How to use
Once you have done the above, you should see 'Bot is running!' in your console. Then you can type !help in any text channel and the bot will explain its own usage.

![Example Usage](https://imgur.com/a/okrGoEg)

![Example of History](https://imgur.com/a/YNCQYhL)