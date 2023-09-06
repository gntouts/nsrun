# execns

`execns` is a simple CLI utility used to run commands inside a specific process's network namespace.
It can be useful when debugging processes with unnamed network namespaces (eg containers).

## Build from source

To build from source and install `execns`, you can follow the instructions:

```bash
make
sudo make install
```

## Usage

`execns` usage is really simple:

```text
NAME:
        execns - execute command in given process's network namespace
                         
 
 USAGE:
        execns [PID] [COMMAND]
```

For example:

```bash
gntouts@ktb:~$ docker run -d --rm  ubuntu:latest sleep infinity
e7551fa8f7ab955b167d8728d5f0a7e2b8234c2bfff3465aa8f7d2b9ff96a96d

gntouts@ktb:~$ ps -ef | grep sleep
root      353396  353375  0 20:14 ?        00:00:00 sleep infinity

gntouts@ktb:~$ sudo execns 353396 ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
11: eth0@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```