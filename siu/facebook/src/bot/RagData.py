from langchain_google_genai import GoogleGenerativeAIEmbeddings



model_embedding="models/gemini-embedding-001"
collection = "paraiso_resort"
directory="./db_chroma"
embedding = GoogleGenerativeAIEmbeddings(model=model_embedding)
archive = "./documents/resortAzure.txt"
