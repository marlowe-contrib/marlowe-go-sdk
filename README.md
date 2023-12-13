# marlowe-go-sdk

## Update automatically generated code Marlowe Runtime Web REST API

Modify the `openapi.json`, making sure it's a valid OpenAPI spec, commit and push. The github action defined in `.github/workflows/codegen.yml` will do the rest using [openapi-generator-cli](https://openapi-generator.tech/)

## Check OpenAPI specs validity

Run the following command to check if the `openapi.json` is valid.

```sh
npx @openapitools/openapi-generator-cli validate -i openapi.json
```
