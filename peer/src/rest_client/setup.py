import requests
import json
import os

class SetUp():
    
    def logIn(self, url, LogIn_data):
        specurl = url + "api/v1/login"
        try:
            login_response = requests.post(specurl, json=LogIn_data)
            login_response.raise_for_status()
            return login_response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return False
        
    def get_files_in_folder():
        files_array = []
        folder_path = os.path.join(os.path.dirname(__file__), "files")
        for file_name in os.listdir(folder_path):
            if os.path.isfile(os.path.join(folder_path, file_name)):
                files_array.append(file_name)
        return files_array
    
    def do_sendIndex(self, url, file_data, username, authToken):
        specurl = url + "api/v1/sendIndex"
        headers = {
            "Content-Type": "application/json",
            "authToken": authToken
        }

        Index_data = {
            "username": username,
            "files": file_data
        }

        try:
            sendIndex_response = requests.post(specurl, json=Index_data, headers=headers)
            print(sendIndex_response)

            sendIndex_response.raise_for_status()
            return sendIndex_response.json()
        
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None