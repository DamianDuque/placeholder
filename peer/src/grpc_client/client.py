import grpc
import time
from protobuf_files.filesystem_pb2_grpc import FileSystemStub
from protobuf_files.filesystem_pb2 import Filename
class Client_Remote:
    def _create_client(self,socket: str):
         channel: grpc.Channel = grpc.insecure_channel(socket)
         return  FileSystemStub(channel)
        
    def upload(self, socket: str) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        client=self._create_client(socket)
        req=  Filename("test.txt")
        response= client.Upload(req)
        print(response)

    
    def download(self, socket, filenamestr) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        print(socket)
        client = self._create_client(socket)
        req =  Filename(id=1,name=filenamestr)
        print(req)
        response = client.Download(req)
        print(response)    
    
        
