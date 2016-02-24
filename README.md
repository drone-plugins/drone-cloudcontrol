# drone-cloudcontrol

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-cloudcontrol/status.svg)](http://beta.drone.io/drone-plugins/drone-cloudcontrol)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-cloudcontrol/coverage.svg)](https://aircover.co/drone-plugins/drone-cloudcontrol)
[![](https://badge.imagelayers.io/plugins/drone-cloudcontrol:latest.svg)](https://imagelayers.io/?images=plugins/drone-cloudcontrol:latest 'Get your own badge on imagelayers.io')

Drone plugin to deploy or update a project on cloudControl. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-cloudcontrol <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "app": "helloworld",
        "deployment": "default",
        "email": "octocat@github.com",
        "password": "my_password"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-cloudcontrol <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "app": "helloworld",
        "deployment": "default",
        "email": "octocat@github.com",
        "password": "my_password"
    }
}
EOF
```
