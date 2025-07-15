<img src=".github/assets/siker.im.svg" alt="siker.im logo" width="200">

🔗 A simple URL shortener service.

## Self-Hosting
This guide explains how to self-host the backend services of **siker.im** using Docker Compose. The frontend (Next.js 15) is hosted on Vercel, while the backend stack — built with Go (Fiber), MongoDB, and Redis — is deployed on your own machine.

### Stack Overview
- **Next.js 15** – Frontend app, hosted on Vercel.
- **Go (Fiber)** – REST API backend.
- **MongoDB** – Encrypted NoSQL database.
- **Redis** – In-memory store used for rate-limiting.
- **Docker Compose** – Container orchestration for local development.

## Back-end Setup Instructions
To set up the server, follow these steps:

### 1. Clone the repository
```bash
git clone git@github.com:lareii/siker.im.git
```
### 2. Navigate to the server directory
```bash
cd siker.im/server
```

### 3. Create a .env file
```bash
cp .env.example .env
```
Update the values in the file according to your environment.

### 4. Start the services:
```bash
make run
```
This will build the Docker images, generate TLS certificates, and start all backend services.

### 5. Access the services
- API: `http://localhost:1209`
- MongoDB: `mongodb://localhost:27017`
- Redis: `redis://localhost:6379`

> [!NOTE]
> MongoDB and Redis services are configured to persist data across restarts.

### Back-end Environment Variables
Located in `.env` inside `/server`.

| Name                        | Description                                               | Default Value                         |
| --------------------------- | --------------------------------------------------------- | ------------------------------------- |
| `MONGODB_URI`               | Full MongoDB URI (with credentials and TLS if applicable) | `mongodb://localhost:27017`           |
| `MONGODB_NAME`              | Name of the MongoDB database                              | `db`                                  |
| `MONGODB_USERNAME`          | MongoDB username                                          | –                                     |
| `MONGODB_PASSWORD`          | MongoDB password                                          | –                                     |
| `REDIS_HOST`                | Redis hostname                                            | `localhost`                           |
| `REDIS_PORT`                | Redis port                                                | `6379`                                |
| `REDIS_PASSWORD`            | Redis password (if any)                                   | –                                     |
| `REDIS_DB`                  | Redis DB index                                            | `0`                                   |
| `PORT`                      | Port that the API server listens on                       | `1209`                                |
| `ALLOWED_ORIGINS`           | CORS origin whitelist (comma-separated)                   | `*`                                   |
| `LOG_LEVEL`                 | Logging level (`debug`, `info`, `warn`, `error`)          | `info`                                |
| `RATE_LIMIT_ENABLED`        | Enable rate limiting                                      | `true`                                |
| `RATE_LIMIT_REQUESTS`       | Max requests per window                                   | `100`                                 |
| `RATE_LIMIT_WINDOW_MINUTES` | Rate-limiting window in minutes                           | `1`                                   |
| `RATE_LIMIT_BLOCK_MINUTES`  | Block time (in minutes) after rate limit exceeded         | `5`                                   |
| `TURNSTILE_SECRET`          | Cloudflare Turnstile secret key                           | `1x0000000000000000000000000000000AA` |

> [!IMPORTANT]
> MongoDB is encrypted and exposed to the public in default setup.
> For local-only use, remove certificate mounts and expose via local ports in `docker-compose.yml`.

## Front-end Setup Instructions
### 1. Clone the repository
```bash
git clone git@github.com:lareii/siker.im.git
```

### 2. Navigate to the client directory
```bash
cd siker.im/client
```

### 3. Install dependencies
```bash
npm install
```

### 4. Create a `.env` file
```bash
cp .env.example .env
```
Edit it as needed for your environment.

### 5. Start the development server
```bash
npm run dev
```

### 6. Access the client
Open your browser and go to `http://localhost:3000`.

### Front-end Environment Variables
Located in `.env` inside `/client`.

| Name                             | Description                                         |
| -------------------------------- | --------------------------------------------------- |
| `NEXT_PUBLIC_BASE_URL`           | Public-facing URL of your site                      |
| `NEXT_PUBLIC_API_URL`            | Base URL of the API (e.g., `http://localhost:1209`) |
| `NEXT_PUBLIC_TURNSTILE_SITE_KEY` | Cloudflare Turnstile site key                       |

## Additional Notes
- Deploy the frontend easily via [Vercel](https://vercel.com/) and set the API URL in their environment settings.
- Set up HTTPS with Nginx for reverse proxying to the backend.
- Secure MongoDB and Redis ports — do not expose them to the public unless behind firewalls or authentication.

## License
This project is licensed under the AGPL-3.0 License. See the [LICENSE](./LICENSE) file for details.
