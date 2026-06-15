<p align="center">
    <img width="250" height="250" src="assets/Jenny.png">
</p>
<h1 align ="center">
    Taskjrnl
</h1>

Taskjrnl is a simple command line tool created to help manage tasks.

### How to install
Make sure you have **GO 1.25.0 or later** installed on your system.

1. *Clone the repository in whichever way that you want or simply do the following.*
```bash
git clone https://github.com/aldoeve/taskjrnl.git
```
2. *Install the application direclty with GO.*
```bash
go install ./cmd/taskjrnl/
```
You can also run the following Makefile command to do the same.
```bash
make install
``` 
3. *Recommended: Create a symlink. You can use the following Makefile command:*
```bash
make symlink
```

### Usage
To get help or more information on the different functions of the application run:
```bash
taskjrnl help
```
Create a task with `add`. Pass it the name the of task then give it a priority and tag to in any order.

See all tasks with `list`.

Add an additional note to a task with `jrnl`.

See more about a task and its notes with `info`.

Modify a task with `modify`.

Give more importance to a task beyond just a priority tag with `weight`.

Create a note between two tasks with `link`.

# Art Work Credits
The main app mascot, Jenny, was created by 
Momomitsuko. You can commission them [here](https://www.etsy.com/shop/Momomitsuko). They do great work on custom profile images, icons, and vtuber models. They have both an [instagram](https://www.instagram.com/momo.mitsuko) and [X](https://x.com/momo_mitsuko).
