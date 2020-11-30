# jamf-api-client-go

This repository contains an unoffical Go API client for Jamf REST API's.
  - [Jamf Classic API](https://www.jamf.com/developers/apis/classic/overview/)
    - [API Reference](https://www.jamf.com/developers/apis/classic/reference/)
    - [Code Samples](https://www.jamf.com/developers/apis/classic/code-samples/)
  - [Jamf Pro API](https://www.jamf.com/developers/apis/jamf-pro/overview/) **(TBD)**
    - [API Reference](https://www.jamf.com/developers/apis/jamf-pro/reference/)
    - **Note:** Development on the pro client has not been started, if an endpoint is to be added here please keep in mind that the endpoints are prefaced with `v1` per the API Reference below and therfore the file structure should reflect `/pro/v1/*.go`

## Disclaimers

The API client remains in active development and there is no affiliation with [Jamf](https://github.com/jamf)

This is **not** an official [Jamf](https://github.com/jamf) API client and the client is **not** formally
supported and the code is available as-is.

[Contribution](https://github.com/DataDog/jamf-api-client-go/CONTRIBUTING.md) is welcome and appreciated! 🚀 💜
## Usage

```go
import  jamf "github.com/DataDog/jamf-api-client-go/classic"

// You can optionally setup a custom HTTP client to use which can
// include any settings you desire. If you would like to use the 
// default client configuration just pass nil. This will default 
// to a client that is simply configured with a timeout of 1 minute
myCustomHTTPClient := &http.Client{
  Timeout: time.Minute,
}

// Create a client instance to interact with API
j, err := jamf.NewClient("https://jamf.example.com", "example.username", "super-secret-password", myCustomHTTPClient)
if err != nil {
  fmt.Println(err.Error())
  os.Exit(1)
}

// Example: Get All Computers
computers, err := j.Computers()
if err != nil {
  os.Exit(1)
}

// Example: Create Script
newScript := &jamf.ScriptContents{
  Name: "Script with API Creation",
}
s, err := j.CreateScript(newScript)
if err != nil {
  os.Exit(1)
}

// Example: Get Script Details
scriptDetails, err := j.ScriptDetails(37)
if err != nil {
  os.Exit(1)
}
```
> Note: It is recommended to use environment variables or a KMS for Jamf credentials

## Developing & Contribution

When building out more API client functionality it's helpful to use the [Postman Collection](https://github.com/jamf/Classic-API-Postman-Collection) provided by Jamf for testing endpoint responses and payloads.
### Tests

Unit tests should be added for all endpoints and passing prior to being checked into the `main` branch

 `go test -v ./...` or `make test`
### Releasing

When a release is ready to be pushed:
- Ensure all intended changes have been merged to `main`
- Create a release on the releases page.
- Specify the version you want to release, following [Semantic Versioning](https://semver.org/spec) principles.
  
  > If the tag isn’t meant for production use, add a pre-release version after the version name. Some good pre-release versions might be v0.2-alpha or v5.9-beta.3. and check the `This is a pre-release` box at the bottom

- Add release title containing the relase version if desired `ex v1.0.0-beta1 Initial Beta Release`
- Add sufficient changelog contents into the description of the release. (`git log` may be helpful)
- Create/Publish the release, which will automatically create a tag on the HEAD commit. (no binaries should be uploaded)

Once a versioned package has been released, the contents of that version **MUST NOT** be modified. Any modifications must be released as a new version.

## API Coverage

The functionality below represents what the current API client is capable of:

#### Classic
  - `/computers`
    - [x] [Get all computers](https://www.jamf.com/developers/apis/classic/reference/#/computers/findComputers)
    - [x] Get specific computer by [ID](https://www.jamf.com/developers/apis/classic/reference/#/computers/findComputersById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/computers/findComputersByName)
    - [x] Update computer by [ID](https://www.jamf.com/developers/apis/classic/reference/#/computers/updateComputerById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/computers/updateComputerByName)

  - `/scripts`
    - [x] [Get all scripts](https://www.jamf.com/developers/apis/classic/reference/#/scripts/findScripts)
    - [x] Get specific script by [ID](https://www.jamf.com/developers/apis/classic/reference/#/scripts/findScriptsById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/scripts/findScriptsByName) **Note:** only JSON response available for GET requests (XML Unmarshalling not currently configured)
    - [x] Update script by [ID](https://www.jamf.com/developers/apis/classic/reference/#/scripts/updateScriptById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/scripts/updateScriptByName)
    - [x] Create new script by [ID](https://www.jamf.com/developers/apis/classic/reference/#/scripts/createScriptById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/scripts/createScriptByName)
    - [x] Delete script by [ID](https://www.jamf.com/developers/apis/classic/reference/#/scripts/deleteScriptById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/scripts/deleteScriptByName)

  - `/policies`
    - [x] [Get all policies](https://www.jamf.com/developers/apis/classic/reference/#/policies/findPolicies)
    - [x] Get policy by [ID](https://www.jamf.com/developers/apis/classic/reference/#/policies/findPoliciesById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/policies/findPoliciesByName)
    - [x] Update policy by [ID](https://www.jamf.com/developers/apis/classic/reference/#/policies/updatePolicyById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/policies/updatePolicyByName)
    - [x] Create new policy by [ID](https://www.jamf.com/developers/apis/classic/reference/#/policies/createPolicyById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/policies/updatePolicyByName)
    - [x] Delete policy by [ID](https://www.jamf.com/developers/apis/classic/reference/#/policies/deletePolicyById) or [Name](https://www.jamf.com/developers/apis/classic/reference/#/policies/deletePolicyByName)

  - `/osxconfigurationprofiles` **(In Progress)**
    - [ ] [Get all configuration profiles](https://www.jamf.com/developers/apis/classic/reference/#/osxconfigurationprofiles/findOsxConfigurationProfiles)
    - [ ] [Get configuration profile by ID or Name](https://www.jamf.com/developers/apis/classic/reference/#/osxconfigurationprofiles/findOsxConfigurationProfilesById)
    - [ ] [Update configuration profile by ID or Name](https://www.jamf.com/developers/apis/classic/reference/#/osxconfigurationprofiles/updateOsxConfigurationProfileById)
    - [ ] [Create configuration profile by ID or Name](https://www.jamf.com/developers/apis/classic/reference/#/osxconfigurationprofiles/createOsxConfigurationProfileById)
