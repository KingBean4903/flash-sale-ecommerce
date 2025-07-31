# ⚡ Flash Sale E-commerce Platform

This project simulates a high-traffic e-commerce flash sale architecture with order deduplication, stock contention handling, and real-time observability.

## 📦 Features

- High-throughput **Order Service** with Redis deduplication
- Kafka event streaming for decoupled communication
- Redis-based inventory reservation with TTL expiration
- K6 load testing for stress and chaos simulation
- Prometheus + Grafana monitoring dashboard
- Kubernetes-ready deployment (AWS EKS optimized)

## 🛠 Tech Stack

- Go (Order Service, Inventory Service)
- Kafka (Bitnami/Strimzi)
- Redis
- K6 for load testing
- Prometheus + Grafana for metrics
- Kubernetes (EKS Auto Mode)

## 🚀 Getting Started (Docker Compose)

```bash
docker-compose up --build

