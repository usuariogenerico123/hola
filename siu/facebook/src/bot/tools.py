from langchain.tools import tool

from .rag import init_db

from . import RagData


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