A simple repo to show how to use `crane` to pull an image without relying on the
`docker` daemon.

Run the following command to see it in action:

```bash
docker run -v $PWD:/go/src -it --rm golang /bin/bash /go/src/run.sh
```

The last line of `run.sh` checks that the docker daemon isn't available inside
the container.
