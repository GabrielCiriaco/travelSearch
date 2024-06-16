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
    
    

if __name__ == "__main__":
    server = server("localhost", 5000)
    server.initServer()

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