# Gator

A command-line RSS feed aggregator built with Go and PostgreSQL. Gator allows you to track and read posts from your favorite RSS feeds directly from the terminal.

## Features

- **User Management**: Register and login to manage your personal feed collection
- **Feed Management**: Add and list RSS feeds
- **Follow System**: Follow and unfollow feeds to curate your reading list
- **Automated Aggregation**: Periodically fetch and store new posts from followed feeds
- **Browse Posts**: Read collected posts from your followed feeds
- **PostgreSQL Storage**: Persistent storage using PostgreSQL database
- **Type-Safe Queries**: Generated database code using sqlc

## Prerequisites

- Go 1.25 or higher
- PostgreSQL database
- Database connection string

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd gator
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your PostgreSQL database and apply migrations from `sql/schema/`

4. Configure your database connection:
   - The application reads configuration from `~/.gatorconfig.json`
   - Ensure your config file contains the database URL

## Usage

### User Commands

**Register a new user:**
```bash
gator register <username>
```

**Login as an existing user:**
```bash
gator login <username>
```

**List all users:**
```bash
gator users
```

### Feed Commands

**Add a new RSS feed:**
```bash
gator addfeed <feed_name> <feed_url>
```

**List all feeds:**
```bash
gator feeds
```

**Follow a feed:**
```bash
gator follow <feed_url>
```

**List feeds you're following:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow <feed_url>
```

### Aggregation

**Start the feed aggregator:**
```bash
gator agg <duration>
```

The aggregator will fetch new posts from feeds at the specified interval.

Example:
```bash
gator agg 1m  # Fetch every 1 minute
gator agg 30s # Fetch every 30 seconds
gator agg 1h  # Fetch every hour
```

### Browse Posts

**Browse collected posts:**
```bash
gator browse
```

### Reset

**Reset the database:**
```bash
gator reset
```

## Project Structure

```
gator/
├── main.go              # Application entry point
├── command.go           # Command handling
├── users.go             # User-related handlers
├── feeds.go             # Feed-related handlers
├── aggregator.go        # RSS feed scraping logic
├── rss.go              # RSS parsing
├── middleware.go        # Authentication middleware
├── reset.go            # Database reset handler
├── internal/
│   ├── config/         # Configuration management
│   └── database/       # Database models and queries (sqlc generated)
└── sql/
    ├── queries/        # SQL queries for sqlc
    └── schema/         # Database migrations
```

## Database Schema

The application uses the following main tables:

- **users**: User accounts
- **feeds**: RSS feed sources
- **feed_follows**: User-feed relationships
- **posts**: Collected RSS posts

## Technologies Used

- **Go**: Primary programming language
- **PostgreSQL**: Database
- **sqlc**: Type-safe SQL code generation
- **github.com/lib/pq**: PostgreSQL driver
- **github.com/google/uuid**: UUID generation

## Development

This project uses sqlc for generating type-safe database code. To regenerate the database code after modifying queries:

```bash
sqlc generate
```

## License

This is a training project.

## Acknowledgments

Built as part of the Boot.dev backend development course.