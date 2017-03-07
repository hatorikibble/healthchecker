# healthchecker

simple script for testing web services. Adapted after [healthchecker](https://github.com/callistaenterprise/goblog/blob/master/healthchecker/main.go) by @eriklupander

## Usage

    ./healthchecker:
      -ip string
        	ip address to check (default "127.0.0.1")
      -path string
        	path to check (default "/health")
      -port string
        	port to check (default "80")

## Example

    healthchecker -ip 127.0.0.1 -port 5000 -path /main

    bash> healthchecker -ip 127.0.0.1 -path /index && echo "everything works!"`
