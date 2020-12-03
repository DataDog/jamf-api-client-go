### Releasing

To prep a new release create a release branch in the format `release-{release number}` and push it to the remote. 

Once the branch is available create a feature branch off of the relase branch, this will allow us to group new functionality together for each release and `main` will always represent the latest official release.

When a release is ready to be pushed:
- Create a PR between the release branch and `main` and ensure all intended changes have been merged into the release branch and all tests are passing, at this point all changes should have been previously reviewed and will just need a once over before making the release official. Only admins will be allowed to push changes directly to a release branch as they should not change once the PR is open.
- Merge the release branch into `main`
- Create a release on the releases page.
- Specify the version you want to release, following [Semantic Versioning](https://semver.org/spec) principles.
  
  > If the tag isn’t meant for production use, add a pre-release version after the version name. Some good pre-release versions might be v0.2-alpha or v5.9-beta.3

- Add release title containing the relase version if desired `ex v1.0.0-beta1 Initial Beta Release`
- Add sufficient changelog contents into the description of the release. (`git log` may be helpful)
- Create/Publish the release, which will automatically create a tag on the HEAD commit. (no binaries should be uploaded)

Once a versioned package has been released, the contents of that version **MUST NOT** be modified. Any modifications must be released as a new version.
