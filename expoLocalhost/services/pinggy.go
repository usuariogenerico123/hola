package services

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)



func ExposeOnlyPort(port string, system string){
	cancel := make(chan os.Signal, 1)
	cmd, info, err := pinggyTunnel(port, system)
	if(err != nil){
		fmt.Println(err)
		return 
	}
	data := strings.Split(info, "\n")
	fmt.Println("Tunnel pinggy inciado:")
	fmt.Println("Puerto: "+port)
	for _, v := range data[len(data)-5:len(data)-1]{
		fmt.Println("Link: "+"\033[0;33m"+v+"\033[0m")
	}

	signal.Notify(cancel, syscall.SIGTERM, os.Interrupt)
	<-cancel 
	fmt.Println("Pinggy cerrado")
	cmd.Process.Kill()
	
}

func ExposeServer(port string, path string){
	cancel := make(chan os.Signal, 1)
	ok := make(chan bool)
	go Server(port, path, ok)
	serverOk :=<- ok
	
	if(!serverOk){
		fmt.Println("No se pudo iniciar el servidor en el puerto:"+port)
		return
	}
	fmt.Println("Servidor iniciado en el puerto: "+port)

	cmd, info, err := pinggyTunnel(port, "linux")
	if(err != nil){
		fmt.Println(err)
		return
	}
	data := strings.Split(info, "\n")
	fmt.Println("Tunnel pinggy inciado:")
	fmt.Println("Servidor iniciado: http://localhost:"+port)
	for _, v := range data[len(data)-5:len(data)-1]{
		fmt.Println("Link: "+"\033[0;33m"+v+"\033[0m")
	}
	signal.Notify(cancel, syscall.SIGINT, os.Interrupt)
	<- cancel
	cmd.Process.Kill()
	


}


 




func pinggyTunnel(port string, system string)(*exec.Cmd, string, error){
	var info bytes.Buffer
	var errors bytes.Buffer

	command := fmt.Sprintf(`ssh -i ./pinggyKey -T -p 443 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -R0:127.0.0.1:%s free.pinggy.io`, port)
	if(system == "windows"){
		command = fmt.Sprintf(`ssh -i .\pinggyKey -T -p 443 -o StrictHostKeyChecking=no -R0:127.0.0.1:%s free.pinggy.io`, port)
	}
	c := strings.Split(command, " ")
	
	_, err := os.Stat("./pinggyKey")
	if(err != nil){
		generaK()
	}
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Stdout = &info
	cmd.Stderr = &errors
	go func (){
		cmd.Run()

	}()
	fmt.Println("Espera consiguiendo datos...")
	time.Sleep(10 * time.Second)
	
	return cmd, info.String(), nil
}


func generaK() error {
	var info bytes.Buffer
	var err bytes.Buffer
	fmt.Println("Generando keys para usar pinggy.io....")
	cmd := exec.Command("ssh-keygen", "-f", "pinggyKey", "-N", "", "-t", "rsa")
	cmd.Stdout = &info 
	cmd.Stderr = &err
	fmt.Printf("\t%s", "------------hecho---------------")
	cmd.Run()
	return nil
}
