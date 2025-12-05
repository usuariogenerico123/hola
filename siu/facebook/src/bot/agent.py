from langchain.chat_models import init_chat_model
from langchain.agents import create_agent
import os

if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = userdata.get("GOOGLE_API_KEY")

class Agent:
    __MODEL:str
    __SYSTEM_PROMPT:str
    __tools:list

    def __init__(self, model, _system_prompt=None, _tools :list =None):
        self.model = init_chat_model(model)
        self.__MODEL=model 
        self.__SYSTEM_PROMPT = _system_prompt
        self.__tools = _tools
        self.agent = self.create_agent()


    def create_agent(self):
        agent_ia = create_agent(
            model=self.model,
            tools=self.__tools,
            system_prompt = self.__SYSTEM_PROMPT
        )

        return agent_ia
        

    def 


    def __str__(self):
        return f"Agente {self.__MODEL} iniciado con exito"


