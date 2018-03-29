from concurrent import futures
import time

import grpc

import nessus_ingest_pb2
import nessus_ingest_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class NessusIngest(nessus_ingest_pb2_grpc.NessusIngestServicer):

    def NessusRaw(self, request, context):
        return nessus_ingest_pb2.HelloReply(message='Hello, %s!' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nessus_ingest_pb2_grpc.add_NessusIngestServicer_to_server(NessusIngest(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
