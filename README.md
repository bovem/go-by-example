## Running Docker Container
```bash
docker run -it --entrypoint="/bin/bash" -v $(pwd):/tmp go:latest
```