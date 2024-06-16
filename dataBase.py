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

    def search(self, connection, table):
        cursor = connection.cursor()
        cursor.execute(f'SELECT * FROM {table}')
        results = cursor.fetchall()
        cursor.close()
        return results

if __name__ == "__main__":
    dbInstance = dataBase('localhost', 'teste', 'postgres', 'admin')
    connection = dbInstance.startConnection()
    pessoas = dbInstance.search(connection, 'pessoa')
    print(pessoas)
    dbInstance.closeConnection(connection)
