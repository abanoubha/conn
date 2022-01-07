package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {
	whichPath, lookErr := exec.LookPath("/usr/bin/which")
	if lookErr != nil {
		panic(lookErr)
	}

	nmapPath, err := exec.Command(whichPath, "nmap").CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	// os.Stderr.WriteString(string(nmapPath)) // exact as below
	// fmt.Println(string(nmapPath))

	nmapStr := string(nmapPath)
	err = runNmap(nmapStr)
	if err != nil {
		panic(err.Error())
	}

	// output, err := exec.Command("nmap -sn 192.168.1.0/24").CombinedOutput()
	// if err != nil {
	// 	os.Stderr.WriteString(err.Error())
	// }
	// fmt.Println(string(output))

	// env := os.Environ()
	// execErr := syscall.Exec(string(nmapPath), args, env)
	// if execErr != nil {
	// 	panic(execErr)
	// }
}

func runNmap(nmap string) error {
	cmd := exec.Command(nmap, "-sn", "192.168.1.0/24")
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			os.Stdout.WriteString(scanner.Text())
		}
		done <- true
	}()
	cmd.Start()
	<-done

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// func runIt(nmap string) ([]byte, error) {
// 	var out bytes.Buffer
// 	cmd := exec.Command(nmap, "-sn", "192.168.1.0/24")
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out.Bytes(), nil
// }
