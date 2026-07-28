---
title: Edit SSL Certificate Pack Validation Method
page_id: operation-patch-zones-zone-id-ssl-verification-certificate-pack-id-fbac86e4
path: operations/ssl-verification
description: Edit SSL validation method for a certificate pack. A PATCH request will request an immediate validation check on any certificate, and return the updated status. If a validation method is provided, the validation will be immediately attempted using that method.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/ssl/verification/{certificate_pack_id}
operation_ids:
    - ssl-verification-edit-ssl-certificate-pack-validation-method
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit SSL Certificate Pack Validation Method

`PATCH /zones/{zone_id}/ssl/verification/{certificate_pack_id}`

Operation ID: `ssl-verification-edit-ssl-certificate-pack-validation-method`

Edit SSL validation method for a certificate pack. A PATCH request will request an immediate validation check on any certificate, and return the updated status. If a validation method is provided, the validation will be immediately attempted using that method.

## Definition

```yaml
{"operationId": "ssl-verification-edit-ssl-certificate-pack-validation-method", "summary": "Edit SSL Certificate Pack Validation Method", "description": "Edit SSL validation method for a certificate pack. A PATCH request will request an immediate validation check on any certificate, and return the updated status. If a validation method is provided, the validation will be immediately attempted using that method.", "parameters": [{"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_cert_pack_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validation_method-3"}}}}, "responses": {"200": {"description": "Edit SSL Certificate Pack Validation Method response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_validation_method_response_collection"}}}}, "4XX": {"description": "Edit SSL Certificate Pack Validation Method response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_validation_method_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["SSL Verification"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read", "SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.verification", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
