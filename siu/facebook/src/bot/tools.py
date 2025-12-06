from langchain.tools import tool

from .rag import init_db
from .RagData import collection, embedding, directory

db = init_db(collection, embedding, directory)
@tool
def retrieve_document(dbv:object, query:str)->str:

    """Retrieve information to help answer a query."""
    
    print("iniciando retrieve")
    print(collection)
    print(directory)
    print(db)
    print(query)

    resp = db.similarity_search(query=query, k=2)
   
    answer = "\n".join([i.page_content for i in resp])
    
    return answer

# e=retrieve_document("Que servicios ofrece la empresa")
# print(e)