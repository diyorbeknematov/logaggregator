package server

import (
	"logaggregator/models"
	"logaggregator/storage"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func StartTCPServer(address string, storage storage.IStorage) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("TCP serverni boshlashda xatolik: %w", err)
	}
	defer listener.Close()
	log.Printf("Server %s adresida ishlayapti...\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Bog'lanishda xatolik: %v", err)
			continue
		}
		go handleConnection(conn, storage)
	}
}

// handleConnection TCP ulanishini boshqarish
func handleConnection(conn net.Conn, storage storage.IStorage) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		logLine := scanner.Bytes()

		log.Printf("Log qabul qilindi: %s", logLine)

		var logJson models.Log
		err := json.Unmarshal(logLine, &logJson)
		if err != nil {
			log.Println("Json datani unmarshal qilganda xatolik: ", err.Error())
			continue
		}
		// Logni SQLite bazasiga saqlash
		err = storage.LogRepository().SaveLog(logJson)
		if err != nil {
			log.Printf("Logni saqlashda xatolik: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ma'lumotlarni o'qishda xatolik: %v", err)
	}
}
