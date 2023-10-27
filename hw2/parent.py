#!/usr/bin/python3
from os import fork, getpid, getppid, wait, execv
from sys import argv
from random import randint

n = int(argv[1])
child = fork()
i = 0
while True:
    if child == 0:
        number = randint(5, 10)
        execv('/usr/bin/python3', ['python3', 'child.py', str(number)])
    else:
        print('Parent[{}]: I ran children process with PID {}.'.format(getppid(), getpid()))
        ret = wait()
        print('Child with PID {0} terminated. Exit Status {1}.'.format(ret[0], ret[1]))
        i += 1
        if i < n:
            child = fork()
        else:
            break
