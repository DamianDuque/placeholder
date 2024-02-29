from src.server import server

if __name__ == "__main__":
    server_test= server.Server("localhost","50051")
    server_test.start_server()
    server_test.wait_for_termination()
    
