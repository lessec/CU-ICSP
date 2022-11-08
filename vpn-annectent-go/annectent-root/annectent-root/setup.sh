sudo echo -e "\n1. Update CentOS"
sudo dnf update -y 

echo -e "\n2. Installing basic dependency library"
sudo dnf install -y curl wget make unzip gcc gcc-c++ go sqlite git git-lfs && git lfs install

echo -e "\n3. Installing Firewall and add start at boot"
sudo dnf install -y firewalld && sudo systemctl enable firewalld && sudo systemctl start firewalld

echo -e "\n4. Installing dependency library for Annectent\n\tGo, NodeJS, MySQL"
sudo dnf install -y go nodejs sqlite mysql-server && sudo systemctl enable mysqld && sudo systemctl start mysqld

echo -e "\n5. Installing WireGuard"
sudo dnf install -y wireguard-tools && sudo mkdir -p /etc/wireguard && sudo chmod -R 600 /etc/wireguard

echo -e "\n6. Setup enable IP forwarding for WireGuard"
sudo echo '# for enabling IPv4/IPv6 forwarding (WG)' >> /etc/sysctl.conf
sudo echo 'net.ipv4.ip_forward = 1' >> /etc/sysctl.conf
sudo echo 'net.ipv6.conf.all.forwarding = 1' >> /etc/sysctl.conf
sudo sysctl -p

echo -e "\n7. Installing backend dependency library"
cd ~/vpn-annectent-go/annectent-root/backend && sudo rm -rf wireguard && go build .

echo -e "\n8. Installing frontend dependency library"
sudo npm i -g @aws-amplify/cli
cd ~/vpn-annectent-go/annectent-root/frontend && rm -rf amplify && npm i && npn audit fix && npm run build

cd ~ && clear
echo -e "\n\nFinished to enviroment setup!\nYou can start Annectent or WG."
echo -e "\nPlease flow next setup guidline"
echo -e "\t1. sudo mysql_secure_installation -> Make root password VPN-annectent-go-root0 (Only test)"
echo -e "\t2. mysql -u root -p -> Create database"
echo -e "\t\tCREATE DATABASE annectent_root;"
echo -e "\t\texit;"
echo -e "\t3. Setup Amplify"
echo -e "\t\t1) npx amplify configure"
echo -e "\t\t2) npx amplify init"
echo -e "\t\t3) npx amplify add auth"
echo -e "\t\t4) npx amplify push"
echo -e "\t\tPlease flow Google Sign-In / Login with Amazon"
echo -e "\t\t : https://docs.amplify.aws/lib/auth/social/q/platform/js/"
echo -e "\t4. Ready to start. Open 2 terminal."
echo -e "\t\t1) 1st termial: Go to backend directory and run"
echo -e "\t\t - Location: ~/vpn-annectent-go/annectent-root/backend"
echo -e "\t\t - On termial: ./backend or go run ."
echo -e "\t\t2) 2nd termial: Go to frontend directory and run"
echo -e "\t\t - Location: ~/vpn-annectent-go/annectent-root/frontend"
echo -e "\t\t - On termial: npm run dev"