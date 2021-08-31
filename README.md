#### simple http server
- simple http server used for k8s testing.
- /time returns current time
- /status returns status:okay
- /exit returns os.Exit(1) crashing the pod causing it to fail inside of k8s
- currently using Google best practice distroless image. This means there is no shell in the docker container. See more at https://github.com/GoogleContainerTools/distroless
