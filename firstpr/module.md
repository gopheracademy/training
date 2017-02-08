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
```

---

# Resources

- [Your First PR](http://yourfirstpr.github.io)
- [Fork A Repository](https://help.github.com/articles/fork-a-repo/)
- [Github Pull Request Tutorial](https://www.thinkful.com/learn/github-pull-request-tutorial/)
t Etiquette

- Links to articles on Go PR's and forking, etc
