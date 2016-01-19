__author__ = 'Lars und Zach'

import thread
import time


i = 0

def thread_1():
    for count in range(0, 1000000):
        global i
        i += 1

def thread_2():
    for count in range(0, 1000000):
        global i
        i -= 1

thread.start_new_thread(thread_1, ())
thread.start_new_thread(thread_2, ())
time.sleep(1)
print i
