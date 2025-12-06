from langchain.chat_models import init_chat_model
from langchain.agents import create_agent
from langgraph.checkpoint.memory import MemorySaver

import os

if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = userdata.get("GOOGLE_API_KEY")

class Agent:
    model:str
    tools:list
    agent:str
    def __init__(self, model, system_prompt=None, tools :list =None):
        self.model = init_chat_model(model)
        self.model=model 
        self.system_prompt = system_prompt
        self.tools = tools
        self.checkpointer = MemorySaver()
        self.agent = self.__create_agent()
        


    def __create_agent(self):
        agent_ia = create_agent(
            model=self.model,
            tools=self.tools,
            checkpointer=self.checkpointer,
            system_prompt = self.system_prompt
        )

        return agent_ia
        
    def chat(self, text:str,  id:str):
        print("iniciandp chat")
        
        config = {"configurable":{"thread_id":id}}

        resp = ""

        for i in self.agent.stream({"messages":[{"role":"user", "content":text}]}, config, stream_mode="values"):
            resp = i["messages"][-1].content

        #print(resp)
        return resp
        

    def __str__(self):
        return f"Agente {self.model} iniciado con exito"


