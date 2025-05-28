package server

import (
    "bufio"
    "encoding/json"
    "log"
    "net"
)

type StratumMessage struct {
    ID     int           `json:"id"`
    Method string        `json:"method"`
    Params []interface{} `json:"params"`
}

func HandleConnection(conn net.Conn) {
    defer conn.Close()
    client := &Client{Conn: conn}
    reader := bufio.NewReader(conn)

    for {
        line, err := reader.ReadBytes('\n')
        if err != nil {
            log.Printf("Erreur de lecture: %v", err)
            return
        }

        var msg StratumMessage
        if err := json.Unmarshal(line, &msg); err != nil {
            log.Printf("JSON invalide reçu: %s", line)
            continue
        }

        switch msg.Method {
        case "mining.subscribe":
            handleSubscribe(client, msg)
        case "mining.authorize":
            handleAuthorize(client, msg)
        default:
            log.Printf("Méthode non supportée: %s", msg.Method)
        }
    }
}

func handleSubscribe(client *Client, msg StratumMessage) {
    response := map[string]interface{}{
        "id":     msg.ID,
        "result": [3]interface{}{[2]string{"mining.set_difficulty", "1"}, "08000002", 4},
        "error":  nil,
    }
    sendJSON(client.Conn, response)
}

func handleAuthorize(client *Client, msg StratumMessage) {
    if len(msg.Params) < 1 {
        log.Println("Paramètres insuffisants pour mining.authorize")
        return
    }
    username, ok := msg.Params[0].(string)
    if !ok {
        log.Println("Nom d'utilisateur invalide")
        return
    }
    client.Username = username
    client.Authorized = true

    response := map[string]interface{}{
        "id":     msg.ID,
        "result": true,
        "error":  nil,
    }
    sendJSON(client.Conn, response)
}

func sendJSON(conn net.Conn, data interface{}) {
    bytes, err := json.Marshal(data)
    if err != nil {
        log.Printf("Erreur lors du marshalling JSON: %v", err)
        return
    }
    conn.Write(append(bytes, '\n'))
}