# COMP385: Go Exercises

This project consists of a sequence of exercises to help you learn Go.

Each directory in this project contains an exercise or set of exercises to be
completed:

* `sqrt/` -- Implement a square root function using Newton's method. This is the
  first exercise that should be completed, and will practice using variables and
  control structures.
* `modules/` -- Implement an array-based implementation of a list interface.
  This is the second exercise that should be completed, and practices defining
  methods for structs and implementing/using interfaces.
* `concurrency-exercises/` -- This is a collection of 5 exercises where you will
  practice using concurrency primitives, like creating threads/goroutines using
  `go`, syncrhonizing threads using locks/mutexes, and using channels. Each
  exercise contains its own README with instructions for how to complete the
  exercise.
* `webchat/` -- This final exercise asks you to implement an RPC-based
  multi-user chat server. You will implement a simple RPC server and then call
  these RPCs from a client. A browser-based interface will allow you to test and
  use the chat server.

## Working with Git

You __MUST__ work in a branch other than main for this project.

Each exercise asks you to commit and push your work to your GitHub repository
after you complete it in order to checkpoint your work. Do not create a pull
request until you are done with the __entire__ project.

To clone the repo and create and checkout a new branch, you will run the
following:

```
git clone <your repo URL>
git checkout -b <your branch name>
```

When you are ready to commit and push, you'll run the following commands to 1)
list the files that have been changed, 2) add the files that you want to commit,
3) make the commit, and finally 4) push your changes to GitHub.

```
git status
git add <files you modified>
git commit -m '<commit message>'
git push --set-upstream origin <your branch name>
```

Note that you only need the `--set-upstream origin` part _the first time_ you
push, subsequent pushes can be done with `git push <your branch name>`.

## Using GoLand and Git

You can also use GoLand or another IDE that integrates with GitHub. GoLand
actually has a browser plugin
([Chrome](https://chrome.google.com/webstore/detail/jetbrains-toolbox-extensi/offnedcbhjldheanlbojaefbfbllddna?hl=en),
[Firefox](https://addons.mozilla.org/en-US/firefox/addon/jetbrains-toolbox/))
that will allow you to clone a repository directly from the GitHub page for the
repository.

![Browser
plugin](https://storage.googleapis.com/distributed-files/jetbrains-plugin.png)

Since your repository is private, you'll have to log in to your account from
your IDE in order to clone it.

If you are choosing to use git from the command line to clone your repository,
you can still use Goland with your project. Just open the directory where you
cloned your project in GoLand.

The __very first__ thing that you do when you open the project in GoLand should
be to create a new branch.

![New branch
menu](https://storage.googleapis.com/distributed-files/git-branch.png)

Then, you can choose to commit and push from the toolbar:

![Toolbar git
buttons](https://storage.googleapis.com/distributed-files/git-buttons.png)

Before you commit, you should verify that you are committing to the branch that
you checked out and __not main__.

![Changelist
branch](https://storage.googleapis.com/distributed-files/git-changelist-branch.png)

## Building and running in GoLand

__First, will need to enable modules support in GoLand.__ In your project, go to
preferences, and under Go, you will find a "Go Modules" menu item. Make sure
"Enable Go modules integration" is checked.

![Preferences
menu](https://storage.googleapis.com/distributed-files/goland-modules.png)

To run parts of the project within GoLand, you should right click on the folder
that you want to run or test, and then choose the appropriate option from the
"Run" menu. Each part of the project also contains instructions for how to run
it.

__If GoLand has trouble finding external libraries like `glog`__ you will need
to enable GOPATH indexing. You should get a tooltip to enable this when hovering
over the import statement of `glog`. __Note that you won't encounter this until
part 3 (concurrency exercises)__.
