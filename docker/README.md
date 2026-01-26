# Deploy Focalboard with Docker

## Docker

The Dockerfile gives a quick and easy way to build the latest Focalboard server and deploy it locally. In the example below,
the Focalboard database and files will be persisted in a named volume called `fbdata`.

From the Focalboard project root directory:

```bash
docker build -f docker/Dockerfile -t focalboard .
docker run -it -v "fbdata:/opt/focalboard/data" -p 80:8000 focalboard
```

Open a browser to [localhost](http://localhost) to start

## Alternative architectures

From the Focalboard project root directory:

```bash
docker build -f docker/Dockerfile --platform linux/arm64 -t focalboard .
docker run -it -v "fbdata:/opt/focalboard/data" -p 80:8000 focalboard
```

## Docker-Compose

Docker Compose provides the option to automate the build and run step when you just need a disposable local test environment.

To start the server using the default local settings, run the following from the project root:

```bash
docker compose up --build
```

This command builds the Focalboard image, exposes it on `localhost:8000`, and creates the persistent named volume `fbdata`.

If you need to bind the server to port 80 for manual testing, use the helper compose file in `docker/docker-compose.yml`, which reuses the same service definition but updates the port mapping and environment variables needed for that scenario.
