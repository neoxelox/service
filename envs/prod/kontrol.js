module.exports = {
  docker: {},
  jobs: {
    api: {
      cron: "0 * * * * *",
      delay: 15000,
      url: "http://api:1111/health",
      method: "GET",
      retry: 5,
      timeout: 5000,
      notify: (state, error) => {
        switch (state) {
          case "UNHEALTHY":
            let message = error.response && error.response.body ? error.response.body : error;
            return {
              attachments: [
                {
                  color: "danger",
                  mrkdwn_in: ["text"],
                  text: "*`Api`* containers are *unhealthy*, restarting...\n```" + message + "```",
                },
              ],
            };
          case "HEALED":
            return {
              attachments: [
                {
                  color: "good",
                  mrkdwn_in: ["text"],
                  text: "*Healed* *`api`* containers.",
                },
              ],
            };
          case "NOT_HEALED":
            return {
              attachments: [
                {
                  color: "warning",
                  mrkdwn_in: ["text"],
                  text: "*Failed* to heal *`api`* containers, retrying...\n```" + error + "```",
                },
              ],
            };
        }
      },
      heal: async (docker, _) => {
        const find = (container, name) => _.find(container.Names, (cname) => cname.includes(name));
        const containers = await docker.listContainers();
        const api = await docker.getContainer(_.find(containers, (container) => find(container, "api")).Id);
        await api.restart();
      },
    },
    worker: {
      cron: "0 * * * * *",
      delay: 15000,
      url: "http://worker:1112/health",
      method: "GET",
      retry: 5,
      timeout: 5000,
      notify: (state, error) => {
        switch (state) {
          case "UNHEALTHY":
            let message = error.response && error.response.body ? error.response.body : error;
            return {
              attachments: [
                {
                  color: "danger",
                  mrkdwn_in: ["text"],
                  text: "*`Worker`* containers are *unhealthy*, restarting...\n```" + message + "```",
                },
              ],
            };
          case "HEALED":
            return {
              attachments: [
                {
                  color: "good",
                  mrkdwn_in: ["text"],
                  text: "*Healed* *`worker`* containers.",
                },
              ],
            };
          case "NOT_HEALED":
            return {
              attachments: [
                {
                  color: "warning",
                  mrkdwn_in: ["text"],
                  text: "*Failed* to heal *`worker`* containers, retrying...\n```" + error + "```",
                },
              ],
            };
        }
      },
      heal: async (docker, _) => {
        const find = (container, name) => _.find(container.Names, (cname) => cname.includes(name));
        const containers = await docker.listContainers();
        const worker = await docker.getContainer(_.find(containers, (container) => find(container, "worker")).Id);
        await worker.restart();
      },
    },
  },
};
