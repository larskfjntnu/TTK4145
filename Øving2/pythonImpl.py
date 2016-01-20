__author__ = 'Lars und Zach'

import thread
import time
import threading


i = 0
mutex = threading.Lock()

def thread_1():
    mutex.acquire()
    for count in range(0, 100000):
        global i
        i += 1
    mutex.release()

def thread_2():
    mutex.acquire()
    for count in range(0, 1000000):
        global i
        i -= 1
    mutex.release()

thread.start_new_thread(thread_1, ())
thread.start_new_thread(thread_2, ())
time.sleep(1)
print i
