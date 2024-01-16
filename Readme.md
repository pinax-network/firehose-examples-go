## Firehose for Antelope Example

This is an exemplary application on how to stream blocks from Firehose and transform them into Antelope blocks.

Create yourself an api key on https://app.pinax.network and then run:

```
export PINAX_KEY=<your_api_key>
export SUBSTREAMS_API_TOKEN=$(curl https://auth.pinax.network/v1/auth/issue -s --data-binary '{"api_key":"'$PINAX_KEY'"}' | jq -r .token)
go run main.go
```

The program will print the block numbers it receives. 
