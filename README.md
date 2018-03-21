# Integration-Nessus-Ingest

# Usage:

Use the ```nessus_ingest_pb2_grpc``` package:

```python
	"github.com/Project-Prismatica/integration-nessus-ingest/nessus_ingest_pb2_grpc"
```

## Compiling go bindings

1. Add the ```protoc``` plugin for python using pip ```python -m pip install grpcio```.

2. run:
```bash
$ python -m grpc_tools.protoc -I..\nessus-ingest\ --python_out=. --grpc_python_out=. nessus_ingest.proto
```

3. Invoke ```python nessus_ingest_server.py```
