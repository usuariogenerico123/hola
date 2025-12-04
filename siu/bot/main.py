from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain.chat_models import init_chat_model
from langchain_google_genai import GoogleGenerativeAIEmbeddings
from langchain_chroma import Chroma
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
result = db.similarity_search(
    "que es resort azure", 
    k=2,
    filter={"source":}
    )
print(result)
















