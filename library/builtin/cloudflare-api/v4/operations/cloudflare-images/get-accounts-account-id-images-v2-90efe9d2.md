---
title: List images V2
page_id: operation-get-accounts-account-id-images-v2-3cf32756
path: operations/cloudflare-images
description: |-
    List up to 10000 images from CF Images, with up to 1000 results per page. Use the optional parameters below to get a specific range of images.
    Pagination is supported via continuation_token.

    **Metadata Filtering (Optional):**

    You can optionally filter images by custom metadata fields using the `meta.<field>[<operator>]=<value>` syntax.

    **Supported Operators:**
    - `eq` / `eq:string` / `eq:number` / `eq:boolean` - Exact match
    - `gt` / `gt:number` - Greater than (number only)
    - `gte` / `gte:number` - Greater than or equal (number only)
    - `lt` / `lt:number` - Less than (number only)
    - `lte` / `lte:number` - Less than or equal (number only)
    - `in` / `in:string` / `in:number` - Match any value in list (pipe-separated)

    **Metadata Filter Constraints:**
    - Maximum 5 metadata filters per request
    - Maximum 5 levels of nesting (e.g., `meta.first.second.third.fourth.fifth`)
    - Maximum 10 elements for list operators (`in`)
    - Supports string, number, and boolean value types
    - Range operators (`gt`, `gte`, `lt`, `lte`) only accept numeric values

    **Filter Consistency:**
    Filters are combined with AND logic. The system does not validate whether filter combinations are logically consistent. For example, `meta.priority[eq:number]=5&meta.priority[lte:number]=3` will return zero results because no value can satisfy both conditions simultaneously. It is the caller's responsibility to ensure filter combinations make sense.

    **Examples:**
    ```
    # List all images
    /images/v2

    # Filter by metadata [eq]
    /images/v2?meta.status[eq:string]=active

    # Filter by metadata [in]
    /images/v2?meta.status[in]=pending|deleted|flagged

    # Filter by metadata [in:number]
    /images/v2?meta.ratings[in:number]=4|5

    # Filter by metadata range [gte:number]
    /images/v2?meta.priority[gte:number]=1

    # Filter by bounded range
    /images/v2?meta.priority[gte:number]=1&meta.priority[lte:number]=5

    # Filter by nested metadata
    /images/v2?meta.region.name[eq]=eu-west

    # Combine metadata filters with creator
    /images/v2?meta.status[eq]=active&creator=user123

    # Multiple metadata filters (AND logic)
    /images/v2?meta.status[eq]=active&meta.priority[eq:number]=5
    ```
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2
operation_ids:
    - cloudflare-images-list-images-v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List images V2

`GET /accounts/{account_id}/images/v2`

Operation ID: `cloudflare-images-list-images-v2`

List up to 10000 images from CF Images, with up to 1000 results per page. Use the optional parameters below to get a specific range of images.
Pagination is supported via continuation_token.

**Metadata Filtering (Optional):**

You can optionally filter images by custom metadata fields using the `meta.<field>[<operator>]=<value>` syntax.

**Supported Operators:**
- `eq` / `eq:string` / `eq:number` / `eq:boolean` - Exact match
- `gt` / `gt:number` - Greater than (number only)
- `gte` / `gte:number` - Greater than or equal (number only)
- `lt` / `lt:number` - Less than (number only)
- `lte` / `lte:number` - Less than or equal (number only)
- `in` / `in:string` / `in:number` - Match any value in list (pipe-separated)

**Metadata Filter Constraints:**
- Maximum 5 metadata filters per request
- Maximum 5 levels of nesting (e.g., `meta.first.second.third.fourth.fifth`)
- Maximum 10 elements for list operators (`in`)
- Supports string, number, and boolean value types
- Range operators (`gt`, `gte`, `lt`, `lte`) only accept numeric values

**Filter Consistency:**
Filters are combined with AND logic. The system does not validate whether filter combinations are logically consistent. For example, `meta.priority[eq:number]=5&meta.priority[lte:number]=3` will return zero results because no value can satisfy both conditions simultaneously. It is the caller's responsibility to ensure filter combinations make sense.

**Examples:**
```
# List all images
/images/v2

# Filter by metadata [eq]
/images/v2?meta.status[eq:string]=active

# Filter by metadata [in]
/images/v2?meta.status[in]=pending|deleted|flagged

# Filter by metadata [in:number]
/images/v2?meta.ratings[in:number]=4|5

# Filter by metadata range [gte:number]
/images/v2?meta.priority[gte:number]=1

# Filter by bounded range
/images/v2?meta.priority[gte:number]=1&meta.priority[lte:number]=5

# Filter by nested metadata
/images/v2?meta.region.name[eq]=eu-west

# Combine metadata filters with creator
/images/v2?meta.status[eq]=active&creator=user123

# Multiple metadata filters (AND logic)
/images/v2?meta.status[eq]=active&meta.priority[eq:number]=5
```

## Definition

