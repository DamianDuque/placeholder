
from src.protobuf_files.filesystem_pb2 import Response
from  src.protobuf_files.filesystem_pb2_grpc import FileSystem

class FileSystemService(FileSystem):
   def Upload(self, request, context):
      print("Initializing uploading of file: " + str(request.filename)+ "to peer")
      return Response("uploaded file "+str(request.filename))
   def Download(self, request, context):
      print("Initializing downloadinf of file: " + str(request.filename)+ "to peer")
      return Response("uploaded file "+str(request.filename))