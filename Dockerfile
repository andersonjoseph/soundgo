FROM golang:1.22

WORKDIR /usr/src/app

ENV PATH "$PATH:/usr/src/bin"
ENV SHELL "bash"

SHELL ["/bin/bash", "-c"]

# COPY go.mod go.sum ./

# RUN go mod download && go mod verify

COPY . .

# install dlv (https://github.com/go-delve/delve)
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1

# install node (https://github.com/nvm-sh/nvm)
ENV NODE_VERSION 22.6.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
ENV NVM_DIR /root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH "/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

# install redocly cly (https://redocly.com/docs/cli/)
RUN npm i -g @redocly/cli@latest