```yaml
{"operationId": "cloudflare-images-list-images-v2", "summary": "List images V2", "description": "List up to 10000 images from CF Images, with up to 1000 results per page. Use the optional parameters below to get a specific range of images.\nPagination is supported via continuation_token.\n\n**Metadata Filtering (Optional):**\n\nYou can optionally filter images by custom metadata fields using the `meta.<field>[<operator>]=<value>` syntax.\n\n**Supported Operators:**\n- `eq` / `eq:string` / `eq:number` / `eq:boolean` - Exact match\n- `gt` / `gt:number` - Greater than (number only)\n- `gte` / `gte:number` - Greater than or equal (number only)\n- `lt` / `lt:number` - Less than (number only)\n- `lte` / `lte:number` - Less than or equal (number only)\n- `in` / `in:string` / `in:number` - Match any value in list (pipe-separated)\n\n**Metadata Filter Constraints:**\n- Maximum 5 metadata filters per request\n- Maximum 5 levels of nesting (e.g., `meta.first.second.third.fourth.fifth`)\n- Maximum 10 elements for list operators (`in`)\n- Supports string, number, and boolean value types\n- Range operators (`gt`, `gte`, `lt`, `lte`) only accept numeric values\n\n**Filter Consistency:**\nFilters are combined with AND logic. The system does not validate whether filter combinations are logically consistent. For example, `meta.priority[eq:number]=5&meta.priority[lte:number]=3` will return zero results because no value can satisfy both conditions simultaneously. It is the caller's responsibility to ensure filter combinations make sense.\n\n**Examples:**\n```\n# List all images\n/images/v2\n\n# Filter by metadata [eq]\n/images/v2?meta.status[eq:string]=active\n\n# Filter by metadata [in]\n/images/v2?meta.status[in]=pending|deleted|flagged\n\n# Filter by metadata [in:number]\n/images/v2?meta.ratings[in:number]=4|5\n\n# Filter by metadata range [gte:number]\n/images/v2?meta.priority[gte:number]=1\n\n# Filter by bounded range\n/images/v2?meta.priority[gte:number]=1&meta.priority[lte:number]=5\n\n# Filter by nested metadata\n/images/v2?meta.region.name[eq]=eu-west\n\n# Combine metadata filters with creator\n/images/v2?meta.status[eq]=active&creator=user123\n\n# Multiple metadata filters (AND logic)\n/images/v2?meta.status[eq]=active&meta.priority[eq:number]=5\n```\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "continuation_token", "in": "query", "schema": {"$ref": "#/components/schemas/images_images_list_continuation_token"}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of items per page", "type": "number", "default": 1000, "maximum": 10000, "minimum": 10}}, {"name": "sort_order", "in": "query", "schema": {"description": "Sorting order by upload time", "type": "string", "default": "desc", "enum": ["asc", "desc"]}}, {"name": "creator", "in": "query", "schema": {"description": "Internal user ID set within the creator field. Setting to empty string \"\" will return images where creator field is not set", "type": "string", "nullable": true}}, {"name": "meta.<field>[<operator>]", "in": "query", "description": "Optional metadata filter(s). Multiple filters can be combined with AND logic.\n\n**Operators:**\n- `eq`, `eq:string`, `eq:number`, `eq:boolean` - Exact match\n- `gt`, `gt:number` - Greater than (number only)\n- `gte`, `gte:number` - Greater than or equal (number only)\n- `lt`, `lt:number` - Less than (number only)\n- `lte`, `lte:number` - Less than or equal (number only)\n- `in`, `in:string`, `in:number` - Match any value in pipe-separated list\n\n**Examples:**\n- `meta.status[eq]=active`\n- `meta.priority[eq:number]=5`\n- `meta.enabled[eq:boolean]=true`\n- `meta.priority[gte:number]=1`\n- `meta.score[lt:number]=100`\n- `meta.region[in]=us-east|us-west|eu-west`\n\n**Note:** Filter consistency is not validated. Contradictory filters (e.g., `meta.priority[eq:number]=5&meta.priority[lte:number]=3`) will return zero results.\n", "schema": {"type": "string"}, "explode": true, "style": "form"}], "responses": {"200": {"description": "List images response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_images_list_response_v2"}}}}, "400": {"description": "Bad request", "content": {"application/json": {"examples": {"invalid_operator": {"summary": "Invalid metadata operator", "value": {"errors": [{"code": 5400, "message": "Unsupported metadata filter operator: 'not-eq'"}], "messages": [], "result": null, "success": false}}, "too_many_filters": {"summary": "Too many metadata filters", "value": {"errors": [{"code": 5400, "message": "Too many metadata filters: 6 provided, maximum 5 allowed"}], "messages": [], "result": null, "success": false}}}, "schema": {"allOf": [{"$ref": "#/components/schemas/images_images_list_response_v2"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}, "4XX": {"description": "List images response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_images_list_response_v2"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "list", "x-forge-hidden": false}
```
