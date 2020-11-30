# Contributing

First of all, thanks for contributing!

This document provides some basic guidelines for contributing to this repository. To propose improvements, feel free to submit a pull request.
## Getting Started

When building out more API client functionality it's helpful to use the [Postman Collection](https://github.com/jamf/Classic-API-Postman-Collection) provided by Jamf for testing endpoint responses and payloads.

The best way to interact with this project is and will be a work in progress. However there are
some core tenets to keep in mind.

1. Code review is important, but it is not important enough to block progress on anyone's end.
   What is more important is architectural review and unit tests.

2. Splitting your changes in easy to consume pieces that are self contained is important, but
   not important enough to impose a constant burden on the code author.
### Submitting Issues

Github issues are welcome, feel free to submit error reports and feature requests! Make sure to add enough details to explain your use case.

### Features & Fixes

- Fork a repository
- Add new functionality or apply a fix
- Check that tests are passing
- Create PR against `main` and await review/approval. Please review the outlined process below to ensure your PR is good to go.

## Pull Request Process

1. Ensure PR template is filled out and all relevant action items in the checklist are covered. PR titles should include one of the following: 
     - `[FEATURE]`   - To be used with new features such as adding a new endpoint
     - `[BUG FIX]`   - To be used with any bug fixes
     - `[REFACTOR]`  - To be used when code change is specific to refactoring and no new features are added
     - `[MISC]`      - To be used with documentation, dependency updates, etc.
  
2. Unless what you are doing is absolutely trivial, add unit tests. Good unit tests come in bundles,
   and usually test for both the expected and how one handles the unexpected case. To ensure your tests 
   pass please run `go test -v ./...` or `make test` from the root of this repo.

3. If you are making changes that add new components, new data structures, or reorganize an existing
   flow, it is helpful to discuss your architecture first. That discussion is better had over an
   issue.
   
4. Invest in making your commit messages descriptive. If you fix a bug, tell us more about how you found
   it, in which circumnstances it appears, etc. That is important for others as well as for future
   you. Split big changes in smaller commits so it is easier for others and future you to follow what
   you are doing. `git add -p` is a really powerful tool and you should use it.
   However, once you added a pull request, and that PR started to see lively discussion,
   do not rewrite history: keep adding code on top.

5. Code formatting and organization is not something worth arguing about. It has value, but not
   nearly enough to justify the time spent. Ensure formatting is consistent and run `go fmt ./...` 
   before requesting a review.
