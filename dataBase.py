import psycopg2

class dataBase:
    def __init__(self, host, database, user, password):
        self.host = host
        self.database = database
        self.user = user
        self.password = password

    def startConnection(self):
        connection = psycopg2.connect(
            host=self.host,
            database=self.database,
            user=self.user,
            password=self.password)
        return connection

    def closeConnection(self, connection):
        if connection:
            connection.close()
    
    def insertVoo(self, connection, id_voo, id_pesquisa, horario_partida, horario_chegada, preco, detalhes):
        cursor = connection.cursor()
        cursor.execute(f"INSERT INTO voos (id_voo, id_pesquisa, horario_partida, horario_chegada, preco, detalhes) VALUES ({id_voo}, {id_pesquisa}, '{horario_partida}', '{horario_chegada}', '{preco}', '{detalhes}')")
        connection.commit()
        cursor.close()

    def insertReservaHotel(self, connection, id_reserva, id_pesquisa, nome_hotel, cidade, endereco, preco_total, descricao_quarto, checkin_date, checkout_date, quantidade_adultos, categoria_quarto):
        cursor = connection.cursor()
        cursor.execute(f"INSERT INTO reserva_hotel (id_reserva, id_pesquisa, nome_hotel, cidade, endereco, preco_total, descricao_quarto, checkin_date, checkout_date, quantidade_adultos, categoria_quarto) VALUES ({id_reserva}, {id_pesquisa}, '{nome_hotel}', '{cidade}', '{endereco}', {preco_total}, '{descricao_quarto}', '{checkin_date}', '{checkout_date}', {quantidade_adultos}, '{categoria_quarto}')")
        connection.commit()
        cursor.close()
    
    def insertPesquisaDePacote(self, connection, id_pesquisa, cidade, checkin_date, checkout_date, adultos, voo_origem, voo_destino, data_pesquisa):
        cursor = connection.cursor()
        cursor.execute(f"INSERT INTO pesquisa_de_pacote (id_pesquisa, cidade, checkin_date, checkout_date, adultos, voo_origem, voo_destino, data_pesquisa) VALUES ({id_pesquisa}, '{cidade}', '{checkin_date}', '{checkout_date}', {adultos}, '{voo_origem}', '{voo_destino}', '{data_pesquisa}')")
        connection.commit()
        cursor.close()

    def pesquisaVoo(self, connection, dataLimite, cidade_origem, cidade_destino, data_ida, data_volta):
        cursor = connection.cursor()
        query = '''
        SELECT v.* FROM pesquisa_de_pacote p
        JOIN voos v ON v.id_pesquisa = p.id_pesquisa
        WHERE p.voo_origem = %s
        AND p.voo_destino = %s
        AND p.checkin_date = %s
        AND p.checkout_date = %s
        AND p.data_pesquisa > %s
        '''
        cursor.execute(query, (cidade_origem, cidade_destino, data_ida, data_volta, dataLimite))
        results = cursor.fetchall()
        cursor.close()
        return results
    
    def pesquisaHotel(self, connection, dataLimite, cidade, data_ida, data_volta, adultos):
        cursor = connection.cursor()
        query = '''
        SELECT r.* FROM pesquisa_de_pacote p
        JOIN reserva_hotel r ON r.id_pesquisa = p.id_pesquisa
        WHERE p.cidade = %s
        AND p.checkin_date = %s
        AND p.checkout_date = %s
        AND p.adultos = %s
        AND p.data_pesquisa > %s
        '''
        cursor.execute(query, (cidade, data_ida, data_volta, adultos, dataLimite))
        results = cursor.fetchall()
        cursor.close()
        return results
    
    def mapeiaVoos(self, voos):
        def format_voo_tuple(voo_tuple):
            id_voo, id_pesquisa, horario_partida, horario_chegada, preco, detalhes = voo_tuple
            return {
                "horario_partida": f"Horário de partida: {horario_partida}.",
                "horario_chegada": f"Horário de chegada: {horario_chegada}.",
                "preco": preco,
                "detalhes": detalhes
            }
        return [format_voo_tuple(voo) for voo in voos]

    def mapeiaHoteis(self, hoteis):
        def format_hotel_tuple(hotel_tuple):
            id_reserva, id_pesquisa, nome_hotel, cidade, endereco, preco_total, descricao_quarto, checkin_date, checkout_date, quantidade_adultos, categoria_quarto = hotel_tuple
            return {
                "nome_hotel": nome_hotel,
                "cidade": cidade,
                "endereco": endereco,
                "preco_total": float(preco_total),
                "descricao_quarto": descricao_quarto,
                "checkin_date": str(checkin_date),
                "checkout_date": str(checkout_date),
                "quantidade_adultos": quantidade_adultos,
                "categoria_quarto": categoria_quarto
            }
        return [format_hotel_tuple(hotel) for hotel in hoteis]

    
    
    

if __name__ == "__main__":
    dbInstance = dataBase('localhost', 'teste', 'postgres', 'admin')
    connection = dbInstance.startConnection()
    
 
    # dbInstance.insertPesquisaDePacote(connection, 4, 'LON', '2024-07-01', '2024-07-05', 2, 'Aeroporto Internacional de Guarulhos', 'Aeroporto Santos Dumont', '2024-07-01')
    # dbInstance.insertVoo(connection, 4, 4, '07:35', '08:30', '2411 Reais brasileiros', 'Preço total a partir de 2880 Reais brasileiros, ida e volta. Voo sem escalas da Azul. Sai do aeroporto Aeroporto Internacional de Guarulhos às 07:35 do dia segunda-feira, junho 17 e chega no aeroporto Aeroporto Santos Dumont às 08:30 do dia segunda-feira, junho 17. Duração total: 55 min.')
    # dbInstance.insertReservaHotel(connection, 4, 4, 'TEST CONTENT', 'LON', 'TEST CONTENT, LON', 5319.00, 'Prepay Non-refundable, Non changeable, includes WiFi\nSup Rm Top floor, 40 sqm king bed, marble\nbathroom, elegant seating area, Floor to', '2024-07-01', '2024-07-05', 2, 'SUPERIOR_ROOM')

    dataLimite = '2024-06-1'
    origem = 'Aeroporto Internacional de Guarulhos'
    destino = 'Aeroporto Santos Dumont'
    cidade = 'LON'
    data_ida = '2024-07-01'
    data_volta = '2024-07-05'
    adultos = 2

    resultadosVoos = dbInstance.pesquisaVoo(connection, dataLimite, origem, destino, data_ida, data_volta)
    resultadosHoteis = dbInstance.pesquisaHotel(connection, dataLimite, cidade, data_ida, data_volta, adultos)
    
    print(resultadosVoos)
    # print(resultadosHoteis)