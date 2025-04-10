FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    GO_VERSION=1.23.0

# Install dependencies and PostgreSQL
RUN apt-get update && \
    apt-get install -y wget curl git postgresql postgresql-contrib \
    build-essential ca-certificates

# Install Go
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Set Go paths
ENV PATH=/usr/local/go/bin:$PATH \
    GOPATH=/go \
    PATH=$GOPATH/bin:$PATH

# Create a postgres data dir and default user/database
RUN service postgresql start && \
    su - postgres -c "psql -c \"CREATE USER forum_user WITH PASSWORD 'forum_pass';\"" && \
    su - postgres -c "psql -c \"CREATE DATABASE forumdb;\"" && \
    su - postgres -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE forumdb TO forum_user;\""

# Create working directory for Go app
WORKDIR /app

# Copy app files (optional placeholder)
COPY . .

# Expose PostgreSQL and Go ports (if needed)
EXPOSE 5432 8080

# Set default command (you can replace this with your Go server command)
CMD service postgresql start && tail -f /dev/null
