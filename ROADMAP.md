# Roadmap: tunnlrX

This document outlines the development roadmap for tunnlrX, a self-hostable Ngrok alternative.

## Phase 1: Basic HTTP Tunneling

**Goal:** Implement the core functionality of creating a tunnel and forwarding HTTP traffic from the public server to the local client.

**Tasks:**

*   [ ] **1.1: Implement Tunnel Creation:**
    *   [ ] Define a gRPC service for the client to request a tunnel with a specific subdomain.
    *   [ ] The server should handle the request and register the tunnel, mapping the subdomain to the client.
    *   [ ] The server should provide the client with a unique URL for the tunnel.

*   [ ] **1.2: Implement HTTP Forwarding:**
    *   [ ] The public HTTP server should look up the subdomain of incoming requests.
    *   [ ] If a tunnel is registered for the subdomain, the server should forward the request to the corresponding client.
    *   [ ] The client should receive the forwarded request and send it to the local service.
    *   [ ] The client should send the response from the local service back to the server, which then sends it to the original caller.

*   [ ] **1.3: Basic Client UI:**
    *   [ ] The client should display the public URL of the tunnel.
    *   [ ] The client should show the status of the tunnel (connected, disconnected).
    *   [ ] The client should log requests and responses.

## Phase 2: Advanced Networking (P2P)

**Goal:** Implement direct P2P connections between the server and client using UDP hole punching and QUIC for lower latency.

**Tasks:**

*   [ ] **2.1: UDP Hole Punching:**
    *   [ ] Implement a mechanism for the client and server to discover each other's public UDP addresses.
    *   [ ] Implement the UDP hole punching process to establish a direct connection.

*   [ ] **2.2: QUIC Integration:**
    *   [ ] Use the established P2P connection to create a QUIC stream.
    *   [ ] Forward traffic over the QUIC stream instead of the gRPC connection for lower latency.

## Phase 3: Features and Polishing

**Goal:** Add features to make tunnlrX more user-friendly and powerful.

**Tasks:**

*   [ ] **3.1: Custom Domains:**
    *   [ ] Allow users to use their own custom domains for tunnels.
    *   [ ] Implement a verification process to ensure the user owns the domain.

*   [ ] **3.2: Authentication:**
    *   [ ] Add user authentication to the server.
    *   [ ] Only allow authenticated users to create tunnels.

*   [ ] **3.3: Web UI:**
    *   [ ] Create a web-based UI for the server to manage tunnels and users.
    *   [ ] Display real-time statistics and logs.

## Phase 4: Production Readiness

**Goal:** Make tunnlrX stable, secure, and easy to deploy.

**Tasks:**

*   [ ] **4.1: Security Hardening:**
    *   [ ] Implement TLS for all communication.
    *   [ ] Add rate limiting and other security measures to the public server.
    *   [ ] Conduct a security audit of the codebase.

*   [ ] **4.2: Observability:**
    *   [ ] Add structured logging to all components.
    *   [ ] Implement metrics and tracing to monitor the health and performance of the system.

*   [ ] **4.3: Deployment:**
    *   [ ] Create Docker images for the client and server.
    *   [ ] Write a Helm chart for deploying the server to Kubernetes.
    *   [ ] Improve documentation with clear instructions for deployment and configuration.
