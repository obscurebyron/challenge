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
    * Install the Kubernetes CLI: `brew install kubectl`
    * Install the KNative CLI: `brew install knative/client/kn`
    * Install the KNative quickstart plugin: `brew install knative-sandbox/kn-plugins/quickstart`
    * When running `kn quickstart kind`, got: `kn-quickstart: bad CPU type in executable`
        * built binary of quickstart only targets amd64.  Try building it from source.
        * source requires bash > 4, install bash: `brew install bash`, run `/opt/Homebrew/Cellar/bash/5.2.15/bin/bash`
        * source requires go, install: `brew install go`, got version `go version go1.20.2 darwin/arm64`
        * build quickstart with `./hack/build.sh` in its source directory
        * copy binary to ~/bin
        * add `export PATH=$PATH:/Users/byron/bin` to ~/.zshrc
        * run `kn quickstart kind` and it appears to work:

            byron@Byrons-MBP challenge % kn quickstart kind
            Running Knative Quickstart using Kind
            ✅ Checking dependencies...
                Kind version is: 0.17.0

            ☸ Creating Kind cluster...
            Creating cluster "knative" ...
            ✓ Ensuring node image (kindest/node:v1.25.3) 🖼 
            ✓ Preparing nodes 📦  
            ✓ Writing configuration 📜 
            ✓ Starting control-plane 🕹️ 
            ✓ Installing CNI 🔌 
            ✓ Installing StorageClass 💾 
            ✓ Waiting ≤ 2m0s for control-plane = Ready ⏳ 
            • Ready after 19s 💚
            Set kubectl context to "kind-knative"
            You can now use your cluster with:

            kubectl cluster-info --context kind-knative

            Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/

            🍿 Installing Knative Serving v1.9.2 ...
                CRDs installed...
                Core installed...
                Finished installing Knative Serving
            🕸️ Installing Kourier networking layer v1.9.2 ...
                Kourier installed...
                Ingress patched...
                Finished installing Kourier Networking layer
            🕸 Configuring Kourier for Kind...
                Kourier service installed...
                Domain DNS set up...
                Finished configuring Kourier
            🔥 Installing Knative Eventing v1.9.7 ... 
                CRDs installed...
                Core installed...
                In-memory channel installed...
                Mt-channel broker installed...
                Example broker installed...
                Finished installing Knative Eventing
            🚀 Knative install took: 2m25s 
            🎉 Now have some fun with Serverless and Event Driven Apps!
    * Install helm: `brew install helm`


