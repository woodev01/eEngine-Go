package main

func main() {
	go func() {
		serveRoom(":8080", "chatroom", "chat")
	}()

	go func() {
		serveRoom(":8081", "battleroom", "battle")
	}()

	x := 0
	for true {
		x = (x + 1) % 100
	}
}
