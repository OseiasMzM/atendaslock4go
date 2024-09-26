package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kardianos/service"
)

const service_name = "atendaslock4goServer"

type Config struct {
	Name, DisplayName, Description string

	Dir  string
	Exec string
	Args []string
	Env  []string

	Stderr, Stdout string
}

type program struct {
	exit    chan struct{}
	service service.Service

	*Config

	cmd *exec.Cmd
}

func (p *program) Start(s service.Service) error {
	//fullExec, err := exec.LookPath(p.Exec)
	//if err != nil {
	//return fmt.Errorf("failed to find executable %q: %v", p.Exec, err)
	//}

	//p.cmd = exec.Command(fullExec, p.Args...)
	//p.cmd.Dir = p.Dir
	//p.cmd.Env = append(os.Environ(), p.Env...)

	go p.run()
	return nil
}

func (p *program) run() {
	defer func() {
		if service.Interactive() {
			p.Stop(p.service)
		} else {
			p.service.Stop()
		}
	}()

	if p.Stderr != "" {
		f, err := os.OpenFile(p.Stderr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			//logger.Warningf("Failed to open std err %q: %v", p.Stderr, err)
			return
		}
		defer f.Close()
		p.cmd.Stderr = f
	}
	if p.Stdout != "" {
		f, err := os.OpenFile(p.Stdout, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			//logger.Warningf("Failed to open std out %q: %v", p.Stdout, err)
			return
		}
		defer f.Close()
		p.cmd.Stdout = f
	}

	//err := p.cmd.Run()
	//if err != nil {
	//logger.Warningf("Error running: %v", err)
	//}

	InitializeServer()
}
func (p *program) Stop(s service.Service) error {
	close(p.exit)
	if p.cmd.Process != nil {
		p.cmd.Process.Kill()
	}
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

func getConfigPath() (string, error) {
	fullexecpath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, execname := filepath.Split(fullexecpath)
	ext := filepath.Ext(execname)
	name := execname[:len(execname)-len(ext)]

	return filepath.Join(dir, name+".json"), nil
}

func getConfig(path string) (*Config, error) {
	//f, err := os.Open(path)
	//if err != nil {
	//	return nil, err
	//}
	//defer f.Close()

	conf := &Config{}

	//r := json.NewDecoder(f)
	//err = r.Decode(&conf)
	//if err != nil {
	//	return nil, err
	//}
	return conf, nil
}

func ParseParams() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return ""
	}
}

func InstallService() bool {
	fmt.Println("Installing Service....")

	svcConfig := &service.Config{
		Name:        service_name,
		DisplayName: "Atendas Backend",
		Description: "Main Handler",
	}

	fmt.Println("Service Set: " + service_name)
	fmt.Println("Setting Config....")

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("ServiceConfig Error...." + err.Error())
		return false
	}

	fmt.Println("Config Set!")
	fmt.Println("Going to install....")

	err = s.Install()
	if err != nil {
		fmt.Println("InstallError...." + err.Error())
		return false
	}

	fmt.Println("Installed Successfully!")
	fmt.Println("Going to start....")

	err = s.Start()
	if err != nil {
		fmt.Println("StartError...." + err.Error())
		return false
	}

	fmt.Println("Started Successfully!")
	return true
}

func UninstallService() bool {
	fmt.Println("Uninstalling Service....")

	svcConfig := &service.Config{
		Name:        service_name,
		DisplayName: "Atendas Backend",
		Description: "Main Handler",
	}

	fmt.Println("Service Set: " + service_name)
	fmt.Println("Setting Config....")

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("ServiceConfig Error...." + err.Error())
		return false
	}

	fmt.Println("Config Set!")
	fmt.Println("Going to stop....")

	err = s.Stop()
	if err != nil {
		fmt.Println("StopError...." + err.Error())
		fmt.Println("Failed to stop but will try to uninstall anyway....")
	} else {
		fmt.Println("Stopped Successfully!")
	}

	fmt.Println("Going to uninstall....")

	err = s.Uninstall()
	if err != nil {
		fmt.Println("UninstallError...." + err.Error())
		return false
	}

	fmt.Println("Uninstalled Successfully!")

	return true
}

func RunService() bool {
	fmt.Println("Running Service....")

	//svcFlag := flag.String("service", "", "Control the system service.")
	//flag.Parse()

	configPath, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	config, err := getConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	svcConfig := &service.Config{
		Name:        service_name,
		DisplayName: "Atendas Backend",
		Description: "Main Handler",
	}

	fmt.Println("Service Set: " + service_name)
	fmt.Println("Setting Config....")

	prg := &program{
		exit:   make(chan struct{}),
		Config: config,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("ServiceConfig Error...." + err.Error())
		return false
	}

	fmt.Println("Config Set!")
	fmt.Println("Going to run....")

	err = s.Run()
	if err != nil {
		fmt.Println("RunError...." + err.Error())
		return false
	}

	fmt.Println("Running Successfully!")

	return true
}
