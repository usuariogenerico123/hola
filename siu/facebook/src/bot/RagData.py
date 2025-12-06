from langchain_google_genai import GoogleGenerativeAIEmbeddings

from pathlib import Path



model_embedding="models/gemini-embedding-001"
collection = "paraiso_resort"
directory= f"{Path().resolve()}/src/bot/db_chroma"
embedding = GoogleGenerativeAIEmbeddings(model=model_embedding)
archive = "./documents/resortAzure.txt"
