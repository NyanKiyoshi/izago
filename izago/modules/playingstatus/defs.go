package playingstatus

import "github.com/NyanKiyoshi/izago/izago/dispatcher"

type _internal struct{}

var module = dispatcher.New(_internal{})

func init() {
	module.AddListener(onConnect)
}
