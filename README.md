# 🛠 Site Reliability Engineering (SRE) Projects (Phase-by-Phase)

Each project will be documented in GitHub (code + README + screenshots).  

---

## Phase 1: Foundation (0–3 months)

### Project 1: Linux Health Check Service
- Write a minimal go program to monitor CPU, RAM, disk, and key logs.  
- Configure logging and alerting (stdout + file).  

### Project 2: Custom systemd Unit
- Create a systemd service that runs the healthcheck program.  
- Enable auto-start on boot and test with `systemctl` and `journalctl`.  

### Project 3: HAProxy Load Balancer
- Deploy two simple backend servers (Python Flask or Go).  
- Configure HAProxy with round-robin load balancing.  
- Add HTTPS termination (self-signed certificate).  

---

## Phase 2: Infrastructure as Code & Containers (3–6 months)

### Project 4: Containerized Web App
- Containerize a Python/Go web app with Docker.  
- Push to DockerHub or GitHub Container Registry.  

### Project 5: Terraform Infrastructure
- Use Terraform to provision EC2 + VPC + security groups on AWS.  
- Deploy Linux servers with networking and storage.  

### Project 6: Ansible Automation
- Write an Ansible playbook to configure the servers (packages, users, configs).  
- Automate deployment of your containerized app.  

### Project 7: CI/CD Pipeline
- Create a GitHub Actions or Jenkins pipeline.  
- Automate build + test + deploy of your Dockerized app.  

---

## Phase 3: Observability & Reliability (6–9 months)

### Project 8: Monitoring Stack
- Deploy Prometheus + Grafana to monitor CPU, memory, HTTP errors.  
- Build dashboards with alerts.  

### Project 9: Logging Stack
- Deploy ELK or EFK stack (Elasticsearch/Fluentd/Kibana).  
- Centralize logs from your app and servers.  

### Project 10: Runbook
- Write a structured runbook (step-by-step outage recovery guide).  
- Example: “High CPU alert at 95% for 10 min → check pod logs → restart service.”  

### Project 11: Chaos Engineering
- Use LitmusChaos or Gremlin in a test cluster.  
- Simulate failure (kill a pod, network latency).  
- Document results and improvements.  

---

## Phase 4: Security + Cloud Specialization (9–12 months)

### Project 12: AWS 3-Tier Application
- Deploy frontend + API + DB on AWS using Terraform.  
- Configure networking (VPC, subnets, routing).  

### Project 13: IAM Security Hardening
- Apply least-privilege IAM roles to your app and users.  
- Test access restrictions.  

### Project 14: Auto-Scaling & HA
- Implement Auto Scaling Groups in AWS.  
- Deploy app across multiple AZs for high availability.  

### Project 15: Postmortem Report
- Simulate an outage (DB down, API crash).  
- Write a professional postmortem: timeline, root cause, action items.  
