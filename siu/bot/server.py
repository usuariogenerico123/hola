from flask import Flask 


app = Flask(__name__)

@app.route("/", methods=["GET"])
def get():
    return "Hello, World! GET", 200




@app.route("/", methods=["POST"])
def post():
    return "Hello, World! POST", 200



if __name__ == "__main__":
    app.run(port=3001, debug=True)