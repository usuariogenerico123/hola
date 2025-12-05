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
