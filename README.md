# 🚀 TunnlrX

> The simplest self-hosted tunneling platform with automatic domains, HTTPS, and secure defaults — built for developers and teams.

---

## 🌍 Vision

TunnlrX makes it easy to expose local or private services to the internet securely, without vendor lock-in. Unlike existing solutions, TunnlrX is **open source, self-hostable, automated, and user-friendly**.

**Core principles:**
- **Ease of use:** One command to expose any service.
- **Automation:** Domains, DNS, and HTTPS managed for you.
- **Security:** Encrypted by default, with access control.
- **Flexibility:** Works for developers, self-hosters, and teams.
- **Extensibility:** CLI, API, and GUI interfaces.

---

## 📦 Features

- ✅ Self-hosted control server (Go)
- ✅ Lightweight client (Go)
- ✅ Auto HTTPS via Let’s Encrypt
- ✅ Automatic domain provisioning (Cloudflare/Route53 integrations)
- ✅ HTTP, TCP, UDP, WebSockets, QUIC support
- ✅ CLI tool (`tunnlrx up 3000`)
- ✅ Web Dashboard (React/Next.js)
- ✅ GUI Client (Tauri/Electron)
- ✅ Multi-tenant support (teams/orgs)
- ✅ OAuth2/SSO Authentication
- ✅ REST/gRPC APIs for automation

---

## 🛠 Architecture

**Components:**

- **Control Plane (Server)**
  - Manages domains, certs, users, tunnel lifecycle
  - Exposes APIs + dashboard

- **Data Plane (Agent)**
  - Lightweight binary that establishes tunnels
  - Multiplexes connections

- **Clients**
  - CLI for developers
  - GUI for non-technical users

**Flow:**
1. User runs client (`tunnlrx up 8080`)
2. Client connects to server → assigns domain + cert
3. Secure tunnel established (QUIC/HTTPS)
4. Service exposed at `https://<subdomain>.tunnlr.me`

---


## 📂 Repository Structure

```
/tunnlrx
 ├── server/         # Go backend (control + data plane)
 ├── client/         # CLI client
 ├── dashboard/      # React/Next.js web UI
 ├── docs/           # Documentation site (MkDocs/Docusaurus)
 ├── scripts/        # DevOps, setup, deploy
 ├── .github/        # Actions CI/CD, issue templates
 ├── README.md       # Project overview
 ├── ROADMAP.md      # Development roadmap
 ├── CONTRIBUTING.md # Contribution guidelines
 ├── CODE_OF_CONDUCT.md
 └── LICENSE         # Apache 2.0
```
## 📜 License

TunnlrX is open-source under the **Apache 2.0 License**.

---

## 📊 Feature Matrix (Competitors vs TunnlrX)

| Tool              | Open Source | Self-Host | Auto HTTPS | Auto DNS | Multi-Protocol (TCP/UDP/QUIC) | GUI | API/Automation | Multi-Tenant |
|-------------------|-------------|-----------|-------------|----------|-------------------------------|-----|----------------|--------------|
| **ngrok**         | ❌          | ❌        | ✅          | ❌       | HTTP/TCP only                 | ❌  | Limited         | ✅ (paid)    |
| **Cloudflare Tunnel** | ❌     | ❌        | ✅          | ✅ (Cloudflare only) | HTTP/TCP | ❌ | Limited         | ✅ (Cloudflare account) |
| **frp**           | ✅          | ✅        | ❌          | ❌       | ✅ (TCP/UDP)                   | ❌  | Partial         | ❌           |
| **boringproxy**   | ✅          | ✅        | ✅          | ❌       | HTTP/TCP only                 | Web UI | ✅             | ❌           |
| **localtunnel**   | ✅          | ❌        | ❌          | ❌       | HTTP only                     | ❌  | ❌             | ❌           |
| **sish**          | ✅          | ✅        | ❌          | ❌       | HTTP/TCP/SSH                  | ❌  | Partial         | ❌           |
| **TunnlrX (planned)** | ✅     | ✅        | ✅          | ✅       | ✅ (HTTP/TCP/UDP/QUIC)        | ✅  | ✅             | ✅           |

## 🗺 Roadmap

### **MVP (0–3 months)**
- [ ] Go server + client for TCP/HTTP tunnels
- [ ] CLI (`tunnlrx up`)
- [ ] Let’s Encrypt HTTPS

### **Phase 2 (3–6 months)**
- [ ] Auto DNS/domain provisioning (Cloudflare/Route53)
- [ ] Web Dashboard (React + Go)
- [ ] OAuth2 authentication

### **Phase 3 (6–12 months)**
- [ ] GUI client
- [ ] Multi-protocol support (UDP, WebSockets, QUIC)
- [ ] Team/org support
- [ ] SaaS offering (managed hosting)


---
## 🤝 Contributing

We ❤️ contributions! To get started:

1. Fork the repo
2. Create a feature branch
3. Commit changes
4. Open a PR

Please read [CONTRIBUTING.md](./CONTRIBUTING.md) before submitting.

