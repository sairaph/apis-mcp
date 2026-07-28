---
title: mq_consumer-request
page_id: schema-mq-consumer-request-e0ea25d5
path: schemas
description: Request body for creating or updating a consumer
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_consumer-request

Request body for creating or updating a consumer

```yaml
{"description": "Request body for creating or updating a consumer", "type": "object", "discriminator": {"mapping": {"http_pull": "#/components/schemas/mq_http-consumer-request", "worker": "#/components/schemas/mq_worker-consumer-request"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/mq_worker-consumer-request"}, {"$ref": "#/components/schemas/mq_http-consumer-request"}]}
```
