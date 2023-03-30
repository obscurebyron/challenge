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

Basic setup
-----------

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

The frontend
------------

* Install node: `brew install node`
* Create NextJS project using TypeScript and Tailwind: `npx create-next-app next-tailwind-typescript-starter --example with-typescript`

The User microservice
---------------------

* Install pip: 

python3 -m pip install --upgrade setuptools
python3 -m pip install --upgrade pip

* Install fastapi: `python3 -m pip install fastapi`
* download existing project with fastapi, motor, and mongodb: `git clone https://github.com/mongodb-developer/mongodb-with-fastapi.git`, `python -m pip install -r requirements.txt`
* pull the MongoDB docker image: `docker pull mongo`

The Articles microservice
-------------------------

* Example project using suggested tech: Go, Fiber, Ent, Postgres

git clone https://github.com/TheBeachMaster/golang-ent.git

* get a Postgres image: `docker pull postgres`
* start postgres: `docker run -e POSTGRES_PASSWORD=password postgres`
* install pgadmin for mac to ease working with db: `brew install --cask pgadmin4`
* review [this site](https://levelup.gitconnected.com/lets-go-and-build-an-application-with-ent-b45909b3aa90)
* installing openapi spec - reading [this](https://pkg.go.dev/entgo.io/contrib/entoas#section-readme)
    * per instructions, installing entoas: `go get -u entgo.io/contrib/entoas`
* installing ogent, reading [this](https://github.com/ariga/ogent)
    * install: `go get ariga.io/ogent@main`
    * This didn't really do what I hoped/expected (which was to show a nice UI
    for the endpoints).  Instead, it built an OpenAPI spec and then generated
    the whole set of CRUD operations for the class I have currently ("Article"))

The Comments microservice
-------------------------

* Install rust: `brew install rust`
* Using quickstart from [here](https://rocket.rs/v0.5-rc/guide/quickstart/), carry out the following steps to create a project for Quickstart in Rust:
    * git clone https://github.com/SergioBenitez/Rocket
    * cd Rocket
    * git checkout v0.5-rc
    * cd examples/hello
    * cargo run
* get the Redis image: `docker pull redis`
* start the Redis database: `docker run redis`

MVP interfaces
--------------

* Frontend - raw prototype running, have login/register/article listing
    * Login page
    * Article listing page
    * Article detail (with comments)
    * User Details
* User auth API - raw prototype running, have register and login functioning (bad security - not even using hashes right now)
    * register user (username) -> return password
    * login (username, password) -> set Cookie
    * authorization ???
* Articles API - raw prototype running, 
    * create article (article text)
    * list -> list of articles
    * view article -> contents of one article
    * favorability
* Comments API
    * create comment (related URL [article or comment] - linked list?)
    * get comments (related URL) -> list of comments
