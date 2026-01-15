#  🛜RescueMesh

**RescueMesh** is an offline-first, peer-to-peer (P2P) emergency communication system designed for disaster-prone areas where internet connectivity is non-existent. 

When traditional networks fail, RescueMesh allows any group of devices (rescue teams or civilians) to communicate as long as they are connected to the same local network via Wi-Fi, a mobile hotspot, or Ethernet LAN.



## ❓ Why RescueMesh?

In the immediate aftermath of natural disasters, cellular towers and ISPs are often the first things to go down. This creates a "blackout" that hinders rescue efforts. 

RescueMesh fills this gap by:
* **Eliminating Internet Dependency:** Communicates directly between devices.
* **Rapid Deployment:** Requires no configuration; just join the same local network and start chatting.
* **Resilience:** No central server means no single point of failure.

---

## 🛠️ Technical Architecture

### Peer-to-Peer Core
The system is built using the [libp2p-go](https://github.com/libp2p/go-libp2p) library. It establishes encrypted P2P connections and utilizes a **Chat Room abstraction** to manage group communications.

### Zero-Config Discovery (mDNS)
To find other devices without an internet connection, RescueMesh uses **Multicast DNS (mDNS)**. 
> mDNS resolves hostnames to IP addresses within small networks that do not have a local name server. It allows nodes to broadcast their presence and "discover" peers automatically on the same LAN.

### File Structure
* `📂 /cmd/rescuemesh/main.go` — The core logic for host creation and peer connection.
* `📂 /cmd/rescuemesh/logs.txt` — A persistent backup of all messages received.
* `📂 /internal/p2p/host.go` — Network host initialization.
* `📂 /internal/p2p/mdns.go` — mDNS discovery implementation.
* `📂 /internal/p2p/pubsub.go` — PubSub (GossipSub) implementation for the chat rooms.

---


## 🚀 Installation & Usage

### Step to Start the Backend 
* ## Navigate to the library
```bash 
cd RescueMesh
```
* # Info about the Flags
* `--Port` — The port where your host runs..
* `--same_string` — Used by MDNS to discover peers; this must be the same among peers wanting to connect.
* `--nick` —Your display name visible to other peers.
* `--room` —The name of the chat room (must be the same for all peers in a group).
* # Run command in another terminal/device connceted together via Wifi/Ethernet LAN to create another peer.
  
* PORT A:
 ```bash
  go run main.go --port 9000 --same_string xyz --room myroom --nick Akshay
```
* PORT B:
 ```bash
  go run main.go --port 9001 --same_string xyz --room myroom --nick Aryan
```

---

