# Honest Backend Engineer Technical Assessment

## Objective

This project is a Decision Engine to approve/decline credit card applicants. It is implemented in Go and uses the `net/http` package to handle HTTP requests.

## High level

- local setup

![image](https://github.com/pcm708/credit-card-eligibilty-backend/assets/52307892/88a93500-fb21-4f14-ab1c-819e6f20a1ee)

- for cloud, redis and db running on windows laptop
- V2:
- using RDS mysql server to storage data on cloud rather than storing on personal windows laptop,
- Run the project locally on terminal, go run main.go
- Uncommented redis implementation for the /process api as im not using redis cloud aws service for now.

- 


## Project Structure

The project is structured as follows:

- `main.go`: This is the entry point of the application. It sets up the HTTP server and handles graceful shutdown on receiving a termination signal.
- `routes/`: This directory contains the setup for the HTTP routes.
- `controllers/`: This directory contains the application's controllers. They handle incoming HTTP requests, delegate business processing to the services, and return HTTP responses.
- `model/`: This directory contains the data models used in the application.
- `services/`: This directory contains the business logic of the application. Services are used by controllers to perform specific tasks related to the business requirements.

## API


0th: (for setting up preapproved numbers)
curl --location 'http://localhost:8080/add' \
--header 'Content-Type: application/json' \
--data '[
    "023-7548-8548",
    "223-7548-8549",
    "223-7548-8541",
    "523-7548-8542"
]'

1st: The application exposes a `POST` HTTP endpoint at `/process`:

curl --location 'http://localhost:8080/process' \
--header 'Content-Type: application/json' \
--data '{
    "income": 800000,
    "number_of_credit_cards": 1,
    "age": 26,
    "politically_exposed": true,
    "phone_number": "823-758-8559"
}'

### Request Body

| Fields                   | Type        |
| -----------              | ----------- |
| income                   | number      |
| number_of_credit_cards   | number      |
| age                      | number      |
| politically_exposed      | bool        |
| job_industry_code        | string      |
| phone_number             | string      |

##### Example

```json
{
  "income": 82428,
  "number_of_credit_cards": 3,
  "age": 9,
  "politically_exposed": true,
  "job_industry_code": "2-930 - Exterior Plants",
  "phone_number": "486-356-0375"
}
```

#### Response Body

| Fields                   | Type        |
| -----------              | ----------- |
| status                   | string      |

##### Example

###### Approved:

```json
{
  "status": "approved"
}
```

###### Declined:

```json
{
  "status": "declined"
}
```

## Running the Application

### Command-line
To run the application(make sure go in installed in your system and go Path is set by default), use the go run command:
```
go run main.go
```
The application will start and listen on the port specified in the .env file.

Simply make a POST request to the / endpoint on port mentioned in .env file with a valid request body. The structure of the request body is described in the API section.

### Docker

To run the application in a Docker container, use the following commands:
```
docker-compose up --build
```
This will start the application and listen on the port specified in the docker-compose.yml file. 

Simply make a POST request to the / endpoint on port mentioned in docker-compose.yml file with a valid request body. The structure of the request body is described in the API section.

## Logging

Every json entry in the request body is logged in the file `log.json` in the following example format:
```
{
    "phone_number": "086-326-0220",
    "status": "approved",
    "message": "number logged",
    "timestamp": "2024-02-21 03:16:44"
}
```
If the request is approved, the phone number is logged in `numbers.txt` file.

## Testing

```
go test -v ./tests
  ```
This command will run all the tests in the project present under tests directory.


### Decision Rules

The application is approved if it evaluates as `true` on the following rules:

1. The applicant must earn more than 100000.
1. The applicant must be at least 18 years old.
1. The applicant must not hold more than 3 credit cards and their `credit_risk_score` must be `LOW`.
1. The applicant must not be involved in any political activities (must not be a Politically Exposed Person or PEP).
1. The applicant's phone number must be in an area that is allowed to apply for this product. The area code is denoted by first digit of phone number. The allowed area codes are `0`, `2`, `5`, and `8`.
1. A pre-approved list of phone numbers should cause the application to be automatically approved without evaluation of the above rules. This list must be able to be updated at runtime without needing to restart the process.
