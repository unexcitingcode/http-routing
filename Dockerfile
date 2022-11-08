FROM debian:bullseye
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ARG TARGETOS
ARG TARGETARCH
RUN apt-get update -y && \
    apt-get install --no-install-recommends -y \
        curl=7.74.0-1.3+deb11u3 \
        git=1:2.30.2-1 \
        pre-commit=2.10.1-1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
RUN curl -sL https://go.dev/dl/go1.19.3.${TARGETOS:-linux}-${TARGETARCH:-amd64}.tar.gz -o /opt/go.tar.gz && \
    mkdir -p /opt/go && \
    tar -C /opt -xf /opt/go.tar.gz && \
    rm /opt/go.tar.gz
ENV GOROOT="/opt/go"
ENV GOPATH="/opt/go/packages"
ENV PATH="${GOROOT}/bin:${PATH}"
COPY ["./scripts", "/opt/scripts"]
ENTRYPOINT ["/opt/scripts/entrypoint.sh"]
