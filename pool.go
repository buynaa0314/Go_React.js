type Pool struct {
	Register   chan *client
	Unregister chan *client
	Clients    map[*Client]bool
	Broadcast  chan Message
}
NewPool()
