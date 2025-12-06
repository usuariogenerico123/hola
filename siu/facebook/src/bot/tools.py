from langchain.tools import tool
import RagData
from rag import init_db




@tool
def retrieve_document(query:str):
    """Informacion Retrieve para responder preguntas"""

    db = init_db(RagData.collection, RagData.embedding, RagData.directory)

    resp = db.similarity_search(
        query,
        k=2
    )
    answer = "\n".join([i.page_content for i in resp])
    
    return answer