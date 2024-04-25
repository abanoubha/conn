# conn

list all WiFi connected devices

## branches

- `tui` : terminal UI
- `gtk` : gtk4 GUI
- `gio` : Gio GUI
- `flutter` : Flutter GUI

## Idea

- get ip address : `ip addr`
- list connected devices : `nmap -sn 192.168.1.0/24`

output of nmap :

```plain
Starting Nmap 7.92 ( https://nmap.org ) at 2022-01-07 01:14 EET
Nmap scan report for _gateway (192.168.1.1)
Host is up (0.011s latency).
Nmap scan report for 192.168.1.4
Host is up (0.054s latency).
Nmap scan report for 192.168.1.6
Host is up (0.11s latency).
Nmap scan report for 192.168.1.7
Host is up (0.083s latency).
Nmap scan report for ubuntu (192.168.1.8)
Host is up (0.000085s latency).
Nmap scan report for 192.168.1.10
Host is up (0.014s latency).
Nmap done: 256 IP addresses (6 hosts up) scanned in 8.08 seconds
```

## commands

```sh
go mod tidy && go build -o conn main.go && ./conn
```
