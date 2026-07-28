---
title: Retrieve a portal configuration
page_id: operation-get-v1-billing-portal-configurations-configuration-50d84b8e
path: operations/untagged
description: <p>Retrieves a configuration that describes the functionality of the customer portal.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/billing_portal/configurations/{configuration}
operation_ids:
    - GetBillingPortalConfigurationsConfiguration
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a portal configuration

`GET /v1/billing_portal/configurations/{configuration}`

Operation ID: `GetBillingPortalConfigurationsConfiguration`

<p>Retrieves a configuration that describes the functionality of the customer portal.</p>

## Definition

```yaml
{"summary": "Retrieve a portal configuration", "description": "<p>Retrieves a configuration that describes the functionality of the customer portal.</p>", "operationId": "GetBillingPortalConfigurationsConfiguration", "parameters": [{"name": "configuration", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing_portal.configuration"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
