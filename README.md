# aggreGATOR (Gator CLI)

A terminal RSS reader that stores feeds and posts in Postgres.

## Prerequisites

- Go 1.21+ installed: https://go.dev/dl/
- PostgreSQL 14+ running and accessible
- A Postgres database and user you can connect with

## Installation

```sh
go install github.com/Ikit24/aggreGATOR@latest
```

Ensure `$HOME/go/bin` is on your PATH so the `gator` command is available.

## Configuration

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://USERNAME:PASSWORD@localhost:5432/DBNAME?sslmode=disable",
  "current_user_name": ""
}
```
## Usage

### Register and set current user

```sh
gator users register <username>
gator users login <username>
```

### Add and follow a feed

```sh
gator feeds add https://example.com/feed.xml "Example Feed"
gator follow https://example.com/feed.xml
```

### Aggregate/scrape feeds

```sh
gator agg
```

### Browse recent posts

Default limit is 2 posts:

```sh
gator browse
```

View more posts:

```sh
gator browse --limit 10
```

### List feeds and follows

```sh
gator feeds list
gator follows list
```

### Unfollow a feed

```sh
gator unfollow https://example.com/feed.xml
```

### Reset (dangerous)

```sh
gator reset
```

## Troubleshooting

- "command not found": Add `$HOME/go/bin` to your PATH.
- DB connection issues: Verify `db_url` in your config file and ensure Postgres is running.
- Permission errors: Ensure your Postgres user has permissions to create tables.
