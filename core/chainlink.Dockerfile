# MAKE ALL CHANGES WITHIN THE DEFAULT WORKDIR FOR YARN AND GO DEP CACHE HITS
FROM node:12-buster
WORKDIR /chainlink

COPY GNUmakefile VERSION ./
COPY tools/bin/ldflags tools/bin/ldflags
ARG COMMIT_SHA

# Install yarn dependencies
COPY yarn.lock package.json .yarnrc ./
COPY solc_bin solc_bin
COPY .yarn .yarn
COPY operator_ui/package.json ./operator_ui/
COPY evm-test-helpers/package.json ./evm-test-helpers/
COPY contracts/package.json ./contracts/
COPY tools/bin/restore-solc-cache ./tools/bin/restore-solc-cache
RUN make yarndep

COPY contracts ./contracts
COPY evm-test-helpers ./evm-test-helpers
COPY tsconfig.cjs.json tsconfig.es6.json ./
COPY operator_ui ./operator_ui

# Build operator-ui and the smart contracts
RUN make contracts-operator-ui-build

# Build the golang binary

FROM golang:1.16-buster
WORKDIR /chainlink

COPY GNUmakefile VERSION ./
COPY tools/bin/ldflags ./tools/bin/

# Env vars needed for chainlink build
ADD go.mod go.sum ./
RUN go mod download

# Env vars needed for chainlink build
ARG COMMIT_SHA
ARG ENVIRONMENT

COPY --from=0 /chainlink/operator_ui/dist ./operator_ui/dist
COPY core core
COPY packr packr

RUN make chainlink-build

# Final layer: ubuntu with chainlink binary
FROM ubuntu:18.04

ARG CHAINLINK_USER=root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y ca-certificates wget gnupg lsb-release

# Install Postgres for CLI tools, needed specifically for DB backups
RUN wget --quiet -O - https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
 && wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - \
 && echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" |tee /etc/apt/sources.list.d/pgdg.list \
 && apt-get update && apt-get install -y postgresql-client-13 \
 && apt-get clean all

COPY --from=1 /go/bin/chainlink /usr/local/bin/

RUN if [ ${CHAINLINK_USER} != root ]; then \
  useradd --uid 14933 --create-home ${CHAINLINK_USER}; \
  fi
USER ${CHAINLINK_USER}
WORKDIR /home/${CHAINLINK_USER}

EXPOSE 6688
ENTRYPOINT ["chainlink"]
CMD ["local", "node"]
