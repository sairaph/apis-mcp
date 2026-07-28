---
title: Delete an ASN
page_id: operation-delete-accounts-account-id-botnet-feed-configs-asn-asn-id-09b4caac
path: operations/botnet-threat-feed
description: Delete an ASN from botnet threat feed for a given user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/botnet_feed/configs/asn/{asn_id}
operation_ids:
    - botnet-threat-feed-delete-asn
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an ASN

`DELETE /accounts/{account_id}/botnet_feed/configs/asn/{asn_id}`

Operation ID: `botnet-threat-feed-delete-asn`

Delete an ASN from botnet threat feed for a given user.

## Definition

```yaml
{"operationId": "botnet-threat-feed-delete-asn", "summary": "Delete an ASN", "description": "Delete an ASN from botnet threat feed for a given user.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "asn_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dos_asn"}}], "responses": {"200": {"description": "Delete ASN response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dos_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"asn": {"type": "integer", "example": 13335}}}}, "type": "object"}]}}}}, "4XX": {"description": "Delete ASN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Botnet Threat Feed"], "x-api-token-group": ["DDoS Botnet Feed Write"]}
```
