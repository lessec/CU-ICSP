# WireGuard for CentOS

## WireGuard auto-setup

## System Specifications
1) OS: CentOS 9 Stream (20220113 aarch64)
2) Cloud: AWS EC2 (t4g.micro)
3) Location: N. Vierginia
4) CPU: 2v
5) Ram: 1GB
6) Disk: 10GiB (SSD)
7) Server: Apache/2.4.51
8) Lang: Go
9) Database: MySQL 8.0.27

## System Information
1) VM ID: `ec2-user`
2) DB PW: `uni.Coventry.ac.uk-annectent-2022`
3) DB Access: `asdf-asdf-asdf-asdf-asdf`
4) SSH Port: `22`
5) Admin email: `-`
6) Send email: `-`

## setup.sh
You can use `setup.sh` to install part. It is on `./annectent-root/setup.sh` and you just run `sh ./setup.sh`


## 1. System update & dependency

##### On shell ($)
```bash
sudo dnf update -y && sudo dnf install -y curl wget make unzip vim gcc gcc-c++ git git-lfs && git lfs install
```


## 2. System Security Setting

### 1) Setup Firewall
##### On shell ($)
```bash
sudo dnf install -y firewalld && sudo systemctl enable firewalld && sudo systemctl start firewalld
```

### 2) Setup Ping Block
##### On shell ($)
```bash
sudo vi /etc/sysctl.conf
```
##### On editor (vi)
```apache
net.ipv4.icmp_echo_ignore_all = 1
```
##### On shell ($)
```bash
sudo sysctl -p
```


## 3. Install WireGuard for Server

### 1) Install WireGuard
##### On shell ($)
```bash
sudo dnf install -y wireguard-tools && sudo mkdir -p /etc/wireguard/
```

### 2) Key generate and get key
##### On shell ($)
```bash
wg genkey | sudo tee /etc/wireguard/root_private.key && sudo chmod go= /etc/wireguard/root_private.key && sudo cat /etc/wireguard/root_private.key | wg pubkey | sudo tee /etc/wireguard/root_public.key
#wg genkey | tee root_private.key | wg pubkey > root_public.key
```
You can check
##### On shell ($)
```bash
sudo cat /etc/wireguard/root_private.key && sudo cat /etc/wireguard/root_public.key
```

### 3) Create WireGuard Configuration File
##### On shell ($)
```bash
sudo vi /etc/wireguard/wg-root0.conf
```
##### On editor (vi) - example
```apache
[Interface]
# Interface IP Address of WG server (WG IP)
Address = 10.119.0.1/24
# Save the configuration when a new client will add
SaveConfig = true
# port number of WG server  (default: 51820)
ListenPort = 1194
# Private Key of WG server
PrivateKey = [Sever's_private.key]
# Command to be executed before starting the interface (default: 51820/udp)
PostUp = firewall-cmd --zone=public --add-port 1194/udp && firewall-cmd --zone=public --add-masquerade
# Command to be executed before turning off the interface (default: 51820/udp)
PostDown = firewall-cmd --zone=public --remove-port 1194/udp && firewall-cmd --zone=public --remove-masquerade
```
- Address – the private IP address for the Interface (wg-root0.conf).
- SaveConfig = true – saves the state of the interface on restart or shutdown of the server. If don't wnat, change true to false.
- PrivateKey – the key that we have just generated.
- ListenPort – the port where the WireGuard daemon listens.
- PostUp – this command will be executed before firing up the interface
- PostDown – this command will be executed before turning off the interface.

##### On shell ($)
```bash
sudo chmod -R 600 /etc/wireguard
```

### 4) Start WireGuard
##### On shell ($)
```bash
sudo wg-quick up wg-root0 && sudo wg
```

### 5) Enable IP Forwarding on the Server
##### On shell ($)
```bash
sudo vi /etc/sysctl.conf
```
##### On editor (vi)
```bash
# for enabling IPv4/IPv6 forwarding (WG)
net.ipv4.ip_forward = 1
net.ipv6.conf.all.forwarding = 1
```
##### On shell ($)
```bash
sudo sysctl -p
```

### 6) WireGuard Start & Stop
#### a) Start WG
##### On shell ($)
```bash
sudo wg-quick up /etc/wireguard/wg-root0.conf

sudo systemctl start wg-quick@wg-root0.service
```
#### b) Stop WG
##### On shell ($)
```bash
sudo wg-quick down /etc/wireguard/wg-root0.conf

sudo systemctl stop wg-quick@wg-root0.service
```
#### c) Run WG on system boot
##### On shell ($)
```bash
sudo systemctl enable wg-quick@wg-root0.service
```
#### d) Check WG Operation (Traffic)
##### On shell ($)
```bash
sudo wg
```


## 4. Install WireGuard for Client

### 1) Install WireGuard
##### On shell ($)
```bash
sudo dnf install -y wireguard-tools && sudo mkdir -p /etc/wireguard/
```
### 2) Key generate and get key
##### On shell ($)
```bash
wg genkey | sudo tee /etc/wireguard/door_private.key && sudo chmod go= /etc/wireguard/door_private.key && sudo cat /etc/wireguard/door_private.key | wg pubkey | sudo tee /etc/wireguard/door_public.key
#wg genkey | tee door_private.key | wg pubkey > root_public.key
```
You can check
##### On shell ($)
```bash
sudo cat /etc/wireguard/door_private.key && sudo cat /etc/wireguard/door_public.key
```
### 3) Create WireGuard Configuration File
##### On shell ($)
```bash
sudo vi /etc/wireguard/wg-door0.conf
```
##### On editor (vi) - example
```apache
[Interface]
# Interface IP Address of WG client (WG IP / Begin 1.119.1.1)
Address = 10.119.1.1/24
# Save the configuration when a new client will add
SaveConfig = true
# Private Key of WG client
PrivateKey = [Client's_private.key]

[Peer]
# Public Key of WG server
PublicKey = [Server's_public.key]
# set ACL
AllowedIPs = 0.0.0.0/0
# IP address and Port of WG server
Endpoint = [Server's_Public_IP]:1194
# Packet refresh every 30s
PersistentKeepalive = 30
```
- PrivateKey – the key generated on the client machine.
- Address – the IP address for the Interface (wg-door0.conf).
- PublicKey – the public key of the VPN server machine to which we want to connect.
- AllowedIPs – all allowed IP addresses for traffic flow using the VPN.
- Endpoint – we will provide the IP address and port number of the WG server machine to which we want to connect.
##### On shell ($)
```bash
sudo chmod -R 600 /etc/wireguard
```
### 4) Start WireGuard
##### On shell ($)
```bash
sudo wg-quick up wg-door0 && sudo wg
```

### 5) Add Peer on WG Server
##### On shell ($)
```bash
sudo wg set wg-root0 peer [Client's_public.key] allowed-ips [Client's_interface_IP]/32
```

### 6) WireGuard Start & Stop
#### a) Start WG
##### On shell ($)
```bash
sudo wg-quick up /etc/wireguard/wg-root0.conf

sudo systemctl start wg-quick@wg-root0.service
```
#### b) Stop WG
##### On shell ($)
```bash
sudo wg-quick down /etc/wireguard/wg-root0.conf

sudo systemctl stop wg-quick@wg-root0.service
```
#### c) Run WG on system boot
##### On shell ($)
```bash
sudo systemctl enable wg-quick@wg-root0.service
```
#### d) Check WG Operation (Traffic)
##### On shell ($)
```bash
sudo wg
```
