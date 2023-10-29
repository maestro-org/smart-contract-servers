# Go server for off-chain queries via the Maestro SDK
## Usage
```bash
git clone https://github.com/maestro-org/smart-contract-servers.git
cd servers/go
go mod tidy
go run main.go
```

## Configuration
Generate a [free API key](https://docs.gomaestro.org/docs/Getting-started/Sign-up-login). Pass that key and the desired network into the `.env` file in this directory:
```
PORT=8080
CARDANO_NETWORK=preprod
MAESTRO_API_KEY=
```
The following networks are supported: `preview`, `preprod`, and `mainnet`
