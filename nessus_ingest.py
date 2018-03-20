from __future__ import print_function

import grpc

import nessus_ingest_pb2
import nessus_ingest_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = nessus_ingest_pb2_grpc.NessusIngestStub(channel)
    response = stub.SayHello(nessus_ingest_pb2.HelloRequest(name='you'))
    print("NessusIngest client received: " + response.message)


if __name__ == '__main__':
    run()
