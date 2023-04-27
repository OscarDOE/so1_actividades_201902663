#Clase del 22/2/23, empiez lo bueno como a la hora exacta
#comando para correr locust =
# locust -f <ruta de archivo>
# Si la terminal se encuentra en la misma ruta, sería:
# locust -f <archivo>
import json
from random import randrange
from locust import HttpUser,task,between

class readFile():
    def __init__(self):
        self.data = []

    def getData(self):
        size = len(self.data)
        if size > 0:
            index = randrange(0, size - 1) if size > 1 else 0
            return self.data.pop(index)

    def loadFile(self):
        print("LOADING .....")
        try:
            with open("MOCK_DATA.json","r") as file:
                self.data = json.loads(file.read())
        except Exception:
            print(f'Error : {Exception}')

class ExampleTest(HttpUser):

    wait_time = between(0.1,0.9)
    reader = readFile()
    reader.loadFile()

    def on_start(self):
        print("On Start")

    @task
    def index(self):
        data = self.reader.getData()
        headers = {'Content-Type': 'application/json'}
        if data is not None:
            res = self.client.post("/add",headers=headers,data=json.dumps(data))

        else:
            print("Empty")
            self.stop(True)
        # data = {
        #     "name": "Doom",
        #     "gender": "Shooter",
        #     "multiplayer": 0
        # }
        # self.client.post("/add", headers=headers, data=json.dumps(data))
        #self.client.post("<endpoint>", <Datos a enviar, si es que se reciben en formato json, se pone en formato jason>{"llave":"valor"})




#from locust import HttpUser, between, task
#
#Ejemplo de la documentacion de Locust
# class WebsiteUser(HttpUser):
#     wait_time = between(5, 15)
    
#     def on_start(self):
#         self.client.post("/login", {
#             "username": "test_user",
#             "password": ""
#         })
    
#     @task
#     def index(self):
#         self.client.get("/")
#         self.client.get("/static/assets.js")
        
#     @task
#     def about(self):
#         self.client.get("/about/")