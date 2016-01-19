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

void *thread_1(){
	for (int count = 0; count < 1000000; count++)
	{
		i++;
	}
	pthread_exit(NULL);
}

void *thread_2(){
	for (int count = 0; count < 1000000; count++){
		i--;
	}
	pthread_exit(NULL);
}

int main(void){
	pthread_t thr[2];
	pthread_create(&thr[1], NULL, thread_1, NULL);
	pthread_create(&thr[2], NULL, thread_2, NULL);

	printf("%s\n", "Waiting to join threads" );
	/* block until all threads complete */
  for (int count = 0; count < 2; count++) {
    	pthread_join(thr[count], NULL);
  }

  printf("%i\n",i);

}