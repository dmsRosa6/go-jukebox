# SubJuke

SubJuke is a Go-based mockup jukebox service. It provides a REST API for operations like enqueuing songs, retrieving the current song, and more. Additionally, it integrates with Subsonic-compatible music servers (via an existing Go library) to fetch music metadata and stream audio.

## Features

- **REST API:**
  - Endpoints for managing the song queue (enqueue, dequeue, get current song, etc.)
  - Built using Gorilla Mux for flexible routing

- **Subsonic Integration:**
  - Consumes a Subsonic API via an existing client library to retrieve metadata, stream music, and control playback on your Subsonic server.
  - Easily configurable authentication (using token or password)

- **Modular Design:**
  - Clean separation between API routing, service/business logic, and configuration.
  - Dependency injection for easy testing and future enhancements.

## Requirements

- Go 1.22.4 or later
- [Gorilla Mux](https://github.com/gorilla/mux)
- A Subsonic API client library (e.g. [delucks/go-subsonic](https://github.com/delucks/go-subsonic) or [mdlayher/gosubsonic](https://pkg.go.dev/github.com/mdlayher/gosubsonic))
- A Subsonic-compatible music server for streaming music

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dmsrosa/jukebox.git
   cd jukebox
   ```


*Working on it*