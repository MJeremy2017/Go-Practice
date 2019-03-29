package main

import "log"

func init()  {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)  // full file name and line number
}

func main() {
	log.Println("This is a message")

	log.Fatalln("Fatal message")

	log.Panicln("Panic message")
}


