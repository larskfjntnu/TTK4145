#include <stdio.h>
#include <string.h>
#include <pthread.h>
#include <stdlib.h>
#include <unistd.h>

/*
	To compile : gcc -o OutputFile cFile
	To run: ./OutputFile
*/
	
int i = 0;
pthread_mutex_t lock; // The mutex thread that is the lock.

void *thread_1(){
	pthread_mutex_lock(&lock);
	for (int count = 0; count < 1000001; count++)
	{
		i++;
	}
	printf("%s\n", "Done incrementing");
	pthread_mutex_unlock(&lock);
	pthread_exit(NULL);
}

void *thread_2(){
	pthread_mutex_lock(&lock);
	for (int count = 0; count < 1000000; count++){
		i--;
	}
	printf("%s\n", "Done decrementing");
	pthread_mutex_unlock(&lock);
	pthread_exit(NULL);
}

int main(void){
	// Initialize the lock
	if(pthread_mutex_init(&lock, NULL) != 0){
		printf("%s\n", "Error initializing mutex" );
		return 1;
	}

	// Create the threads
	pthread_t thr[2];
	pthread_create(&thr[0], NULL, thread_1, NULL);
	pthread_create(&thr[1], NULL, thread_2, NULL);

	printf("%s\n", "Waiting to join threads" );
	/* block until all threads complete */
  for (int count = 0; count < 2; count++) {
    	pthread_join(thr[count], NULL);
  }
  pthread_mutex_destroy(&lock);
  printf("%i\n",i);

}