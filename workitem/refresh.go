package workitem

type callback func()

var wiclients map[string][]callback

// RegisterRefreshClient register clients for refresh
func RegisterRefreshClient(id string, callback func()) {
	wiclients[id] = append(wiclients[id], callback)
}
