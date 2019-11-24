# News scraper Backend rest

This repository contains the code of the service that stores the news sent from the scrapers. This code only stores the news, for the scrapers you will also need the [scrapers repo](https://github.com/news-scrapers/news-scraper-workers-go)

## Installation
You will need an instance of a **mongodb database** running to store the news.

To easily run a mongo instance you can use docker.

    docker pull mongo
    sudo docker run -p 27017:27017 -it -d  --restart always mongo

To install docker go  [here](https://runnable.com/docker/install-docker-on-windows-10)


You will also need **golang** (version 1.12 at least). Follow [this](https://golang.org/doc/install) to install it.

## Configuration
* Clone this repository to a directory:

     git clone https://github.com/news-scrapers/backend-rest.git

* Move to the cloned directory and create a file named **.env** . Inside this file you will need to add two variables, one with the mongo url and another with an invented password for encripting tokens. If you are running mongo locally (as explained in the installation), you can use the following .env file:
  
        token_password = thisIsTheJwtPassword
        database_url =mongodb://localhost:27017

* Run the golang code that starts the rest service on port 8000. Inside the project folder run:
  
        go run main.go