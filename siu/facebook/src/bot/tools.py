from langchain.tools import tool

from .rag import init_db

from . import RagData


@tool
def retrieve_document(query:str):
    
    """Retrieve information to help answer a query."""
    db = init_db(RagData.collection, RagData.embedding, RagData.directory)
    print("iniciando retrieve")
    resp = db.similarity_search(
        query,
        k=2
    )
    print(resp)
    answer = "\n".join([i.page_content for i in resp])

    return answer