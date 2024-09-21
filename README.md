# MicroKey

<img src="/api/placeholder/800/200" alt="MicroKey Banner" />

MicroKey is an open-source, microservices-based application that integrates a Django backend with a Golang Gin service, providing a seamless and scalable architecture for authentication and API management.

## Project Structure

```
MicroKey/
├── Dockerfile.django           # Dockerfile for Django service
├── Dockerfile.gin              # Dockerfile for Golang service
├── docker-compose.yml          # Docker Compose configuration
├── auth-service-django/        # Django project directory
│   ├── apps/                   # Django apps directory
│   │   └── accounts/           # Accounts app
│   │       ├── migrations/     # Database migrations
│   │       ├── __init__.py
│   │       ├── admin.py        # Admin configuration
│   │       ├── apps.py         # App configuration
│   │       ├── authentication.py
│   │       ├── models.py       # Database models
│   │       ├── permissions.py  # Custom permissions
│   │       ├── serializers.py  # API serializers
│   │       ├── tests.py        # Unit tests
│   │       ├── urls.py         # URL routing
│   │       └── views.py        # API views
│   ├── config/                 # Django project configuration
│   │   ├── __init__.py
│   │   ├── asgi.py             # ASGI configuration
│   │   ├── settings.py         # Django settings
│   │   ├── urls.py             # Main URL routing
│   │   └── wsgi.py             # WSGI configuration
│   ├── manage.py               # Django management script
│   ├── db.sqlite3              # SQLite database
│   ├── requirements.txt        # Python dependencies
│   └── myenv/                  # Virtual environment (not tracked in git)
└── gin-auth-api/               # Golang Gin project directory
    ├── go.mod                  # Go module file
    ├── go.sum                  # Go module checksum
    ├── main.go                 # Main application entry
    └── middleware.go           # Gin middleware

```

## Features

<img src="/api/placeholder/600/300" alt="MicroKey Features" />

- **Django Backend**: Handles user authentication and serves REST APIs.
- **Golang Gin Service**: Lightweight middleware for authentication.
- **SQLite Database**: Lightweight, serverless database for easy setup and testing.
- **Docker Support**: Easy to set up and run using Docker Compose.
- **Open Source**: Contributions are welcome and encouraged!

## Technology Stack

<img src="/api/placeholder/700/150" alt="Technology Stack" />

- **Backend**: 
  - Django (Python)
  - Golang (Gin framework)
- **Database**: SQLite
- **Containerization**: Docker

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop) installed on your machine.
- [Docker Compose](https://docs.docker.com/compose/) for orchestration.

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
   - Django API: [http://localhost:8000](http://localhost:8000)
   - Golang Gin API: [http://localhost:9000](http://localhost:9000)

## Usage

- To authenticate a user via the Django API, use the following CURL command:
  ```bash
  curl --location 'http://127.0.0.1:8000/auth/service/third_party' \
  --header 'Authorization: Basic YOUR_BASE64_ENCODED_CREDENTIALS'
  ```

## Development

### Django Development

1. Navigate to the Django project directory:
   ```bash
   cd auth-service-django
   ```

2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

3. Run migrations:
   ```bash
   python manage.py migrate
   ```

4. Start the development server:
   ```bash
   python manage.py runserver
   ```

### Golang Gin Development

1. Navigate to the Gin project directory:
   ```bash
   cd gin-auth-api
   ```

2. Run the Gin server:
   ```bash
   go run main.go
   ```

## Contributing

We welcome and encourage contributions from the open-source community! Here's how you can contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all the contributors who have helped shape MicroKey!
- Gratitude to the open-source community for their continuous support and inspiration.
- Special thanks to the Django, Golang Gin, SQLite, and Docker teams for their excellent tools and frameworks.

## Get Involved

<img src="/api/placeholder/600/150" alt="Get Involved" />

We're always looking for new contributors and ideas! If you're interested in getting involved:

- Check out our [issue tracker](https://github.com/vivekjha1213/MicroKey/issues) for open issues and feature requests.
- Join our [community chat](https://discord.gg/microkey) to discuss development and get help.
- Follow us on [Twitter](https://twitter.com/MicroKeyProject) for updates and announcements.

Let's build something amazing together!
