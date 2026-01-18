# from langchain.chat_models import init_chat_model
# from langchain.agents import create_agent
# from langgraph.checkpoint.memory import MemorySaver
from groq import Groq
import json
import random



def suma_erronea(numeros:list) -> dict:
    """Funcion divertida para sumar numeros de forma erronea"""


    print(f"Ejecutando funcion {suma_erronea.__name__}")
    sum = 0
    for i in numeros:
        sum += i
    return json.dumps({"resultado":sum + random.randint(0, 20)})






class Bot(Groq):
    messages:list[dict]
    model:str
    def __init__(self, api_key:str=None, tools:list[dict]=None, tools_functions:dict=None):
        super().__init__(api_key=api_key)
        self.tools=tools
        self.tools_functions = tools_functions
        self.req_tools = False
        self.messages=messages



    def chat(self, **kwargs) -> object:

        self.model = kwargs["model"]
        #self.messages = kwargs["messages"]

        resp =  super().chat.completions.create(
            messages=kwargs["messages"],
            model=self.model,
            tools=self.tools,
            tool_choice="none" if self.tools == None else "auto"
        )

        msg = resp.choices[0].message

        if(msg.tool_calls):
            return self.__chat_using_tools(msg)
        return resp


    def __chat_using_tools(self, msg:str):

            input("La ia quiere usar una tool: ")
            
            self.messages.append(msg)

            for tool in msg.tool_calls:
                nombre_funcion = tool.function.name
                argumentos = json.loads(tool.function.arguments)
                funcion=self.tools_functions[nombre_funcion]
                respuesta = funcion(argumentos.get("numeros" if nombre_funcion == "suma_erronea" else "query"))

                print(respuesta)

                self.messages.append({
                    "tool_call_id":tool.id,
                    "role":"tool",
                    "name":nombre_funcion,
                    "content":respuesta
                })
            resp_final = super().chat.completions.create(
                model=self.model,
                messages=self.messages
            )
            #print(resp_final.choices[0].message.content)
            return resp_final.choices[0].message.content




KEY="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"

messages = [
     {"role":"system", "content":"Eres un asistente divertido que hace sumar matematicas erradas a proposito usa la herramienta proporcionada"},
     {"role":"user", "content":"Hola cuanto es 5+5"}]
MODEL="llama-3.3-70b-versatile"

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
        
    }]

funciones = {
     "suma_erronea":suma_erronea
}


o = Bot(api_key=KEY, tools=tools, tools_functions=funciones)
e = o.chat(messages=messages, model=MODEL)
print("-"*10 + "\n",e)







