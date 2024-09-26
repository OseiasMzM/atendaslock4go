package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Paths struct {
	servicePath    string
	frontPath      string
	monitoringPath string
	databasePath   string
	filesPath      string
	logsPath       string
	sslCertFile    string
	sslKeyFile     string
	dbDataFile     string
	dbConfigFile   string
}

var (
	pathInst *Paths
	pathOnce sync.Once
)

func getPathInstance() *Paths {
	pathOnce.Do(func() {
		pathInst = newPaths()
	})
	return pathInst
}

func newPaths() *Paths {
	var p = Paths{}
	p.servicePath = filepath.Dir(os.Args[0]) + "\\"
	p.frontPath = p.servicePath + "www"
	p.monitoringPath = p.servicePath + "Monitor\\"
	p.databasePath = p.servicePath + "DB\\"
	p.filesPath = p.servicePath + "Files\\"
	p.logsPath = p.servicePath + "Logs\\"

	p.sslCertFile = p.servicePath + "certs\\cert.crt"
	p.sslKeyFile = p.servicePath + "certs\\cert.key"

	if runtime.GOOS == "windows" {
		p.monitoringPath = "C:\\Makrolock\\Monitor\\"
		p.databasePath = "C:\\ProgramData\\Makrolock\\DB\\"
		p.filesPath = "C:\\ProgramData\\Makrolock\\Files\\"
		p.logsPath = "C:\\ProgramData\\Makrolock\\Logs\\"

	}

	p.dbDataFile = p.databasePath + "data.mkl"
	p.dbConfigFile = p.databasePath + "config.mkl"

	createDirectory(p.monitoringPath)
	createDirectory(p.databasePath)
	createDirectory(p.filesPath)
	createDirectory(p.logsPath)
	createFileDB(p.dbDataFile)
	createFileDB(p.dbConfigFile)

	// TODO: Check docker volume

	return &p
}
func pathsDataBase() *Paths {
	var p = Paths{}

	p.databasePath = "C:\\ProgramData\\Makrolock\\DB\\"
	p.dbDataFile = p.databasePath + "data.mkl"
	p.dbConfigFile = p.databasePath + "config.mkl"

	return &p
}

func createDirectory(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		return err == nil
	}
	return true
}

func createFileDB(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, _ = os.Create(path)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("file exists", path)
	}
	return true
}

// Public =============================================

func InstanceDataBase() string {
	return newPaths().databasePath
}
func InstancePathDbData() string {
	p := pathsDataBase()
	return p.dbDataFile
}
func InstancePathDbConfig() string {
	p := pathsDataBase()
	return p.dbConfigFile
}

func ServicePath() string {
	return getPathInstance().servicePath
}

func FrontPath() string {
	return getPathInstance().frontPath
}

func SslCertFile() string {
	return getPathInstance().sslCertFile
}

func SslKeyFile() string {
	return getPathInstance().sslKeyFile
}

func GenerateMonitoringFile(computer, user string) string {
	var now = time.Now()
	var fileName = computer + "." + user + "." + now.Format("2006-01-02.15-04-05") + ".jpeg" //Computerzor.Userzor.50514-03-46.15-314-146
	createDirectory(getPathInstance().monitoringPath + computer)
	return getPathInstance().monitoringPath + computer + "\\" + fileName
}
