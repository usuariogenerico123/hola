from langchain_text_splitters import RecursiveCharacterTextSplitter 
from langchain_google_genai import GoogleGenerativeAIEmbeddings
from langchain_chroma import Chroma
import os



if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = userdata.get("GOOGLE_API_KEY")


def read_file(file) -> str:
    with open(file, "r") as f:
        print("Archivo cargado")
        return f.read()


collection = "paraiso_resort"
directory="./db_chroma"
embedding = GoogleGenerativeAIEmbeddings(model="models/gemini-embedding-001")



def init_db(collectionName:str, name_embedding: object, dir_db:str):
    
    db = Chroma(
        collection_name=collectionName,
        embedding_function=name_embedding,
        persist_directory=dir_db
    )
    print("Db iniciada..")
    return db 


def chunk_text(data, chunksize=1000, chunkoverlap=100):

    doc = RecursiveCharacterTextSplitter(
        chunk_size=chunksize,
        chunk_overlap=chunkoverlap,
        add_start_index=True
    )
    print("Chunks creados..")
    return doc.split_text(data)


texto = read_file("./documents/resortAzure.txt")
chunks = chunk_text(texto)
db = init_db(collectionName=collection, name_embedding=embedding, dir_db=directory)