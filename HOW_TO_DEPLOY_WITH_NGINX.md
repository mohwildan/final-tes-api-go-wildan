# Golang API Deployment Guide

This guide outlines the steps to deploy your Golang-based API, utilizing Nginx as a reverse proxy and configuring firewall rules to ensure secure operation.

---

## 1. Prerequisites

Make sure the following software is installed on your server:

- **Go** for compiling and running the API
- **Nginx** as the web server/reverse proxy
- **UFW** (Uncomplicated Firewall) to manage server security

---

## 2. Project Setup

Before you start the API, ensure that all project dependencies are installed using `go mod`:

```bash
go mod tidy
```

This command will automatically retrieve and tidy up all necessary dependencies for your project.

---

## 3. Key Commands for API Management

### Database Migration

To manage database migrations, use the following commands to apply or roll back schema changes:

```bash
# Apply database updates
go run migrate.go up

# Revert last migration
go run migrate.go down
```

### Adding New Features

To create a new feature in your API, you can generate the necessary code files:

```bash
# Generate feature module
go run feature_generator.go --name "feature_name"
```

---

## 4. Running the Application

### Development Mode

In the development environment, you can directly run the API with this command:

```bash
go run main.go
```

### Production Mode

To build and run the application in production, follow these steps:

```bash
# Build the API
go build -o final-api-be .

# Run the API binary
./final-api-be
```

---

## 5. pm2 for Process Management

To keep your API running in the background, use `pm2` for process management. Here’s how to install and configure it.

### Installing pm2

To install `pm2` globally on your server:

```bash
# Update packages
sudo apt update

# Install pm2
sudo npm install pm2@latest -g
```

### Running the API with pm2

You can now use `pm2` to run your API and keep it alive in the background:

```bash
# Start API with pm2
pm2 start ./final-api-be --name "final-api-be"

# Alternatively, use go run command with pm2
pm2 start "go run main.go" --name "final-api-be"
```

### Managing pm2 Processes

Here are some common pm2 commands to manage your application:

```bash
# View list of running processes
pm2 list

# View logs
pm2 logs final-api-be

# Restart the API process
pm2 restart final-api-be

# Stop the API process
pm2 stop final-api-be

# Ensure pm2 restarts on system boot
pm2 startup
```

---

## 6. Nginx Setup as a Reverse Proxy

Nginx will forward client requests to your API, running on a different port (e.g., 3000). Here’s how to set it up.

### Installing Nginx

To install Nginx on your server:

```bash
sudo apt update
sudo apt install nginx
```

### Configuring Nginx

Create a configuration file for your API application in `/etc/nginx/sites-available/` and link it to `/etc/nginx/sites-enabled/`.

```bash
# Open the Nginx configuration file
sudo nano /etc/nginx/sites-available/final-api-be

# Add the following content
server {
    listen 80;
    server_name _;

    location / {
        proxy_pass http://localhost:5050;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

### Activating Nginx Configuration

Enable the Nginx site configuration, check for syntax errors, and restart Nginx:

```bash
# Enable site configuration
sudo ln -s /etc/nginx/sites-available/final-api-be /etc/nginx/sites-enabled/

# Test configuration for syntax errors
sudo nginx -t

# Restart Nginx to apply changes
sudo systemctl restart nginx
```

---

## 7. Configuring the Firewall

Secure your server by managing firewall rules. Here, we’ll ensure only necessary ports (HTTP, HTTPS, and SSH) are open.

### Checking Firewall Status

Use the following command to check the firewall status:

```bash
sudo ufw status
```

### Allowing Traffic for HTTP, HTTPS, and SSH

Run these commands to open ports 80 (HTTP), 443 (HTTPS), and 22 (SSH):

```bash
# Allow HTTP (port 80)
sudo ufw allow 80

# Allow HTTPS (port 443)
sudo ufw allow 443

# Allow SSH (port 22)
sudo ufw allow 22
```

### Enabling Firewall

Enable the firewall to apply your rules:

```bash
# Enable the firewall
sudo ufw enable

# Check firewall status again
sudo ufw status
```
