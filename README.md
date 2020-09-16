# Our school project

The goal of this project is to create a simple email manager such as Mailchimp.
We used Golang for our API, and Postgresql as a database.

## Features:

- User authentication with JWT.
- Marketing Automation : Creation of marketing campaigns with sending emails to a specific mailing list

## Installation

The project requires Golang v 1.14.4

Install the dependencies and start the server.

```sh
$ git clone https://github.com/myouuu/Authentication-in-Golang-with-JWT.git
$ cd src
$ go get
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/cosmtrek/air
$ go get gopkg.in/gomail.v2
$ go mod vendor
$ air
```

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.
Go version : 1.14.4
PostgreSQL : 12.3

# Licence

The code is available under the MIT license.
