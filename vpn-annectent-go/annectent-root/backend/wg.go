package main

import (
	"fmt"
	cmdchain "github.com/rainu/go-command-chain"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	WGPort        = "1194"
	WGDirPath     = "wireguard/"
	WGPeerDirPath = "wireguard/peer/"
	DatKey        = ".key"
	RootPriKey    = "root_private"
	RootPubKey    = "root_public"
	PeerPriKey    = "door_private"
	PeerPubKey    = "door_public"
	DatConf       = ".conf"
	WGRootConf    = "wg-root"
	PeerConf      = "wg-peer"
)

func WGRootKeys() {
	WGDirCreate()
	PeerDirCreate()
	RootGenPrivateKey()
	RootGenPublicKey()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func RootGenPrivateKey() {
	err := cmdchain.Builder().
		Join("wg", "genkey").
		Join("tee", WGDirPath+RootPriKey+DatKey).
		Finalize().Run()
	checkError(err)

	RootPrivateKeySec()
}

func RootGenPublicKey() {
	err := cmdchain.Builder().
		Join("cat", WGDirPath+RootPriKey+DatKey).
		Join("wg", "pubkey").
		Join("tee", WGDirPath+RootPubKey+DatKey).
		Finalize().Run()
	checkError(err)
}

func RootPrivateKeySec() {
	_, err := os.Stat(WGDirPath + RootPriKey + DatKey)
	checkError(err)

	err = os.Chmod(WGDirPath+RootPriKey+DatKey, 0600)
	checkError(err)
}

func ReadRootPrivateKey() string {
	var plaintextPrivate, err = os.ReadFile(WGDirPath + RootPriKey + DatKey)
	checkError(err)

	_, err = os.Open(WGDirPath + RootPriKey + DatKey)
	checkError(err)

	return string(plaintextPrivate)
}

func ReadRootPublicKey() string {
	var plaintextPublic, err = os.ReadFile(WGDirPath + RootPubKey + DatKey)
	checkError(err)

	_, err = os.Open(WGDirPath + RootPubKey + DatKey)
	checkError(err)

	return string(plaintextPublic)
}

func WGPeerKeys(devicePk uint) {
	PeerGenPrivateKey(devicePk)
	PeerGenPublicKey(devicePk)
}

func PeerGenPrivateKey(devicePk uint) {
	devicePkString := strconv.Itoa(int(devicePk))

	err := cmdchain.Builder().
		Join("wg", "genkey").
		Join("tee", WGPeerDirPath+PeerPriKey+devicePkString+DatKey).
		Finalize().Run()
	checkError(err)

	PeerPrivateKeySec(devicePk)
}

func PeerGenPublicKey(devicePk uint) {
	devicePkString := strconv.Itoa(int(devicePk))

	err := cmdchain.Builder().
		Join("cat", WGPeerDirPath+PeerPriKey+devicePkString+DatKey).
		Join("wg", "pubkey").
		Join("tee", WGPeerDirPath+PeerPubKey+devicePkString+DatKey).
		Finalize().Run()
	checkError(err)
}

func PeerPrivateKeySec(devicePk uint) {
	devicePkString := strconv.Itoa(int(devicePk))
	_, err := os.Stat(WGPeerDirPath + PeerPriKey + devicePkString + DatKey)
	checkError(err)

	err = os.Chmod(WGPeerDirPath+PeerPriKey+devicePkString+DatKey, 0600)
	checkError(err)
}

func ReadPeerPrivateKey(devicePk uint) string {
	devicePkString := strconv.Itoa(int(devicePk))

	var plaintextPrivate, err = os.ReadFile(WGPeerDirPath + PeerPriKey + devicePkString + DatKey)
	checkError(err)

	_, err = os.Open(WGPeerDirPath + PeerPriKey + devicePkString + DatKey)
	checkError(err)

	return string(plaintextPrivate)
}

func ReadPeerPublicKey(devicePk uint) string {
	devicePkString := strconv.Itoa(int(devicePk))

	var plaintextPublic, err = os.ReadFile(WGPeerDirPath + PeerPubKey + devicePkString + DatKey)
	checkError(err)

	_, err = os.Open(WGPeerDirPath + PeerPubKey + devicePkString + DatKey)

	return string(plaintextPublic)
}

func WGStatus() {
	err := cmdchain.Builder().
		Join("sudo", "wg").
		Finalize().Run()
	checkError(err)
}

func WGUp() {
	cwd, _ := os.Getwd()

	err := cmdchain.Builder().
		Join("sudo", "wg-quick", "up", cwd+"/"+WGDirPath+WGRootConf+"0"+DatConf).
		Finalize().Run()
	checkError(err)
}

func WGDown() {
	cwd, _ := os.Getwd()

	err := cmdchain.Builder().
		Join("sudo", "wg-quick", "down", cwd+"/"+WGDirPath+WGRootConf+"0"+DatConf).
		Finalize().Run()
	checkError(err)
}

func WGDirCreate() {
	//Create a WireGuard directory
	err := os.Mkdir(WGDirPath, 744)
	checkError(err)
}

func WGDirRemove() {
	//Create a WireGuard directory
	err := os.RemoveAll(WGDirPath)
	checkError(err)
}

func PeerDirCreate() {
	//Create a Peer directory in WireGuard directory
	err := os.Mkdir(WGPeerDirPath, 744)
	checkError(err)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func confError(e error) bool {
	if e != nil {
		fmt.Println(e.Error())
	}
	return (e != nil)
}

func WGConfCreate() {
	// check if file exists
	var _, err = os.Stat(WGDirPath + WGRootConf + "0" + DatConf)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(WGDirPath + WGRootConf + "0" + DatConf)
		if confError(err) {
			return
		}
		defer file.Close()
	}
}

func WGConfRemove() {
	var err = os.Remove(WGDirPath + WGRootConf + "0" + DatConf)
	if confError(err) {
		return
	}
}

func WGConfWriteInterface() {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(WGDirPath+WGRootConf+"0"+DatConf, os.O_RDWR, 0644)
	if confError(err) {
		return
	}
	defer file.Close()

	RootInterfaceAddress := "10.119.0.1"

	_, err = file.WriteString("[Interface]" +
		"\nAddress = " + RootInterfaceAddress + "/24" +
		//"\nSaveConfig = true" +
		//"\nPostUp = firewall-cmd --zone=public --add-port 1194/udp && firewall-cmd --zone=public --add-masquerade" +
		//"\nPostDown = firewall-cmd --zone=public --remove-port 1194/udp && firewall-cmd --zone=public --remove-masquerade" +
		"\nListenPort = " + WGPort +
		"\nPrivateKey = " + ReadRootPrivateKey())
	if confError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if confError(err) {
		return
	}
}

func WGConfWritePeer(devicePk uint, ip string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(WGDirPath+WGRootConf+"0"+DatConf, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if confError(err) {
		return
	}
	defer file.Close()

	RootPeerAddress := ip

	_, err = file.WriteString("\n[Peer]" +
		"\nAllowedIPs = " + RootPeerAddress + "/32" +
		"\nPublicKey = " + ReadPeerPublicKey(devicePk))
	if confError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if confError(err) {
		return
	}
}

func WGConfRead(data *string) {
	// Open file for reading.
	var file, err = os.OpenFile(WGDirPath+WGRootConf+"0"+DatConf, os.O_RDWR, 0644)
	if confError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var contents = make([]byte, 1024)
	for {
		_, err = file.Read(contents)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			confError(err)
			break
		}
	}

	fmt.Println(string(contents))
	*data = string(contents)
}

func PeerConfCreate(devicePk uint) {
	devicePkString := strconv.Itoa(int(devicePk))

	// check if file exists
	var _, err = os.Stat(WGPeerDirPath + PeerConf + devicePkString + DatConf)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(WGPeerDirPath + PeerConf + devicePkString + DatConf)
		if confError(err) {
			return
		}
		defer file.Close()
	}
}

