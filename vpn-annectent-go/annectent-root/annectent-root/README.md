# Annectent Root by Go

## Finished to enviroment setup!\nYou can start Annectent or WG.

### You can easly setup through run `setup.sh`
### or flow this
##### 1. Update CentOS
```bash
sudo dnf update -y 
```
##### 2. Installing basic dependency library
```bash
sudo dnf install -y curl wget make unzip gcc gcc-c++ go sqlite git git-lfs && git lfs install
```
##### 3. Installing Firewall and add start at boot
```bash
sudo dnf install -y firewalld && sudo systemctl enable firewalld && sudo systemctl start firewalld
```
##### 4. Installing dependency library for Annectent\n\tGo, NodeJS, MySQL"
```bash
sudo dnf install -y go nodejs sqlite mysql-server && sudo systemctl enable mysqld && sudo systemctl start mysqld
```
##### 5. Installing WireGuard"
```bash
sudo dnf install -y wireguard-tools && sudo mkdir -p /etc/wireguard && sudo chmod -R 600 /etc/wireguard
```
##### 6. Setup enable IP forwarding for WireGuard
```bash
sudo vi /etc/sysctl.conf
```
```apache
# for enabling IPv4/IPv6 forwarding (WG)
net.ipv4.ip_forward = 1
net.ipv6.conf.all.forwarding = 1
```
```bash
sudo sysctl -p
```
##### 7. Installing backend dependency librar
```bash
cd ~/vpn-annectent-go/annectent-root/backend && sudo rm -rf wireguard && go build .
```
##### 8. Installing frontend dependency library"
```bash
sudo npm i -g @aws-amplify/cli
cd ~/vpn-annectent-go/annectent-root/frontend && rm -rf amplify && npm i && npn audit fix && npm run build
```

## If tou finish, please flow next setup guidline
1. sudo mysql_secure_installation -> Make root password VPN-annectent-go-root0 (Only test)
2. mysql -u root -p -> Create database
```sql
CREATE DATABASE annectent_root;
exit;
```
3. Setup Amplify
    1) npx amplify configure
    2) npx amplify init
    3) npx amplify add auth
    4) npx amplify push

Please flow Google Sign-In / Login with Amazon: https://docs.amplify.aws/lib/auth/social/q/platform/js/

4. Ready to start. Open 2 terminal.
    1) 1st termial: Go to backend directory and run
        - Location: ~/vpn-annectent-go/annectent-root/backend
        - On termial: ./backend or go run .
    2) 2nd termial: Go to frontend directory and run
        - Location: ~/vpn-annectent-go/annectent-root/frontend
        - On termial: npm run dev

### As a way to connect an endpoint, easily install the WireGuard application on your mobile and scan it with the QR code.