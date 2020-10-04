from flask import Flask, render_template, request, redirect, jsonify
from flask.json import jsonify
import subprocess
app = Flask(__name__, static_url_path='/static')

@app.route("/update/", methods =["POST", "GET"])
def update_data():
    print("Been pinged, let's update!")
    print("Beginning git pull")
    process = subprocess.Popen(["git", "pull"], stdout=subprocess.PIPE)
    gitpull = process.communicate()[0]
    print("Git pull is done, now let's run the docker container!")
                
               
@app.route("/", methods =["POST", "GET"])
def index():
        return "hi"
        
    
if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)