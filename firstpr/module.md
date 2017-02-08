class: center, middle
name: Slide Name
video: 
description: 
level: Beginner
topic: Go

# Your First PR
## Contributing to a Go OSS Project
---
video: 

# Agenda

- Git Etiquette
- Mechanics of a Pull Request
- Go Specifics
- Resources


---

# Hooray!

You've decided to start contributing to Open Source projects.  Great decision!

Open Source has a lot of benefits, and you'll soon enjoy them too.  My favorite:

- Giving Back
- Making Friends
- Learning New Things
- New Ideas - MASHUPS!

---

# Git Etiquette

Before you dive in to an Open Source project and start submitting pull requests it's important to consider the social aspect of contributing to Open Source.


We'll discuss a few scenarios briefly, and the generally accepted way to communicate in each scenario.

---
# Git Etiquette

## Scenario 1:  You found a bug!

You've found a bug and you want to fix it.  

YOU'RE AWESOME!

- Step 1: Open an Issue, suggest a fix, and that you're willing to submit a PR
- Step 2: Wait for a response to your issue. 
- Step 3: Read the response, and if green-lighted, make the fix.

---
# Git Etiquette

## Scenario 2:  You want to make an enhancement.

Every Open Source project's favorite thing -- someone else wants to write code!

YOU'RE AWESOME!

- Step 1: Open an Issue, suggest your enhancement, and that you're willing to submit a PR.
- Step 2: Wait for a response to your issue. 
- Step 3: Read the response, and if green-lighted, make the fix.  Pay close attention to any direction the maintainers give you, and follow their guidance.

---
# Git Etiquette

## For ALL Pull Requests:

Check the documentation in the README.md or CONTRIBUTING.md files.  Often they'll have very explicit instructions on how to proceed with a pull request.  Follow these guidelines, and you'll have a higher chance of a successfully merged pull request.  Ignore them, and prepare to have your PR deleted :)

This is the most important step of all.  Pay attention when the project tells you how to interact with them.


---
# Mechanics of a Pull Request

Now that you've done the communication part of your open source adventure, it's time to start making changes to source code.  Let's talk about how a Pull Request works.

