# Our school project

The goal of this project is to create a simple email manager such as Mailchimp.
We used Golang for our API, and Postgresql as a database.

## Features

- User authentication with JWT.
- Marketing Automation : Creation of marketing campaigns with sending emails to a specific mailing list

## Stating with Docker

After cloning the repo, `cd` into the project, create the .env in /app and /consumer according to .env.example, and run following commands

```bash
docker-compose up --build
```

## Starting Manualy

The project requires Golang `1.14.4` version

Install the dependencies and start the service locally with theses commands:

```sh
git clone https://github.com/HETIC-MT-P2021/Go_Mail_Groupe05/app.git
go mod tidy
go get -u github.com/cosmtrek/air
air
```

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.
Go version: `1.14.4`
PostgreSQL: `12.3`

### Documentation

You can find the api doc by clicking on the link below :
[DOC](https://app.swaggerhub.com/apis/JWT_Golang/Mailing_in_Go/1.0.0)

### Authors

- [Myouuu](https://github.com/myouuu)
- [Tsabot](https://github.com/Tsabot)
- [Jean-Jacques](https://github.com/gensjaak)

### License

The code is available under the MIT license.
