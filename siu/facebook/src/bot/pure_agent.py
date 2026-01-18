from groq import Groq
import time
import json
import random
from tools import retrieve_document



KEY="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"



client = Groq(
    api_key=KEY
)



def suma_erronea(numeros:list) -> dict:
    """Funcion divertida para sumar numeros de forma erronea"""


    print(f"Ejecutando funcion {suma_erronea.__name__}")
    sum = 0
    for i in numeros:
        sum += i
    return json.dumps({"resultado":sum + random.randint(0, 20)})


tools=[
    {
        "type":"function",
        "function":{
            "name":"suma_erronea",
            "description":"Recibe una lista de numeros",
            "parameters":{
                "type":"object",
                "properties":{
                    "numeros":{"type":"array", "description":"Numeros,  ejemplo: [2,4,5,1,4]"}
                },
                "required":["numeros"]
            }

        }
        
    },
    {
        "type":"function",
        "function":{
            "name":"retrieve_document",
            "description": "Recibe una pregunta ",
            "parameters":{
                "type":"object",
                "properties":{
                    "query":{"type":"string","description":"Un texto en formato string"}
                },
                "required":["query"]
            }
        }
    }
]


MODEL="llama-3.3-70b-versatile"

def init(message):
    
    messages = [
        {"role":"system", "content":"Eres un asistente de AzureResort que responde preguntas frecuentes, usa la herramienta para tener un contexto de lo que debes responder"},
        {
            "role":"user",
            "content":message
        }
    ]

    
    chat = client.chat.completions.create(
    messages=messages,
    model=MODEL,
    tools=tools,   
    tool_choice="auto"
    )
    
    
    #mappear funciones 
    funciones = {
        "suma_erronea":suma_erronea,
        "retrieve_document": retrieve_document
    }
   

    message = chat.choices[0].message
    req_tools = message.tool_calls


    if(req_tools):
        input("La ia quiere usar una tool: ")
        
        messages.append(message)

        for tool in req_tools:
            nombre_funcion = tool.function.name
            argumentos = json.loads(tool.function.arguments)
            funcion=funciones[nombre_funcion]
            respuesta = funcion(argumentos.get("numeros" if nombre_funcion == "suma_erronea" else "query"))

            print(respuesta)

            messages.append({
                "tool_call_id":tool.id,
                "role":"tool",
                "name":nombre_funcion,
                "content":respuesta
            })
        resp_final = client.chat.completions.create(
            model=MODEL,
            messages=messages
        )
        print(resp_final.choices[0].message.content)
        return True
    print(message.content)
            

preg = input("Pregunta al bot: ")
init(preg)

# while True:
#     o = input("Mensaje: ")
#     if(o == 0):break;print("Adios")
    
#     message = init(o).choices[0].message
#     time.sleep(2)
#     if(message.tool_calls):

    


