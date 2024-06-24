# Golang Load Balancer

## Problem Statement

As microservices architectures become increasingly popular, the need for efficient load balancing grows. Traditional load balancers can be expensive and complex to configure. There's a need for a simple, cost-effective load balancing solution that can be easily configured and deployed, ensuring high availability and reliability of services.

## Purpose

The purpose of this project is to provide a lightweight, configurable load balancer written in Golang that supports multiple load balancing strategies and health checks for service instances. This solution aims to simplify the deployment and management of load balancing in microservices environments.

## Use Case

This load balancer is designed for scenarios where:

- Microservices need to be distributed across multiple instances.
- Traffic needs to be evenly distributed or weighted based on instance capabilities.
- High availability is critical, and only healthy instances should receive traffic.
- A simple, easily configurable load balancer is preferred over more complex solutions.

## How It Works

The load balancer works by reading a configuration file (YAML) that defines the services, their endpoints, and the load balancing strategies to be used. It supports two strategies:

- **Weighted Round Robin**: Distributes traffic based on weights assigned to each instance.
- **Round Robin**: Distributes traffic evenly among instances.

Health checks are performed periodically to ensure that only healthy instances receive traffic. Unhealthy instances are temporarily removed from the rotation until they recover.

## Configuration

Create a config.yaml file with the following structure:

```yaml
services:
  - name: "Test"
    matcher: "/api/v1"
    strategy: "WeightedRoundRobin"
    replicas:
      - host: "http://localhost:8081"
        metadata:
          weight: 10
      - host: "http://localhost:8082"
        metadata:
          weight: 5
  - name: "Test2"
    matcher: "/api/v2"
    strategy: "RoundRobin"
    replicas:
      - host: "http://localhost:8083"
      - host: "http://localhost:8084"
defaultStrategy: "RoundRobin"
```

**Fields**:

- **services**: List of services to load balance.
- **name**: Name of the service.
- **matcher**: URL path matcher for routing.
- **strategy**: Load balancing strategy (WeightedRoundRobin or RoundRobin).
- **replicas**: List of service instances.
- **host**: URL of the service instance.
- **metadata**: Additional metadata (e.g., weight for WeightedRoundRobin).
- **defaultStrategy**: Default load balancing strategy if not specified for a service.

## Deployment

1. **Clone the repository**:

```sh
git clone https://github.com/yourusername/golang-load-balancer.git
cd golang-load-balancer
```

2. **Build the project**:

```sh
go build -o load-balancer
```

3. **Run the load balancer with the configuration file and port**:

```sh
./load-balancer -config-path=config.yaml -port=8080
```

**Flags**:

- **-port**: Port to listen on (default: 8080)
- **-config-path**: Path to the configuration file

## Future Enhancements

- **Additional Load Balancing Strategies**: Implement strategies such as Least Connections, IP Hash, etc.
- **Dynamic Configuration Reload**: Allow reloading configuration without restarting the service.
- **Enhanced Health Checks**: Support for more complex health checks (e.g., checking specific endpoints).
- **Monitoring and Metrics**: Integrate monitoring tools to track load balancer performance and health.

## Contributing

We welcome contributions from the community. Please submit your pull requests with detailed descriptions of the changes and the problem they solve. Make sure to run tests and follow the code style of the project.

## Code of Conduct

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported to the community leaders responsible for enforcement at [gnana097@gmail.com] . All complaints will be reviewed and investigated promptly and fairly.

All community leaders are obligated to respect the privacy and security of the reporter of any incident.

---

Feel free to customize this template further based on your specific requirements and deployment details.

---
