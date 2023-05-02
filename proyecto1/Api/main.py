from fastapi import FastAPI
import mysql.connector
from pydantic import BaseModel
import pandas as pd
import csv

app = FastAPI()
app.title = "SOPES Proyecto 2 Api"

config = {
    'user': 'elmoco13',
    'password': 'elmoco13',
    'host':'34.134.48.99',
    'database':'sopestarea2',
    'raise_on_warnings':True,
}

class Votes(BaseModel):
    ##sede,municipio,departamento,papeleta,partido
    sede:int
    municipio:str
    departamento:str
    papeleta:str
    partido:str

cnx = mysql.connector.connect(**config)

@app.post("/add")
async def create(item:Votes):
    insert_sql = f"INSERT INTO "