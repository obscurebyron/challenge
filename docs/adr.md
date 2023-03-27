Architectural Decision Record
=============================

What's driving us
-----------------

Review Github contents and history
• Review pipeline contents (not history)
• Review SonarQube Score (not history)
• Review best practices

    Execute local demo:

    kn quickstart kind -k 1.24.3 -n kind
    helm install demo <github helm chart artifact link>
    git clone <github repo>
    cd <repo>/tests
    npx playwright test

_March 27, 2023_

* Install Visual Studio Code for an IDE to begin `brew install --cask visual-studio-code`

* Reading through https://knative.dev/docs/getting-started/quickstart-install/

* Install "kind" (Kubernetes in Docker) since testing this will use that, in the following command: `kn quickstart kind -k 1.24.3 -n kind`
  * Install kind from [here](https://kind.sigs.k8s.io/docs/user/quick-start): `brew install kind`
  * Install docker for Mac: `brew install --cask docker` and run docker desktop
  * run the command `kind create cluster`



