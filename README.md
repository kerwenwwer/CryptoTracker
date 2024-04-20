# CryptoTracker

CryptoTracker is a cryptocurrency tracking application that provides real-time price information. The project consists of a Go backend and a Vue.js frontend.

## Project Structure

- `cmd/proxy/`: Contains the Go server entry point.
- `pkg/`: Contains Go packages for handling business logic.
  - `api/`: Handles API requests to external cryptocurrency price services.
  - `trader/`: Contains logic for automated trading algorithms.
- `.crypto-tracker-frontend/`: Contains the Vue.js frontend project.
- `Makefile`: Defines commands for building and running the project.
- `README.md`: This file.

## Prerequisites

Before you can run this project, you need to install the following:

- Go (version 1.16 or later)
- Node.js and npm (Node 14.x or later recommended)
- Vue CLI (for managing the Vue.js frontend)

## Setup

To set up the project for development:

1. Clone the repository:
   ```bash
   git clone https://yourrepositoryurl.com/CryptoTracker.git
   cd CryptoTracker
   ```
2. Install dependencies for the Go backend:
    ```bash
    go mod tidy
    ```
3. Install dependencies for the Vue.js frontend:
    ```bash
    cd frontend
    yarn install
    cd ..
    ```

## Running the Application
To run the entire application (both frontend and backend):
```bash
make all
make go-run
```



## Contributing
Contributions are welcome! Please read the CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.