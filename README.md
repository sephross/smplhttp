#### simple http server
- simple http server used for k8s testing.
- <code>/</code> returns status:okay
- <code>/time</code> returns current time
- <code>/exit</code> returns os.Exit(1) crashing the pod causing it to fail inside of k8s
- currently using Google distroless image. This means there is no bash in the container image.
