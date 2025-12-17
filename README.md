# tunnlrX 🚀
### The Unified Connectivity Engine: Public Ingress & P2P Mesh

**tunnlrX** is a next-generation networking tool that combines the instant accessibility of **Ngrok** with the secure, high-performance P2P mesh of **WireGuard**. Whether you need a public URL for a webhook or a private, zero-latency tunnel to your home lab, tunnlrX handles it through a single, lightweight Go agent.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## ⚡ Two Modes. One Agent.

tunnlrX automatically chooses the best transport for your specific use case:

### 1. v1.0: Public Ingress (QUIC Relay)
Expose local ports to the public internet instantly. 
* **Tech:** QUIC (UDP) multiplexing.
* **Best for:** Webhooks, sharing live demos, and public API testing.
* **No Install:** Visitors use a standard browser to hit your `*.tunnlr.me` domain.

### 2. v1.1: Private Mesh (WireGuard P2P)
Connect your devices in a secure, encrypted "flat" network.
* **Tech:** WireGuard + STUN Hole Punching.
* **Best for:** SSH, Database access, and secure Team collaboration.
* **Pure P2P:** Data flows directly between nodes with 0ms added relay latency.

---

## 🏗 Architecture



1.  **Control Plane:** Orchestrates WireGuard keys and manages public subdomains.
2.  **Relay Plane:** Uses QUIC to tunnel traffic for clients behind restrictive Symmetric NATs.
3.  **P2P Plane:** Facilitates UDP hole punching for direct device-to-device communication.

---

## 🚀 Quick Start

### 1. Start the Control Plane
Deploy on a VPS with a public IP to manage your mesh and public tunnels.
```bash
tunnlrx-server start --domain yourdomain.com
```

### 2. Public Ingress Mode (v1.0)
Expose your local web server to the world:
```bash
# Map localhost:8080 to [https://dev-app.yourdomain.com](https://dev-app.yourdomain.com)
tunnlrx up 8080 --subdomain dev-app
```

### 3. Mesh Mode (v1.1)
Join your machine to your private, encrypted P2P mesh:
```bash
# Join the mesh and get a private IP (e.g., 10.0.0.5)
tunnlrx join --alias my-laptop
```

---

## 🛠 Technical Differentiators

| Feature | tunnlrX | Ngrok / Cloudflare |
| :--- | :--- | :--- |
| **Transport** | **QUIC (UDP)** | TCP (HTTP/2) |
| **P2P Support** | **Native WireGuard** | Relay Only |
| **Latency** | **Minimal (Hole Punching)** | High (Server Hop) |
| **Privacy** | **E2E Encrypted** | Provider can decrypt |
| **Auth** | **Passkeys / WebAuthn** | Email / SSO |

---

## 📦 Installation

**Using Go:**
```bash
go install [github.com/eyanshu1997/tunnlrx@latest](https://github.com/eyanshu1997/tunnlrx@latest)
```

---

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.