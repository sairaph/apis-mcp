---
title: Batch DNS Records
page_id: operation-post-zones-zone-id-dns-records-batch-fc6b7379
path: operations/dns-records-for-a-zone
description: |-
    Send a Batch of DNS Record API calls to be executed together.

    Notes:
    - Although Cloudflare will execute the batched operations in a single database transaction, Cloudflare's distributed KV store must treat each record change as a single key-value pair. This means that the propagation of changes is not atomic. See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/batch-record-changes/ "Batch DNS records") for more information.
    - The operations you specify within the /batch request body are always executed in the following order:

        - Deletes
        - Patches
        - Puts
        - Posts
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records/batch
operation_ids:
    - dns-records-for-a-zone-batch-dns-records
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch DNS Records

`POST /zones/{zone_id}/dns_records/batch`

Operation ID: `dns-records-for-a-zone-batch-dns-records`

Send a Batch of DNS Record API calls to be executed together.

Notes:
- Although Cloudflare will execute the batched operations in a single database transaction, Cloudflare's distributed KV store must treat each record change as a single key-value pair. This means that the propagation of changes is not atomic. See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/batch-record-changes/ "Batch DNS records") for more information.
- The operations you specify within the /batch request body are always executed in the following order:

    - Deletes
    - Patches
    - Puts
    - Posts

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-batch-dns-records", "summary": "Batch DNS Records", "description": "Send a Batch of DNS Record API calls to be executed together.\n\nNotes:\n- Although Cloudflare will execute the batched operations in a single database transaction, Cloudflare's distributed KV store must treat each record change as a single key-value pair. This means that the propagation of changes is not atomic. See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/batch-record-changes/ \"Batch DNS records\") for more information.\n- The operations you specify within the /batch request body are always executed in the following order:\n\n    - Deletes\n    - Patches\n    - Puts\n    - Posts\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}, {"$ref": "#/components/parameters/dns-records_include_shadow_metadata"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns-request-batch-object"}}}}, "responses": {"200": {"description": "Batch DNS Records response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_batch"}}}}, "4XX": {"description": "Batch DNS Records response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:batch"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "batch"}
```
