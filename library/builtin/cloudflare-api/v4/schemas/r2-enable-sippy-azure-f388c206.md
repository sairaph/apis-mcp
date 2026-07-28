---
title: r2_enable_sippy_azure
page_id: schema-r2-enable-sippy-azure-f388c206
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_enable_sippy_azure

```yaml
{"properties": {"destination": {"description": "R2 bucket to copy objects to.", "type": "object", "properties": {"accessKeyId": {"description": "ID of a Cloudflare API token.\nThis is the value labelled \"Access Key ID\" when creating an API.\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n", "type": "string"}, "provider": {"type": "string", "enum": ["r2"], "x-auditable": true}, "secretAccessKey": {"description": "Value of a Cloudflare API token.\nThis is the value labelled \"Secret Access Key\" when creating an API.\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n", "type": "string", "x-sensitive": true}}}, "source": {"description": "Azure Blob Storage container to copy objects from.", "type": "object", "properties": {"accountKey": {"description": "Access key for the Azure Storage account. Mutually exclusive with sasToken.", "type": "string", "x-sensitive": true}, "accountName": {"description": "Name of the Azure Storage account.", "type": "string", "x-auditable": true}, "container": {"description": "Name of the Azure Blob Storage container.", "type": "string", "x-auditable": true}, "provider": {"type": "string", "enum": ["azure"], "x-auditable": true}, "sasToken": {"description": "Shared Access Signature token for the Azure Storage account. Mutually exclusive with accountKey.", "type": "string", "x-sensitive": true}}}}}
```
