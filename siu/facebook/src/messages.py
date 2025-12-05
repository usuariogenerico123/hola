import requests
from data import *


msg_method = "v21.0/me/messages"

def sendMessage(text:str, id=None)->bool:
    #print(BEARER_TOKEN_MESSENGER)
    header = {
        "Authorization": f"Bearer {BEARER_TOKEN_MESSENGER}",
        "Content-Type": "application/json"
    }
    data = {
        "message":{"text":text},
        "recipient": {"id": f"{id}"}
    }

    try:
        resp = requests.post(URL_FACEBOOK_API + msg_method, headers=header, json=data)
        print(resp.status_code)
        return True
    except Exception as e:
        print(e)
        return False 
