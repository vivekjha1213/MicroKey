# MicroKey

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Python](https://img.shields.io/badge/python-3.8%2B-blue.svg)
![Go](https://img.shields.io/badge/go-1.16%2B-blue.svg)

MicroKey is an open-source, microservices-based application that integrates a Django backend with a Golang Gin service, providing a seamless and scalable architecture for authentication and API management.

```
  _____________________
 /  Microservices       \
|   ____     ____        |
|  |    |   |    |       |
|  |Django  | Gin|       |
|  |____|   |____|       |
|                        |
 \______________________/
        Spinning...
```

## Microservices Architecture

MicroKey leverages a microservices architecture to ensure scalability, maintainability, and flexibility:

1. **Django Auth Service**: Handles user authentication and management.
2. **Golang Gin API Service**: Provides lightweight, high-performance API endpoints.

## Features

![Django](https://img.shields.io/badge/django-backend-green.svg)
![Golang](https://img.shields.io/badge/golang-gin-blue.svg)
![SQLite](https://img.shields.io/badge/database-sqlite-orange.svg)
![Docker](https://img.shields.io/badge/container-docker-blue.svg)
![GitHub](https://img.shields.io/badge/version%20control-github-black.svg)

- **Django Backend**: Handles user authentication and serves REST APIs.
- **Golang Gin Service**: Lightweight middleware for authentication.
- **SQLite Database**: Lightweight, serverless database for easy setup and testing.
- **Docker Support**: Easy to set up and run using Docker Compose.
- **GitHub Integration**: Version control and collaboration.
- **Open Source**: Contributions are welcome and encouraged!

## Technology Stack

![Backend](https://img.shields.io/badge/backend-django%20%7C%20golang-blue.svg)
![Database](https://img.shields.io/badge/database-sqlite-orange.svg)
![Containerization](https://img.shields.io/badge/containerization-docker-blue.svg)
![Version Control](https://img.shields.io/badge/version%20control-git-black.svg)

## Project Structure

```
MicroKey/
├── Dockerfile.django
├── Dockerfile.gin
├── docker-compose.yml
├── auth-service-django/
│   ├── apps/
│   │   └── accounts/
│   ├── config/
│   ├── manage.py
│   ├── db.sqlite3
│   ├── requirements.txt
│   └── myenv/
├── gin-auth-api/
│   ├── main.go
│   ├── middleware.go
│   ├── go.mod
│   └── go.sum
└── README.md
```

## GitHub Integration

[![GitHub repo](https://img.shields.io/badge/github-MicroKey-brightgreen.svg)](https://github.com/vivekjha1213/MicroKey)
[![GitHub issues](https://img.shields.io/github/issues/vivekjha1213/MicroKey.svg)](https://github.com/vivekjha1213/MicroKey/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/vivekjha1213/MicroKey.svg)](https://github.com/vivekjha1213/MicroKey/pulls)

## Getting Started

### Prerequisites

![Docker](https://img.shields.io/badge/required-docker-blue.svg)
![Docker Compose](https://img.shields.io/badge/required-docker--compose-blue.svg)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/vivekjha1213/MicroKey.git
   cd MicroKey
   ```

2. Build and run the services:
   ```bash
   docker-compose up --build
   ```

3. Access the services:
   - Django API: http://localhost:8000
   - Golang Gin API: http://localhost:9000

## Development

### Django Development

```bash
cd auth-service-django
pip install -r requirements.txt
python manage.py migrate
python manage.py runserver
```

### Golang Gin Development

```bash
cd gin-auth-api
go run main.go
```

## Contributing

We welcome contributions! Here's how you can contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

[![MIT License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## Get Involved

[![GitHub Stars](https://img.shields.io/github/stars/vivekjha1213/MicroKey.svg?style=social)](https://github.com/vivekjha1213/MicroKey)
[![Discord](https://img.shields.io/discord/XXXXXX.svg?label=&logo=discord&logoColor=ffffff&color=7389D8&labelColor=6A7EC2)](https://discord.gg/vivekjha1213)
[![Twitter Follow](https://img.shields.io/twitter/follow/MicroKeyProject.svg?style=social)](https://twitter.com/vivekjha1213)

Let's build something amazing together!
