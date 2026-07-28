---
title: billable-usage-api_usage_record
page_id: schema-billable-usage-api-usage-record-5239514b
path: schemas
description: Represents a single billable usage record.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_usage_record

Represents a single billable usage record.

```yaml
{"description": "Represents a single billable usage record.", "type": "object", "properties": {"BillingCurrency": {"description": "Specifies the billing currency code (ISO 4217).", "type": "string", "example": "USD"}, "BillingPeriodStart": {"description": "Indicates the start of the billing period.", "type": "string", "format": "date-time", "example": "2025-02-01T00:00:00Z"}, "ChargePeriodEnd": {"description": "Indicates the end of the charge period.", "type": "string", "format": "date-time", "example": "2025-02-02T00:00:00Z"}, "ChargePeriodStart": {"description": "Indicates the start of the charge period.", "type": "string", "format": "date-time", "example": "2025-02-01T00:00:00Z"}, "ConsumedQuantity": {"description": "Specifies the quantity consumed during this charge period.", "type": "number", "example": 150000}, "ConsumedUnit": {"description": "A display name for the unit of measurement used for the product (for example, \"GB-months\", \"GB-seconds\"). May be empty when the unit is implicit in the service name.", "type": "string", "example": "GB-months"}, "ContractedCost": {"description": "Specifies the cost for this charge period in the billing currency.", "type": "number", "example": 0.75}, "CumulatedContractedCost": {"description": "Specifies the cumulated cost for the billing period in the billing currency.", "type": "number", "example": 2.25}, "CumulatedPricingQuantity": {"description": "Specifies the cumulated pricing quantity for the billing period.", "type": "integer", "example": 4500000}, "PricingQuantity": {"description": "Specifies the pricing quantity for this charge period.", "type": "integer", "example": 150000}, "ServiceFamilyName": {"description": "Identifies the product family for the Cloudflare service.", "type": "string", "example": "Workers"}, "ServiceName": {"description": "Identifies the Cloudflare service.", "type": "string", "example": "Workers Standard"}, "SubscriptionId": {"description": "The identifier for the Cloudflare subscription.", "type": "string", "example": "3F3CD4CQ6N7FXO7IK6NVFJBOYA", "nullable": true}, "ZoneId": {"description": "The identifier for the Cloudflare zone (zone tag).", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "nullable": true}, "ZoneName": {"description": "The display name of the Cloudflare zone.", "type": "string", "example": "example.com", "nullable": true}}, "required": ["ChargePeriodStart", "ChargePeriodEnd", "BillingPeriodStart", "ServiceName", "ConsumedQuantity", "ConsumedUnit", "PricingQuantity", "CumulatedPricingQuantity", "ContractedCost", "CumulatedContractedCost", "BillingCurrency"]}
```
