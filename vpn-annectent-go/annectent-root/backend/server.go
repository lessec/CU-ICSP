package main

import (
	"backend.com/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/skip2/go-qrcode"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	dataDir = "/var/lib/wireguard-ui"

	listenAddr            = ":8080"
	natEnabled            = true
	natLink               = "wlp2s0"
	clientIPRange         = ""
	authUserHeader        = "X-Forwarded-User"
	maxNumberClientConfig = 0

	wgLinkName   = "wg0"
	wgListenPort = 1194
	wgEndpoint   = "127.0.0.1:1194"
	wgAllowedIPs = []string{"0.0.0.0/0", "::/0"}
	wgDNS        = ""
	wgKeepAlive  = ""

	devUIServer = ""
	filenameRe  = regexp.MustCompile("[^a-zA-Z0-9]+")
)

type contextKey string

const key = contextKey("user")

type Server struct {
	serverConfigPath string
	mutex            sync.RWMutex
	ipAddr           net.IP
	clientIPRange    *net.IPNet
	assets           http.Handler
	db               *gorm.DB
}

// Setup Server

func NewServer() *Server {
	// prepare wg
	//WGDown()
	WGRootKeys()
	WGConfCreate()
	WGConfWriteInterface()
	//WGUp()

	s := Server{}

	// get db instance
	//s.db, _ = gorm.Open(sqlite.Open("my.db"), &gorm.Config{})
	dsn := "root:VPN-annectent-go-root0@tcp(127.0.0.1:3306)/annectent_root?charset=utf8mb4&parseTime=True&loc=Local"
	s.db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// migrate db
	s.db.AutoMigrate(&models.User{})
	s.db.AutoMigrate(&models.UserGroup{})
	s.db.AutoMigrate(&models.UserDevice{})
	s.db.AutoMigrate(&models.UUDBridge{})
	s.db.AutoMigrate(&models.UUGBridge{})

	return &s
}

// HTTP Server

func (s *Server) Start() error {
	// create routers
	router := httprouter.New()
	router.GET("/", s.Index)

	// UserGroup Routers
	router.POST("/api/v1/group/add", s.AddUserGroup)
	router.GET("/api/v1/group/get/all", s.GetAllUserGroups)
	router.GET("/api/v1/group/get/id/:id", s.GetUserGroup)
	router.POST("/api/v1/group/update/:id", s.UpdateUserGroup)
	router.POST("/api/v1/group/delete/:id", s.DeleteUserGroup)
	router.GET("/api/v1/group/users/:id", s.GetUsersRelatedUserGroup)

	// User Routers
	router.POST("/api/v1/user/add", s.AddUser)
	router.GET("/api/v1/user/get/id/:id", s.GetUser)
	router.GET("/api/v1/user/get/u/:username", s.GetUserByUserName)
	router.GET("/api/v1/user/get/all", s.GetAllUsers)
	router.POST("/api/v1/user/update/:id", s.UpdateUser)
	router.POST("/api/v1/user/delete/:id", s.DeleteUser)
	router.GET("/api/v1/user/validation", s.ValidationAccount)
	router.GET("/api/v1/user/connect/:id/g/:gid", s.ConnectUserGroup)
	router.GET("/api/v1/user/connect/:id/d/:did", s.ConnectUserDevice)
	router.GET("/api/v1/user/disconnect/:id/g/:gid", s.DisconnectUserGroup)
	router.GET("/api/v1/user/disconnect/:id/d/:did", s.DisconnectUserDevice)

	// UserDevice Routers
	router.POST("/api/v1/device/add/:name/ip/:ip/u/:uid", s.AddUserDeviceByName)
	router.GET("/api/v1/device/get/id/:id", s.GetUserDevice)
	router.GET("/api/v1/device/get/all", s.GetAllUserDevices)
	router.POST("/api/v1/device/delete/:id", s.DeleteUserDevice)
	router.POST("/api/v1/device/update/:id/n/:name/ip/:ip", s.UpdateUserDeviceName)
	router.POST("/api/v1/device/update/:id/rpri", s.RegeneratePrivateKey)
	router.POST("/api/v1/device/update/:id/rpub", s.RegeneratePublicKey)
	router.POST("/api/v1/device/update/:id/rqrc", s.RegenerateQRCode)
	router.POST("/api/v1/device/update/:id/rconf", s.RegenerateConfigFile)

	// Get private IP
	router.GET("/api/v1/device/newip", s.GetNewIp)

	// Get wireguard control in frontend
	router.GET("/api/v1/wiregaurd/up", s.WireGuardStart)
	router.GET("/api/v1/wiregaurd/down", s.WireGuardStop)

	// start server
	return http.ListenAndServe(":5555", router)
}

