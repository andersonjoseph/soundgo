FROM golang:1.22

WORKDIR /usr/src/app

ENV PATH "$PATH:/usr/src/bin"
ENV SHELL "bash"

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# install goose (https://github.com/pressly/goose)
RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh | \
    GOOSE_INSTALL=/usr/src sh -s v3.19.2

# install air (https://github.com/cosmtrek/air)
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/src/bin

# install nvm (https://github.com/nvm-sh/nvm) and redocly
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash && \
    export NVM_DIR="$HOME/.nvm" && \
  [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  && \
  nvm install node && \
  npm i -g @redocly/cli@latest

# install hurl (https://hurl.dev/)
RUN curl -O --output-dir /usr/src/bin --location --remote-name \
  https://github.com/Orange-OpenSource/hurl/releases/download/4.3.0/hurl_4.3.0_amd64.deb -o /usr/src/bin/hurl.deb && \
  apt update && apt -y install /usr/src/bin/hurl_4.3.0_amd64.deb

# install dlv (https://hurl.dev/)
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1
