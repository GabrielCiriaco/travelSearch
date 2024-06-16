import requests
from time import sleep

class Api:
    def __init__(self):
        self.urlToken = 'https://test.api.amadeus.com/v1/security/oauth2/token'
        self.urlFlights = f'https://test.api.amadeus.com/v1/shopping/flight-destinations'
        self.client_id = 'AjdPpr0yEaASWSDXrH7RT3SdTgFKSJWM'
        self.client_secret = 'AoHpAq4x8OPmwADd'



    def obterToken(self):
        headers = {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
        dados = {
            'grant_type': 'client_credentials',
            'client_id': self.client_id,
            'client_secret': self.client_secret
        }

        resposta = requests.post(self.urlToken, headers=headers, data=dados)
        dados_resposta = resposta.json()
        return dados_resposta['access_token']


    def obter_destinos_voo(self, token_acesso, origem, preco_maximo):
        
        headers = {
            'Authorization': f'Bearer {token_acesso}'
        }
        params = {
            'origin': origem,
            'maxPrice': preco_maximo
        }
        resposta = requests.get(self.urlFlights, headers=headers, params=params)
        return resposta.json()

# Função principal
if __name__ == '__main__':
    apiInstance= Api()
    token_acesso = apiInstance.obterToken()
    
    # Definir parâmetros de consulta
    origem = 'PAR'
    preco_maximo = 200
    
    # Obter destinos de voo
    destinos_voo = apiInstance.obter_destinos_voo(token_acesso, origem, preco_maximo)
    
    # Imprimir os resultados
    print(destinos_voo)