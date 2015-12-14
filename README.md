# drone-cloudcontrol

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-cloudcontrol/status.svg)](http://beta.drone.io/drone-plugins/drone-cloudcontrol)
[![](https://badge.imagelayers.io/plugins/drone-cloudcontrol:latest.svg)](https://imagelayers.io/?images=plugins/drone-cloudcontrol:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to cloudControl

## Usage

```sh
./drone-cloudcontrol <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
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

Build the Docker container using `make`:

```
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-cloudcontrol <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
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