func (s *Server) CORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Credentials", "true")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	(*w).Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {
		http.Error(*w, "No Content", http.StatusNoContent)
		return
	}
}

// Index TODO: it is not meaningful.
func (s *Server) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// AddUserGroup Router
func (s *Server) AddUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// for solve CORS
	s.CORS(&w, r)

	// parse JSON
	decoder := json.NewDecoder(r.Body)
	newUserGroup := &models.UserGroup{}
	err := decoder.Decode(newUserGroup)
	if err != nil {
		panic(err)
	}

	// create model
	errForDB := s.db.Create(&newUserGroup).Error

	if errForDB != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return
}

// GetUserGroup Get a UserGroup Router
func (s *Server) GetUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// for solve CORS
	s.CORS(&w, r)

	// container
	var group models.UserGroup

	// parse url and get group from db
	id, _ := strconv.Atoi(ps.ByName("id"))
	err := s.db.First(&group, "id = ?", id).Error

	// db error handling
	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// object serialized
		serialized := models.SerializedUserGroup{
			ID:            group.ID,
			Name:          group.Name,
			Describe:      group.Describe,
			GrantedServer: group.GrantedServer,
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetAllUserGroups get all groups router
func (s *Server) GetAllUserGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// container for existing groups
	var groups []models.UserGroup

	// get all groups
	err := s.db.Find(&groups).Error

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// create container for serialized
		serialized := make([]models.SerializedUserGroup, len(groups))

		// add serialized objects
		for i, group := range groups {
			serialized[i] = models.SerializedUserGroup{
				ID:            group.ID,
				Name:          group.Name,
				Describe:      group.Describe,
				GrantedServer: group.GrantedServer,
			}
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteUserGroup delete group router
func (s *Server) DeleteUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, err := strconv.Atoi(ps.ByName("id"))
	if err == nil {
		// delete in db
		err2 := s.db.Delete(&models.UserGroup{}, id).Error

		if err2 == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

// UpdateUserGroup update group router
func (s *Server) UpdateUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// parse JSON
	decoder := json.NewDecoder(r.Body)

	// create new group
	newUserGroup := &models.UserGroup{}
	err := decoder.Decode(newUserGroup)
	if err != nil {
		panic(err)
	}

	// get existing model
	var group models.UserGroup
	s.db.First(&group, "id = ?", id)

	// update object
	group.Name = newUserGroup.Name
	group.Describe = newUserGroup.Describe
	group.GrantedServer = newUserGroup.GrantedServer

	// commit
	s.db.Save(&group)

}

// AddUser add user router
func (s *Server) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse JSON
	decoder := json.NewDecoder(r.Body)
	newUser := &models.User{}
	err := decoder.Decode(newUser)
	if err != nil {
		panic(err)
	}

	// create model
	errForDB := s.db.Create(&newUser).Error

	if errForDB != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return
}

// GetUser get user router
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// container
	var user models.User

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// get existing model
	err := s.db.First(&user, "id = ?", id).Error

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// created serialized object
		serialized := models.SerializedUser{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Type:    user.Type,
			Devices: s.GetUserRelatedUserDevice(user.ID),
			Groups:  s.GetUserRelatedUserGroup(user.ID),
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteUser delete user router
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, err := strconv.Atoi(ps.ByName("id"))
	if err == nil {
		// delete in db
		err2 := s.db.Delete(&models.User{}, id).Error

		if err2 == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

// UpdateUser update user router
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// parse JSON
	decoder := json.NewDecoder(r.Body)
	newUser := &models.User{}
	err := decoder.Decode(newUser)
	if err != nil {
		panic(err)
	}

	// update existing model
	var user models.User
	s.db.First(&user, "id = ?", id)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Type = newUser.Type

	// commit
	s.db.Save(&user)

}

// GetAllUsers get all users router
func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// container for existing models
	var users []models.User

	// get all model
	err := s.db.Find(&users).Error

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// create container for serialized object
		serialized := make([]models.SerializedUser, len(users))

		// create and add serialized models
		for i, user := range users {
			serialized[i] = models.SerializedUser{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Type:  user.Type,
			}
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// AddUserDeviceByName add device router
func (s *Server) AddUserDeviceByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// server lock itself
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// create new device but not commited
	newUserDevice := &models.UserDevice{}

	// get name from url and add it
	newUserDevice.Name = ps.ByName("name")
	newUserDevice.IP = ps.ByName("ip")

	// commit
	errForDB := s.db.Create(&newUserDevice).Error

	uid, _ := strconv.Atoi(ps.ByName("uid"))
	s.db.Create(&models.UUDBridge{
		UID:  uint(uid),
		UDID: newUserDevice.ID,
	})

	// wg
	WGDown()
	defer WGUp()

	WGPeerKeys(newUserDevice.ID)
	WGConfWritePeer(newUserDevice.ID, newUserDevice.IP)
	PeerConfCreate(newUserDevice.ID)
	PeerConfWriteInterface(newUserDevice.ID, newUserDevice.IP)
	PeerConfWritePeer(newUserDevice.ID)

	if errForDB != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return
}

// GetUserDevice get device router
func (s *Server) GetUserDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// container
	var device models.UserDevice

	// parse url and get device from s.db.
	id, _ := strconv.Atoi(ps.ByName("id"))
	err := s.db.First(&device, "id = ?", id).Error

	if err == nil {
		// if it request other type; qr, config
		format := r.URL.Query().Get("format")

		configData := s.generateConfigData()

		// if it request qr code
		if format == "qr" {
			// generate QR image
			qr := PeerConfRead(id)
			println(qr)
			png, err := qrcode.Encode(qr, qrcode.Medium, 220)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// write png in http body
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(png)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		}

		if format == "config" {
			filename := "test.conf"
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
			w.Header().Set("Content-Type", "application/config")
			w.WriteHeader(http.StatusOK)
			_, err := fmt.Fprint(w, configData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		// if it request just JSON
		w.Header().Set("Content-Type", "application/json")

		// create serialized object
		serialized := models.SerializedUserDevice{
			ID:         device.ID,
			IP:         device.IP,
			Name:       device.Name,
			PublicKey:  device.PublicKey,
			ConfigFile: device.ConfigFile,
			QRCode:     device.QRCode,
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetAllUserDevices get all devices router
func (s *Server) GetAllUserDevices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// container for existing model
	var devices []models.UserDevice

	// get all models
	err := s.db.Find(&devices).Error

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// create container for serialized models
		serialized := make([]models.SerializedUserDevice, len(devices))

		// create and add serialized model
		for i, device := range devices {
			serialized[i] = models.SerializedUserDevice{
				ID:         device.ID,
				IP:         device.IP,
				Name:       device.Name,
				PublicKey:  device.PublicKey,
				ConfigFile: device.ConfigFile,
				QRCode:     device.QRCode,
			}
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteUserDevice delete device router
func (s *Server) DeleteUserDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// wg
	s.mutex.Lock()
	defer s.mutex.Unlock()

	WGDown()
	defer WGUp()

	PeerConfRemove(ps.ByName("id"))

	// parse url
	id, err := strconv.Atoi(ps.ByName("id"))
	if err == nil {
		// delete model in db
		err2 := s.db.Delete(&models.UserDevice{}, id).Error

		if err2 == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// UpdateUserDeviceName update device router
func (s *Server) UpdateUserDeviceName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))
	ip := ps.ByName("ip")
	name := ps.ByName("name")

	// update just name only
	s.db.Model(&models.UserDevice{}).Where("id = ?", id).Update("Name", name).Update("IP", ip)
}

// RegeneratePrivateKey request regenrate private key router
func (s *Server) RegeneratePrivateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// update db
	id, _ := strconv.Atoi(ps.ByName("id"))
	s.db.Model(&models.UserDevice{}).Where("id = ?", id).Update("PrivateKey", "Regenerated!")
}

// RegeneratePublicKey regenerate public key router
func (s *Server) RegeneratePublicKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// update db
	s.db.Model(&models.UserDevice{}).Where("id = ?", id).Update("PublicKey", "Regenerated!")
}

// RegenerateQRCode regenerate qr code router
func (s *Server) RegenerateQRCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// update db
	s.db.Model(&models.UserDevice{}).Where("id = ?", id).Update("QRCode", "Regenerated!")
}

// RegenerateConfigFile regenerate config file router
func (s *Server) RegenerateConfigFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	// update db
	s.db.Model(&models.UserDevice{}).Where("id = ?", id).Update("ConfigFile", "Regenerated!")
}

func (s *Server) ValidationAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	var isVaild bool

	// parse JSON
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	usertype := r.URL.Query().Get("type")
	newUser := &models.User{
		Name:  name,
		Email: email,
		Type:  usertype,
	}

	// data container for find
	var userContainer models.User

	// find object in users include soft deleted
	s.db.Unscoped().Where("name = ? AND email = ?", newUser.Name, newUser.Email).Find(&userContainer)
	if userContainer.Model.ID == 0 {
		// non existed
		newUser.Type = models.Staff
		s.db.Create(newUser)
		isVaild = true
	} else if userContainer.DeletedAt.Valid {
		// invalidate
		isVaild = false
	} else {
		// normal
		isVaild = true
	}

	// if it request just JSON
	w.Header().Set("Content-Type", "application/json")

	// create serialized object
	var serialized struct {
		IsVaild bool `json:"isVaild"`
	}
	serialized.IsVaild = isVaild

	json.NewEncoder(w).Encode(serialized)
	w.WriteHeader(http.StatusOK)
}

