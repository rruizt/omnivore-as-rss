# Omnivore as RSS

This application queries the Omnivore API to get the latest articles and exposes an RSS feed with them. 

You can use it for applications or devices that don't support integration with Omnivore directly, but do support RSS. ie. [KOReader](https://koreader.rocks/)

## Usage

### Using Docker

1. Get your [Omnivore authentication token](https://docs.omnivore.app/integrations/api.html#getting-an-api-token)

2. Run the Docker Image

```shell
docker run -p 8090:8090 -e OMNIVORE_AUTH_TOKEN='[YOUR-AUTH-TOKEN]' /omnivore-to-rss:latest-amd64
```

3. Access the endpoint with your preferred software

```shell
curl http://localhost:8090/rss
```

You can use Docker Compose as well:

```yaml
services:
  omnivore-as-rss:
    image: rruizt/omnivore-as-rss:latest
    restart: unless-stopped
    environment:
      - OMNIVORE_AUTH_TOKEN=[YOUR_AUTH_TOKEN]
    ports:
      - "8090:8090"
    healthcheck:
      test: ["CMD", "wget" ,"--no-verbose", "--tries=1", "http://localhost:8090/rss"]
      interval: 1m
      timeout: 3s
```

### Using the binaries

#### Linux/MacOS

1. Download the released tar.gz for Linux or MacOS (Darwin) from [here](https://github.com/rruizt/omnivore-as-rss/releases/latest)
2. Decompress the file and run the binary with the right flags

```
./omnivore-as-rss -t=00000000-0000-0000-0000-000000000000 
```

#### Windows

1. Download the released zip for Windows from [here](https://github.com/rruizt/omnivore-as-rss/releases/latest)
2. Decompress the file and run the .exe file with the right flags

ie.
```
omnivore-as-rss.exe -t=00000000-0000-0000-0000-000000000000 
```


## Configuration

You can configure `omnivore-as-rss` with environment variables or command line flags

| Environment Var / Flag  | Description | Example |
| ------------- | ------------- |-----------|
| OMNIVORE_AUTH_TOKEN / -t  | The API token from Omnivore  | `00000000-0000-0000-0000-000000000000` |
| OMNIVORE_AUTH_TOKEN_FILEPATH / -tf | The filepath of the file containing the API token from Omnivore | `/run/secrets/omnivore` |



