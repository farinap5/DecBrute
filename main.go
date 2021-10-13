package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheynewallace/tabby"
	"net/http"
	"os"
)

//------// Global Variables
var port string = "4444"
var lhost string = "0.0.0.0"
var n int = 0 // Number of the actual credential
var list []string
var user string
var verboso bool = false
var hn int = 1 // Host number
//-----//

func _help() {
	t := tabby.New()
	t.AddHeader("COMMAND","DESCRIPTION","DEFAULT CONF","REQUIRED")
	t.AddLine("-h","Help menu.","false","NO")
	t.AddLine("-H","Local Host.","0.0.0.0","NO")
	t.AddLine("-P","Local Port.","4444","NO")
	t.AddLine("-u","User.","admin","YES")
	t.AddLine("-l","Wordlist.","NO","YES")
	t.AddLine("-e","Enable service.","false","YES")
	t.AddLine("-v","Verbose Mode.","false","NO")
	println(" ")
	t.Print()
	os.Exit(1)
}

// Confirm password state
func conf(w http.ResponseWriter, r *http.Request) {
	// http://0.0.0.0:4444/conf?state=no&credent=sapato
	// http://0.0.0.0:4444/conf?state=ye&credent=sapato
	var resp string
	// State - yes/no
	confirm,ok := r.URL.Query()["state"]
	if !ok || len(confirm) < 1 {
		resp = ""
	} else {
		resp = confirm[0]
	}
	
	// Credential
	var resp2 string
	confirm2,ok := r.URL.Query()["credent"]
	if !ok || len(confirm2) < 1 {
		resp2 = ""
	} else {
		resp2 = confirm2[0]
	}

	// If confirmed password is equal to the actual password,
	// it can be passed.
	if resp == "no" && list[n] == resp2 {
		if verboso {
			println("[\u001B[1;31m!\u001B[0;0m]- No:",list[n],"-",n)
		}
		n++
	}
	if resp == "ye" {
		println("[\u001B[1;32m!\u001B[0;0m]-",n,"Credential Found! -",resp2)
	}
}

// Catch user and password
func cred(w http.ResponseWriter, r *http.Request) {
	tos := user+":"+list[n]
	w.Write([]byte(tos))
}

// Root directory
func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It Works!"))
	if (verboso == true) {
		fmt.Println("[\u001B[1;32m!\u001B[0;0m]",hn,"- New Host:",r.Host)
		hn++
	}
}

// Start service
func start_server() {
	http.HandleFunc("/",root)
	http.HandleFunc("/conf", conf)
	http.HandleFunc("/cred", cred)
	lis := lhost+":"+port
	println("[\u001B[1;32mOK\u001B[0;0m]- INIT service -",lis)
	http.ListenAndServe(lis,nil)
}


func main() {
	var h = flag.String("H","0.0.0.0","Local Host.")
	var p = flag.String("P","4444","Local Port.")
	var l = flag.String("l","","List of passwords.")
	var u = flag.String("u","admin","User.")
	var stsv = flag.Bool("e",false,"Enable service.")
	var verb = flag.Bool("v",false,"Verbose mode.")
	var help = flag.Bool("h",false,"Help Menu.")
	flag.Parse()

	if *help {
		_help()
	}

	//----//
	verboso = *verb
	lhost = *h
	port = *p
	user = *u
	//----//

	if len(*l) > 2 {
		file,err := os.Open(*l)
		if err != nil {
			println("[!]- Error while trying to open file!")
			os.Exit(1)
		}
		println("[\u001B[1;32mOK\u001B[0;0m]- File:",*l)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			list = append(list, scanner.Text())
		}
		file.Close()
		/*for i := 0;i < len(list);i++ {
			println(list[i])
		}*/
	}

	if *stsv {
		println("[\u001B[1;32mOK\u001B[0;0m]- Local Host Listener:", lhost+":"+port)
		if verboso {
			println("[\u001B[1;32mOK\u001B[0;0m]- Verbose mode: true")
		}
		println("[\u001B[1;32mOK\u001B[0;0m]- User:",user)
		start_server()
	}
}
