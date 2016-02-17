package main

import (
	"flag"
	"fmt"
	"time"
	"udp"
	"os/exec"
)

func main() {
	var count int = 0
	var backupCount int = count
	sendTimer := time.NewTimer(time.Second)
	masterTimer := time.NewTimer(3000 * time.Millisecond)
	receiveChannel := make(chan int)
	sendChannel := make(chan int)
	killChannel := make(chan struct{})
	killedChannel := make(chan struct{})
	
	// Do some checking to see if booted to master or backup
	backup := flag.Bool("backup", false, "Set backup or not")
	flag.Parse()
	// Do the right thing depending on if master or backup
	if *backup {
		fmt.Println("Backup mode..")
		go udp.ReadUdp(receiveChannel, killChannel, killedChannel)
	} else {
		fmt.Println("Master mode..")
		startBackup()
		go udp.SendUdp(sendChannel)
	}
	for {
		select {
		case <-sendTimer.C:
			if !*backup {
				fmt.Printf("Count: %d\n", count)
				count++
				sendChannel<-count
				sendTimer.Reset(time.Second)
			}
		case <-masterTimer.C:
			// Master has timed out.
			if *backup {
				*backup = false
				count = backupCount
				sendTimer.Reset(time.Second)
				masterTimer.Stop();
				// Kill the UDP listener.
				close(killChannel)
				time.Sleep(2000*time.Millisecond)
				//<-killedChannel
				go udp.SendUdp(sendChannel)
				startBackup()
			}
		case mes := <-receiveChannel:
			if *backup{
				backupCount = mes
				masterTimer.Reset(1500*time.Millisecond)
			}
		}
	}
}

func startBackup(){
	(exec.Command("osascript", "-e","tell app \"Terminal\" to do script \"/Users/Lars/Dropbox/NTNU/Semester6/TTK4145/TTK4145--vinger/oving6/phoenix -backup\"")).Start()
}
