# gRPC: predictable Go interfaces

## Run it

```shell
% go run .
```

## Use it

```shell
% curl localhost:8080/v1/list?limit=3&page=5 | jq
{
  "entities": [
    {
      "id": 13,
      "name": "entity_13"
    },
    {
      "id": 14,
      "name": "entity_14"
    },
    {
      "id": 15,
      "name": "entity_15"
    }
  ]
}
```
