---
title: Revoke Certificate
page_id: operation-delete-certificates-certificate-id-5fd36194
path: operations/origin-ca
description: Revoke an existing Origin CA certificate by its serial number. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /certificates/{certificate_id}
operation_ids:
    - origin-ca-revoke-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke Certificate

`DELETE /certificates/{certificate_id}`

Operation ID: `origin-ca-revoke-certificate`

Revoke an existing Origin CA certificate by its serial number. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).

## Definition

```yaml
{"operationId": "origin-ca-revoke-certificate", "summary": "Revoke Certificate", "description": "Revoke an existing Origin CA certificate by its serial number. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Revoke Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_revoke_response"}}}}, "4XX": {"description": "Revoke Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_revoke_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"user_service_key": []}], "tags": ["Origin CA"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-ca-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
