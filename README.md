# **Load Balancer**

## **Overview**

This project is a simple **Application Layer (Layer 7) Load Balancer** written in **Golang**. It accepts HTTP requests from clients and distributes them across a pool of backend servers. The load balancer ensures fault tolerance, minimizes response time, and avoids overloading any single server.

---

## **Features**

- **Round-Robin Load Balancing**: Distributes requests evenly across all backend servers.
- **Health Checks**: Periodically checks backend server availability and excludes unhealthy servers.
- **Dynamic Backend Pool**: Add or remove backend servers at runtime without restarting the load balancer.
- **Fault Tolerance**: Automatically redirects requests to healthy servers if one server goes offline.
- **Concurrency**: Handles multiple client requests concurrently using **Goroutines**.

---

## **Technologies Used**

- **Golang**: Language for its simplicity and concurrency support.
- **net/http**: For HTTP server and client handling.
- **Goroutines**: To handle multiple requests and health checks simultaneously.
- **Channels**: For communication between goroutines (e.g., health checks).

---
