from flask import Flask, request 
from src.messages import sendMessage
from src.bot.agent import Agent
from src.bot.tools import retrieve_document
import os

app = Flask(__name__)




# @app.route("/", methods=["GET"])
# def get():
    
#     hub = ""
#     if(request.args):
#         hub=request.args.get("hub.challenge")
#         return hub, 200

#     return "siu", 200


TOOLS=[retrieve_document]
MODEL="google_genai:gemini-2.5-flash-lite"
SYSTEM_PROMPT=(
        "Eres un asistente amable de azure resort",
        "Eres un asistente útil que ayuda a responder preguntas utilizando "
        "una base de conocimientos. Utiliza las herramientas proporcionadas "
        "Usa la herramienta para ayudarte con preguntas y respuestas"
    )

agent = Agent(MODEL, SYSTEM_PROMPT, TOOLS)
def siu():

    while True:
        o = input(":  ")
        if(o == "salir"): break 

        re = agent.chat(o, "1")
        print(re)
siu()
# @app.route("/", methods=["POST"])
# def post():
#     data=request.get_json()
#     id=data["entry"][0]["messaging"][0]["sender"]["id"]
#     message=data["entry"][0]["messaging"][0]["message"]["text"]
    
#     #sendMessage(message, id)
#     print(agent)
#     resp = agent.chat(message, id)
#     sendMessage(resp, id)
#     print(id + " "+ message)
#     return "Hello, World! POST", 200



# if __name__ == "__main__":
#     app.run(port=3003, debug=True)