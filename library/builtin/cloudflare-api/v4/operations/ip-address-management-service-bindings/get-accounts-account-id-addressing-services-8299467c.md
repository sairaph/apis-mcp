---
title: List Services
page_id: operation-get-accounts-account-id-addressing-services-b0af62d8
path: operations/ip-address-management-service-bindings
description: Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a service running on the Cloudflare network to enable a Cloudflare product on the IP addresses. This endpoint can be used as a reference of available services on the Cloudflare network, and their service IDs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/services
operation_ids:
    - ip-address-management-service-bindings-list-services
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Services

`GET /accounts/{account_id}/addressing/services`

Operation ID: `ip-address-management-service-bindings-list-services`

Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a service running on the Cloudflare network to enable a Cloudflare product on the IP addresses. This endpoint can be used as a reference of available services on the Cloudflare network, and their service IDs.

## Definition

```yaml
{"operationId": "ip-address-management-service-bindings-list-services", "summary": "List Services", "description": "Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a service running on the Cloudflare network to enable a Cloudflare product on the IP addresses. This endpoint can be used as a reference of available services on the Cloudflare network, and their service IDs.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "Service names and IDs", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/addressing_service_identifier"}, "name": {"$ref": "#/components/schemas/addressing_service_name"}}}}}}]}}}}, "4XX": {"description": "List Services response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Service Bindings"], "x-api-token-group": ["IP Prefixes: Write", "IP Prefixes: Read"]}
```
