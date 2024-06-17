from datetime import datetime, timedelta

import socket
import travelSearch_pb2
from crawler import scrap
from api import Api
from dataBase import dataBase

class server:
    def __init__(self, host, port):
        self.host = host
        self.port = port


    def initServer(self):

        serverSocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM) #definindo ipv4 e tcp
        serverSocket.bind((self.host, self.port)) 
        serverSocket.listen(5)
        
        print("Servidor escutando em {}:{}".format(self.host, self.port))
        
        while True:
            # Aceitar conexão de um cliente
            client_socket, client_address = serverSocket.accept()
            print("Conexão estabelecida com:", client_address)
            
            # Lidar com as mensagens recebidas do cliente
            self.handleClient(client_socket)
            
            # Fechar o socket do cliente
            client_socket.close()


    def handleClient(self, client_socket):
        # Receber os dados do cliente
        data = client_socket.recv(1024)
    
        # Verificar se os dados foram recebidos corretamente
        if not data:
            return
        
        requestInstance = travelSearch_pb2.Request()  

        requestInstance.ParseFromString(data)

        request = {
            "origem": requestInstance.origem,
            "destino": requestInstance.destino,
            "cidade": requestInstance.cidade,
            "data_ida": requestInstance.data_ida,
            "data_volta": requestInstance.data_volta,
            "quantidade_adultos": requestInstance.quantidade_adultos,
            "dataLimite": (datetime.today() - timedelta(days=15)).strftime('%Y-%m-%d')
        }

        # Inicializar o banco de dados
        dbInstance = dataBase('localhost', 'teste', 'postgres', 'admin')
        connection = dbInstance.startConnection()

        resultadosVoos = dbInstance.pesquisaVoo(connection, request.dataLimite, request.origem, request.destino, request.data_ida, request.data_volta)
        resultadosHoteis = dbInstance.pesquisaHotel(connection, request.dataLimite, request.cidade, request.data_ida, request.data_volta, request.adultos)

        
        
        if(resultadosVoos != []):
            resultadosVoos = dbInstance.mapeiaVoos(resultadosVoos)
        if(resultadosHoteis != []):
            resultadosHoteis = dbInstance.mapeiaHoteis(resultadosHoteis)
            

        if(resultadosVoos == []):
            # Inicializar o crawler
            scrapInstance = scrap(request.origem, request.destino, request.data_ida, request.data_volta)
            resultadosVoos = scrapInstance.buscar_passagens()
            dbInstance.salvarVoos(connection, resultadosVoos, request.origem, request.destino, request.data_ida, request.data_volta)

        if(resultadosHoteis == []):
            # Inicializar a API
            apiInstance= Api()
            token_acesso = apiInstance.obterToken()
            resultadosHoteis = apiInstance.obter_ofertas_hoteis(token_acesso, request.cidade, request.quantidade_adultos, request.data_ida, request.data_volta)
            dbInstance.salvarHoteis(connection, resultadosHoteis, request.cidade, request.data_ida, request.data_volta, request.quantidade_adultos)
        
        dbInstance.closeConnection(connection)
        responseInstance = travelSearch_pb2.TravelResponse()

        for voo in resultadosVoos:
            vooInstance = responseInstance.voos.add()
            vooInstance.horario_partida = voo["horario_partida"]
            vooInstance.horario_chegada = voo["horario_chegada"]
            vooInstance.preco = voo["preco"]
            vooInstance.detalhes = voo["detalhes"]
        
        for hotel in resultadosHoteis:
            hotelInstance = responseInstance.hoteis.add()
            hotelInstance.nome_hotel = hotel["nome_hotel"]
            hotelInstance.cidade = hotel["cidade"]
            hotelInstance.endereco = hotel["endereco"]
            hotelInstance.preco_total = hotel["preco_total"]
            hotelInstance.descricao_quarto = hotel["descricao_quarto"]
            hotelInstance.checkin_date = hotel["checkin_date"]
            hotelInstance.checkout_date = hotel["checkout_date"]
            hotelInstance.quantidade_adultos = hotel["quantidade_adultos"]
            hotelInstance.categoria_quarto = hotel["categoria_quarto"]

        client_socket.send(responseInstance.SerializeToString())

        
        

    
    

if __name__ == "__main__":
    # server = server("localhost", 5000)
    # server.initServer()

    # Teste do crawler
    # scrapInstance = scrap("CGH", "GIG", "2024/06/16", "2024/07/15")
    # scrapInstance.buscar_passagens()

    # Teste da API
    # apiInstance= Api()
    # token_acesso = apiInstance.obterToken()
    # origem = 'PAR'
    # preco_maximo = 200
    # destinos_voo = apiInstance.obter_destinos_voo(token_acesso, origem, preco_maximo)
    # print(destinos_voo)

    # Teste do dataBase
    # dbInstance = dataBase('localhost', 'teste', 'postgres', 'admin')
    # connection = dbInstance.startConnection()
    # pessoas = dbInstance.search(connection, 'pessoa')
    # print(pessoas)
    # dbInstance.closeConnection(connection)

    request = {
            "origem": "São Paulo",
            "destino": "Rio de Janeiro",
            "cidade": 'LON',
            "data_ida": '2024-07-01',
            "data_volta": '2024-07-05',
            "quantidade_adultos": 2,
            "dataLimite": "2021-1-1"
        }

    # Inicializar o banco de dados
    dbInstance = dataBase('localhost', 'teste', 'postgres', 'admin')
    connection = dbInstance.startConnection()

    resultadosVoos = dbInstance.pesquisaVoo(connection, request['dataLimite'], request['origem'], request['destino'], request['data_ida'], request['data_volta'])
    resultadosHoteis = dbInstance.pesquisaHotel(connection, request['dataLimite'], request['cidade'], request['data_ida'], request['data_volta'], request['quantidade_adultos'])
   
    
    
    if(resultadosVoos != []):
        resultadosVoos = dbInstance.mapeiaVoos(resultadosVoos)
    if(resultadosHoteis != []):
        resultadosHoteis = dbInstance.mapeiaHoteis(resultadosHoteis)
        

    if(resultadosVoos == []):
        # Inicializar o crawler
        scrapInstance = scrap(request['origem'], request['destino'], request['data_ida'], request['data_volta'])
        resultadosVoos = scrapInstance.buscar_passagens()
        
    if(resultadosHoteis == []):
        # Inicializar a API
        apiInstance= Api()
        token_acesso = apiInstance.obterToken()
        resultadosHoteis = apiInstance.obter_ofertas_hoteis(token_acesso, request.cidade, request.quantidade_adultos, request.data_ida, request.data_volta)
        
    dbInstance.closeConnection(connection)
    
    print(resultadosVoos)   
    print(resultadosHoteis)
