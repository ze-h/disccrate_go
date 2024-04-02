# DiscCrate

DiscCrate provides a convenient and efficient way to manage your record collection, with support for multiple users and Discogs integration to make it even easier to catalog your music!

## Building

This project is written in Go, and building can be done using Go's inbuilt tools.

There are three required configuration files for the project:

- db.cfg: the DSN for your MySQL Server
- admin.cfg: the MD5 hash of your administrator password
- discogs.cfg: the user agent and key of your Discogs account

Each are in the format KEY=value, with keys separated by newlines

The location of these files is determined by config.cfg, which is the local directory by default
