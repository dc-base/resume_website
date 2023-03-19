package main

import (
	"fmt" //Print package
	"log"
	"net/http"
	"time"
	"os/signal"
	"syscall"
	"os"
	"context"
)
func checkMethod(w http.ResponseWriter, r *http.Request) bool{
	if r.Method != "GET" {
		http.Error(w, "Only get is support", http.StatusNotFound)
		return false
	} else {
		return true
	}
}
func aboutHdlr(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/about" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	if !checkMethod(w, r){return}
	fmt.Fprintf(w, "We are in the About Page")
}
func chadHdlr(w http.ResponseWriter, r *http.Request){
//	if r.URL.Path != "" || r.URL.Path != "/" {
//		http.Error(w, "c04", http.StatusNotFound)
//		return
//	}
	if !checkMethod(w, r){return}
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "\t<head>")
	fmt.Fprintf(w, "\t\t<title>I am normal</title>")
	fmt.Fprintf(w, "\t\t\t<img src='media/test.jpg' alt='fail' style='width:600px;height:600px;'>")
	fmt.Fprintf(w, "\t\t\t<strong>surprisingly im not high</strong>")
	//fmt.Fprintf(w, "\t\t\t<img src='media/retard.jpg' alt='fail'>")
	fmt.Fprintf(w, "\t<head/>")
	fmt.Fprintf(w, "<html>")


	//fmt.Fprintf(w, "We are in the Landing Page")
}
func landingHdlr(w http.ResponseWriter, r *http.Request){
//	if r.URL.Path != "" || r.URL.Path != "/" {
//		http.Error(w, "c04", http.StatusNotFound)
//		return
//	}
	if !checkMethod(w, r){return}
	fmt.Fprintf(w, "We are in the Landing Page")
}
//func shutdownHdlr(w http.ResponseWriter, r *http.Request){
//	if r.URL.Path != "/shutdown" {
//		http.Error(w, "404", http.StatusNotFound)
//		return
//	}
//	s := http.Server{Addr: ":8080", nil}
//	s.Shutdown(context.Background())
//	fmt.Printf("Shutdown")
//	fmt.Fprintf(w, "Understood shutdown")
//}
const port = 8080

//func exitInterrupt() {
//	done := make(chan os.Signal, 0)
//	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
//	<-done
//	ctxsd, cancel := context.WithTimeout(context.Background(), 0 * time.Second)
//	defer cancel() //Does not run until function returns
//	server.Shutdown(ctxsd)
//	fmt.Println("Stopped")
//
//}
func main() {
	//Listen for 8080/about
	//http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
	//	fmt.Fprintf(w, "Req for about page")
	//})
	//fs := http.FileServer(http.Dir("./media"))
	//http.Handle("/retard.jpg", http.FileServer(http.Dir("./")))
	mux := http.NewServeMux()
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	mux.HandleFunc("/about", aboutHdlr)
	//mux.HandleFunc("/", landingHdlr)
	mux.HandleFunc("/chad", chadHdlr)
	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./media"))))
	fmt.Println("Server started on: ", port)


	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error launching server")
		log.Fatal(err)
	}
	//http.HandleFunc("/about", aboutHdlr)
	//http.HandleFunc("/shutdown", shutdownHdlr)
	//var portNum = ":8080"
	//fmt.Printf("Start at: 8080\n")
	//if err := http.ListenAndServe(portNum, nil); err != nil {
	//	log.Fatal(err)
	//}
	//HTTPs Startup time"
	go func() {
		done := make(chan os.Signal, 0)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		ctxsd, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
		defer cancel() //Does not run until function returns
		fmt.Println("Stopped")
		server.Shutdown(ctxsd)
		fmt.Println("Stopped")
	}()
	time.Sleep(100 * time.Millisecond)

}
