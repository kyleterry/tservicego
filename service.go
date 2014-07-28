package tservicego

/*
    "target":"#tenyks",
    "command":"PRIVMSG",
    "mask":"unaffiliated/vhost-",
    "direct":true,
    "nick":"vhost-",
    "host":"unaffiliated/vhost-",
    "fullmsg":":vhost-!~vhost@unaffiliated/vhost- PRIVMSG #tenyks :tenyks-demo: weather 97217",
    "full_message":":vhost-!~vhost@unaffiliated/vhost- PRIVMSG #tenyks :tenyks-demo: weather 97217",
    "user":"~vhost",
    "fromchannel":true,
    "from_channel":true,
    "connection":"freenode",
    "payload":"weather 97217",
    "meta":{
        "name":"Tenyks",
        "version":"1.0"
    }
*/
type Message struct {
	Target string
	Command string
	Mask string
	Direct string
	Nick string
	Host string
	Full_message string
	User string
	From_channel bool
	Connection string
	Payload string
	Meta *TenyksMeta
}

type TenyksMeta struct {
	Name string
	Version string
}

type Connection struct {
	
}
