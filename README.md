# Gator

Gator is a command-line RSS feed aggregator written in Go. It allows users to add and follow RSS feeds, periodically fetch posts, store them in PostgreSQL, and browse posts directly from the terminal.

## Requirements

To run Gator, you need:

- Go
- PostgreSQL

Make sure PostgreSQL is installed and running before using Gator.

## Installation

Install the Gator CLI using:

```bash
go install github.com/Agron123/blogaggregator@latest
```

After installation, you can run Gator directly from the terminal:

```bash
gator
```

## Configuration

Create a `.gatorconfig.json` file in your home directory:

```json
{
  "db_url": "postgres://USERNAME:PASSWORD@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `USERNAME` and `PASSWORD` with your PostgreSQL credentials.

Make sure the `gator` PostgreSQL database exists before running the program.

## Usage

### Register a user

```bash
gator register username
```

### Login

```bash
gator login username
```

### List users

```bash
gator users
```

### Add a feed

```bash
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
```

The user who adds the feed will automatically follow it.

### List feeds

```bash
gator feeds
```

### Follow a feed

```bash
gator follow "https://news.ycombinator.com/rss"
```

### View followed feeds

```bash
gator following
```

### Unfollow a feed

```bash
gator unfollow "https://news.ycombinator.com/rss"
```

### Aggregate posts

Start fetching posts from the feeds:

```bash
gator agg 30s
```

The duration controls how often Gator fetches the next feed. The aggregator will continue running until it is stopped with `Ctrl+C`.

### Browse posts

Show the latest posts from feeds followed by the current user:

```bash
gator browse
```

By default, Gator displays 2 posts.

You can specify a different limit:

```bash
gator browse 10
```

### Reset the database

```bash
gator reset
```

This removes the application's stored user data.