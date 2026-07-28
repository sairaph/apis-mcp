---
title: Get list of ASNs
page_id: operation-get-accounts-account-id-botnet-feed-configs-asn-47841b1c
path: operations/botnet-threat-feed
description: Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/botnet_feed/configs/asn
operation_ids:
    - botnet-threat-feed-list-asn
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get list of ASNs

`GET /accounts/{account_id}/botnet_feed/configs/asn`

Operation ID: `botnet-threat-feed-list-asn`

Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.

## Definition

```yaml
{"operationId": "botnet-threat-feed-list-asn", "summary": "Get list of ASNs", "description": "Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Get list of ASNs response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dos_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"asn": {"type": "integer", "example": 13335}}}}, "type": "object"}]}}}}, "4XX": {"description": "Get list of ASNs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Botnet Threat Feed"], "x-api-token-group": ["DDoS Botnet Feed Write", "DDoS Botnet Feed Read"]}
```
