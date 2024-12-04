# Welcome to the AggreGATOR 🐊

The goal of this blog aggregator project is to allow you to follow your favorite blog feeds and be up to date to the latest posts

To run this project you will need to have Go and PostgreSQL installed

## Installing Go

- Simply run the follow command in your terminal
```
curl -sS https://webi.sh/golang | sh
```

Or 

- Follow the [official installation guide](https://go.dev/doc/install)

## Installing PostgreSQL

If you use Linux you can run the following command in your terminal

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

If you use MacOS you can run
```
brew install postgresql@15
```


## Finish Gator installation

Navigate to the location you cloned the repo through your terminal and run the following:

```
go install .
```

after that you can use the commands through 

```
blog_aggregator <command> <param>
```

# How to use

There are serveral commands available, below you can find the complete list:

- `login`: `blog_aggregator login <user_name>` (keep in mind user_name must be registered)
- `register`: `blog_aggregator register <user_name>` (you can't register same user twice)
- `reset`: `blog_aggregator reset` - this will reset the program to initial state (you will lose all your data)
- `users`: `blog_aggregator users` - will list all the users registered and the current user logged in
- `agg`: `blog_aggregator agg <time_between_request>` - this is the command you will wanna keep running, to always keep your posts up to date
- `addfeed`: `blog_aggregator addfeed <feed_name> <feed_url>` - will add a feed by name and url to the available feeds
- `feeds`: `blog_aggregator ` - 
- `follow`: `blog_aggregator ` - 
- `following`: `blog_aggregator ` - 
- `unfollow`: `blog_aggregator ` - 
- `browse`: `blog_aggregator ` - 


## Suggestions / Issues

In case you have any suggestion or any problem, please contact me through my email gabrieldsschneider@gmail.com