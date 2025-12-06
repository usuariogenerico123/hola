from flask import Flask, request 
from src.messages import sendMessage
from src.bot.agent import Agent
from src.bot.tools import retrieve_document
import os

app = Flask(__name__)



@app.route("/", methods=["GET"])
def get():
    args= request.args
    hub=request.args.get("hub.challenge")
    return hub, 200



TOOLS=[retrieve_document]
MODEL="google_genai:gemini-2.5-flash-lite"
SYSTEM_PROMPT=(
        "Eres un asistente amable de Paraiso azure resort un servicio de alojamientos y vacaciones",
        "Tienes una base de conocimientos, utiliza las herramientas proporcionadas"
        "No responderas preguntas que no sean relacionadas a la informacion, se amable al decirles que estas limitado"
        "No digas que fuiste creado por google eres un asistente de paraiso azure resort"
    )

agent = Agent(MODEL, SYSTEM_PROMPT, TOOLS)
    
@app.route("/", methods=["POST"])
def post():
    data=request.get_json()
    id=data["entry"][0]["messaging"][0]["sender"]["id"]
    message=data["entry"][0]["messaging"][0]["message"]["text"]
    
    #sendMessage(message, id)
    print(agent)
    resp = agent.chat(message, id)
    sendMessage(resp, id)
    print(id + " "+ message)
    return "Hello, World! POST", 200



if __name__ == "__main__":
    app.run(port=3003, debug=True)