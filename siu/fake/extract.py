
import json

with open("cloud.json", "r") as f:

    jsson = json.loads(f.read()) 

print(jsson["creationTime"])


with open("gool.txt", "w") as f:
    cont = 0
    for i in jsson["prefixes"]:
        cont += 1
        try:
            print(i["ipv4Prefix"])
            f.write(i["ipv4Prefix"]+ " ")
        except:
            pass
        #f.write(i["ipv4Prefix"]+ " ")

    f.close()
print("Total " , cont)
print("Hecho")


