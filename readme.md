# High-Performance Image Storage Solution

## Overview

This project is designed to be a high-performance, scalable, and resilient image storage solution leveraging the principles of distributed systems. It utilizes Kubernetes for orchestration and MinIO for distributed object storage, providing an efficient and reliable way to store and manage large volumes of images.

## Architecture

The architecture is built on a microservices-based approach, ensuring high availability and fault tolerance. It consists of:

- **API Gateway**: Serves as the entry point for all client requests, directing them to the appropriate service.
- **Storage Service**: Handles image upload, download, and management, backed by MinIO for storage.
- **MinIO**: A high-performance distributed object storage server, configured for automatic sharding and erasure coding to ensure data durability and high availability.

The entire system is deployed on Kubernetes, which automates deployment, scaling, and operations of application containers across clusters of hosts. We are using namespaces in kubernetes to simulate servers in different locations and using minIO's server side bucket replication to ensure data consistency across 4 different "servers". By default, minIO uses erasure coding to provide data redundancy and availability. For each write operation, minIO partitions the obejct into data and parity shards, by default are using EC: 4, meaning that we are setting aside 25% of storage space for parity, and require a minimum of 75% drive space available for both reading and writing.

## Technologies Used

- **Go**: For writing backend services, offering high performance and easy concurrency.
- **React**: Powers the front-end application, providing a responsive user interface.
- **PostgreSQL**: Database for managing users and authentication
- **Docker & Docker Compose**: For containerizing the application and its dependencies, simplifying deployment and development.
- **Kubernetes**: For orchestrating container deployment, scaling, and management.
- **MinIO**: For distributed object storage, offering scalability, security, and performance.
