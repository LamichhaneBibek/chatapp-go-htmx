# Chat Web App with Golang, Echo, htmx, and WebSockets

## Overview
This is a simple chat web application built using Golang, the Echo web framework, htmx for seamless updates, and WebSockets for real-time communication. The application allows users to join chat rooms, send messages, and experience a dynamic and responsive interface.

## Features
- **Real-Time Chat:** Utilizes WebSockets for instant message updates.
- **Multiple Chat Rooms:** Users can join different chat rooms for separate conversations.
- **Responsive Design:** Built with htmx for seamless and fast UI updates without full-page reloads.
- **Golang and Echo:** Backed by the Go programming language and the Echo web framework for a robust backend.

## Prerequisites
Make sure you have the following installed on your machine before running the application:
- Go (Golang): [Install Go](https://golang.org/doc/install)
- Echo: Install using `go get -u github.com/labstack/echo/v4`
- htmx: Include the htmx library in your project or use a CDN. [htmx Documentation](https://htmx.org/docs/)
- WebSocket library for Golang: Install using `go get github.com/gorilla/websocket`

## Getting Started
1. Clone the repository: `git clone [repository_url]`
2. Navigate to the project directory: `cd chat-web-app`
3. Run the Golang server: `go run main.go`
4. Open your web browser and go to `http://localhost:8080` to access the chat application.

## Configuration
- Update the WebSocket endpoint in the frontend JavaScript to match your server configuration.

## Structure
- `main.go`: Main Golang server file.
- `static/`: Directory for static assets (CSS, JavaScript).
- `templates/`: HTML templates for the web application.
- `handlers/`: Golang handlers for different routes.

## Contributing
Feel free to contribute to this project by opening issues or submitting pull requests. Your feedback and suggestions are highly appreciated.

## License
This project is licensed under the [MIT License](LICENSE).

## Acknowledgments
- Thanks to the Golang, Echo, htmx, and WebSocket communities for their excellent tools and libraries.
