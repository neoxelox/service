import ssl

import superinvoke
from superinvoke.constants import Platforms

from . import tasks
from .envs import Envs
from .tools import Tools

# Temporal fix very annoying error: certificate verify failed: unable to get local issuer certificate
ssl._create_default_https_context = ssl._create_unverified_context

root = superinvoke.init(tools=Tools, envs=Envs)

root.configure({
    "run": {
        "pty": (
            Platforms.CURRENT != Platforms.WINDOWS
            and Envs.Current != Envs.Ci
        ),
    },
})


root.add_task(tasks.list)
root.add_task(tasks.start)
root.add_task(tasks.clean)
root.add_task(tasks.logs)
root.add_task(tasks.prune)
root.add_task(tasks.build)
root.add_task(tasks.deploy)
root.add_task(tasks.test)
root.add_task(tasks.lint)
root.add_task(tasks.format)
root.add_task(tasks.migrate)
root.add_task(tasks.run)