// GetUserRelatedUserGroup get groups related an user
func (s *Server) GetUserRelatedUserGroup(id uint) []models.UserGroup {
	var bridges []models.UUGBridge
	var uidArray []uint

	// get bridges
	s.db.Find(&bridges, models.UUGBridge{UID: id})
	for uug := range bridges {
		// add group id to array
		uidArray = append(uidArray, bridges[uug].UGID)
	}

	// get groups
	var groups []models.UserGroup
	if uidArray != nil {
		s.db.Find(&groups, uidArray)
	}

	return groups
}

// GetUserRelatedUserDevice get devices related an user
func (s *Server) GetUserRelatedUserDevice(id uint) []models.UserDevice {
	var bridges []models.UUDBridge
	var uidArray []uint

	// get bridges
	s.db.Find(&bridges, models.UUDBridge{UID: id})
	for uud := range bridges {
		// add device id to array
		uidArray = append(uidArray, bridges[uud].UDID)
	}

	// get devices
	var devices []models.UserDevice
	if uidArray != nil {
		s.db.Find(&devices, uidArray)
	}

	return devices
}

func (s *Server) ConnectUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	uid, _ := strconv.Atoi(ps.ByName("id"))
	gid, _ := strconv.Atoi(ps.ByName("gid"))

	// create bridge
	bridge := models.UUGBridge{
		UID:  uint(uid),
		UGID: uint(gid),
	}

	err := s.db.Create(&bridge).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func (s *Server) ConnectUserDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	uid, _ := strconv.Atoi(ps.ByName("id"))
	did, _ := strconv.Atoi(ps.ByName("did"))

	// create bridge
	bridge := models.UUDBridge{
		UID:  uint(uid),
		UDID: uint(did),
	}

	err := s.db.Create(&bridge).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func (s *Server) DisconnectUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	uid, _ := strconv.Atoi(ps.ByName("id"))
	gid, _ := strconv.Atoi(ps.ByName("gid"))

	var bridge models.UUGBridge
	s.db.First(&bridge, models.UUGBridge{UID: uint(uid), UGID: uint(gid)})

	// get bridges
	err := s.db.Delete(&bridge).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) DisconnectUserDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	uid, _ := strconv.Atoi(ps.ByName("id"))
	did, _ := strconv.Atoi(ps.ByName("did"))

	// get bridges
	var bridge models.UUDBridge
	s.db.First(&bridge, models.UUDBridge{UID: uint(uid), UDID: uint(did)})

	err := s.db.Delete(&bridge).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) GetUserByUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	username := ps.ByName("username")

	var user models.User
	err := s.db.First(&user, "name = ?", username).Error

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		// created serialized object
		serialized := models.SerializedUser{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Type:    user.Type,
			Devices: s.GetUserRelatedUserDevice(user.ID),
			Groups:  s.GetUserRelatedUserGroup(user.ID),
		}

		json.NewEncoder(w).Encode(serialized)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

