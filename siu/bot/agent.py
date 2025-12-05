from langchain.agents import create_agent
from tools.tools import crea_archivo, elimina_archivo
import os  
from langchain.chat_models import init_chat_model


if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = userdata.get("GOOGLE_API_KEY")


gemini = init_chat_model("google_genai:gemini-2.5-flash-lite")

tool = [crea_archivo, elimina_archivo]

systemprompt=(
    "Eres un asistente útil que crea archivos "
    "Utiliza las herramientas proporcionadas "
    "No hagas otras cosas y si te preguntan cosas que no vienen al tema, responde que esa no es tu labor"
)

agent = create_agent(
    model=gemini,
    tools=tool,
    system_prompt=systemprompt

)


#pregunta = "Elimina mi archivo documentos "

while True:
    print("Escribe el prompt")
    pregunta = input(": ")
    if(pregunta == "Salir" ):
        print("Adios")
        break
    else:

        for i in agent.stream({"messages": [{"role":"user", "content":pregunta }]}, stream_mode="values"):
            i["messages"][-1].pretty_print()