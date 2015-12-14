Use this plugin for deplying an application to cloudControl. You can override
the default configuration with the following parameters:

* `app` - Application name, defaults to repo name
* `deployment` - Deployment name, defaults to `default`
* `email` - Email address to authenticate
* `password` - Password to authenticate
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

```yaml
deploy:
  cloudcontrol:
    app: helloworld
    deployment: default
    email: octocat@github.com
    password: my_password
```
