# 🚀 RescueMesh: High-Impact Implementation Roadmap

This roadmap focuses on the **Must-Have** features that demonstrate backend maturity, system design skills, and real-world utility for engineering recruiters and hackathon evaluators.

---

## 📦 Phase 1: Data Integrity & Reliability
*Goal: Ensure no data is lost and communication is optimized for low-bandwidth.*

- [ ] **SQLite Persistence:** Replace `logs.txt` with an SQLite database to handle message history, metadata, and crash recovery.
- [ ] **Protobuf Migration:** Replace JSON/String payloads with **Protocol Buffers** to reduce packet size by ~40% and optimize radio bandwidth.
- [ ] **Message Acknowledgments (ACKs):** Implement a feedback loop where nodes send a "received" signal back to the sender to verify delivery in the mesh.
- [ ] **Store-and-Forward:** Create a local queue to store messages for peers that are currently out of range, automatically re-sending once they are rediscovered.

---

## 🔋 Phase 2: Power Optimization (The Interview "Hook")
*Goal: Solve the #1 problem in P2P—battery drain.*

- [ ] **Burst Transmission Algorithm:** Implement a "Sleep-Wake" cycle where mDNS and GossipSub radios only activate for a 15-second "Burst" every 2 minutes.
- [ ] **Battery-Aware Logic:** Develop a trigger to automatically increase the "Sleep" interval when the device battery drops below 20%.

---

## 🛡️ Phase 3: Security & Identity
*Goal: Establish trust in a decentralized, anonymous environment.*

- [ ] **End-to-End Encryption (E2EE):** Wrap message payloads in an encryption layer (Noise Protocol or AES-GCM) so only the intended room members can read them.
- [ ] **Verified Responder Tags:** Implement a simple public-key verification system to distinguish official rescue personnel from civilian users.

---

## 📍 Phase 4: Disaster-Specific Intelligence
*Goal: Turn "chat" into "actionable rescue data."*

- [ ] **GPS Metadata Tagging:** Automatically append latitude/longitude to every SOS packet.
- [ ] **Priority Levels:** Add a header to messages (e.g., `CRITICAL`, `MEDICAL`, `INFO`) to allow the system to prioritize urgent packets during a "Burst."
- [ ] **Mesh Strength Indicator:** A backend logic that calculates the number of nearby hops and signal quality (RSSI) for UI visualization.

---

## ☁️ Phase 5: The Hybrid Bridge (Gateway)
*Goal: Bridge the gap between the offline mesh and the online command center.*

- [ ] **Internet Detection Worker:** A background Go routine that pings a public DNS (8.8.8.8) to detect WAN availability.
- [ ] **Cloud Sync Engine:** Once internet is detected, automatically POST all "unsynced" SQLite records to a centralized **PostgreSQL** cloud database.
- [ ] **Heartbeat to HQ:** Send a "Node Active" signal to the cloud dashboard when a gateway is reached to map the mesh's coverage area.

---

## 🖥️ Phase 6: Frontend (React) Essentials
*Goal: A high-contrast, low-power UI for emergency use.*

- [ ] **Peer Discovery Dashboard:** List all active nearby nodes and their last-seen status.
- [ ] **One-Tap SOS:** A high-priority broadcast button that sends location + battery status + name.
- [ ] **Sync Status Bar:** Clear visual feedback showing if the node is "Mesh Only" or "Cloud Synced."

---
### **Resume Talking Points Enabled by this Roadmap:**
* "Architected a **Hybrid Cloud-Mesh** system using **Go** and **libp2p** to maintain communication in 100% offline environments."
* "Optimized network overhead by **40%** via **Protobuf** serialization and implemented **SQLite** for edge-case data persistence."
* "Engineered a **Burst Transmission Algorithm** to reduce radio power consumption, extending mobile battery life for disaster scenarios."

**Would you like me to generate the Go code for the SQLite database setup or the Protobuf schema to get started?**