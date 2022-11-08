module backend

go 1.17

require (
	github.com/julienschmidt/httprouter v1.3.0
	gorm.io/gorm v1.22.5
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/rainu/go-command-chain v0.1.0
	golang.org/x/crypto v0.0.0-20220208050332-20e1d8d225ab // indirect
	golang.org/x/sys v0.0.0-20220207234003-57398862261d // indirect
)

require (
	backend.com/models v0.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20220208144051-fde48d68ee68
	gorm.io/driver/mysql v1.2.3
)

replace backend.com/models => ./models
