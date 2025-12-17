# tunnlrX 🚀
### High-Performance, P2P-First Ingress for Developers

**tunnlrX** is a modern, self-hostable alternative to Ngrok and Cloudflare Tunnel. Built from the ground up on **QUIC** and **UDP Hole Punching**, it provides a blazing-fast, secure bridge between the public internet and your local machine.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## ⚡ Why tunnlrX?

Most tunneling tools rely on legacy TCP-based relays, which suffer from **TCP Meltdown** and high latency. **tunnlrX** solves this with a next-gen networking stack:

* **P2P-First Architecture:** Uses STUN and UDP Hole Punching to establish direct connections between users and your local machine, bypassing the server entirely for 10x lower latency.
* **QUIC Native:** For connections that require a relay (Symmetric NATs), tunnlrX uses QUIC to eliminate Head-of-Line blocking and ensure stability on poor Wi-Fi.
* **Zero-Knowledge Privacy:** Unlike SaaS providers, tunnlrX supports E2E TLS passthrough. Your data remains encrypted even while passing through the relay.
* **Integrated Inspector:** A built-in local dashboard at `localhost:4040` to view, filter, and replay HTTP requests for effortless debugging.

---

## 🏗 How it Works



1.  **Discovery:** The Agent uses STUN to discover its public-facing NAT mapping.
2.  **Signaling:** The Agent maintains a persistent QUIC control channel to the tunnlrX server.
3.  **The Punch:** When a public request arrives, the server coordinates a "rendezvous" between the peer and the agent.
4.  **Data Flow:** Traffic flows directly (P2P). If the NAT is too restrictive, it automatically falls back to a high-speed QUIC relay.

---

## 🚀 Quick Start

### 1. Deploy the Server (Public VPS)
The server acts as the control plane and STUN rendezvous point.
```
tunnlrx-server start --domain yourdomain.com
```

### 2. Run the Agent (Local Machine)
Expose your local development server to the world instantly:
```
# Map port 8080 to [https://my-app.yourdomain.com](https://my-app.yourdomain.com)
tunnlrx up 8080 --subdomain my-app
```

---

## 🛠 Features

* **Automatic SSL:** Seamless Let's Encrypt integration for all subdomains.
* **Protocol Agnostic:** Tunnel HTTP, TCP (SSH/DB), or UDP (Game Servers).
* **Passwordless Auth:** Authenticate agents using Passkeys (WebAuthn).
* **Request Replay:** Re-send webhooks or API calls with one click from the Inspector UI.

---

## 📦 Installation

**Using Go:**
```
go install [github.com/eyanshu1997/tunnlrx@latest](https://github.com/eyanshu1997/tunnlrx@latest)
```

---

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.