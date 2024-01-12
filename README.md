
---

## Project Description: Microservice for Retrieving Latest Trades

### Introduction:
This project is designed as a microservice to efficiently fetch the latest trades for each instrument from a PostgreSQL database and deliver them in JSON format to clients. Emphasizing a microservices architecture, the system integrates Protobuf for efficient serialization, Swagger for API documentation, and Docker Compose for containerization. The primary goal is to optimize latency and enhance throughput, utilizing appropriate technologies and packages for a streamlined and efficient service.

### Features:
1. **Technology Stack**:
    - **Programming Language**: Developed using the Go programming language.
    - **Framework**: Utilizes the Gin framework for constructing the HTTP API.
    - **Database Connectivity**: Connects to a PostgreSQL database, managed through the GORM package.
    - **Data Serialization**: Incorporates Protobuf for optimized data serialization.

2. **Microservices Architecture**:
    - The project is structured as a microservice, focusing on modularity, scalability, and maintainability.
    - **Documentation**: Features Swagger documentation to facilitate API exploration, interaction, and integration.

3. **Optimization and Efficiency**:
    - Prioritizes throughput and service efficiency through tailored optimizations in database interactions and request handling.

4. **Containerization**:
    - Ensures seamless deployment and scalability by containerizing the application using Docker Compose, fostering consistency across diverse environments.

### Guidance:
- **For Question 1**: Refer to the `./assets/report.sql` file for the SQL script to address the first question.
- **For Question 2**: Execute the command `make run trades` or `make docker trades` to initiate the service. You can then test the API using Swagger(by visiting http://localhost:8080/api/docs/index.html) to interact and validate its functionality.
- **For Question 3**: Execute the command `make cli trades <YOUR_NUMBER>` to run the CLI component, generating and processing the specified number of trades.

### Goals:
As a dedicated microservice, the project aims to deliver a lightweight and efficient solution for retrieving the latest trades. By incorporating Protobuf for serialization, Swagger for documentation, and Docker Compose for containerization, the microservice is structured to support streamlined development, deployment, and ongoing maintenance, aligning with best practices in microservices architecture.

---
