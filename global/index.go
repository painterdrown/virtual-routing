package global

var Host string
var Port string
var Lock chan bool = make(chan bool, 1)
