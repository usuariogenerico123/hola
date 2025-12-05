from langchain.tools import tool

import os


@tool 
def crea_archivo(file: str):
    
    """Crea un archivo"""

    os.system(f"touch {file}")

    return f"Archivo {file} creado con exito en {os.system('pwd')}"

@tool 
def elimina_archivo(file :str):

    """Elimina un archivo"""
    
    os.system(f"rm {file}")

    return f"archivo {file} eliminado con exito"


@tool
def crea_archivo_en_carpeta(file: str, path: str):

    """Crea un archivo que esta ubicado en una carpeta"""
    os.system(f"touch {path}/{file}")

    return f"Archivo {file} creado correctamente"

@tool 
def elimina_archivo_en_carpeta(file: str, path:str):
    """Elimina un archivo que esta ubicado en una carpeta"""
    
    os.system(f"rm {path}/{file}")

    return f"Archivo {file} eliminado correctamente"

@tool 
def escribe_en_archivo(text:str, file:str, path=None):
    """Escribe datos dentro de un archivo si no hay ruta:path omite y si hay, usa la ruta/archivo"""
    

    if(path == None):
        with open(file, "w") as f:
            f.write(text)
    if(path != None):

        with open(f"{path}/{file}", "w") as f:
            f.write(text)

    return f"Archivo escrito en {file}"