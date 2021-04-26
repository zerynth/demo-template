# zDeviceManager Grafana Demo

The repository contains a Docker Compose based demo that:

- receives data from a zDeviceManager webhook data stream at the endpoint `zdm/data`
- receives data from a zDeviceManager webhook condition stream at the endpoint `zdm/condition`
- insert data and conditions into a TimescaleDB instance
- fires up a Grafana instance with access to TimescaleDB as data source


## Usage

Clone the repository and type

```docker-compose up```

Grafana will be avaiable at `localhost`
