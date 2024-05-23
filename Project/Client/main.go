package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type Medication struct {
	ID           int
	Name         string
	Dosage       string
	Manufacturer string
	Price        float64
}

func main() {
	for {
		// Connect to server
		conn, err := net.Dial("tcp", "192.168.43.31:8080")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}
		defer conn.Close()

		fmt.Print("Enter command (getAll, createSendToServer, delete, or update): ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "getAll":
			getAll(conn)
		case "createSendToServer":
			createSendToServer(conn)
		case "delete":
			delete(conn)
		case "update":
			update(conn)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}

func createSendToServer(conn net.Conn) {
	// Send request to server
	var (
		name, dosage, manufacturer string
		price                      float64
	)

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter dosage: ")
	fmt.Scanln(&dosage)
	fmt.Print("Enter manufacturer: ")
	fmt.Scanln(&manufacturer)
	fmt.Print("Enter price: ")
	fmt.Scanln(&price)

	item := Medication{ID: 1, Name: name, Dosage: dosage, Manufacturer: manufacturer, Price: price}
	create(conn, item)
}

func create(conn net.Conn, medication Medication) {
	request := []byte("insert ")

	x := []byte(strconv.Itoa(medication.ID))
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Name)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Dosage)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Manufacturer)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(strconv.FormatFloat(medication.Price, 'f', -1, 64))
	request = append(request, x...)

	conn.Write(request)
}

func getAll(conn net.Conn) {
	request := []byte("getAll")
	conn.Write(request)
	get := make([]byte, 1024)
	conn.Read(get)
	fmt.Println(string(get))
}

func delete(conn net.Conn) {

	var ID int
	fmt.Print("Enter id: ")
	fmt.Scanln(&ID)

	request := []byte("delete ")
	x := []byte(strconv.Itoa(ID))
	request = append(request, x...)

	conn.Write(request)
	get := make([]byte, 1024)
	conn.Read(get)
	fmt.Println(string(get))
}

func update(conn net.Conn) {

	var (
		name, dosage, manufacturer string
		price                      float64
		id                         int
	)

	fmt.Print("Enter id: ")
	fmt.Scanln(&id)
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter dosage: ")
	fmt.Scanln(&dosage)
	fmt.Print("Enter manufacturer: ")
	fmt.Scanln(&manufacturer)
	fmt.Print("Enter price: ")
	fmt.Scanln(&price)

	medication := Medication{ID: id, Name: name, Dosage: dosage, Manufacturer: manufacturer, Price: price}

	request := []byte("update ")

	x := []byte(strconv.Itoa(medication.ID))
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Name)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Dosage)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(medication.Manufacturer)
	x = append(x, " "...)
	request = append(request, x...)

	x = []byte(strconv.FormatFloat(medication.Price, 'f', -1, 64))
	request = append(request, x...)

	conn.Write(request)

	get := make([]byte, 1024)
	conn.Read(get)
	fmt.Println(string(get))
}