// GetUsersRelatedUserGroup get Users related a group
func (s *Server) GetUsersRelatedUserGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	// parse url
	id, _ := strconv.Atoi(ps.ByName("id"))

	var bridges []models.UUGBridge
	var uidArray []uint

	// get bridges
	s.db.Find(&bridges, models.UUGBridge{UGID: uint(id)})
	for uug := range bridges {
		// add user id to array
		uidArray = append(uidArray, bridges[uug].UID)
	}

	// get users
	var users []models.User
	if uidArray != nil {
		s.db.Find(&users, uidArray)
	}

	// create container for serialized object
	serialized := make([]models.SerializedUser, len(users))

	// create and add serialized models
	for i, user := range users {
		serialized[i] = models.SerializedUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Type:  user.Type,
		}
	}

	json.NewEncoder(w).Encode(serialized)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetNewIp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	var device models.UserDevice
	s.db.Unscoped().Order("IP desc").First(&device)

	var newIp string
	splited := strings.Split(device.IP, ".")

	if len(splited) < 4 {
		newIp = "10.119.1.1"
	} else {
		forth, _ := strconv.Atoi(splited[3])
		if forth >= 255 {
			third, _ := strconv.Atoi(splited[2])
			splited[2] = strconv.Itoa(third + 1)
			splited[3] = "1"
			newIp = strings.Join(splited, ".")
		} else {
			splited[3] = strconv.Itoa(forth + 1)
			newIp = strings.Join(splited, ".")
		}

	}

	serialized := models.SerializedUserDevice{
		IP: newIp,
	}

	json.NewEncoder(w).Encode(serialized)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) WireGuardStart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	WGUp()

	w.WriteHeader(http.StatusOK)
}

func (s *Server) WireGuardStop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// to solve CORS
	s.CORS(&w, r)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	WGDown()

	w.WriteHeader(http.StatusOK)
}

// generateConfigData create config data
func (s *Server) generateConfigData() string {
	return `[Interface]
Address = %s
PrivateKey = %s
%s
[Peer]
PublicKey = %s
AllowedIPs = %s
Endpoint = %s
%s
%s
`
}
