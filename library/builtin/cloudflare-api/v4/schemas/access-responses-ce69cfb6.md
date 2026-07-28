---
title: access_responses
page_id: schema-access-responses-ce69cfb6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_responses

```yaml
{"type": "object", "properties": {"cf_resource_id": {"description": "The unique Cloudflare-generated Id of the SCIM resource.", "type": "string", "example": "bd97ef8d-7986-43e3-9ee0-c25dda33e4b0"}, "error_description": {"description": "The error message which is generated when the status of the SCIM request is 'FAILURE'.", "type": "string", "example": "Invalid JSON body"}, "idp_id": {"description": "The unique Id of the IdP that has SCIM enabled.", "type": "string", "example": "df7e2w5f-02b7-4d9d-af26-8d1988fca630"}, "idp_resource_id": {"description": "The IdP-generated Id of the SCIM resource.", "type": "string", "example": "all_employees"}, "logged_at": {"$ref": "#/components/schemas/access_timestamp"}, "request_body": {"description": "The JSON-encoded string body of the SCIM request.", "type": "string", "example": "{}}"}, "request_method": {"description": "The request method of the SCIM request.", "type": "string", "example": "DELETE"}, "resource_group_name": {"description": "The display name of the SCIM Group resource if it exists.", "type": "string", "example": "ALL_EMPLOYEES"}, "resource_type": {"description": "The resource type of the SCIM request.", "type": "string", "example": "GROUP"}, "resource_user_email": {"description": "The email address of the SCIM User resource if it exists.", "type": "string", "format": "email", "example": "john.smith@example.com"}, "status": {"description": "The status of the SCIM request.", "type": "string", "example": "FAILURE"}}}
```
