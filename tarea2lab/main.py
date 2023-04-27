import mysql.connector
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()
# uvicorn main:app --reload

config = {
    'user': 'elmoco13',
    'password': 'elmoco13',
    'host': '34.134.48.99',
    'database': 'sopestarea2',
    'raise_on_warnings': True,
}

class Item(BaseModel):
    name: str
    gender: str
    multiplayer: int


cnx = mysql.connector.connect(**config)

@app.post("/add")
async def create_games(item: Item):
    insert_sql = f"INSERT INTO games(name, gender, multiplayer) VALUES ('{item.name}','{item.gender}',{item.multiplayer})"
    # insert_sql = f"INSERT INTO games(name, gender, multiplayer) VALUES('Doom','Shooter',2) "
    cursor = cnx.cursor()
    cursor.execute(insert_sql)
    inserts = cursor.fetchall()
    cursor.close()
    cnx.commit()
    return {"item":inserts}


# @app.get("/")
# def get_main():
#     cursor = cnx.cursor()
#     query = "SELECT * FROM games"
#     cursor.execute(query)
#     users = cursor.fetchall()
#     cursor.close()
#     print("juegos",users)
#     return {"juegos":users}


# @app.get("/users")
# def get_users():
#     cursor = cnx.cursor()
#     query = "SELECT * FROM users"
#     cursor.execute(query)
#     users = cursor.fetchall()
#     cursor.close()
#     return {"users": users}

# @app.get("/dbtest")
# def test_db_connection():
#     cursor = cnx.cursor()
#     cursor.execute("SELECT VERSION()")
#     result = cursor.fetchone()
#     cursor.close()
#     return {"db_version": result[0]}

