package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"mangahub/pkg/database"
)

type Progress struct {
	Username string `json:"username"`
	MangaID  string `json:"manga_id"`
	Chapter  int    `json:"chapter"`
}

func StartTCPServer() {
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TCP Server running on :7070")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	message, err := reader.ReadBytes('\n')
	if err != nil {
		return
	}

	var progress Progress

	err = json.Unmarshal(message, &progress)
	if err != nil {
		fmt.Println("invalid json")
		return
	}

	fmt.Println("=== Reading Progress ===")
	fmt.Println("User:", progress.Username)
	fmt.Println("Manga:", progress.MangaID)
	fmt.Println("Chapter:", progress.Chapter)
	
	saveProgress(progress)
}

func saveProgress(progress Progress) {
	query := `
	INSERT INTO reading_progress (
		username,
		manga_id,
		chapter
	)
	VALUES (?, ?, ?)
	`

	_, err := database.DB.Exec(
		query,
		progress.Username,
		progress.MangaID,
		progress.Chapter,
	)

	if err != nil {
		fmt.Println("save progress error:", err)
		return
	}

	fmt.Println("progress saved")
}