package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone-plugins/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"github.com/fern4lvarez/gocclib/cclib"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone cloudControl Plugin built from %s\n", buildCommit)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Email) == 0 {
		fmt.Println("Please provide an email")

		os.Exit(1)
		return
	}

	if len(vargs.Password) == 0 {
		fmt.Println("Please provide a password")

		os.Exit(1)
		return
	}

	if len(vargs.Application) == 0 {
		vargs.Application = repo.Name
	}

	if len(vargs.Deployment) == 0 {
		vargs.Deployment = "default"
	}

	err := run(&workspace, &build, &vargs)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func run(workspace *drone.Workspace, build *drone.Build, vargs *Params) error {
	repo.GlobalName(build).Run()
	repo.GlobalUser(build).Run()

	api := cclib.NewAPI()
	api.SetUrl("https://api.cloudcontrolled.com")

	if err := api.CreateToken(vargs.Email, vargs.Password); err != nil {
		return errors.New("Failed to authenticate with email/password")
	}

	if _, err := api.ReadApplication(vargs.Application); err != nil {
		return errors.New("Failed to find the requested application")
	}

	if workspace.Keys != nil && len(workspace.Keys.Public) > 0 {
		publicKey, err := api.CreateUserKey(
			vargs.Email,
			workspace.Keys.Public)

		if err != nil {
			return errors.New("Failed to create a deployment key")
		}

		defer func() {
			api.DeleteUserKey(
				vargs.Email,
				publicKey.Id)
		}()
	}

	if err := repo.WriteKey(workspace); err != nil {
		return err
	}

	defer func() {
		execute(
			repo.RemoteRemove(
				"deploy"),
			workspace)
	}()

	cmd := repo.RemoteAdd(
		"deploy",
		remote(vargs))

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	if vargs.Commit {
		if err := execute(repo.ForceAdd(), workspace); err != nil {
			return err
		}

		if err := execute(repo.ForceCommit(), workspace); err != nil {
			return err
		}
	}

	cmd = repo.RemotePush(
		"deploy",
		vargs.Deployment,
		vargs.Force)

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	if _, err := api.UpdateDeployment(vargs.Application, vargs.Deployment, "", "", "", 0, 0); err != nil {
		return errors.New("Failed to trigger the final deployment")
	}

	return nil
}

func remote(vargs *Params) string {
	return fmt.Sprintf(
		"ssh://%s@cloudcontrolled.com/repository.git",
		vargs.Application)
}

func execute(cmd *exec.Cmd, workspace *drone.Workspace) error {
	trace(cmd)

	cmd.Dir = workspace.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
