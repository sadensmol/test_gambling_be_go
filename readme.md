# Test task description

Your task is to create a backend system for a gambling services platform that utilizes multiple communication protocols for various features. The system should manage user wallets and implement real-time updates using WebSockets. Additionally, you should add a gRPC endpoint to demonstrate the use of multiple communication transports.

# Requirements

User Wallet API:
• Implement the following API endpoints using HTTP:
• POST /api/wallet/deposit: Accepts a JSON payload with the user ID and deposit amount. Updates the user's wallet balance accordingly.
• POST /api/wallet/withdraw: Accepts a JSON payload with the user ID and withdrawal amount. Updates the user's wallet balance if sufficient funds are available.
• GET /api/wallet/balance/:user_id: Retrieves the current balance of the user's wallet.
• Use an in-memory data store to manage user wallets.
Real-time Updates with WebSockets:
• Implement a WebSocket server to notify users of game outcomes and leaderboard changes in real-time.
• Clients should be able to subscribe to specific events (e.g., game outcomes, leaderboard changes) using WebSocket connections.
gRPC Endpoint:
• Add a gRPC endpoint to the backend that allows users to retrieve their wallet balances.
• Define a gRPC service with a single RPC method to retrieve the balance.
• Implement the service using Go and demonstrate communication via gRPC.
Bonus (Optional):
• Implement input validation for API endpoints and gRPC calls.
• Add logging to record API and gRPC interactions.
• Dockerize the application for easy setup and deployment.

# Implementation details

All amounts are in currency cents to avoid floating point precision issues, thus Wallet is a single currency.
There is no proper error management - server returns only 200 and 500 in case of any error.

Directory /tests contains e2d tests with http and grpc clients (grpc just for example a single test). To start e2e tests you need to start application  with `make run` command.

Project structure is based on [clean architecture article](https://medium.com/@sadensmol/my-clean-architecture-go-application-e4611b1754cb)