.center[![Pull Request](http://www.gulistanboylu.com/wp-content/uploads/pullrequest1-611x400.png)]

---
# Step One

## Create a Fork or a Branch

Generally, you're going to create a `fork` of a repository and make your changes there.  It's the easiest and cleanest way to contribute changes.

If you are a member of the project team, you may have permission to create a branch instead.  A `branch` works in the same git repository, a `fork` uses a different repository.

.center[![Fork It](https://help.github.com/assets/images/help/repository/fork_button.jpg)]

---
# Step Two

## Download Your Fork

If you hit the `fork` button on Github, you'll end up with a copy of the project under your Github account.

Now you need a local copy of that fork in order to make changes.

---
# Step Two

## Download Your Fork
Go to the Github page of `your copy` of the project and find the big green "clone or download" link:

.center[![clone](https://help.github.com/assets/images/help/repository/clone-repo-clone-url-button.png)]

Click that, and choose "SSH" if given the option, then copy the repository location to your clipboard.

.center[![clone2](https://help.github.com/assets/images/help/repository/https-url-clone.png)]

---
# Step Two

## Download Your Fork

Now open your terminal and navigate to the directory where you keep your projects.  Type this:

```shell
> git clone (pasted url here)
```

---
# Step Three

## Get to Work!

Start by making a branch to keep your changes self-contained.

```
> git checkout -b my_branch_name
```
It's very common to use a descriptive name for your branch, and polite too.  Good names:

- update-homepage
- fix-model-errors
- add-download-feature


---
# Step Three

## Get to Work!

Make your changes.  Follow the style of the code around you as a guide.  If other variable names are short, use short variable names.  Use the same patterns in code and filenames that the rest of the project uses.  

When in doubt, ask a question in the issue you raised in the first step!  Communication helps to prevent frustration.


---
# Step Three

## Tests, Tests, and Tests

If the project you've forked has a test suite, run the test suite after you've made your changes to ensure you haven't broken anything.  You may need to update tests to reflect the new code you've written.

Add tests that cover the new code you've written, or the bug you've fixed.  

*A specific test that highlights the fix to a bug is a wonderful, beautiful thing to add to a project.*


---
# Step Four

## Check In Your Code

Look for the files you've modified with `git status`.

``` shell
> git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   firstpr/module.md

no changes added to commit (use "git add" and/or "git commit -a")
```


---
# Step Four

## Add Your Changes to a Commit

Use `git add` to add your changes to a commit.  Use a well written commit message that explains what you did, why you did it, and what it adds/fixes.  Mention the Github issue you opened in the commit message:

```
> git add firstpr/module.md
```
```
> git commit -m"This change adds a new button to the home page 
allowing users to download their pictures as a zip file.  Fixes #12"
```

---
# Step Four

## Push Your Changes to Your Fork

Now that you've added your changes to git's staging area, you need to push them up to your Github fork.

```
> git push origin master
Counting objects: 4, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 2.82 KiB | 0 bytes/s, done.
Total 4 (delta 2), reused 0 (delta 0)
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To github.com:gopheracademy/training.git
   e61a980..1e7fd72  master -> master
Branch master set up to track remote branch master from origin.
```

---
# Step Five

## Open Your Pull Request

Navigate to your fork on Github and click on the `Pull Requests` tab.

Click the "Pull Requests" tab, then the green "New pull request" button.

---

# Step Five

Give it a clear title, and a well worded description.  Then click "Send pull request".

.center[![PR](http://www.pontikis.net/blog/media/2013/03/how-to-collaborate-on-github-open-source-projects/post/pull_request_submit_github_thumb.png)]

---

# Step Five

Wait.  Open Source maintainers are people with lives just like you.  If you're lucky they may act on your request right away.  Sometimes it might take a few days.  Be patient, never rude! 

They may suggest changes.  If they do, make the changes in your fork and push them again.  The PR will automatically be updated.  *magic*

---

# Go Specifics

If only life were always simple...

Because of the `GOPATH`, when you fork a Go project, it's almost guaranteed not to work if you put it in `your` Github source directory.  Why?

The `go` toolchain uses the directory structure of your `GOPATH` to determine a library's import path.  My fork of `spf13/viper` lives in `bketelsen/viper`.  That will never work.  

---

# Go Specifics

You have two choices, and one of them is really bad.

- Change all the import statements in your fork to match the GOPATH of your fork
- Put your fork in the original project's import path.

Don't do the first one.  Ever.  Just Don't.


---

# Go Specifics

So what does `put your fork in the original project's import path` mean?

It means your Github repository is checked out in the upstream project's import path.  Using my example above, I would checkout `bketelsen/viper` at `github.com/spf13/viper`.

Now the project lives where it expects to be, and everything is right with the world!

ALMOST.


---

# Go Specifics

The last problem is what to do about the original project.  You don't want your fork to be used as `spf13/viper` when you build your apps that depend on it.  

Again, you have two choices.

- Add a git remote to the upstream project
- Delete your fork when you're done with your fix, and replace with the upstream original project.


---

# Go Specifics

The choice you make for this problem depends on how often you think you'll be working on that project.  If you are making a single change and expect no further changes, you should probably just delete your fork and `go get` the original project when your change is merged.

If you intend to be making a lot of changes to this project in the future, you need to use some more git magic.

---

# Go Specifics
## Git Remotes

Git is a distributed version control system.  It's decentralized in nature... unless everyone is using Github, then it's centralized in it's decentralization :)

You can have many different remote repositories as targets for pulling and pushing. 

To see the remotes configured for a git repository type:

```
> git remote -v
origin	git@github.com:bketelsen/spf13.git (fetch)
origin	git@github.com:bketelsen/spf13.git (push)
```
This shows that the `origin` remote of my repository is my fork of `spf13/viper`.  But I said  you can have many remotes.  Let's add another!

---

# Go Specifics
## Add New Git Remote

To add a new remote representing the original project `spf13/viper`, issue the `git remote` command:

```
> git remote add upstream https://github.com/spf13/viper
```
This adds a new remote named `upstream` that points to the original project.  Now I can have my fork of `spf13/viper` living in the correct import path, but cloned from MY Github repository, and STILL keep track of changes to the upstream branch.

---

# Go Specifics
## Managing Remotes

Now when you want to work on a change, you can create a new branch, make your change, send your pull request and follow the workflow we've shown.  When you're done with your change, just change to the `upstream master` branch, which represents the original project.

```
> git fetch upstream master
> git checkout upstream/master
```
`fetch` pulls the changes from the upstream repository.

`git checkout upstream/master` makes the upstream/master branch the current branch in your workspace.

---

# Resources

- [http://yourfirstpr.github.io](http://yourfirstpr.github.io)
- [https://help.github.com/articles/fork-a-repo/](https://help.github.com/articles/fork-a-repo/)
- [https://www.thinkful.com/learn/github-pull-request-tutorial/](https://www.thinkful.com/learn/github-pull-request-tutorial/)
t Etiquette

