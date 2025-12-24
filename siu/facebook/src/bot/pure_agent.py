
# To run this code you need to install the following dependencies:
# pip install google-genai

import base64
import os
from google import genai
from google.genai import types



client = genai.Client(
        api_key=os.environ.get("GOOGLE_API_KEY"),
    )
model = "gemini-2.0-flash-lite"
def generate(prompt):
    

    contents = []


    # contents = [
    #     types.Content(
    #         role="user",
    #         parts=[
    #             types.Part.from_text(text=prompt),
    #         ],
    #     ),
    # ]
    contents.append({"role":"user","parts": [{"text":prompt}]})

    # tools = [
    #     types.Tool(googleSearch=types.GoogleSearch(
    #     )),
    # ]
    generate_content_config = types.GenerateContentConfig(
        # thinking_config=types.ThinkingConfig(
        #     thinking_level="HIGH",
        # ),
        #tools=tools,
    )

    #print(contents)
    resp = ""
    for chunk in client.models.generate_content_stream(
        model=model,
        contents=contents,
        config=generate_content_config,
    ):
        resp += chunk.text
        print(chunk.text)

    contents.append({"role":"model","parts":[{"text":resp}]})


if __name__ == "__main__":
    while True:
        u = input(": ")
        if(u == "0"): print("Adios");break 
        else:

            generate(u)