func PeerConfRemove(devicePk string) {
	var err = os.Remove(WGPeerDirPath + PeerConf + devicePk + DatConf)
	if confError(err) {
		return
	}
}

func PeerConfWriteInterface(devicePk uint, ip string) {
	devicePkString := strconv.Itoa(int(devicePk))

	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(WGPeerDirPath+PeerConf+devicePkString+DatConf, os.O_RDWR, 0644)
	if confError(err) {
		return
	}
	defer file.Close()

	DoorInterfaceAddress := ip

	_, err = file.WriteString("[Interface]" +
		"\nAddress = " + DoorInterfaceAddress + "/24" +
		"\nListenPort = " + WGPort +
		"\nDNS = 9.9.9.9" +
		"\nPrivateKey = " + ReadPeerPrivateKey(devicePk))
	if confError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if confError(err) {
		return
	}
}

func PeerConfWritePeer(devicePk uint) {
	devicePkString := strconv.Itoa(int(devicePk))

	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(WGPeerDirPath+PeerConf+devicePkString+DatConf, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if confError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("\n[Peer]" +
		"\nAllowedIPs = 0.0.0.0/0" +
		"\nEndpoint = " + GetIP() + ":1194" +
		"\nPublicKey = " + ReadRootPublicKey())
	if confError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if confError(err) {
		return
	}
}

func PeerConfRead(devicePk int) string {
	devicePkString := strconv.Itoa(devicePk)

	var peerConfContents, err = os.ReadFile(WGPeerDirPath + PeerConf + devicePkString + DatConf)
	checkError(err)

	_, err = os.Open(WGPeerDirPath + PeerConf + devicePkString + DatConf)

	return string(peerConfContents)
}

func GetIP() string {
	ipify, _ := http.Get("https://api.ipify.org")
	ipaddr, _ := ioutil.ReadAll(ipify.Body)
	return string(ipaddr)
}
