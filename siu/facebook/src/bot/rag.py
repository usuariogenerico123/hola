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



# def verify_collection(name_collection:str, name_embedding:str, dir_db:str):
#     db = init_db(name_collection, name_embedding, dir_db)
#     print(db.)

def add_to_vectordb(text_archive:str, collection_name:str, name_embedding:object, name_db_directory:str):

    texto = read_file(text_archive)
    chunks = chunk_text(texto)
    db = init_db(collectionName=collection_name, name_embedding=name_embedding, dir_db=name_db_directory)

    db.add_texts(chunks)

    print(f"Proceso completado  coleccion: {collection_name} texto: {text_archive} directorioDb: {name_db_directory}")



collection = "paraiso_resort"
directory="./db_chroma"
embedding = GoogleGenerativeAIEmbeddings(model="models/gemini-embedding-001")
archive = "./documents/resortAzure.txt"


add_to_vectordb(archive, collection, embedding, directory)