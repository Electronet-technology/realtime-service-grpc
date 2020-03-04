#commands for generating client server code with protoc

#go
export PATH=$PATH:~/go/bin

protoc --go_out=plugins=grpc:. *.proto 

#python
pip install grpcio-tools

python -m grpc_tools.protoc -I . --python_out=python --grpc_python_out=python *.proto

#js
protoc -I=. *.proto \
  --js_out=import_style=commonjs:js