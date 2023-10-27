#!/usr/bin/python3
from os import fork, getpid, getppid, wait, execv
from sys import argv
from time import sleep
from random import randint

s = int(argv[1])
print('Сhild[{}]: I am started. My PID {}. Parent PID {}.'.format(getpid(),getpid(),getppid()))
sleep(s)
print('Child[{}]: I am ended. PID {}. Parent PID {}.'.format(getpid(),getpid(),getppid()))
_exit(randint(0, 1))
