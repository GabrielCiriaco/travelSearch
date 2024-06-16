from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium import webdriver
from time import sleep

class scrap:
    def __init__(self, cidade_origem, cidade_destino, data_ida, data_volta):
        self.cidade_origem = cidade_origem
        self.cidade_destino = cidade_destino
        self.data_ida = data_ida
        self.data_volta = data_volta
        
    def encontrar_elementos(self, navegador, seletor_css):
        return navegador.find_elements(By.CSS_SELECTOR, seletor_css)

    def preencher_campo_texto(self, elemento, texto):
        elemento.send_keys(Keys.CONTROL, 'A') 
        elemento.send_keys(texto)  
        sleep(1)

    def clicar_elemento(self, navegador, seletor_css):
        navegador.find_element(By.CSS_SELECTOR, seletor_css).click()
        sleep(1)  
    
    def aguardar_elemento(self, navegador, seletor_css):
        while len(navegador.find_elements(By.CSS_SELECTOR, seletor_css)) < 1:
            sleep(0.3) 

    def buscar_passagens(self):
        
        # Inicializa o Chrome WebDriver com opções padrão
        opcoes = webdriver.ChromeOptions()
        navegador = webdriver.Chrome(options=opcoes)
        
        # Abre a página do Google Flights
        navegador.get('https://www.google.com/flights?hl=pt&curr=BRL#flt=')

        # Aguarda até que os campos de texto estejam disponíveis
        self.aguardar_elemento(navegador, '[type="text"]')

        # Encontra todos os elementos do tipo texto
        campos_texto = self.encontrar_elementos(navegador, '[type="text"]')

        # Preenche a data de ida
        campo_data_ida = campos_texto[4]
        self.preencher_campo_texto(campo_data_ida, self.data_ida)

         # Preenche a cidade de origem
        campo_cidade_origem = campos_texto[0]
        self.preencher_campo_texto(campo_cidade_origem, self.cidade_origem)
        # Confirma a cidade de origem
        self.clicar_elemento(navegador, '[class="P1pPOe"]')  
        
        # Preenche a data de volta
        campo_data_volta = campos_texto[5]
        self.preencher_campo_texto(campo_data_volta, self.data_volta)

       

        # Preenche a cidade de destino
        campo_cidade_destino = campos_texto[2]
        self.preencher_campo_texto(campo_cidade_destino, self.cidade_destino)
        # Confirma a cidade de destino
        self.clicar_elemento(navegador, '[class="P1pPOe"]')  

        # Inicia a pesquisa clicando no botão de pesquisa
        self.clicar_elemento(navegador, '[aria-label="Pesquisar"]')

        # Aguarda até que os resultados apareçam
        self.aguardar_elemento(navegador, '[class="JMc5Xc"]')
        
        # Coleta informações sobre a melhor passagem
        melhor_voo_info = self.encontrar_elementos(navegador, '[class="JMc5Xc"]')[0]
        detalhes_melhor_voo = melhor_voo_info.get_attribute("aria-label")

        # Coleta informações adicionais sobre o preço da passagem
        # informacao_adicional = self.encontrar_elementos(navegador, '[class="frOi8 AdWm1c fVSoi"]')[0].text

        # Coleta informações sobre passagens mais baratas em períodos próximos
        # passagens_baratas_info = [elemento.text.split('\n') for elemento in self.encontrar_elementos(navegador, '[class="mrywM"]')]
        
        # Exibe as informações coletadas
        print(f'Melhor voo:\n')
        print(detalhes_melhor_voo, '\n')
        # print(informacao_adicional)
        # print(passagens_baratas_info)
        
        # Fecha o navegador
        navegador.quit()

if __name__ == "__main__":
    scrap = scrap("São Paulo", "Rio de Janeiro", "2024/06/16", "2024/07/15")
    scrap.buscar_passagens()