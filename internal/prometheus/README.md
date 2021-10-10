# Prometheus
This directory contains the required Dockerfiles and scripts to test the prometheus integration with microservices.

## prometheus.yml
Contains the scrape configs to collect metrics from any number of services.

It currently only scrapes the microservice running on the specified address along with Prometheus itself for testing purposes. This allows us to check if any issues are due to a configuration/code error on the microservice side or if Prometheus itself is failing to import.

## Dockerfile
Build a Dockerfile with prometheus.yml inside already. It is possible to skip this step and mount prometheus.yml to /etc/prometheus/prometheus.yml on startup. However, building the Docker image with the yaml file directly inside allows us to avoid any path issues.

## docker-compose.yml
Setup for connecting prometheus and any microservice for local testing. This is required for the microservice to be able to run on the same port that prometheus is listening on.

```image: my-prometheus```
Should contain the name of the image built with the Dockerfile in this directory.
Build with : ```docker build -t my-prometheus .```

## up, down, clean
Basic scripts for starting up and shutting down the docker-compose.
