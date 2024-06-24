package main

import (
    "encoding/json"
    "fmt"
    "net"
    "net/http"
    "os"
    "strconv"

    "github.com/FelipeFAlves/viagem/contact"
    "google.golang.org/protobuf/proto"
)

type TravelRequestData struct {
    Origem            string `json:"origem"`
    Destino           string `json:"destino"`
    Cidade            string `json:"cidade"`
    DataIda           string `json:"dataIda"`
    DataVolta         string `json:"dataVolta"`
    QuantidadeAdultos int    `json:"quantidadeAdultos"`
}

func consultaHandler(w http.ResponseWriter, r *http.Request) {
    var genericData map[string]interface{}

    err := json.NewDecoder(r.Body).Decode(&genericData)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Convertendo e validando `quantidadeAdultos`
    quantidadeAdultosStr, ok := genericData["quantidadeAdultos"].(string)
    if !ok {
        http.Error(w, "Invalid quantidadeAdultos", http.StatusBadRequest)
        fmt.Println("Invalid quantidadeAdultos")
        return
    }

    quantidadeAdultos, err := strconv.Atoi(quantidadeAdultosStr)
    if err != nil {
        http.Error(w, "Invalid quantidadeAdultos value", http.StatusBadRequest)
        fmt.Println("Error converting quantidadeAdultos:", err)
        return
    }

    requestData := TravelRequestData{
        Origem:            genericData["origem"].(string),
        Destino:           genericData["destino"].(string),
        Cidade:            genericData["cidade"].(string),
        DataIda:           genericData["dataIda"].(string),
        DataVolta:         genericData["dataVolta"].(string),
        QuantidadeAdultos: quantidadeAdultos,
    }

    // Log de confirmação dos dados recebidos
    fmt.Println("Dados recebidos do front-end:")
    fmt.Printf("Origem: %s, Destino: %s, Cidade: %s, DataIda: %s, DataVolta: %s, QuantidadeAdultos: %d\n",
        requestData.Origem, requestData.Destino, requestData.Cidade, requestData.DataIda, requestData.DataVolta, requestData.QuantidadeAdultos)

    serverAddr := "192.168.182.26:8080"
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        http.Error(w, "Failed to connect to server", http.StatusInternalServerError)
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    request := &contact.TravelRequest{
        Origem:            requestData.Origem,
        Destino:           requestData.Destino,
        Cidade:            requestData.Cidade,
        DataIda:           requestData.DataIda,
        DataVolta:         requestData.DataVolta,
        QuantidadeAdultos: int32(requestData.QuantidadeAdultos),
    }

    data, err := proto.Marshal(request)
    if err != nil {
        http.Error(w, "Failed to serialize request", http.StatusInternalServerError)
        fmt.Println("Error serializing request:", err)
        return
    }

    _, err = conn.Write(data)
    if err != nil {
        http.Error(w, "Failed to send request to server", http.StatusInternalServerError)
        fmt.Println("Error sending data:", err)
        return
    }

    modifiedData := make([]byte, 65536)

    n, err := conn.Read(modifiedData)
    if err != nil {
        http.Error(w, "Failed to receive response from server", http.StatusInternalServerError)
        fmt.Println("Error reading data:", err)
        return
    }

    response := &contact.TravelResponse{}
    err = proto.Unmarshal(modifiedData[:n], response)
    if err != nil {
        http.Error(w, "Failed to deserialize response", http.StatusInternalServerError)
        fmt.Println("Error deserializing response:", err)
        return
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
        fmt.Println("Error serializing response:", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func main() {
    http.HandleFunc("/consulta", consultaHandler)
    http.Handle("/", http.FileServer(http.Dir(".")))

    fmt.Println("Servidor iniciado na porta 8080")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
        os.Exit(1)
    }
}
