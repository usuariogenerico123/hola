from langchain_text_splitters import RecursiveCharacterTextSplitter
import os
from langchain.chat_models import init_chat_model




def load_text(text: str):
    with open(text, "r") as f:
        return f.read()


texto = load_text("./ResortAzure.txt")
doc = RecursiveCharacterTextSplitter()
s=doc.split_text(texto)

























