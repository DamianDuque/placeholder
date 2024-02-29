import sys
from src.rest_client.apiclient import APIClient
from src.rest_client.messages import Messages
from src.rest_client.setup import SetUp
#from peer.src.rest_client.verification import Verify

def query_procedure():
    print("Please write the name of the file you are currently searching for:")
    file2search = input()
    querydata = {"filename": file2search}
    return querydata
    

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Please run the client using the following format:")
        print("python main.py <Username> <Password> <User url>")
        sys.exit()

    client, messagec, setup = APIClient(), Messages(), SetUp()

    username,password,user_url = sys.argv[1],sys.argv[2],sys.argv[3]
    LogIn_data = {"username": username, "password": password, "user_url": user_url}
    
    login_response = setup.logIn(client.url_servidor, LogIn_data)
    
    print("Logged in successfully!!")
    print(login_response)
    authToken = login_response['token']

    list = SetUp.get_files_in_folder()
    #print(list)
    sent_index_files = setup.do_sendIndex(client.url_servidor, list, username, authToken)
    print("All files were successfully added to the index!!")

    Messages.introduction()
    while True:
        Messages.general_message()
        action = input()
        
        if action == "1":
            querydata = query_procedure()
            query_response = client.do_query(client.url_servidor, querydata, authToken)
            print(query_response)

        elif action == "2":
            logoutdata = {"username": username}
            logout_response = client.logOut(client.url_servidor, logoutdata, authToken)
            print(logout_response)
            break
