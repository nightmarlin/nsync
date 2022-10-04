---
title: 001 - Initial Criteria
authors:
  - Lewis W. Miller
description: >
  The initial criteria for this project, transcribed from my notebook
keywords:
  - requirements
---

## The goal

> Connect a Calendar to a Notion DB.

**_What does this mean?_** After all, Calendar Items and Notion DBs are quite different...

| Calendar Items   | Notion DBs |
|------------------|------------|
| Events           | Pages      |
| Tasks            | Properties |
| Reminders        | Views      |
| Out-of-Office    |            |
| Working Location |            |

## What should I be able to do?

- Sign up/in

  - With Google account?
  - With Notion account?
  - With a custom system?
    - Probably not - this needs to be secure...

  > Which is the source of truth?

  Create relations between calendars and databases.

- Sync status

  - Which calendars are synced with which dbs
  - Whether syncs are up-to-date
  - Force a sync

- Sync settings
  - Select calendars and DBs to sync
  - Define relationships/mappings between DBs and calendars

## MVP

> What is the minimum successful product?

Sync events from one Google Calendar to a Notion DB

- Predefined properties, no customization
- Google is source of truth
- One GCal to one Notion DB
- Events only
- Properties to sync
  - Location
  - Date, TIme & TZ
  - Free/Busy status
  - Public/Private visibility
  - Description
  - Colour
- Not to sync
  - Notifications
  - Attendees
  - Repeats
  - Attachments
  - Video links
  - etc...

Connecting will result in the db being fully overwritten.
