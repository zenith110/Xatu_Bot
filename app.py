from os import name
from flask import Flask, render_template, request, redirect, jsonify
from flask.json import jsonify
import subprocess
import docker
import dockerhub_login
app = Flask(__name__, static_url_path='/static')
@app.route("/update/", methods =["POST", "GET"])
def update_data():
    """
    Does some configuring for dockerhub
    """
    client = docker.from_env()
    client.login(username=dockerhub_login.username, password=dockerhub_login.password)

    """
    If there's a docker instance, pull the latest image from the repo
    """
    try:
        client.images.pull(dockerhub_login.repo)
    # Removes the last instance and pulls the new one
    except:
        client.images.remove(dockerhub_login.repo + ":latest")
        client.images.pull(dockerhub_login.repo)

    # Attempts to deploys a docker container
    try:
        docker_container = client.containers.run(dockerhub_login.repo + ":latest", name= "xatu")
    # If a docker container exist with the name, remove it and make a new instance
    except:
        xatu = client.containers.get("xatu")
        xatu.stop()
        xatu.remove()
        docker_container = client.containers.run(dockerhub_login.repo + ":latest", name= "xatu")
    return "Now running Xatu!"
    
@app.route("/", methods =["POST", "GET"])
def index():
        return "Please use the routes to do commands"
        
    
if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)
