# The Void

## Introduction

Every once in a while, every one feels the need to let their thoughts out.
Sometimes, you might want that to be to a friend or family, but other times it
might be something you just want to put out into the world and allow it to
disppear. This idea of a shout into the void is where the idea for this has
come from.

The plan for this is for someone to come along and order the thoughts in their
head through writing them out. The user then has a choice around whether they
want to keep their writings private or public. At the time of writing, the plan
is for this to be truly anonymous. No logins, or usernames, just a place to
come and type. As such, if they choose for what they've written to be private,
it's not saved anywhere, it's just sent nowhere. If they opt for it to be
public, then it will be stored, and allowed to surface elsewhere on the app.

Due to the anonymous nature of this project, it would be optimistic to hope
that it wouldn't be abused. Moderation is going to be key. In the first
instance, I would look to implement this through user reporting, tracking the
number of reports and the reasons provided. This raises three immediate
questions:

1. How to differentiate between frivolous reports versus actual issues?
    * In turn, what constitutes frivolous and actual?
2. What to do when reports are genuinely offensive?

My initial proposal is a shadow block, where the post is quietly removed from
the feed before a review. This should be automated in the first instance, with
more ambiguous cases being flagged for closer inspection. If this grows, then
this will need a more careful review.

## Development

The plan is for the backend and frontend to be written in Go and JavaScript
respectively. For the frontend, I will use Vue as a framework. For the database
I will be using Postgres. Deployment will be handled using Docker compose.

As I work on developing this, the Go backend will have a list of posts to serve
as a database for development. This is because I am using this project to help 
me learn, and I don't want to go down a rabbit hole focusing on getting the
database side to work too early on. I would prefer to have a functioning
backend and frontend, before then focusing on implementing a database.

## Changelog

The format of this is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and may be broken out into its own file further down the line.
