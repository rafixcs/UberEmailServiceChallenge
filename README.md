# Email Microservice
## Uber Backend Challenge

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazonaws&logoColor=white)
[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)

This project is an API built using **Golang and AWS Simple Email Service**


## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)

## Installation

1. Clone the repository:
```bash
git clone git@github.com:rafixcs/UberEmailServiceChallenge.git
```

2. Install dependencies:
```bash
go mod download && go mod verify
```

## Configuration
Set the enviroment variables

```bash
export AWS_ACCESS_KEY_ID=<access-key> \
export AWS_SECRET_ACCESS_KEY=<secret-access-key>\
export AWS_REGION=<aws-region>
```
## Usage
1. Start the application with:
```bash
make all
```
2. The API will be available at http://localhost:3000

## API Endpoints

The API provides the following endpoints
**Send Email**
```markdown
POST /api/email - Send a e-mail from your source to the destination
```

**BODY**
```json
{
    "destination": "test.dest@gmail.com",
    "source": "test.source@gmail.com",
    "content": "hello!",
    "subject": "request throught golang and aws server"
}
```


## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request to the repository.

When contributing to this project, please follow the existing code style, [commit conventions](https://www.conventionalcommits.org/en/v1.0.0/), and submit your changes in a separate branch.
