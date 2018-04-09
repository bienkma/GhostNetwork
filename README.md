# Introduction

I had learnt security information major at Academy of cryptography techniques of Vietnam. After finished college I have a work on some IT company with SysAdmin and currently a DevOpser role;) funny. I have planed make a C&C network for study
very long time but I han't complete because some reasons. GhostNetwork is a Go implementation of the C&C network for study.
You should run the GhostNetwork on Labs environment. I am not responsible if you use the GhostNetwork for illegal purposes.

# Installation

```bash
$ cd GhostNetwork
$ go build Server.go -o CNCServer
$ go build bot/bot.go -o bot
$ sudo mkdir -p /etc/cnc
$ sudo bash -c 'cat <<EOF > /etc/cnc/cnc.json
{
  	"CNCAddress":      "0.0.0.0",
  	"CNCPort":         443,
  	"DBAddress":       "127.0.0.1",
  	"DBName":          "cnc",
  	"DBPort":          "3306",
  	"DBUsername":      "username",
  	"DBPassword":      "password",
  	"DBLogMode":       "true",
  	"DBMaxConnection": 10240,
  	"Storage":         "./BOTDATA",
  	"UseSSL":          true,
  	"SSLCert":         "/etc/ssl/server.crt",
  	"SSLKey":          "/etc/ssl/server.key"
}
EOF'
```

# Donation
Ethereum Wallet: 0x1abfea1c8f22da9c11eb1e6293bdc0a061c70694
