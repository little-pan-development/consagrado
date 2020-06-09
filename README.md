# Consagrado Bot

***Problem:***

Working at a startup, almost every day, we asked for food delivery. But the team grew too fast, and it became a **HUGE** mess to take all orders manually.

***Solution:***

Our main communication channel between the company's areas was Discord. So, why not create a bot to organize orders? And here we are!

***But why Consagrado?***

This project was born in *Brazil*, and there Consagrado is a "*loving*" way to call the waiter.

**The project is being organized to accept new changes. Follow!**

## How to install?

*Requirements:*
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

Clone the project: `git clone git@github.com:little-pan-development/consagrado.git`

Open the folder: `cd consagrado`

Create a configuration file: `cp docker/development/env.example docker/development/env`

Edit the newly created file with your information.

Build an application with command: `make dev-up`

Create the database structure by running: `make dev-migration`

To see other commands, type: `make help`

## Contribute
Very soon!

