import superinvoke

from .tags import Tags
from .envs import Envs


class Tools(superinvoke.Tools):
    Docker = superinvoke.Tool(
        name="docker",
        version=">=20.10.14",
        tags=[*Tags.As("*")],
        path="docker",
    )

    Compose = superinvoke.Tool(
        name="docker compose",
        version=">=2.2.3",
        tags=[*Tags.As("*")],
        path=f"docker compose --env-file envs/{Envs.Current}/.env --file envs/{Envs.Current}/docker-compose.yaml",
    )

    Git = superinvoke.Tool(
        name="git",
        version=">=2.34.1",
        tags=[*Tags.As("*")],
        path="git",
    )

    Go = superinvoke.Tool(
        name="go",
        version=">=1.22.0",
        tags=[Tags.DEV, Tags.CI_INT],
        path="go",
    )

    GoTestSum = superinvoke.Tool(
        name="gotestsum",
        version="1.11.0",
        tags=[Tags.DEV, Tags.CI_INT],
        links={
            superinvoke.Platforms.LINUX: (
                "https://github.com/gotestyourself/gotestsum/releases/download/v1.11.0/gotestsum_1.11.0_linux_amd64.tar.gz",
                "gotestsum",
            ),
            superinvoke.Platforms.MACOS: (
                "https://github.com/gotestyourself/gotestsum/releases/download/v1.11.0/gotestsum_1.11.0_darwin_arm64.tar.gz",
                "gotestsum",
            ),
            superinvoke.Platforms.WINDOWS: (
                "https://github.com/gotestyourself/gotestsum/releases/download/v1.11.0/gotestsum_1.11.0_windows_amd64.tar.gz",
                "gotestsum.exe",
            ),
        },
    )

    GolangMigrate = superinvoke.Tool(
        name="golang-migrate",
        version="4.17.0",
        tags=[Tags.DEV],
        links={
            superinvoke.Platforms.LINUX: (
                "https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz",
                "migrate",
            ),
            superinvoke.Platforms.MACOS: (
                "https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.darwin-arm64.tar.gz",
                "migrate",
            ),
            superinvoke.Platforms.WINDOWS: (
                "https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.windows-amd64.zip",
                "migrate.exe",
            ),
        },
    )

    GolangCILint = superinvoke.Tool(
        name="golangci-lint",
        version="1.55.2",
        tags=[Tags.DEV, Tags.CI_INT],
        links={
            superinvoke.Platforms.LINUX: (
                "https://github.com/golangci/golangci-lint/releases/download/v1.55.2/golangci-lint-1.55.2-linux-amd64.tar.gz",
                "golangci-lint-1.55.2-linux-amd64/golangci-lint",
            ),
            superinvoke.Platforms.MACOS: (
                "https://github.com/golangci/golangci-lint/releases/download/v1.55.2/golangci-lint-1.55.2-darwin-arm64.tar.gz",
                "golangci-lint-1.55.2-darwin-arm64/golangci-lint",
            ),
            superinvoke.Platforms.WINDOWS: (
                "https://github.com/golangci/golangci-lint/releases/download/v1.55.2/golangci-lint-1.55.2-windows-amd64.zip",
                "golangci-lint-1.55.2-windows-amd64/golangci-lint.exe",
            ),
        },
    )

    Squawk = superinvoke.Tool(
        name="squawk",
        version="0.28.0",
        tags=[Tags.DEV, Tags.CI_INT],
        links={
            superinvoke.Platforms.LINUX: (
                "https://github.com/sbdchd/squawk/releases/download/v0.28.0/squawk-linux-x86_64",
                ".",
            ),
            superinvoke.Platforms.MACOS: (
                "https://github.com/sbdchd/squawk/releases/download/v0.28.0/squawk-darwin-x86_64",
                ".",
            ),
        },
    )
