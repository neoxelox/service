import os
import re

from superinvoke import console, rich, task

from .envs import Envs
from .tools import Tools


@task()
def list(context):
    """List containers."""

    context.run(f"{Tools.Compose} ps")


@task(
    help={
        "extra": "Whether extra tooling is included in the infrastructure.",
        "detach": "Whether the infrastructure starts detached.",
    }
)
def start(context, extra=False, detach=False):
    """Start infrastructure."""

    context.run(f"{Tools.Compose} {'--profile extra' if extra else ''} up --build {'--detach' if detach else ''}")


@task()
def clean(context):
    """Remove all containers, volumes and networks."""

    context.run(f"{Tools.Compose} --profile extra down --volumes")


@task(
    help={
        "container": "Container name to show logs.",
    }
)
def logs(context, container):
    """Show logs of a container."""

    context.run(f"{Tools.Compose} logs -f {container}")


@task()
def prune(context):
    """Remove all containers, volumes, networks and images."""

    context.run(f"{Tools.Docker} system prune -a --volumes")


@task()
def build(context):
    """Build images."""

    if Envs.Current != Envs.Ci:
        context.fail(f"build command only available in {Envs.Ci} environment!")

    context.run(
        f"{Tools.Docker} build --file envs/prod/api.Dockerfile "
        f"-t service-api:{context.commit()} "
        f"-t service-api:latest "
        f"."
    )

    context.run(
        f"{Tools.Docker} build --file envs/prod/worker.Dockerfile "
        f"-t service-worker:{context.commit()} "
        f"-t service-worker:latest "
        f"."
    )

    context.run(
        f"{Tools.Docker} build --file envs/prod/cli.Dockerfile "
        f"-t service-cli:{context.commit()} "
        f"-t service-cli:latest "
        f"."
    )


@task()
def deploy(context):
    """Push images to the registry."""

    if Envs.Current != Envs.Ci:
        context.fail(f"deploy command only available in {Envs.Ci} environment!")

    # NOTE: REMOVE
    context.warn("NOTE: REMOVE 'return'")
    return

    context.run(f"{Tools.Docker} push --all-tags service-api")

    context.run(f"{Tools.Docker} push --all-tags service-worker")

    context.run(f"{Tools.Docker} push --all-tags service-cli")


@task(
    help={
        "test": "[<PACKAGE_PATH>]::[<TEST_NAME>]. If empty, it will run all tests.",
        "verbose": "Show stdout of tests.",
        "show": "Show coverprofile page.",
    },
)
def test(context, test="", verbose=False, show=False):
    """Run tests."""

    test_arg = "./..."
    if test:
        test = test.split("::")
        if len(test) == 1 and test[0]:
            test_arg = f"{test[0]}/..."
        if len(test) == 2 and test[1]:
            test_arg += f" -run {test[1]}"

    verbose_arg = ""
    if verbose:
        verbose_arg = "-v"

    parallel_arg = ""
    if os.cpu_count():
        parallel_arg = f"--parallel={os.cpu_count()}"

    coverprofile_arg = ""
    if show:
        coverprofile_arg = "-coverprofile=coverage.out"

    result = context.run(
        f"{Tools.GoTestSum} --format=testname --no-color=False -- {verbose_arg} {parallel_arg} -race -count=1 -cover {coverprofile_arg} {test_arg}",
    )

    if "DONE 0 tests" not in result.stdout:
        packages = 0
        coverage = 0.0

        for cover in re.findall(r"[0-9]+\.[0-9]+(?=%)", result.stdout):
            packages += 1
            coverage += float(cover)

        if packages:
            coverage = round(coverage / packages, 1)

        console.print(
            rich.panel.Panel(
                f"Total Coverage ([bold]{packages} pkg[/bold]): [bold green]{coverage}%[/bold green]",
                expand=False,
            )
        )

    if show:
        context.run(f"{Tools.Go} tool cover -html=coverage.out")
        context.remove("coverage.out")


@task()
def lint(context):
    """Run linter."""

    context.run(f"{Tools.Go} vet ./...")
    context.run(f"{Tools.GolangCILint} run ./... -c .golangci.yaml")
    context.run(f"{Tools.Squawk} -c .squawk.toml migrations/*.sql")


@task()
def format(context):
    """Run formatter."""

    context.run(f"{Tools.Go} fmt ./...")
    context.run(f"{Tools.GolangCILint} run ./... -c .golangci.yaml --fix")


@task(
    help={
        "name": "Migration name.",
    }
)
def migrate(context, name):
    """Create a migration."""

    context.run(f"{Tools.GolangMigrate} create -ext sql -dir migrations -seq -digits 4 {name}")


@task(variadic=True)
def run(context, args):
    """Execute a command."""

    if Envs.Current == Envs.Dev:
        context.run(f"{Tools.Compose} exec -T cli cli {args}", pty=False)
    else:
        context.run(f"{Tools.Compose} run --rm cli {args}", pty=False)
