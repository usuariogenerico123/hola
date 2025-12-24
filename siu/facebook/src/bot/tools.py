from langchain.tools import tool

from .rag import init_db
from .RagData import collection, embedding, directory

db = init_db(collection, embedding, directory)



def retrieve_document( query:str)->str:

    """Retrieve information to help answer a query."""
    
    print("iniciando retrieve")
    

    resp = db.similarity_search(query=query, k=1)
   
    answer = "\n".join([i.page_content for i in resp])
    
    return answer

# e=retrieve_document("Horarios")
# print(e)