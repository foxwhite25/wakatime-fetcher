# Wakatime Fetcher

This is a simple script to fetch your wakatime stats yesterday and save it to a database.

Currently, it only supports sqlite3, but pull requests are welcome.

## Usage

### Standalone

If you want to run it on your own machine, you can do so by running the following commands:

```bash
git clone
cd wakatime-fetcher
go run wakatime <your api key>
```

This will create a file called `wakatime.tar.gz` in the current directory. Which contains the sqlite3 database.

### Workflow

If you want to run it as a workflow, first you need to fork this to a private repository (or if you want your data to be publicly available, you can fork it to a public repository).

#### The GitHub way

You can click the use this template button on the top right of the repository. This will create a new repository with the same files as this one.

#### The Git way

What you need to do is create a new private repo via the [Github UI](https://github.com/new). (Lets call it `private-repo` here)

Then you can run these command to fork this repo to your private repo:

```bash
git clone --bare https://github.com/foxwhite25/wakatime-fetcher.git
cd wakatime-fetcher.git
git push --mirror https://github.com/<your username>/private-repo.git
cd ..
rm -rf wakatime-fetcher.git
```

#### Setting up the workflow
Now that you have your private repo, you can modify the workflow file to your liking. Start by cloning your private repo:

```bash
git clone https://github.com/<your username>/private-repo.git
cd private-repo
```

By default, the workflow file only runs on trigger, but you can change it to run on a schedule by changing the `on` field to `schedule` and adding a cron expression like this:

```yaml
on:
  schedule:
    - cron: '0 0 * * *' # Runs every day at midnight
```

Now you can commit and push the changes to your private repo like any other git repo.

```bash
git add .
git commit -m "Change workflow to run on schedule"
git push origin master
```

You will need to get your wakatime api key. You can get it from [here](https://wakatime.com/api-key). 

When you got the key, create a new secret in your private repo called `WAKATIME_API_KEY` and set the value to your api key.

Now you can go to the actions tab of your private repo and enable the workflow, and go to setting and enable workflow write access to the repo.

#### Pull Changes From This Repo
To pull new hotness from the public repo:
```bash
cd private-repo
git remote add public https://github.com/foxwhite25/wakatime-fetcher.git
git pull public master # Creates a merge commit
git push origin master
```

### Database

The database contains a single table called `heartbeats` with the following schema:

```sqlite
CREATE TABLE heartbeats (
    id TEXT PRIMARY KEY,
    branch TEXT, 
    category TEXT, 
    created_at TEXT,
    cursorpos INTEGER, 
    dependencies TEXT,
    entity TEXT, 
    is_write INTEGER, 
    language TEXT, 
    lineno INTEGER, 
    lines INTEGER, 
    machine_name_id TEXT, 
    project TEXT, 
    project_root_count INTEGER, 
    time REAL, 
    type TEXT, 
    user_agent_id TEXT, 
    user_id TEXT
);
```

### Data Analysis

You can use the data in the database to do some data analysis. 

The data are heartbeats, which means that it only contains data when you are coding, you need to do some data processing to get the total time you spent coding.

What wakatime does is that it groups the heartbeats by project and/or language, and then sum up the time spent on each project and language.

Note that the time here refers to the timestamp of the heartbeat, not the time spent on the project.

A session in wakatime is defined as a period of time when you are coding on the same project and/or language, which stops when heartbeats stops for more than 25 minutes.