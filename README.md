# KubeVision: Log Processing and Visualization for Kubernetes

![Banner Image](images/kubevision-diagram.png)

Welcome to KubeVision! This project is a powerful, intuitive system designed to collect, process, and visualize logs generated within a Kubernetes cluster. Developed with Go and equipped with a CLI, it provides real-time insights into the events occurring within your cluster.

## Components

The system is composed of two main parts:

1. **Event Simulator**: A Kubernetes pod running a simple Go application that periodically generates random logs.

2. **Log Processing and Visualization Application**: An application that collects the logs from the event simulator, processes them, and provides a user-friendly interface for log visualization.

### Event Simulator

The simulator generates logs every 5 seconds, following this pattern:

```
[TIMESTAMP] [EVENT_TYPE] Message
```

where:

- `TIMESTAMP` is the current timestamp.
- `EVENT_TYPE` can be INFO, WARNING, or ERROR.
- `Message` is a predefined text message associated with the event type.

### Log Processing and Visualization Application

The application is divided into the following components:

- **Log Collection**: Uses the official Go client library for Kubernetes to collect logs from the event simulator pod.

- **Log Processing**: Extracts relevant information from each log entry and classifies messages by event type.

- **Command Line Interface (CLI)**: Provides an interactive way for users to work with the application. The available commands are:
  - `mycli get logs`: Fetches all recent logs and displays them in their original format.
  - `mycli process logs`: Fetches logs and shows the number of events of each type that occurred in the last 10 minutes.
  - `mycli stream logs`: Follows the pod logs in real time, showing each new entry as it is generated.

## Deployment

KubeVision is dockerized and can be deployed within any Kubernetes cluster. The repository includes a Dockerfile for creating a Docker image of the application and a Kubernetes manifest file (YAML) for deploying the application to a Kubernetes cluster.

## Development Environment

For local development and testing, Minikube or kind can be used to emulate a Kubernetes cluster.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

## License

Distributed under the MIT License. See `LICENSE` for more information.