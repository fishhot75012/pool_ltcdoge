package server

import (
    "net"
)

type Client struct {
    Conn       net.Conn
    Username   string
    Authorized bool
}