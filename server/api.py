import requests
from time import sleep

class Api:
    def __init__(self):
        self.urlToken = 'https://test.api.amadeus.com/v1/security/oauth2/token'
        self.urlFlights = f'https://test.api.amadeus.com/v1/shopping/flight-destinations'
        self.urlHoteisCidade = 'https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city'
        self.urlHoteisOfertas = 'https://test.api.amadeus.com/v3/shopping/hotel-offers'
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
    
    
    def obter_hoteis_por_cidade(self, token_acesso, cidade):
        headers = {
            'Authorization': f'Bearer {token_acesso}'
        }
        params = {
            'cityCode': cidade,
        }

        resposta = requests.get(self.urlHoteisCidade, headers=headers, params=params)
        
        # idsHoteis = []

        # for hotel in resposta.json()['data']:
        #     idsHoteis.append(hotel['hotelId'])

        # return idsHoteis
        return resposta.json()['data']
    
    
    def obter_ofertas_hoteis(self, token_acesso, cidade, adults, checkin_date, checkout_date):
        Hoteis = self.obter_hoteis_por_cidade(token_acesso, cidade)
        
        idsHoteis = []

        for hotel in Hoteis:
            idsHoteis.append(hotel['hotelId'])


        headers = {
            'Authorization': f'Bearer {token_acesso}'
        }
        ids = []
        for i in range(30):
            ids.append( idsHoteis[i])
        
        params = {
            'hotelIds': ids,
            'adults': adults,
            'checkInDate': checkin_date,
            'checkOutDate': checkout_date,
        }

        resposta = requests.get(self.urlHoteisOfertas, headers=headers, params=params)
        
        if 'data' not in resposta.json():
            print("Não há ofertas disponíveis")

            primeirosHoteis = []
            for i in range(5):
                primeirosHoteis.append(Hoteis[i])

            apenasHoteis = []
            for hotel in primeirosHoteis:
                apenasHoteis.append({
                    'nome_hotel': hotel.get('name', ''),
                    'cidade': hotel.get('iataCode', ''),
                    'endereco': f"{hotel.get('name', '')}, {hotel.get('iataCode', '')}",
                    'preco_total': 1,
                    'descricao_quarto': 'Não há ofertas disponíveis',
                    'checkin_date': checkin_date,
                    'checkout_date': checkout_date,
                    'quantidade_adultos': adults,
                    'categoria_quarto': 'não há ofertas disponíveis',
                })
            return apenasHoteis
        
        ofertas = []
        
        for hotel in resposta.json()['data']:
            nome_hotel = hotel['hotel']['name']
            cidade = hotel['hotel']['cityCode']
            
            for oferta in hotel['offers']:
                endereco = f"{hotel['hotel']['name']}, {hotel['hotel']['cityCode']}"
                preco_total = oferta['price']['total']
                descricao_quarto = oferta['room']['description']['text']
                checkin = oferta['checkInDate']
                checkout = oferta['checkOutDate']
                quantidade_adultos = oferta['guests']['adults']
                categoria_quarto = str(oferta['room']['typeEstimated'])

                
                ofertas.append({
                    'nome_hotel': nome_hotel,
                    'cidade': cidade,
                    'endereco': endereco,
                    'preco_total': preco_total,
                    'descricao_quarto': descricao_quarto,
                    'checkin_date': checkin,
                    'checkout_date': checkout,
                    'quantidade_adultos': quantidade_adultos,
                    'categoria_quarto': categoria_quarto,

                })
        
        return ofertas

# Função principal
if __name__ == '__main__':
    apiInstance = Api()
    token_acesso = apiInstance.obterToken()

    # Definir parâmetros de consulta de voos
    # origem = 'LON'
    # preco_maximo = 200
    # destinos_voo = apiInstance.obter_destinos_voo(token_acesso, origem, preco_maximo)

    
    # Definir parâmetros de consulta de hotéis
    cidade = 'LON'
    checkin_date = '2024-07-01'
    checkout_date = '2024-07-05'
    adultos = 2

    # Obter ofertas de hotel
    
    ofertas_hotel = apiInstance.obter_ofertas_hoteis(token_acesso, cidade, adultos, checkin_date, checkout_date)
    