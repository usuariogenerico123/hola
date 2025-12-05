from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain.chat_models import init_chat_model
from langchain_google_genai import GoogleGenerativeAIEmbeddings
from langchain_chroma import Chroma
from langchain.tools import tool
from langchain.agents import create_agent
import os



if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = userdata.get("GOOGLE_API_KEY")

def load_text(text: str):
    with open(text, "r") as f:
        return f.read()



texto = load_text("./src/rag/document/ResortAzure.txt")
doc = RecursiveCharacterTextSplitter(
    chunk_size=1000,
    chunk_overlap=200,
    add_start_index=True
)
s=doc.split_text(texto)



model = init_chat_model("google_genai:gemini-2.5-flash-lite")


embeddings = GoogleGenerativeAIEmbeddings(model="models/gemini-embedding-001")


db = Chroma(
    collection_name="resortazure",
    embedding_function=embeddings,
    persist_directory="./chroma_db",
)

# siu = db.add_texts(s)
# print(siu)
# result = db.similarity_search(
#     "que es resort azure", 
#     k=2
#     )
# print(result)



#---------------------------- AGENT-------------

#@tool(response_format="content_and_artifact")
@tool
def mi_funcion(query: str) -> str:

    """Retrieve information to help answer a query."""

    resp = db.similarity_search(
        query,
        k=1
    )
    re = "\n".join([i.page_content for i in resp])
    return re


tools = [mi_funcion]
system_prompt = (
    "Eres amable y servicial y siempre trata bien al cliente"
    "Eres un asistente útil que ayuda a responder preguntas utilizando "
    "una base de conocimientos. Utiliza las herramientas proporcionadas "
    "debes ser preciso no mas de dos lineas de parrafo"
    "Solo responde cosas relacionadas a la base de conocimientos de la herramienta, si te preguntan otra cosa di que no puedes ayudar con eso"
)


pregunta = "Hola quien es messi siu"
agent = create_agent(model, tools, system_prompt=system_prompt)

for i in agent.stream(
    {"messages": [{"role": "user", "content": pregunta}]},
    stream_mode = "values"
    ):
    i["messages"][-1].pretty_print()












