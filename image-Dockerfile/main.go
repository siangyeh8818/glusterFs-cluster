package main

import (
  "strconv"
  "strings"
  "bufio"
  "os"
  "os/exec"
  "fmt"
  "log"
  "io"
  "bytes"
)
func main() {
        var PLATFORMMODE string
        PLATFORMMODE = os.Getenv("PLATFORM_MODE")
        fmt.Printf("env for PLATFORM_MODE : " + PLATFORMMODE+"\n")
        if PLATFORMMODE=="host" {
                var list string
                list = read_etchosts()
                fmt.Println("list : "+list)
                list_array := strings.Split(list,";")
		fmt.Print("/etc/hosts 內含有的節點數量 : ")
                fmt.Println(strconv.Itoa(len(list_array)-1))
                cmdnumber,_ := exec_shell("gluster peer status")
		fmt.Println("cmdnumber : " + cmdnumber)
                numberarray := strings.Split(cmdnumber,":")
		fmt.Println("numberarray[0] : "+numberarray[0])
		fmt.Println("numberarray[1] : "+numberarray[1])
                peernumber := strings.Replace(numberarray[1], " ", "", -1)
		re_peernumber :=strings.Split(peernumber,"\n")
		fmt.Println("peernumber : "+re_peernumber[0])
//              集群連接的數量判斷不足 , 需進行連結
		fmt.Println(strconv.Itoa(len(list_array)-2))
                if re_peernumber[0] != strconv.Itoa((len(list_array)-2)) {
                        fmt.Println("Some glusterFS container is not yet join the cluster")
                	for i:=0; i<len(list_array)-1;i++ {
				fmt.Println(i)
				fmt.Println(" value for list_array is " + list_array[i])
				addcluster(list_array[i])				
			}  
		}else {
			fmt.Println("GlusterFS cluster number is correct !")
		}
                _,err_r :=exec_shell("gluster volume info")
		if strings.Contains(err_r, "No volumes present") {
			fmt.Println("You need create volume")
			createvolume(list,strconv.Itoa(len(list_array)-1))
		}else {
			fmt.Println("Volume is already existing , you don't need to create volume")
		}
		fmt.Println("Finish init glusterFS cluster........")
//		exec_shell("mount -t glusterfs localhost:gfs_bfop /opt/gfs")
        }
//      true_result,err_result := exec_shell("gluster peer status")
}

func createvolume(node_list string, replica_nu string) {
	cr_list_array := strings.Split(node_list,";")
	var combin_command string
	numb,_ := strconv.Atoi(replica_nu)
	combin_command = "gluster volume create gfs_bfop replica " + replica_nu +" "
	for k:=0; k< numb;k++ {
			combin_command = combin_command + cr_list_array[k] + ":/data/gfs "
	}
	combin_command = combin_command + " force"
	fmt.Println(combin_command)
        exec_shell(combin_command)
	exec_shell("gluster volume start gfs_bfop")
//	exec_shell("mount -t glusterfs localhost:gfs_bfop /opt/gfs")
}

func addcluster(node_name string) {
	fmt.Println("add node to cluster , node : " + node_name)
	exec_shell("gluster peer probe " + node_name)
}

func exec_shell(s_command string) (string , string ) {
	var stdoutBuf, stderrBuf bytes.Buffer
        cmd := exec.Command("/bin/bash", "-c",s_command)
        stdoutIn, _ := cmd.StdoutPipe()
        stderrIn, _ := cmd.StderrPipe()
        var errStdout, errStderr error
        stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
        stderr := io.MultiWriter(os.Stderr, &stderrBuf)
        err := cmd.Start()
        if err != nil {
                log.Fatalf("cmd.Start() failed with '%s'\n", err)
        }
        go func() {
                _, errStdout = io.Copy(stdout, stdoutIn)
        }()
        go func() {
                _, errStderr = io.Copy(stderr, stderrIn)
        }()
        err = cmd.Wait()
        if err != nil {
                log.Fatalf("cmd.Run() failed with %s\n", err)
        }
        if errStdout != nil || errStderr != nil {
                log.Fatal("failed to capture stdout or stderr\n")
        }
        outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
/*         fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)*/
	return outStr,errStr
}
func read_etchosts() (string){
	var hostnameresult string 
	file ,err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//是否有下一行
	for scanner.Scan() {
		if (strings.Contains(scanner.Text(),"ip6")|| strings.Contains(scanner.Text(),"localhost")) != true {
			fmt.Println("read string:",scanner.Text())
			temp := strings.Split(scanner.Text(),"	")
			fmt.Println("split string 0: ",temp[0])
			fmt.Println("split string 1: ",temp[1])
			hostnameresult = temp[1]+";"+hostnameresult
		}
	}
	fmt.Println("hostnameresult : ",hostnameresult)
	return hostnameresult
}

