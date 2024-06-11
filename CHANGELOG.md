## 1.0.0.beta.6
- Adds backwards compatible support for [classic API auth changes](https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes) using `WithTokenAuth` client option
- Adds basic `AuthToken` struct for storing token and checking if it is expired
- Bump Go version from `1.18` to `1.22.2`
- Patch `gopkg.in/yaml.v3 Denial of Service` vulnerability ([dependabot alert](https://github.com/DataDog/jamf-api-client-go/security/dependabot/2))

## 1.0.0.beta.5
- Adds support for XML nesting extension attributes in `ComputerDetails`
- Renames `ExtensionAttributes` type to `ExtensionAttribute`
- Adds XML field tags for `ExtensionAttribute`

## 1.0.0.beta.4
- Adds support for `/classes` endpoint
- Fixes bug in computer update response
- Adds more intuitive search parameters for `GetComputers` (i.e id, name, serialnumber)
- Adds in `UpdateComputer` method
## 1.0.0.beta.3
- Refactors list methods to return list of objects by default i.e `j.Computers() => []BasicComputerInfo`.
- Refactors list related structs to use `List` key
## 1.0.0.beta.2
- Adds support for `/computerextensionattributes` endpoint
- Fixes minor typo in `client.go` causing tests to fail
## 1.0.0.beta.1
Initial release of Jamf classic API Go client with support for managing the following:
- computers
- scripts
- policies