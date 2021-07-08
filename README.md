# cuddly-vector
:robot: Look at this cute and little robot

## Getting started

#### Go to scripts folder and download dependencies
```shell
cd scripts && go mod tidy
```

### Setup and Authentication

#### Run the login script to get your **Client GUID Token**
Remember to set your robot's IP address using the `-host` flag
```shell
go run login/main.go -host "<robot's ip address>"
```

#### Export the `BOT_TARGET` and `BOT_TOKEN` environment variables
Set `BOT_TARGET` to your robot's IP + port
```shell
export BOT_TARGET="<robot's ip address>:443"
```

Set `BOT_TOKEN` to your GUID token from the login script
```shell
export BOT_TOKEN="<guid token>"
```

### Running Scripts
Once you have done the setup and auth steps, you can simply run the scripts like: 
```shell
go run speak/main.go -talk "hello, world!"
```
