from langchain_text_splitters import RecursiveCharacterTextSplitter


def load_text(text: str):
    with open(text, "r") as f:
        return f.read()


texto = load_text("./ResortAzure.txt")
print(texto)
























