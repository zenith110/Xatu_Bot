from os import name
from flask import Flask, render_template, request, redirect, jsonify
from flask.json import jsonify
import subprocess
import docker
import dockerhub_login
from discord_webhook import DiscordWebhook, DiscordEmbed
import datetime
import discord_key
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
    down = DiscordWebhook(url=discord_key.api_key, content="Xatu is going down for a bit!")
    down_response = down.execute()
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
        client.containers.prune()
        subprocess.Popen("sudo", "killall", "./main.go")
        now = datetime.now()
        month = datetime.date.today()
        time_stamp = str(now.strftime("%b %d %Y %H:%M:%S"))
        up = DiscordWebhook(url=discord_key.api_key, content='Xatu is up again! Done at:\n' + time_stamp)
        up_response = up.execute()
        docker_container = client.containers.run(dockerhub_login.repo + ":latest", name= "xatu")
    return "Now running Xatu!"
    

        
    
if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)
