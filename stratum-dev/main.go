package main

import (
    "log"
    "net"
    "pool_ltcdoge/server"
)

func main() {
    listener, err := net.Listen("tcp", ":3333")
    if err != nil {
        log.Fatalf("Erreur lors de l'écoute sur le port 3333: %v", err)
    }
    defer listener.Close()
    log.Println("✅ Serveur Stratum TCP prêt sur le port 3333")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Erreur lors de l'acceptation de la connexion: %v", err)
            continue
        }
        go server.HandleConnection(conn)
    }
}