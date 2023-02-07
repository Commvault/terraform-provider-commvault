# commvault_usergroup (Data Source)


## Example Usage

```hcl
data "commvault_usergroup" "test-user-group" {
    name = "Test User Group"
    }
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Read-Only

- `id` (String) The ID of this resource.

